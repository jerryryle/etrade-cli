package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/session"
	"golang.org/x/exp/slog"
	"os"
	"time"
)

func createClientWithCredentialCache(
	production bool, consumerKey string, consumerSecret string, cacheFilePath string, logger *slog.Logger,
) (client.ETradeClient, error) {
	cachedCredentials, err := LoadCachedCredentialsFromFile(cacheFilePath, logger)
	if err != nil {
		// Create a new, empty credential cache. It will yield empty strings for the cached token, which
		// will indicate that there are no cached credentials for this customer
		cachedCredentials = &CachedCredentials{}
	}

	var eTradeClient client.ETradeClient
	authSession, err := session.CreateSession(production, consumerKey, consumerSecret, logger)
	if err != nil {
		return nil, err
	}
	var accessToken = cachedCredentials.AccessToken
	var accessSecret = cachedCredentials.AccessSecret
	eTradeClient, err = authSession.Renew(accessToken, accessSecret)
	if err != nil {
		authUrl, err := authSession.Begin()
		if err != nil {
			return nil, err
		}
		_, _ = fmt.Fprintf(os.Stderr, "Visit this URL to get a validation code:\n%s\n\n", authUrl)

		var validationCode string
		_, _ = fmt.Fprintf(os.Stderr, "Enter validation code: ")
		_, err = fmt.Scanln(&validationCode)
		if err != nil {
			return nil, err
		}
		if validationCode == "" {
			return nil, errors.New("no validation code provided")
		}

		eTradeClient, accessToken, accessSecret, err = authSession.Verify(validationCode)
		if err != nil {
			return nil, err
		}
	}
	err = SaveCachedCredentialsToFile(
		cacheFilePath,
		&CachedCredentials{accessToken, accessSecret, time.Now()},
		logger,
	)
	if err != nil {
		return nil, err
	}
	return eTradeClient, nil
}
