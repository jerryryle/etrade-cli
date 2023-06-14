package etradelib

import (
	"errors"
	"github.com/dghubble/oauth1"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"net/url"
)

type ETradeSession interface {
	Renew(accessToken string, accessSecret string) (client.ETradeClient, error)
	Begin() (string, error)
	Verify(verifyKey string) (eTradeClient client.ETradeClient, accessToken string, accessSecret string, err error)
}

type eTradeSession struct {
	config         OAuthConfig
	urls           client.EndpointUrls
	consumerKey    string
	consumerSecret string
	requestToken   string
	requestSecret  string
	accessToken    string
	accessSecret   string
	logger         *slog.Logger
}

func CreateSession(production bool, consumerKey string, consumerSecret string, logger *slog.Logger) (ETradeSession, error) {
	if consumerKey == "" || consumerSecret == "" {
		return nil, errors.New("invalid consumer credentials provided")
	}
	urls := client.GetEndpointUrls(production)

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
		urls:           urls,
		config:         &config,
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
		requestToken:   "",
		requestSecret:  "",
		accessToken:    "",
		accessSecret:   "",
		logger:         logger,
	}, nil
}

func (s *eTradeSession) Renew(accessToken string, accessSecret string) (client.ETradeClient, error) {
	s.accessToken = accessToken
	s.accessSecret = accessSecret

	if s.accessToken == "" || s.accessSecret == "" {
		return nil, errors.New("invalid access credentials provided")
	}
	token := oauth1.NewToken(s.accessToken, oauth1.PercentEncode(s.accessSecret))
	httpClient := s.config.Client(oauth1.NoContext, token)
	response, err := httpClient.Get(s.urls.RenewAccessTokenUrl())
	if response != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				s.logger.Error(err.Error())
			}
		}(response.Body)
	}
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("invalid access token")
	}
	return client.CreateETradeClient(s.urls, httpClient, s.logger), nil
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

func (s *eTradeSession) Verify(verifyKey string) (eTradeClient client.ETradeClient, accessToken string, accessSecret string, err error) {
	s.accessToken, s.accessSecret, err = s.config.AccessToken(s.requestToken, oauth1.PercentEncode(s.requestSecret), verifyKey)
	if err != nil {
		return nil, "", "", err
	}
	token := oauth1.NewToken(s.accessToken, oauth1.PercentEncode(s.accessSecret))
	httpClient := s.config.Client(oauth1.NoContext, token)
	return client.CreateETradeClient(s.urls, httpClient, s.logger), s.accessToken, s.accessSecret, nil
}
