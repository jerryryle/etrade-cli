package cmd

import (
	"github.com/spf13/cobra"
)

type CommandCfg struct {
}

func (c *CommandCfg) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cfg",
		Short: "Configuration actions",
		Long:  "View or create configuration",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandCfgList{}).Command(globalFlags))
	cmd.AddCommand((&CommandCfgCreate{}).Command(globalFlags))
	return cmd
}
