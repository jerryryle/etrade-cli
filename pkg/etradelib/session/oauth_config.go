package session

import (
	"context"
	"github.com/dghubble/oauth1"
	"net/http"
	"net/url"
)

type OAuthConfig interface {
	Client(ctx context.Context, t *oauth1.Token) *http.Client
	RequestToken() (requestToken, requestSecret string, err error)
	AuthorizationURL(requestToken string) (*url.URL, error)
	AccessToken(requestToken, requestSecret, verifier string) (accessToken, accessSecret string, err error)
}
