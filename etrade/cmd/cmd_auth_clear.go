package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/spf13/cobra"
	"os"
)

type CommandAuthClear struct {
	Context *CommandContextWithStore
}

func (c *CommandAuthClear) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear authentication credentials for the current Customer ID",
		Long:  "Clear authentication credentials for the current Customer ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ClearAuth(globalFlags.customerId)
		},
	}
	return cmd
}

func (c *CommandAuthClear) ClearAuth(customerId string) error {
	customerConfig, err := c.Context.CustomerConfigurationStore.GetCustomerConfigurationById(customerId)
	if err != nil {
		return fmt.Errorf("customer id '%s' not found in config file", customerId)
	}
	cacheFilePath := getFileCachePathForCustomer(c.Context.ConfigurationFolder, customerConfig.CustomerConsumerKey)
	err = os.Remove(cacheFilePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("unable to remove auth cache for %s (%w)", customerId, err)
	}

	resultMap := jsonmap.JsonMap{
		"status": "success",
	}

	if err := c.Context.Renderer.Render(resultMap, clearAuthDescriptor); err != nil {
		return err
	}
	return nil
}

var clearAuthDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Status", Path: ".status"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
