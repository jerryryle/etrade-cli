package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccountsTransactions struct {
	Resources *CommandResources
}

func (c *CommandAccountsTransactions) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Transaction actions",
		Long:  "Perform actions on account transactions",
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAccountsTransactionsList{Resources: c.Resources}).Command())
	cmd.AddCommand((&CommandAccountsTransactionsDetails{Resources: c.Resources}).Command())

	return cmd
}
