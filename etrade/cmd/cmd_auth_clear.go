package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type CommandAuthClear struct {
	Context *CommandContext
}

func (c *CommandAuthClear) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear [Customer ID (or * for all customers)]",
		Short: "Clear authentication credentials",
		Long:  "Clear authentication credentials",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ClearAuth(args[0])
		},
	}
	return cmd
}

func (c *CommandAuthClear) ClearAuth(customerId string) error {
	if customerId == "*" {
		for k, v := range c.Context.CustomerConfigurationStore.GetAllConfigurations() {
			cacheFilePath := getFileCachePathForCustomer(c.Context.ConfigurationFolder, v.CustomerConsumerKey)
			err := os.Remove(cacheFilePath)
			if err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("unable to remove auth cache for %s (%w)", k, err)
			}
		}
	} else {
		customerConfig, err := c.Context.CustomerConfigurationStore.GetCustomerConfigurationById(customerId)
		if err != nil {
			return fmt.Errorf("customer id '%s' not found in config file", customerId)
		}
		cacheFilePath := getFileCachePathForCustomer(c.Context.ConfigurationFolder, customerConfig.CustomerConsumerKey)
		err = os.Remove(cacheFilePath)
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("unable to remove auth cache for %s (%w)", customerId, err)
		}
	}
	fmt.Println("Done!")
	return nil
}
