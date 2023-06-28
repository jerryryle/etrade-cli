package cmd

import (
	"github.com/spf13/cobra"
)

type CommandMarket struct {
	context CommandContextWithClient
}

func (c *CommandMarket) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market",
		Short: "Market actions",
		Long:  "Perform market actions",
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
	cmd.AddCommand((&CommandMarketLookup{Context: &c.context}).Command())
	cmd.AddCommand((&CommandMarketQuote{Context: &c.context}).Command())
	cmd.AddCommand((&CommandMarketOptionChains{Context: &c.context}).Command())
	cmd.AddCommand((&CommandMarketOptionExpire{Context: &c.context}).Command())
	return cmd
}
