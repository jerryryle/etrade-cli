package cmd

import (
	"github.com/spf13/cobra"
)

type CommandOrders struct {
	GlobalFlags *GlobalFlags
	context     CommandContext
}

func (c *CommandOrders) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orders",
		Short: "Order actions",
		Long:  "Perform order actions",
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
	cmd.AddCommand((&CommandOrdersList{Context: &c.context}).Command())
	return cmd
}
