package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"golang.org/x/exp/slog"
	"time"
)

func getClientWithCredentialCache(production bool, consumerKey string, consumerSecret string, cacheFilePath string, logger *slog.Logger) (client.ETradeClient, error) {
	cachedCredentials, err := LoadCachedCredentialsFromFile(cacheFilePath, logger)
	if err != nil {
		// Create a new, empty credential cache. It will yield empty strings for the cached token, which
		// will indicate that there are no cached credentials for this customer
		cachedCredentials = &CachedCredentials{}
	}

	var client client.ETradeClient
	session, err := etradelib.CreateSession(production, consumerKey, consumerSecret, logger)
	if err != nil {
		return nil, err
	}
	var accessToken = cachedCredentials.AccessToken
	var accessSecret = cachedCredentials.AccessSecret
	client, err = session.Renew(accessToken, accessSecret)
	if err != nil {
		authUrl, err := session.Begin()
		if err != nil {
			return nil, err
		}
		fmt.Printf("Visit this URL to get a validation code:\n%s\n\n", authUrl)

		var validationCode string
		fmt.Print("Enter validation code: ")
		_, err = fmt.Scanln(&validationCode)
		if err != nil {
			return nil, err
		}
		if validationCode == "" {
			return nil, errors.New("no validation code provided")
		}

		client, accessToken, accessSecret, err = session.Verify(validationCode)
		if err != nil {
			return nil, err
		}
	}
	err = SaveCachedCredentialsToFile(
		cacheFilePath,
		&CachedCredentials{accessToken, accessSecret, time.Now()},
		logger)
	if err != nil {
		return nil, err
	}
	return client, nil
}
