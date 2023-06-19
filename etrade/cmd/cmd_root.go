package cmd

import (
	"github.com/spf13/cobra"
)

type RootCommand struct {
	globalFlags GlobalFlags
}

func (c *RootCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "etrade",
		Short: "E*TRADE CLI",
		Long:  "E*TRADE Command Line Interface",
	}
	// Add Global Flags
	cmd.PersistentFlags().StringVar(&c.globalFlags.customerId, "customerId", "", "customer identifier")
	cmd.PersistentFlags().BoolVar(&c.globalFlags.debug, "debug", false, "debug output")

	// Add Subcommands
	cmd.AddCommand((&CommandAccounts{GlobalFlags: &c.globalFlags}).Command())
	cmd.AddCommand((&CommandAlerts{GlobalFlags: &c.globalFlags}).Command())
	cmd.AddCommand((&CommandMarket{GlobalFlags: &c.globalFlags}).Command())
	cmd.AddCommand((&CommandOrders{GlobalFlags: &c.globalFlags}).Command())
	cmd.AddCommand((&CommandCfg{}).Command())

	return cmd
}
