package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAlerts struct {
	AppContext *ApplicationContext
}

func (c *CommandAlerts) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alerts",
		Short: "Alert actions",
		Long:  "Perform actions on alerts",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAlertsList{AppContext: c.AppContext}).Command())
	return cmd
}
