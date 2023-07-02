package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/spf13/cobra"
	"os"
	"time"
)

type CommandAuthLogin struct {
	Context *CommandContextWithStore
}

func (c *CommandAuthLogin) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authorize with the current Customer ID",
		Long:  "Authorize or renew credentials with the current Customer ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Login(globalFlags.customerId)
		},
	}
	return cmd
}

func (c *CommandAuthLogin) Login(customerId string) error {
	eTradeClient, err := NewETradeClientForCustomer(
		customerId, c.Context.ConfigurationFolder, c.Context.CustomerConfigurationStore, c.Context.Logger,
	)
	if err != nil {
		return err
	}

	authUrl, err := eTradeClient.Authenticate()
	if err != nil {
		return err
	}
	if authUrl != "" {
		// If the Authenticate() method returns an auth url, then show it to
		// the user and get a validation code
		_, _ = fmt.Fprintf(os.Stderr, "Visit this URL to get a validation code:\n%s\n\n", authUrl)
		// Prompt the user to visit the auth URL to get a validation code.
		// Then wait for them to input the code.
		var validationCode string
		_, _ = fmt.Fprintf(os.Stderr, "Enter validation code: ")
		_, err = fmt.Scanln(&validationCode)
		if err != nil {
			return err
		}
		if validationCode == "" {
			return errors.New("no validation code provided")
		}

		// Verify the code.
		err = eTradeClient.Verify(validationCode)
		if err != nil {
			return err
		}
	}
	// Store new or renewed credentials to the cache file.
	consumerKey, _, accessToken, accessSecret := eTradeClient.GetKeys()
	cacheFilePath := getFileCachePathForCustomer(c.Context.ConfigurationFolder, consumerKey)
	err = SaveCachedCredentialsToFile(
		cacheFilePath,
		&CachedCredentials{accessToken, accessSecret, time.Now()},
		c.Context.Logger,
	)
	if err != nil {
		return err
	}

	resultMap := jsonmap.JsonMap{
		"status": "success",
	}
	return c.Context.Renderer.Render(resultMap, loginDescriptor)
}

var loginDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Status", Path: ".status"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
