package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type cfgCreateFlags struct {
	force bool
}

type CommandCfgCreate struct {
	flags cfgCreateFlags
}

func (c *CommandCfgCreate) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create default configuration",
		Long:  "Create a default configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.CreateConfig()
		},
	}
	cmd.Flags().BoolVar(&c.flags.force, "force", false, "overwrite any existing configuration file")

	return cmd
}

func (c *CommandCfgCreate) CreateConfig() error {
	userHomeFolder, err := getUserHomeFolder()
	if err != nil {
		return fmt.Errorf("unable to locate the current user's home folder: %w", err)
	}
	cfgFilePath := getCfgFilePath(userHomeFolder)

	defaultConfig := CustomerConfigurationsStore{
		customerConfigMap: map[string]CustomerConfiguration{
			"CustomerId1": {
				CustomerName:           "Customer Name 1",
				CustomerProduction:     true,
				CustomerConsumerKey:    "consumer key",
				CustomerConsumerSecret: "consumer secret",
			},
			"CustomerId2": {
				CustomerName:           "A human-readable customer name of your choosing. For display purposes.",
				CustomerProduction:     true,
				CustomerConsumerKey:    "The consumer key you got from ETrade. Change the above boolean to reflect whether this is a sandbox (false) or production (true) key",
				CustomerConsumerSecret: "The consumer secret you got from ETrade. Request a sandbox key/secret here: https://us.etrade.com/etx/ris/apikey or request a prod key/secret here: https://us.etrade.com/etx/ris/apisurvey/#/questionnaire",
			},
		},
	}
	err = SaveCustomerConfigurationsStoreToFile(cfgFilePath, c.flags.force, &defaultConfig, nil)
	if err != nil {
		return fmt.Errorf(
			"Unable to create the default configuration file at %s (%w).\nThe file may already exist. To overwrite it, use the --force flag",
			cfgFilePath, err,
		)
	}
	fmt.Printf(
		"Default configuration file successfully created at %s\nPlease update it with your customer information from ETrade.\n",
		cfgFilePath,
	)
	return nil
}
