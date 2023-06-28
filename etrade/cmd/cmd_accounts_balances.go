package cmd

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/spf13/cobra"
)

type commandAccountsBalancesFlags struct {
	realTimeBalance bool
}

type CommandAccountsBalances struct {
	Context *CommandContextWithClient
	flags   commandAccountsBalancesFlags
}

func (c *CommandAccountsBalances) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances [account ID]",
		Short: "Get account balances",
		Long:  "Get account balances",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetAccountBalances(args[0])
		},
	}
	cmd.Flags().BoolVarP(&c.flags.realTimeBalance, "realtime-balance", "r", true, "return real time balance")
	return cmd
}

func (c *CommandAccountsBalances) GetAccountBalances(accountId string) error {
	account, err := GetAccountById(c.Context.Client, accountId)
	if err != nil {
		return err
	}

	response, err := c.Context.Client.GetAccountBalances(account.GetIdKey(), c.flags.realTimeBalance)
	if err != nil {
		return err
	}
	balances, err := etradelib.CreateETradeBalancesFromResponse(response)
	if err != nil {
		return err
	}
	err = c.Context.Renderer.Render(balances.AsJsonMap(), balancesDescriptor)
	if err != nil {
		return err
	}
	return nil
}

var balancesDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Account ID", Path: ".accountId"},
			{Header: "Account Type", Path: ".accountType"},
			{Header: "Option Level", Path: ".optionLevel"},
			{Header: "Account Description", Path: ".accountDescription"},
			{Header: "Quote Mode", Path: ".quoteMode", Transformer: quoteModeTransformer},
			{Header: "Day Trader Status", Path: ".dayTraderStatus"},
			{Header: "Account Mode", Path: ".accountMode"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".cash",
		Values: []RenderValue{
			{Header: "Cash Balance", Path: ".moneyMktBalance"},
			{Header: "Cash Reserved For Open Orders", Path: ".fundsForOpenOrdersCash"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".computed",
		Values: []RenderValue{
			{Header: "Cash Available For Investments", Path: ".cashAvailableForInvestment"},
			{Header: "Cash Available For Withdrawal", Path: ".cashAvailableForWithdrawal"},
			{Header: "Total Available For Withdrawal", Path: ".totalAvailableForWithdrawal"},
			{Header: "Net Cash Balance", Path: ".netCash"},
			{Header: "Current Cash Balance", Path: ".cashBalance"},
			{Header: "Settled Cash For Investments", Path: ".settledCashForInvestment"},
			{Header: "Unsettled Cash For Investments", Path: ".unSettledCashForInvestment"},
			{Header: "Funds Withheld From Purchasing Power", Path: ".fundsWithheldFromPurchasePower"},
			{Header: "Funds Withheld From Withdrawal", Path: ".fundsWithheldFromWithdrawal"},
			{Header: "Margin Buying Power", Path: ".marginBuyingPower"},
			{Header: "Cash Buying Power", Path: ".cashBuyingPower"},
			{Header: "Day Trader Margin Buying Power", Path: ".dtMarginBuyingPower"},
			{Header: "Day Trader Cash Buying Power", Path: ".dtCashBuyingPower"},
			{Header: "Margin Balance", Path: ".marginBalance"},
			{Header: "Short Adjusted Balance", Path: ".shortAdjustBalance"},
			{Header: "Regulation T Equity $", Path: ".regtEquity"},
			{Header: "Regulation T Equity %", Path: ".regtEquityPercent"},
			{Header: "Current Account Balance", Path: ".accountBalance"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".computed.openCalls",
		Values: []RenderValue{
			{Header: "Minimum Equity Call", Path: ".minEquityCall"},
			{Header: "Federal Call", Path: ".fedCall"},
			{Header: "Cash Call", Path: ".cashCall"},
			{Header: "House Call", Path: ".houseCall"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".computed.portfolioMargin",
		Values: []RenderValue{
			{Header: "Cash Open Order Reserve", Path: ".dtCashOpenOrderReserve"},
			{Header: "Margin Open Order Reserve", Path: ".dtMarginOpenOrderReserve"},
			{Header: "Liquidating Equity", Path: ".liquidatingEquity"},
			{Header: "House Excess Equity", Path: ".houseExcessEquity"},
			{Header: "Total House Requirement", Path: ".totalHouseRequirement"},
			{Header: "Excess Equity Minus Portfolio Requirement", Path: ".excessEquityMinusRequirement"},
			{Header: "Total Margin Requirements", Path: ".totalMarginRqmts"},
			{Header: "Available Excess Equity", Path: ".availExcessEquity"},
			{Header: "Excess Equity", Path: ".excessEquity"},
			{Header: "Open Order Reserve", Path: ".openOrderReserve"},
			{Header: "Funds On Hold", Path: ".fundsOnHold"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".computed.realTimeValues",
		Values: []RenderValue{
			{Header: "Realtime Total Account Value", Path: ".totalAccountValue"},
			{Header: "Realtime Net Market Value", Path: ".netMv"},
			{Header: "Realtime Long Net Market Value", Path: ".netMvLong"},
			{Header: "Realtime Short Net Market Value", Path: ".netMvShort"},
			{Header: "Realtime Total Long Value", Path: ".totalLongValue"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}

func quoteModeTransformer(value interface{}) interface{} {
	var quoteModes = []string{
		"REALTIME", "DELAYED", "CLOSING", "AHT REALTIME", "AHT BEFORE OPEN", "AHT CLOSING", "NONE",
	}

	switch t := value.(type) {
	case json.Number:
		qm, err := t.Int64()
		if err != nil {
			return value
		}
		if qm > int64(len(quoteModes)) {
			return value
		}
		return quoteModes[qm]
	default:
		return value
	}
}
