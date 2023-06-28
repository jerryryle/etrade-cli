package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccountsTransactions struct {
	Context *CommandContextWithClient
}

func (c *CommandAccountsTransactions) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Transaction actions",
		Long:  "Perform actions on account transactions",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAccountsTransactionsList{Context: c.Context}).Command())
	cmd.AddCommand((&CommandAccountsTransactionsDetails{Context: c.Context}).Command())

	return cmd
}
