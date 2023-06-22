package cmd

import (
	"github.com/spf13/cobra"
)

type CommandMarket struct {
	GlobalFlags *GlobalFlags
	context     CommandContext
}

func (c *CommandMarket) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market",
		Short: "Market actions",
		Long:  "Perform market actions",
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
	cmd.AddCommand((&CommandMarketLookup{Context: &c.context}).Command())
	cmd.AddCommand((&CommandMarketQuote{Context: &c.context}).Command())
	cmd.AddCommand((&CommandMarketOptionchains{Context: &c.context}).Command())
	cmd.AddCommand((&CommandMarketOptionexpire{Context: &c.context}).Command())
	return cmd
}
