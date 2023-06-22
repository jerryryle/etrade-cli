package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type CommandOrdersList struct {
	Context *CommandContext
}

func (c *CommandOrdersList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [account ID]",
		Short: "List orders",
		Long:  "List orders",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListOrders(args[0])
		},
	}
	return cmd
}

func (c *CommandOrdersList) ListOrders(accountKeyId string) error {
	response, err := c.Context.Client.ListOrders(
		accountKeyId, "", -1, constants.OrderStatusNil, nil, nil, nil, constants.OrderSecurityTypeNil,
		constants.OrderTransactionTypeNil, constants.MarketSessionNil,
	)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(c.Context.OutputFile, string(response))
	return nil
}
