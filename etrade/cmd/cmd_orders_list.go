package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
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

func (c *CommandOrdersList) ListOrders(accountId string) error {
	// This determines how many order items will be retrieved in each request.
	// This should normally be set to the max for efficiency, but can be
	// lowered to test the pagination logic.
	const countPerRequest = constants.OrdersMaxCount

	account, err := GetAccountById(c.Context.Client, accountId)
	if err != nil {
		return err
	}

	response, err := c.Context.Client.ListOrders(
		account.GetIdKey(), "", countPerRequest, constants.OrderStatusNil, nil, nil, nil,
		constants.OrderSecurityTypeNil,
		constants.OrderTransactionTypeNil, constants.MarketSessionNil,
	)
	if err != nil {
		return err
	}

	orderList, err := etradelib.CreateETradeOrderListFromResponse(response)
	if err != nil {
		return err
	}

	for orderList.NextPage() != "" {
		response, err = c.Context.Client.ListOrders(
			account.GetIdKey(), orderList.NextPage(), countPerRequest, constants.OrderStatusNil, nil, nil, nil,
			constants.OrderSecurityTypeNil,
			constants.OrderTransactionTypeNil, constants.MarketSessionNil,
		)
		if err != nil {
			return err
		}
		err = orderList.AddPageFromResponse(response)
		if err != nil {
			return err
		}
	}

	err = c.Context.Renderer.Render(orderList.AsJsonMap(), orderListDescriptor)
	if err != nil {
		return err
	}
	return nil
}

var orderListDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".orders",
		Values: []RenderValue{
			{Header: "Order Id", Path: ".orderId"},
			{Header: "Order Type", Path: ".orderType"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
