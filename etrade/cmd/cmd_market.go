package cmd

import (
	"github.com/spf13/cobra"
)

type CommandMarket struct {
	AppContext *ApplicationContext
}

func (c *CommandMarket) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market",
		Short: "Market actions",
		Long:  "Perform market actions",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandMarketLookup{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandMarketQuote{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandMarketOptionchains{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandMarketOptionexpire{AppContext: c.AppContext}).Command())
	return cmd
}
