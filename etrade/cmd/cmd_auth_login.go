package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/spf13/cobra"
)

type CommandAuthLogin struct {
	Context *CommandContextWithStore
}

func (c *CommandAuthLogin) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authorize with the current Customer ID",
		Long:  "Authorize with the current Customer ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Login(globalFlags.customerId)
		},
	}
	return cmd
}

func (c *CommandAuthLogin) Login(customerId string) error {
	if customerId == "" {
		return errors.New("customer id must be specified with --customer-id flag")
	}
	customerConfig, err := c.Context.CustomerConfigurationStore.GetCustomerConfigurationById(customerId)
	if err != nil {
		return fmt.Errorf("customer id '%s' not found in config file", customerId)
	}
	cacheFilePath := getFileCachePathForCustomer(c.Context.ConfigurationFolder, customerConfig.CustomerConsumerKey)

	// Create an ETrade client that's authorized for the customer
	if _, err = createClientWithCredentialCache(
		customerConfig.CustomerProduction,
		customerConfig.CustomerConsumerKey,
		customerConfig.CustomerConsumerSecret,
		cacheFilePath,
		c.Context.Logger,
	); err != nil {
		return err
	}

	resultMap := jsonmap.JsonMap{
		"status": "success",
	}

	if err := c.Context.Renderer.Render(resultMap, loginDescriptor); err != nil {
		return err
	}
	return nil
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
