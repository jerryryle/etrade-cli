package cmd

import (
	"github.com/spf13/cobra"
)

type CommandCfgList struct {
	context CommandContextWithStore
}

func (c *CommandCfgList) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List configuration",
		Long:  "List all available customers in the app configuration",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContextWithStoreFromFlags(globalFlags)
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
			response := GetCustomerList(c.context.CustomerConfigurationStore)
			return c.context.Renderer.Render(response, cfgListDescriptor)
		},
	}
	return cmd
}

var cfgListDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".customers",
		Values: []RenderValue{
			{Header: "Customer Id", Path: ".customerId"},
			{Header: "Customer Name", Path: ".customerName"},
			{Header: "Production Access", Path: ".productionAccess"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
