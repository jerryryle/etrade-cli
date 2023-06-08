package etradelib

import (
	"errors"
	"github.com/dghubble/oauth1"
	"net/url"
)

type ETradeSession interface {
	Renew(accessToken string, accessSecret string) (ETradeCustomer, error)
	Begin() (string, error)
	Verify(verifyKey string) (customer ETradeCustomer, accessToken string, accessSecret string, err error)
}

type eTradeSession struct {
	customerName   string
	urls           EndpointUrls
	consumerKey    string
	consumerSecret string
	requestToken   string
	requestSecret  string
	accessToken    string
	accessSecret   string
	config         OAuthConfig
}

func CreateSession(customerName string, production bool, consumerKey string, consumerSecret string) (ETradeSession, error) {
	if consumerKey == "" || consumerSecret == "" {
		return nil, errors.New("invalid consumer credentials provided")
	}
	urls := GetEndpointUrls(production)

	authorizeEndpoint := oauth1.Endpoint{
		RequestTokenURL: urls.GetRequestTokenUrl(),
		AuthorizeURL:    urls.AuthorizeApplicationUrl(),
		AccessTokenURL:  urls.GetAccessTokenUrl(),
	}

	config := oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: oauth1.PercentEncode(consumerSecret),
		CallbackURL:    "oob",
		Endpoint:       authorizeEndpoint,
	}

	return &eTradeSession{
		customerName:   customerName,
		urls:           urls,
		config:         &config,
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
		requestToken:   "",
		requestSecret:  "",
		accessToken:    "",
		accessSecret:   "",
	}, nil
}

func (s *eTradeSession) Renew(accessToken string, accessSecret string) (ETradeCustomer, error) {
	s.accessToken = accessToken
	s.accessSecret = accessSecret

	if s.accessToken == "" || s.accessSecret == "" {
		return nil, errors.New("invalid access credentials provided")
	}
	token := oauth1.NewToken(s.accessToken, oauth1.PercentEncode(s.accessSecret))
	httpClient := s.config.Client(oauth1.NoContext, token)
	response, err := httpClient.Get(s.urls.RenewAccessTokenUrl())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// TODO: read body
	return &eTradeCustomer{
		customerName: s.customerName,
		client: &eTradeClientStruct{
			urls:       s.urls,
			httpClient: httpClient,
		},
	}, nil
}

func (s *eTradeSession) Begin() (string, error) {
	var err error
	s.requestToken, s.requestSecret, err = s.config.RequestToken()
	if err != nil {
		return "", err
	}
	// Format and return the authorization string
	authorizeUrl, err := url.Parse(s.urls.AuthorizeApplicationUrl())
	values := authorizeUrl.Query()
	values.Add("key", s.consumerKey)
	values.Add("token", s.requestToken)
	authorizeUrl.RawQuery = values.Encode()
	return authorizeUrl.String(), nil
}

func (s *eTradeSession) Verify(verifyKey string) (customer ETradeCustomer, accessToken string, accessSecret string, err error) {
	s.accessToken, s.accessSecret, err = s.config.AccessToken(s.requestToken, oauth1.PercentEncode(s.requestSecret), verifyKey)
	if err != nil {
		return nil, "", "", err
	}
	token := oauth1.NewToken(s.accessToken, oauth1.PercentEncode(s.accessSecret))
	httpClient := s.config.Client(oauth1.NoContext, token)
	return &eTradeCustomer{
		customerName: s.customerName,
		client: &eTradeClientStruct{
			urls:       s.urls,
			httpClient: httpClient,
		},
	}, s.accessToken, s.accessSecret, nil
}
