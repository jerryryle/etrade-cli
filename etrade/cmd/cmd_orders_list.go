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
		SubObjects: []RenderDescriptor{
			{
				ObjectPath: ".orderDetail",
				Values: []RenderValue{
					{Header: "Placed Time", Path: ".placedTime", Transformer: dateTimeTransformerMs},
					{Header: "Executed Time", Path: ".executedTime", Transformer: dateTimeTransformerMs},
					{Header: "Order Value", Path: ".orderValue"},
					{Header: "Status", Path: ".status"},
					{Header: "Order Type", Path: ".orderType"},
					{Header: "Order Term", Path: ".orderTerm"},
					{Header: "Price Type", Path: ".priceType"},
					{Header: "Price Value", Path: ".priceValue"},
					{Header: "Limit Price", Path: ".limitPrice"},
					{Header: "Stop Price", Path: ".stopPrice"},
					{Header: "Stop Limit Price", Path: ".stopLimitPrice"},
					{Header: "Offset Type", Path: ".offsetType"},
					{Header: "Offset Value", Path: ".offsetValue"},
					{Header: "Market Session", Path: ".marketSession"},
					{Header: "Routing Destination", Path: ".routingDestination"},
					{Header: "Bracketed Limit Price", Path: ".bracketedLimitPrice"},
					{Header: "Initial Stop Price", Path: ".initialStopPrice"},
					{Header: "Trail Price", Path: ".trailPrice"},
					{Header: "Trigger Price", Path: ".triggerPrice"},
					{Header: "Condition Price", Path: ".conditionPrice"},
					{Header: "Condition Symbol", Path: ".conditionSymbol"},
					{Header: "Condition Type", Path: ".conditionType"},
					{Header: "Condition Follow Price", Path: ".conditionFollowPrice"},
					{Header: "Condition Security Type", Path: ".conditionSecurityType"},
					{Header: "Replaced By Order Id", Path: ".replacedByOrderId"},
					{Header: "Replaces Order Id", Path: ".replacesOrderId"},
					{Header: "All Or None", Path: ".allOrNone"},
					{Header: "Investment Amount", Path: ".investmentAmount"},
					{Header: "Position Quantity", Path: ".positionQuantity"},
					{Header: "Aip Flag", Path: ".aipFlag"},
					{Header: "Execution Guarantee", Path: ".egQual"},
					{Header: "Reinvest Option", Path: ".reInvestOption"},
					{Header: "Estimated Commission", Path: ".estimatedCommission"},
					{Header: "Estimated Fees", Path: ".estimatedFees"},
					{Header: "Estimated Total Amount", Path: ".estimatedTotalAmount"},
					{Header: "Net Price", Path: ".netPrice"},
					{Header: "Net Bid", Path: ".netBid"},
					{Header: "Net Ask", Path: ".netAsk"},
					{Header: "GCD", Path: ".gcd"},
					{Header: "Ratio", Path: ".ratio"},
					{Header: "Mutual Fund Price Type", Path: ".mfpriceType"},
				},
				SubObjects: []RenderDescriptor{
					{
						ObjectPath: ".instrument",
						Values: []RenderValue{
							{Header: "Symbol", Path: ".product.symbol"},
							{Header: "SecurityType", Path: ".product.securityType"},
							{Header: "Symbol Description", Path: ".symbolDescription"},
							{Header: "Order Action", Path: ".orderAction"},
							{Header: "Quantity Type", Path: ".quantityType"},
							{Header: "Quantity", Path: ".quantity"},
							{Header: "Cancel Quantity", Path: ".cancelQuantity"},
							{Header: "Ordered Quantity", Path: ".orderedQuantity"},
							{Header: "Filled Quantity", Path: ".filledQuantity"},
							{Header: "Average Execution Price", Path: ".averageExecutionPrice"},
							{Header: "Estimated Commission", Path: ".estimatedCommission"},
							{Header: "Estimated Fees", Path: ".estimatedFees"},
							{Header: "Bid", Path: ".bid"},
							{Header: "Ask", Path: ".ask"},
							{Header: "Last Price", Path: ".lastprice"},
							{Header: "Currency", Path: ".currency"},
							{Header: "Options Symbology Initiative Key", Path: ".osiKey"},
							{Header: "Mutual Fund Transaction", Path: ".mfTransaction"},
							{Header: "Reserve Order", Path: ".reserveOrder"},
							{Header: "Reserve Quantity", Path: ".reserveQuantity"},
						},
						SubObjects:   nil,
						DefaultValue: "",
						SpaceAfter:   false,
					},
				},
				DefaultValue: "",
				SpaceAfter:   true,
			},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
