package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
	"time"
)

type ordersListFlags struct {
	fromDate        string
	toDate          string
	status          enumFlagValue[constants.OrderStatus]
	securityType    enumFlagValue[constants.OrderSecurityType]
	transactionType enumFlagValue[constants.OrderTransactionType]
	marketSession   enumFlagValue[constants.MarketSession]
	symbols         string
}

type CommandOrdersList struct {
	Context *CommandContextWithClient
	flags   ordersListFlags
}

func (c *CommandOrdersList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [account ID] <symbol> ...",
		Short: "List orders",
		Long:  "List orders (with optional list of symbols to filter on)",
		Args:  cobra.MatchAll(cobra.RangeArgs(1, 26)),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountId := args[0]
			symbols := args[1:]
			var fromDate, toDate *time.Time = nil, nil
			if c.flags.fromDate != "" {
				var err error
				*fromDate, err = time.Parse("01022006", c.flags.fromDate)
				if err != nil {
					return errors.New("from date must be in format MMDDYYYY")
				}
			}
			if c.flags.toDate != "" {
				var err error
				*toDate, err = time.Parse("01022006", c.flags.toDate)
				if err != nil {
					return errors.New("to date must be in format MMDDYYYY")
				}
			}
			if response, err := ListOrders(
				c.Context.Client, accountId, c.flags.status.Value(), fromDate, toDate, symbols,
				c.flags.securityType.Value(), c.flags.transactionType.Value(), c.flags.marketSession.Value(),
			); err == nil {
				return c.Context.Renderer.Render(response, orderListDescriptor)
			} else {
				return err
			}
		},
	}

	// Add Flags
	cmd.Flags().StringVarP(&c.flags.fromDate, "from-date", "f", "", "from date (MMDDYYYY)")
	cmd.Flags().StringVarP(&c.flags.toDate, "to-date", "t", "", "to date (MMDDYYYY)")

	// Initialize Enum Flag Values
	c.flags.status = *newEnumFlagValue(orderStatusMap, constants.OrderStatusNil)
	c.flags.securityType = *newEnumFlagValue(orderSecurityTypeMap, constants.OrderSecurityTypeNil)
	c.flags.transactionType = *newEnumFlagValue(orderTransactionTypeMap, constants.OrderTransactionTypeNil)
	c.flags.marketSession = *newEnumFlagValue(marketSessionMap, constants.MarketSessionNil)

	// Add Enum Flags
	cmd.Flags().VarP(
		&c.flags.status, "status", "s",
		fmt.Sprintf("order status (%s)", c.flags.status.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"status",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.status.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.securityType, "security-type", "c",
		fmt.Sprintf("security type (%s)", c.flags.securityType.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"security-type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.securityType.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.transactionType, "transaction-type", "y",
		fmt.Sprintf("transaction type (%s)", c.flags.transactionType.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"transaction-type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.transactionType.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.marketSession, "market-session", "m",
		fmt.Sprintf("market session (%s)", c.flags.marketSession.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"market-session",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.marketSession.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	return cmd
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
