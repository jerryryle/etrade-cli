package cmd

import (
	"github.com/spf13/cobra"
)

type CommandOrders struct {
	context CommandContext
}

func (c *CommandOrders) Command(globalFlags *GlobalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orders",
		Short: "Order actions",
		Long:  "Perform order actions",
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
	cmd.AddCommand((&CommandOrdersList{Context: &c.context}).Command())
	return cmd
}
