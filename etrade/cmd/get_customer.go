package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"time"
)

func getCustomerWithCredentialCache(production bool, consumerKey string, consumerSecret string, cacheFilePath string) (etradelib.ETradeCustomer, error) {
	cachedCredentials, err := LoadCachedCredentialsFromFile(cacheFilePath)
	if err != nil {
		// Create a new, empty credential cache. It will yield empty strings for the cached token, which
		// will indicate that there are no cached credentials for this customer
		cachedCredentials = &CachedCredentials{}
	}

	var customer etradelib.ETradeCustomer
	session, err := etradelib.CreateSession(production, consumerKey, consumerSecret)
	if err != nil {
		return nil, err
	}
	var accessToken = cachedCredentials.AccessToken
	var accessSecret = cachedCredentials.AccessSecret
	customer, err = session.Renew(accessToken, accessSecret)
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

		customer, accessToken, accessSecret, err = session.Verify(validationCode)
		if err != nil {
			return nil, err
		}
	}
	err = SaveCachedCredentialsToFile(
		cacheFilePath,
		&CachedCredentials{accessToken, accessSecret, time.Now()})
	if err != nil {
		return nil, err
	}
	return customer, nil
}
