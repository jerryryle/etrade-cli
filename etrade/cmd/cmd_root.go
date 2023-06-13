package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"
)

type RootCommandFlags struct {
	customerId string
	debug      bool
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
	cmd.PersistentFlags().StringVar(&c.flags.customerId, "customerId", "", "customer identifier")
	_ = cmd.MarkPersistentFlagRequired("customerId")
	cmd.PersistentFlags().BoolVar(&c.flags.debug, "debug", false, "debug output")

	return cmd
}

func (c *RootCommand) RootSetupApplicationContext() error {
	// Set the default log level, based on the verbose flag.
	var logLevel = slog.LevelError
	if c.flags.debug {
		logLevel = slog.LevelDebug
	}

	// Create a logger.
	logHandlerOptions := slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
	}
	c.AppContext.Logger = slog.New(slog.NewJSONHandler(os.Stderr, &logHandlerOptions))

	// Load the configuration file and locate the configuration for the requested customer ID
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cfgFileName := ".etradecfg"
	cfgFilePath := filepath.Join(homeDir, cfgFileName)

	customerConfigStore, err := LoadCustomerConfigurationsStoreFromFile(cfgFilePath, c.AppContext.Logger)
	if err != nil {
		return err
	}

	customerConfig, err := customerConfigStore.GetCustomerConfigurationById(c.flags.customerId)
	if err != nil {
		return errors.New(fmt.Sprintf("customer id '%s' not found in config file at %s", c.flags.customerId, cfgFilePath))
	}

	cacheFileName := "." + customerConfig.CustomerConsumerKey
	cacheFilePath := filepath.Join(homeDir, ".etrade", cacheFileName)

	// Get an ETrade client that's authorized for the customer
	c.AppContext.Client, err = getClientWithCredentialCache(
		customerConfig.CustomerProduction,
		customerConfig.CustomerConsumerKey,
		customerConfig.CustomerConsumerSecret,
		cacheFilePath,
		c.AppContext.Logger)
	if err != nil {
		return err
	}

	// Create an ETrade customer object
	c.AppContext.Customer = etradelib.CreateETradeCustomer(c.AppContext.Client, customerConfig.CustomerName)

	return nil
}
