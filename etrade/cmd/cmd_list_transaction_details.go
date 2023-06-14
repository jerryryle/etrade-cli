package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type ListTransactionDetailsCommand struct {
	AppContext *ApplicationContext
}

func (c *ListTransactionDetailsCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listtransactiondetails [account ID] [transaction ID]",
		Short: "List transaction details",
		Long:  "List transaction details",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListTransactionDetails(args[0], args[1])
		},
	}
	return cmd
}

func (c *ListTransactionDetailsCommand) ListTransactionDetails(accountKeyId string, transactionId string) error {
	response, err := c.AppContext.Client.ListTransactionDetails(accountKeyId, transactionId)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", response)
	return nil
}
