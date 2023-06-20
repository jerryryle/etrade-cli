package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAlerts struct {
	GlobalFlags *GlobalFlags
	resources   CommandResources
}

func (c *CommandAlerts) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alerts",
		Short: "Alert actions",
		Long:  "Perform actions on alerts",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			resources, err := NewCommandResources(
				c.GlobalFlags.customerId, c.GlobalFlags.debug, c.GlobalFlags.outputFileName,
			)
			if err != nil {
				return err
			}
			c.resources = *resources
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return CleanupCommandResources(&c.resources)
		},
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAlertsList{Resources: &c.resources}).Command())
	return cmd
}
