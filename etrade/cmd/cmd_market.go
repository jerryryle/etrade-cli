package cmd

import (
	"github.com/spf13/cobra"
)

type CommandMarket struct {
	GlobalFlags *GlobalFlags
	resources   CommandResources
}

func (c *CommandMarket) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market",
		Short: "Market actions",
		Long:  "Perform market actions",
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
	cmd.AddCommand((&CommandMarketLookup{Resources: &c.resources}).Command())
	cmd.AddCommand((&CommandMarketQuote{Resources: &c.resources}).Command())
	cmd.AddCommand((&CommandMarketOptionchains{Resources: &c.resources}).Command())
	cmd.AddCommand((&CommandMarketOptionexpire{Resources: &c.resources}).Command())
	return cmd
}
