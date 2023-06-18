package cmd

import (
	"github.com/spf13/cobra"
)

type CommandOrders struct {
	AppContext *ApplicationContext
}

func (c *CommandOrders) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orders",
		Short: "Order actions",
		Long:  "Perform order actions",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandOrdersList{AppContext: c.AppContext}).Command())
	return cmd
}
