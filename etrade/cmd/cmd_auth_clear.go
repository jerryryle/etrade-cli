package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type CommandAuthClear struct {
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

	// Load the configuration file and locate the configuration for the requested customer ID
	userHomeFolder, err := getUserHomeFolder()
	if err != nil {
		return fmt.Errorf("unable to locate the current user's home folder: %w", err)
	}
	cfgFilePath := getCfgFilePath(userHomeFolder)
	customerConfigStore, err := LoadCustomerConfigurationStoreFromFile(cfgFilePath, nil)
	if err != nil {
		return fmt.Errorf(
			"configuration file %s is missing or corrupt (error: %w). you can create a default configuration file with the command 'cfg create'",
			cfgFilePath, err,
		)
	}
	if customerId == "*" {
		for k, v := range customerConfigStore.GetAllConfigurations() {
			cacheFilePath := getFileCachePathForCustomer(userHomeFolder, v.CustomerConsumerKey)
			err = os.Remove(cacheFilePath)
			if err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("unable to remove auth cache for %s (%w)", k, err)
			}
		}
	} else {
		customerConfig, err := customerConfigStore.GetCustomerConfigurationById(customerId)
		if err != nil {
			return fmt.Errorf("customer id '%s' not found in config file at %s", customerId, cfgFilePath)
		}
		cacheFilePath := getFileCachePathForCustomer(userHomeFolder, customerConfig.CustomerConsumerKey)
		err = os.Remove(cacheFilePath)
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("unable to remove auth cache for %s (%w)", customerId, err)
		}
	}
	fmt.Println("Done!")
	return nil
}
