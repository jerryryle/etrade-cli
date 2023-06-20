package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CommandAccountsTransactionsDetails struct {
	Resources *CommandResources
}

func (c *CommandAccountsTransactionsDetails) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "details [account ID] [transaction ID]",
		Short: "List transaction details",
		Long:  "List transaction details",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListTransactionDetails(args[0], args[1])
		},
	}
	return cmd
}

func (c *CommandAccountsTransactionsDetails) ListTransactionDetails(accountKeyId string, transactionId string) error {
	response, err := c.Resources.Client.ListTransactionDetails(accountKeyId, transactionId)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(c.Resources.OFile, string(response))
	return nil
}
