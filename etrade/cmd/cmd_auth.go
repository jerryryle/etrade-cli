package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAuth struct {
}

func (c *CommandAuth) Command(globalFlags *GlobalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication actions",
		Long:  "Authentication actions",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAuthClear{}).Command())
	return cmd
}
