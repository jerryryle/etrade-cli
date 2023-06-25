package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccounts struct {
	GlobalFlags *GlobalFlags
	context     CommandContext
}

func (c *CommandAccounts) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "Account actions",
		Long:  "Perform actions on accounts",
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
	cmd.AddCommand((&CommandAccountsList{Context: &c.context}).Command())
	cmd.AddCommand((&CommandAccountsBalances{Context: &c.context}).Command())
	cmd.AddCommand((&CommandAccountsPortfolio{Context: &c.context}).Command())
	cmd.AddCommand((&CommandAccountsTransactions{Context: &c.context}).Command())
	return cmd
}
