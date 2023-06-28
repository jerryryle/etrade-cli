package cmd

import (
	"github.com/spf13/cobra"
)

type CommandCfg struct {
	context CommandContext
}

func (c *CommandCfg) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cfg",
		Short: "Configuration actions",
		Long:  "View or create configuration",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContextFromFlags(globalFlags)
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
	cmd.AddCommand((&CommandCfgList{Context: &c.context}).Command())
	cmd.AddCommand((&CommandCfgCreate{Context: &c.context}).Command())
	return cmd
}
