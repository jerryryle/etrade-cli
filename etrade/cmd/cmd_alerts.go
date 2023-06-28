package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAlerts struct {
	context CommandContext
}

func (c *CommandAlerts) Command(globalFlags *GlobalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alerts",
		Short: "Alert actions",
		Long:  "Perform actions on alerts",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContext(
				globalFlags.customerId, globalFlags.debug, globalFlags.outputFileName, globalFlags.outputFormat.Value(),
			)
			if err != nil {
				return err
			}
			c.context = *context
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return CleanupCommandContext(&c.context)
		},
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAlertsList{Context: &c.context}).Command())
	cmd.AddCommand((&CommandAlertsDetails{Context: &c.context}).Command())
	cmd.AddCommand((&CommandAlertsDelete{Context: &c.context}).Command())
	return cmd
}
