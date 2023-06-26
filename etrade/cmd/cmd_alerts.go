package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAlerts struct {
	GlobalFlags *GlobalFlags
	context     CommandContext
}

func (c *CommandAlerts) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alerts",
		Short: "Alert actions",
		Long:  "Perform actions on alerts",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContext(
				c.GlobalFlags.customerId, c.GlobalFlags.debug, c.GlobalFlags.outputFileName,
				c.GlobalFlags.outputFormat.Value(),
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
