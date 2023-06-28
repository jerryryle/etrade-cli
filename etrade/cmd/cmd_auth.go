package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAuth struct {
	context CommandContextWithStore
}

func (c *CommandAuth) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication actions",
		Long:  "Authentication actions",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContextWithStoreFromFlags(globalFlags)
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
	cmd.AddCommand((&CommandAuthClear{Context: &c.context}).Command())
	return cmd
}
