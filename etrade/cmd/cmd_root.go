package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"
)

type rootCommandFlags struct {
	customerId string
	debug      bool
}

type RootCommand struct {
	AppContext *ApplicationContext
	flags      rootCommandFlags
}

func (c *RootCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "etrade",
		Short: "E*TRADE CLI",
		Long:  "E*TRADE Command Line Interface",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Don't try to set up the application context if we're just
			// displaying help.
			if cmd.CalledAs() == "help" {
				return nil
			}
			return c.RootSetupApplicationContext()
		},
	}
	// Add Global Flags
	cmd.PersistentFlags().StringVar(&c.flags.customerId, "customerId", "", "customer identifier")
	_ = cmd.MarkPersistentFlagRequired("customerId")
	cmd.PersistentFlags().BoolVar(&c.flags.debug, "debug", false, "debug output")

	// Add Subcommands
	cmd.AddCommand((&CommandAccounts{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandAlerts{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandMarket{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandOrders{AppContext: c.AppContext}).Command())

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
		return fmt.Errorf("customer id '%s' not found in config file at %s", c.flags.customerId, cfgFilePath)
	}

	cacheFileName := "." + customerConfig.CustomerConsumerKey
	cacheFilePath := filepath.Join(homeDir, ".etrade", cacheFileName)

	// Get an ETrade client that's authorized for the customer
	c.AppContext.Client, err = getClientWithCredentialCache(
		customerConfig.CustomerProduction,
		customerConfig.CustomerConsumerKey,
		customerConfig.CustomerConsumerSecret,
		cacheFilePath,
		c.AppContext.Logger,
	)
	if err != nil {
		return err
	}

	return nil
}
