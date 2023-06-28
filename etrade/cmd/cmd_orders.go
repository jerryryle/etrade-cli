package cmd

import (
	"github.com/spf13/cobra"
)

type CommandOrders struct {
	context CommandContextWithClient
}

func (c *CommandOrders) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orders",
		Short: "Order actions",
		Long:  "Perform order actions",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContextWithClientFromFlags(globalFlags)
			if err != nil {
				return err
			}
			c.context = *context
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return c.context.Close()
		},
	}
	// Add Subcommands
	cmd.AddCommand((&CommandOrdersList{Context: &c.context}).Command())
	return cmd
}
