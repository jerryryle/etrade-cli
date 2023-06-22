package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CommandAccountsList struct {
	Context *CommandContext
}

func (c *CommandAccountsList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List accounts",
		Long:  "List all accounts for the current customer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListAccounts()
		},
	}
	return cmd
}

func (c *CommandAccountsList) ListAccounts() error {
	response, err := c.Context.Client.ListAccounts()
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(c.Context.OutputFile, string(response))
	return nil
}
