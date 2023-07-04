package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
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

	response, err := eTradeClient.Authenticate()
	if err != nil {
		return err
	}
	authStatus, err := etradelib.CreateETradeAuthenticationStatusFromResponse(response)
	if err != nil {
		return err
	}

	statusMap := authStatus.AsJsonMap()

	if authStatus.NeedAuthorization() {
		// If the Authenticate() method requires authorization, then prompt the
		// user to visit the authorization URL and get a validation code.
		_, _ = fmt.Fprintf(
			os.Stderr, "Visit this URL to get a validation code:\n%s\n\n", authStatus.GetAuthorizationUrl(),
		)
		// Wait for the user to input the code.
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
		response, err = eTradeClient.Verify(validationCode)
		if err != nil {
			return err
		}
		verifyStatus, err := etradelib.CreateETradeStatusFromResponse(response)
		if err != nil {
			return err
		}
		statusMap = verifyStatus.AsJsonMap()
	}
	// Store new or renewed credentials to the cache file.
	consumerKey, _, accessToken, accessSecret := eTradeClient.GetKeys()
	if err = c.Context.ConfigurationFolder.SaveCachedCredentialsToFile(
		consumerKey, &CachedCredentials{accessToken, accessSecret, time.Now()}, c.Context.Logger,
	); err != nil {
		return err
	}

	return c.Context.Renderer.Render(statusMap, loginDescriptor)
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
