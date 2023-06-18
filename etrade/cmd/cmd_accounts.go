package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccounts struct {
	AppContext *ApplicationContext
}

func (c *CommandAccounts) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "Account actions",
		Long:  "Perform actions on accounts",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAccountsList{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandAccountsBalances{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandAccountsPortfolio{AppContext: c.AppContext}).Command())
	return cmd
}
