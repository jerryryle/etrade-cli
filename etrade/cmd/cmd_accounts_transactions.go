package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccountsTransactions struct {
	AppContext *ApplicationContext
}

func (c *CommandAccountsTransactions) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Transaction actions",
		Long:  "Perform actions on account transactions",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAccountsTransactionsList{AppContext: c.AppContext}).Command())
	cmd.AddCommand((&CommandAccountsTransactionsDetails{AppContext: c.AppContext}).Command())

	return cmd
}
