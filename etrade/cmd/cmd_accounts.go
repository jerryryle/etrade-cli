package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccounts struct {
	GlobalFlags *GlobalFlags
	resources   CommandResources
}

func (c *CommandAccounts) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "Account actions",
		Long:  "Perform actions on accounts",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			resources, err := NewCommandResources(
				c.GlobalFlags.customerId, c.GlobalFlags.debug, c.GlobalFlags.outputFileName,
			)
			if err != nil {
				return err
			}
			c.resources = *resources
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return CleanupCommandResources(&c.resources)
		},
	}
	// Add Subcommands
	cmd.AddCommand((&CommandAccountsList{Resources: &c.resources}).Command())
	cmd.AddCommand((&CommandAccountsBalances{Resources: &c.resources}).Command())
	cmd.AddCommand((&CommandAccountsPortfolio{Resources: &c.resources}).Command())
	return cmd
}
