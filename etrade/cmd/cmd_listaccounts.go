package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type ListAccountsCommand struct {
	root *RootCommand
}

func (c *ListAccountsCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listaccounts",
		Short: "List accounts",
		Long:  "List all accounts for the current customer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListAccounts()
		},
	}
	return cmd
}

func (c *ListAccountsCommand) ListAccounts() error {
	fmt.Println("doListAccounts:", c.root.flags.customerId)
	return nil
}
