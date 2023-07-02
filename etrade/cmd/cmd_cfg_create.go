package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/spf13/cobra"
)

type cfgCreateFlags struct {
	force bool
}

type CommandCfgCreate struct {
	context CommandContext
	flags   cfgCreateFlags
}

func (c *CommandCfgCreate) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create default configuration",
		Long:  "Create a default configuration file",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContextFromFlags(globalFlags)
			if err != nil {
				return err
			}
			c.context = *context
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return c.context.Close()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.CreateConfig()
		},
	}
	cmd.Flags().BoolVar(&c.flags.force, "force", false, "overwrite any existing configuration file")

	return cmd
}

func (c *CommandCfgCreate) CreateConfig() error {
	defaultConfig := CustomerConfigurationStore{
		customerConfigMap: map[string]CustomerConfiguration{
			"CustomerId1": {
				CustomerName:           "Customer Name 1",
				CustomerProduction:     true,
				CustomerConsumerKey:    "consumer key",
				CustomerConsumerSecret: "consumer secret",
			},
			"CustomerId2 - a short customer ID that you'll specify with --customer-id to use this configuration": {
				CustomerName:           "A human-readable customer name of your choosing. For display purposes.",
				CustomerProduction:     true,
				CustomerConsumerKey:    "The consumer key you got from ETrade. Change the above boolean to reflect whether this is a sandbox (false) or production (true) key",
				CustomerConsumerSecret: "The consumer secret you got from ETrade. Request a sandbox key/secret here: https://us.etrade.com/etx/ris/apikey or request a prod key/secret here: https://us.etrade.com/etx/ris/apisurvey/#/questionnaire",
			},
		},
	}
	err := c.context.ConfigurationFolder.SaveCustomerConfiguration(
		&defaultConfig, c.flags.force, c.context.Logger,
	)
	if err != nil {
		return fmt.Errorf(
			"Unable to create the default configuration file at %s (%w).\nThe file may already exist. To overwrite it, use the --force flag",
			c.context.ConfigurationFolder.GetConfigurationFilePath(), err,
		)
	}

	resultMap := jsonmap.JsonMap{
		"status": "success",
		"message": fmt.Sprintf(
			"Default configuration file successfully created at %s. Please update it with your customer information from ETrade.",
			c.context.ConfigurationFolder.GetConfigurationFilePath(),
		),
	}
	if err := c.context.Renderer.Render(resultMap, cfgCreateDescriptor); err != nil {
		return err
	}
	return nil
}

var cfgCreateDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Status", Path: ".status"},
			{Header: "Message", Path: ".message"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
