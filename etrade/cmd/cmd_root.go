package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type RootCommandFlags struct {
	customerId string
}

type RootCommand struct {
	AppContext *ApplicationContext
	flags      RootCommandFlags
}

func (c *RootCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "etrade",
		Short: "E*TRADE CLI",
		Long:  "E*TRADE Command Line Interface",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return c.RootSetupApplicationContext()
		},
	}
	// Add Global Flags
	cmd.PersistentFlags().StringVarP(&c.flags.customerId, "customerId", "c", "", "customer identifier")
	_ = cmd.MarkPersistentFlagRequired("customerId")

	return cmd
}

func (c *RootCommand) RootSetupApplicationContext() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cfgFileName := ".etradecfg"
	cfgFilePath := filepath.Join(homeDir, cfgFileName)

	customerConfigStore, err := LoadCustomerConfigurationsStoreFromFile(cfgFilePath)
	if err != nil {
		return err
	}

	customerConfig, err := customerConfigStore.GetCustomerConfigurationById(c.flags.customerId)
	if err != nil {
		return errors.New(fmt.Sprintf("customer id '%s' not found in config file at %s", c.flags.customerId, cfgFilePath))
	}

	cacheFileName := "." + customerConfig.CustomerConsumerKey
	cacheFilePath := filepath.Join(homeDir, ".etrade", cacheFileName)

	c.AppContext.Client, err = getClientWithCredentialCache(
		customerConfig.CustomerProduction,
		customerConfig.CustomerConsumerKey,
		customerConfig.CustomerConsumerSecret,
		cacheFilePath)
	if err != nil {
		return err
	}
	c.AppContext.Customer = etradelib.CreateETradeCustomer(c.AppContext.Client, customerConfig.CustomerName)

	return nil
}
