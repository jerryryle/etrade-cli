package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccounts struct {
	context CommandContext
}

func (c *CommandAccounts) Command(globalFlags *GlobalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "Account actions",
		Long:  "Perform actions on accounts",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContext(
				globalFlags.customerId, globalFlags.debug, globalFlags.outputFileName, globalFlags.outputFormat.Value(),
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
