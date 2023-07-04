package cmd

import (
	"github.com/spf13/cobra"
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
			if response, err := ClearAuth(
				globalFlags.customerId, c.Context.ConfigurationFolder, c.Context.CustomerConfigurationStore,
			); err == nil {
				return c.Context.Renderer.Render(response, clearAuthDescriptor)
			} else {
				return err
			}
		},
	}
	return cmd
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
