package cmd

import (
	"github.com/spf13/cobra"
)

type CommandOrders struct {
	GlobalFlags *GlobalFlags
	resources   CommandResources
}

func (c *CommandOrders) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orders",
		Short: "Order actions",
		Long:  "Perform order actions",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			resources, err := NewCommandResources(c.GlobalFlags.customerId, c.GlobalFlags.debug)
			if err != nil {
				return err
			}
			c.resources = *resources
			return nil
		},
	}
	// Add Subcommands
	cmd.AddCommand((&CommandOrdersList{Resources: &c.resources}).Command())
	return cmd
}
