package etradelib

import (
	"context"
	"github.com/dghubble/oauth1"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/url"
)

type oAuthConfigMock struct {
	mock.Mock
}

func (m *oAuthConfigMock) Client(ctx context.Context, t *oauth1.Token) *http.Client {
	args := m.Called(ctx, t)
	return args.Get(0).(*http.Client)
}

func (m *oAuthConfigMock) RequestToken() (requestToken, requestSecret string, err error) {
	args := m.Called()
	return args.String(0), args.String(1), args.Error(2)
}

func (m *oAuthConfigMock) AuthorizationURL(requestToken string) (*url.URL, error) {
	args := m.Called(requestToken)
	return args.Get(0).(*url.URL), args.Error(1)
}

func (m *oAuthConfigMock) AccessToken(requestToken, requestSecret, verifier string) (
	accessToken, accessSecret string, err error,
) {
	args := m.Called(requestToken, requestSecret, verifier)
	return args.String(0), args.String(1), args.Error(2)
}
