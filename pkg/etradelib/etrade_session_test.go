package etradelib

import (
	"errors"
	"github.com/dghubble/oauth1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestCreateSessionWithEmptyConsumerCredentialsFails(t *testing.T) {
	session, err := CreateSession(false, "", "", CreateNullLogger())
	assert.EqualError(t, err, "invalid consumer credentials provided")
	assert.Nil(t, session)
}

func TestCreateSessionWithConsumerCredentialsSucceeds(t *testing.T) {
	session, err := CreateSession(false, "TestConsumerKey", "TestConsumerSecret", CreateNullLogger())
	assert.Nil(t, err)
	assert.NotNil(t, session)
}

func TestCreateSessionWithEmptyCustomerNameSucceeds(t *testing.T) {
	session, err := CreateSession(false, "TestConsumerKey", "TestConsumerSecret", CreateNullLogger())
	assert.Nil(t, err)
	assert.NotNil(t, session)
}

type ETradeSessionTestSuite struct {
	suite.Suite
	configMock *oAuthConfigMock
	session    ETradeSession
}

func (s *ETradeSessionTestSuite) SetupTest() {
	s.configMock = new(oAuthConfigMock)

	// Create a test session manually, so we can use the mock OAuth config
	s.session = &eTradeSession{
		urls:           GetEndpointUrls(false),
		consumerKey:    "TestConsumerKey",
		consumerSecret: "TestConsumerSecret",
		config:         s.configMock,
	}
}

func (s *ETradeSessionTestSuite) TestNoAccessTokenOrSecretReturnsError() {
	// No Token
	client, err := s.session.Renew("", "TestAccessSecret")
	s.Assert().EqualError(err, "invalid access credentials provided")
	s.Assert().Nil(client)

	// No secret
	client, err = s.session.Renew("TestAccessToken", "")
	s.Assert().EqualError(err, "invalid access credentials provided")
	s.Assert().Nil(client)

	// Neither
	client, err = s.session.Renew("", "")
	s.Assert().EqualError(err, "invalid access credentials provided")
	s.Assert().Nil(client)

	s.configMock.AssertExpectations(s.T())
}

func (s *ETradeSessionTestSuite) TestBadAccessTokenReturnsError() {
	// Create a fake HTTP client that will return 400 (Bad request) for the renewal request
	httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(strings.NewReader(`Bad Request`)),
		}
	})

	// Set up mock Client() call to return the fake HTTP client.
	s.configMock.On("Client", oauth1.NoContext, oauth1.NewToken("TestAccessToken", "TestAccessSecret")).Return(httpClient)

	client, err := s.session.Renew("TestAccessToken", "TestAccessSecret")
	s.Assert().Nil(client)
	s.Assert().Error(err)

	s.configMock.AssertExpectations(s.T())
}

func (s *ETradeSessionTestSuite) TestGoodRenewalSessionReturnsCustomer() {
	// Create a fake HTTP client that will return 200 (Ok) for the renewal request
	httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`Access Token has been renewed`)),
		}
	})

	// Set up mock Client() call to return the fake HTTP client.
	s.configMock.On("Client", oauth1.NoContext, oauth1.NewToken("TestAccessToken", "TestAccessSecret")).Return(httpClient)

	client, err := s.session.Renew("TestAccessToken", "TestAccessSecret")
	s.Assert().Nil(err)
	s.Assert().NotNil(client)

	s.configMock.AssertExpectations(s.T())
}

func (s *ETradeSessionTestSuite) TestBeginNewSessionFailsIfRequestTokenReturnsError() {
	// Set up mock RequestToken() call to return an error
	s.configMock.On("RequestToken").Return("", "", errors.New("mock error"))

	authUrl, err := s.session.Begin()
	s.Assert().Empty(authUrl)
	s.Assert().EqualError(err, "mock error")

	s.configMock.AssertExpectations(s.T())
}

func (s *ETradeSessionTestSuite) TestBeginNewSessionFailsIfAccessTokenReturnsError() {
	// Set up mock RequestToken() call to return a fake token
	s.configMock.On("RequestToken").Return("MockRequestToken", "MockRequestSecret", nil)

	authUrl, err := s.session.Begin()
	s.Assert().Nil(err)
	s.Assert().Equal("https://us.etrade.com/e/t/etws/authorize?key=TestConsumerKey&token=MockRequestToken", authUrl)

	// Set up mock AccessToken() call to return a fake token
	s.configMock.On("AccessToken", "MockRequestToken", "MockRequestSecret", "FakeVerifyKey").Return("", "", errors.New("mock error"))

	client, accessToken, accessSecret, err := s.session.Verify("FakeVerifyKey")
	s.Assert().Nil(client)
	s.Assert().Equal(accessToken, "")
	s.Assert().Equal(accessSecret, "")
	s.Assert().EqualError(err, "mock error")

	s.configMock.AssertExpectations(s.T())
}

func (s *ETradeSessionTestSuite) TestNewSessionSucceeds() {
	s.configMock.On("RequestToken").Return("MockRequestToken", "MockRequestSecret", nil)

	// This is a new session, so Begin() should only return the authUrl
	authUrl, err := s.session.Begin()
	s.Assert().Nil(err)
	s.Assert().Equal("https://us.etrade.com/e/t/etws/authorize?key=TestConsumerKey&token=MockRequestToken", authUrl)

	s.configMock.On("AccessToken", "MockRequestToken", "MockRequestSecret", "FakeVerifyKey").Return("MockAccessToken", "MockAccessSecret", nil)
	s.configMock.On("Client", oauth1.NoContext, oauth1.NewToken("MockAccessToken", "MockAccessSecret")).Return(new(http.Client))

	client, accessToken, accessSecret, err := s.session.Verify("FakeVerifyKey")
	s.Assert().Nil(err)
	s.Assert().NotNil(client)
	s.Assert().Equal(accessToken, "MockAccessToken")
	s.Assert().Equal(accessSecret, "MockAccessSecret")

	s.configMock.AssertExpectations(s.T())
}

func TestETradeSessionTestSuite(t *testing.T) {
	suite.Run(t, new(ETradeSessionTestSuite))
}
