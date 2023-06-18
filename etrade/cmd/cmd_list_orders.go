package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type ListOrdersCommand struct {
	AppContext *ApplicationContext
}

func (c *ListOrdersCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listorders [account ID]",
		Short: "List orders",
		Long:  "List orders",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListOrders(args[0])
		},
	}
	return cmd
}

func (c *ListOrdersCommand) ListOrders(accountKeyId string) error {
	response, err := c.AppContext.Client.ListOrders(
		accountKeyId, "", -1, constants.OrderStatusNil, nil, nil, nil, constants.OrderSecurityTypeNil,
		constants.OrderTransactionTypeNil, constants.MarketSessionNil,
	)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
