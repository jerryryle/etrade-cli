package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
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
			var fromDate, toDate *time.Time = nil, nil
			var err error
			if c.flags.fromDate != "" {
				*fromDate, err = time.Parse("01022006", c.flags.fromDate)
				if err != nil {
					return errors.New("from date must be in format MMDDYYYY")
				}
			}
			if c.flags.toDate != "" {
				*toDate, err = time.Parse("01022006", c.flags.toDate)
				if err != nil {
					return errors.New("to date must be in format MMDDYYYY")
				}
			}

			return c.ListOrders(args[0], fromDate, toDate, args[1:])
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
		&c.flags.status, "security-type", "c",
		fmt.Sprintf("security type (%s)", c.flags.securityType.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"security-type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.securityType.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.status, "transaction-type", "y",
		fmt.Sprintf("transaction type (%s)", c.flags.transactionType.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"transaction-type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.transactionType.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.status, "market-session", "m",
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

func (c *CommandOrdersList) ListOrders(
	accountId string, fromDate *time.Time, toDate *time.Time, symbols []string,
) error {
	// This determines how many order items will be retrieved in each request.
	// This should normally be set to the max for efficiency, but can be
	// lowered to test the pagination logic.
	const countPerRequest = constants.OrdersMaxCount

	account, err := GetAccountById(c.Context.Client, accountId)
	if err != nil {
		return err
	}

	response, err := c.Context.Client.ListOrders(
		account.GetIdKey(), "", countPerRequest, constants.OrderStatusNil, fromDate, toDate, symbols,
		c.flags.securityType.Value(), c.flags.transactionType.Value(), c.flags.marketSession.Value(),
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
			account.GetIdKey(), orderList.NextPage(), countPerRequest, constants.OrderStatusNil, fromDate, toDate,
			symbols, c.flags.securityType.Value(), c.flags.transactionType.Value(), c.flags.marketSession.Value(),
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

var orderStatusMap = map[string]enumValueWithHelp[constants.OrderStatus]{
	"open":            {constants.OrderStatusOpen, "only open orders"},
	"executed":        {constants.OrderStatusExecuted, "only executed orders"},
	"canceled":        {constants.OrderStatusCanceled, "only canceled orders"},
	"individualFills": {constants.OrderStatusIndividualFills, "only orders with individual fills"},
	"cancelRequested": {constants.OrderStatusCancelRequested, "only cancel requested orders"},
	"expired":         {constants.OrderStatusExpired, "only expired orders"},
	"rejected":        {constants.OrderStatusRejected, "only rejected orders"},
}

var orderSecurityTypeMap = map[string]enumValueWithHelp[constants.OrderSecurityType]{
	"equity":          {constants.OrderSecurityTypeEquity, "only equity orders"},
	"option":          {constants.OrderSecurityTypeOption, "only option orders"},
	"mutualFund":      {constants.OrderSecurityTypeMutualFund, "only mutual fund orders"},
	"moneyMarketFund": {constants.OrderSecurityTypeMoneyMarketFund, "only money market fund orders"},
}

var orderTransactionTypeMap = map[string]enumValueWithHelp[constants.OrderTransactionType]{
	"extendedHours":      {constants.OrderTransactionTypeExtendedHours, "only extended hours orders"},
	"buy":                {constants.OrderTransactionTypeBuy, "only buy orders"},
	"sell":               {constants.OrderTransactionTypeSell, "only sell orders"},
	"short":              {constants.OrderTransactionTypeShort, "only short orders"},
	"buyToCover":         {constants.OrderTransactionTypeBuyToCover, "only buy to cover orders"},
	"mutualFundExchange": {constants.OrderTransactionTypeMutualFundExchange, "only mutual fund exchange orders"},
}
