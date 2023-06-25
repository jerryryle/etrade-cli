package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type marketQuotesFlags struct {
	requireEarningsDate, skipMiniOptionsCheck bool
	detail                                    enumFlagValue[constants.QuoteDetailFlag]
}

type CommandMarketQuote struct {
	Context *CommandContext
	flags   marketQuotesFlags
}

func (c *CommandMarketQuote) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote [symbol] ...",
		Short: "Get quotes",
		Long:  "Get quotes for one or more symbols",
		Args:  cobra.MatchAll(cobra.MaximumNArgs(50)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetQuotes(args)
		},
	}
	// Add Flags
	cmd.Flags().BoolVarP(
		&c.flags.requireEarningsDate, "require-earnings-date", "r", true, "include next earning date in output",
	)
	cmd.Flags().BoolVarP(
		&c.flags.skipMiniOptionsCheck, "skip-mini-check", "s", false,
		"skip the check for whether the symbol has mini options",
	)

	// Initialize Enum Flag Values
	c.flags.detail = *newEnumFlagValue(detailMap, constants.QuoteDetailAll)

	// Add Enum Flags
	cmd.Flags().VarP(
		&c.flags.detail, "detail", "d",
		fmt.Sprintf("quote details (%s)", c.flags.detail.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"detail",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.detail.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	return cmd
}

func (c *CommandMarketQuote) GetQuotes(symbols []string) error {
	response, err := c.Context.Client.GetQuotes(
		symbols, c.flags.detail.Value(), c.flags.requireEarningsDate, c.flags.skipMiniOptionsCheck,
	)
	if err != nil {
		return err
	}
	quoteList, err := etradelib.CreateETradeQuoteListFromResponse(response)
	if err != nil {
		return err
	}

	renderDescriptor := quoteListAllDescriptor
	switch c.flags.detail.Value() {
	case constants.QuoteDetailFundamental:
		renderDescriptor = quoteListFundamentalDescriptor
	case constants.QuoteDetailIntraday:
		renderDescriptor = quoteListIntradayDescriptor
	case constants.QuoteDetailOptions:
		renderDescriptor = quoteListOptionDescriptor
	case constants.QuoteDetailWeek52:
		renderDescriptor = quoteListWeek52Descriptor
	case constants.QuoteDetailMutualFund:
		renderDescriptor = quoteListMutualFundDescriptor
	}

	err = c.Context.Renderer.Render(quoteList.AsJsonMap(), renderDescriptor)
	if err != nil {
		return err
	}
	return nil
}

var detailMap = map[string]enumValueWithHelp[constants.QuoteDetailFlag]{
	"all":         {constants.QuoteDetailAll, "get all details"},
	"fundamental": {constants.QuoteDetailFundamental, "get fundamental details"},
	"intraday":    {constants.QuoteDetailIntraday, "get intraday details"},
	"options":     {constants.QuoteDetailOptions, "get options details"},
	"week52":      {constants.QuoteDetailWeek52, "get 52-week details"},
	"mutualFund":  {constants.QuoteDetailMutualFund, "get mutual fund details"},
}

var quoteListAllDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".quotes",
		Values: []RenderValue{
			{Header: "Date", Path: ".dateTimeUTC", Transformer: dateTimeTransformer},
			{Header: "Quote Status", Path: ".quoteStatus"},
			{Header: "After Hours", Path: ".ahFlag"},
			{Header: "Has Mini Options", Path: ".hasMiniOptions"},
			{Header: "Symbol", Path: ".product.symbol"},
			{Header: "Security Type", Path: ".product.securityType"},
			{Header: "Adjusted", Path: ".all.adjustedFlag"},
			{Header: "Ask", Path: ".all.ask"},
			{Header: "Ask Size", Path: ".all.askSize"},
			{Header: "Ask Time", Path: ".all.askTime", Transformer: dateTimeTransformer},
			{Header: "Bid", Path: ".all.bid"},
			{Header: "Bid Exchange", Path: ".all.bidExchange"},
			{Header: "Bid Size", Path: ".all.bidSize"},
			{Header: "Bid Time", Path: ".all.bidTime", Transformer: dateTimeTransformer},
			{Header: "Change Close $", Path: ".all.changeClose"},
			{Header: "Change Close %", Path: ".all.changeClosePercentage"},
			{Header: "Company Name", Path: ".all.companyName"},
			{Header: "Days To Expiration", Path: ".all.daysToExpiration"},
			{Header: "Dir Last", Path: ".all.dirLast"},
			{Header: "Dividend", Path: ".all.dividend"},
			{Header: "Earnings Per Share", Path: ".all.eps"},
			{Header: "Estimated Earnings", Path: ".all.estEarnings"},
			{Header: "Dividend Date", Path: ".all.exDividendDate", Transformer: dateTimeTransformer},
			{Header: "High", Path: ".all.high"},
			{Header: "52-week High", Path: ".all.high52"},
			{Header: "Last Trade", Path: ".all.lastTrade"},
			{Header: "Low", Path: ".all.low"},
			{Header: "52-week Low", Path: ".all.low52"},
			{Header: "Open", Path: ".all.open"},
			{Header: "Open Interest", Path: ".all.openInterest"},
			{Header: "Option Style", Path: ".all.optionStyle"},
			{Header: "Option Underlier", Path: ".all.optionUnderlier"},
			{Header: "Option Underlier Exchange", Path: ".all.optionUnderlierExchange"},
			{Header: "Previous Close", Path: ".all.previousClose"},
			{Header: "Previous Day Volume", Path: ".all.previousDayVolume"},
			{Header: "Primary Exchange", Path: ".all.primaryExchange"},
			{Header: "Symbol Description", Path: ".all.symbolDescription"},
			{Header: "Total Volume", Path: ".all.totalVolume"},
			{Header: "Uniform Practice Code", Path: ".all.upc"},
			{Header: "CashDeliverable", Path: ".all.cashDeliverable"},
			{Header: "Market Cap", Path: ".all.marketCap"},
			{Header: "Shares Outstanding", Path: ".all.sharesOutstanding"},
			{Header: "Next Earning Date", Path: ".all.nextEarningDate"},
			{Header: "Beta", Path: ".all.beta"},
			{Header: "Yield", Path: ".all.yield"},
			{Header: "Declared Dividend", Path: ".all.declaredDividend"},
			{Header: "Dividend Payable Date", Path: ".all.dividendPayableDate", Transformer: dateTimeTransformer},
			{Header: "PE", Path: ".all.pe"},
			{Header: "52-Week Low Date", Path: ".all.week52LowDate", Transformer: dateTimeTransformer},
			{Header: "52-Week High Date", Path: ".all.week52HiDate", Transformer: dateTimeTransformer},
			{Header: "Intrinsic Value", Path: ".all.intrinsicValue"},
			{Header: "Time Premium", Path: ".all.timePremium"},
			{Header: "Option Multiplier", Path: ".all.optionMultiplier"},
			{Header: "Contract Size", Path: ".all.contractSize"},
			{Header: "Expiration Date", Path: ".all.expirationDate"},
			{Header: "Option Previous Bid Price", Path: ".all.optionPreviousBidPrice"},
			{Header: "Option Previous Ask Price", Path: ".all.optionPreviousAskPrice"},
			{Header: "Options Symbology Initiative Key", Path: ".all.osiKey"},
			{Header: "Time Of Last Trade", Path: ".all.timeOfLastTrade", Transformer: dateTimeTransformer},
			{Header: "Average Volume", Path: ".all.averageVolume"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}

var quoteListFundamentalDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".quotes",
		Values: []RenderValue{
			{Header: "Date", Path: ".dateTimeUTC", Transformer: dateTimeTransformer},
			{Header: "Quote Status", Path: ".quoteStatus"},
			{Header: "After Hours", Path: ".ahFlag"},
			{Header: "Has Mini Options", Path: ".hasMiniOptions"},
			{Header: "Symbol", Path: ".product.symbol"},
			{Header: "Security Type", Path: ".product.securityType"},
			{Header: "Company Name", Path: ".fundamental.companyName"},
			{Header: "Earnings Per Share", Path: ".fundamental.eps"},
			{Header: "Estimated Earnings", Path: ".fundamental.estEarnings"},
			{Header: "52-week High", Path: ".fundamental.high52"},
			{Header: "Last Trade", Path: ".fundamental.lastTrade"},
			{Header: "52-week Low", Path: ".fundamental.low52"},
			{Header: "Symbol Description", Path: ".fundamental.symbolDescription"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}

var quoteListIntradayDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".quotes",
		Values: []RenderValue{
			{Header: "Date", Path: ".dateTimeUTC", Transformer: dateTimeTransformer},
			{Header: "Quote Status", Path: ".quoteStatus"},
			{Header: "After Hours", Path: ".ahFlag"},
			{Header: "Has Mini Options", Path: ".hasMiniOptions"},
			{Header: "Symbol", Path: ".product.symbol"},
			{Header: "Security Type", Path: ".product.securityType"},
			{Header: "Ask", Path: ".intraday.ask"},
			{Header: "Bid", Path: ".intraday.bid"},
			{Header: "Change Close $", Path: ".intraday.changeClose"},
			{Header: "Change Close %", Path: ".intraday.changeClosePercentage"},
			{Header: "Company Name", Path: ".intraday.companyName"},
			{Header: "High", Path: ".intraday.high"},
			{Header: "Last Trade", Path: ".intraday.lastTrade"},
			{Header: "Low", Path: ".intraday.low"},
			{Header: "Total Volume", Path: ".intraday.totalVolume"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}

var quoteListOptionDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".quotes",
		Values: []RenderValue{
			{Header: "Date", Path: ".dateTimeUTC", Transformer: dateTimeTransformer},
			{Header: "Quote Status", Path: ".quoteStatus"},
			{Header: "After Hours", Path: ".ahFlag"},
			{Header: "Has Mini Options", Path: ".hasMiniOptions"},
			{Header: "Symbol", Path: ".product.symbol"},
			{Header: "Security Type", Path: ".product.securityType"},
			{Header: "Ask", Path: ".option.ask"},
			{Header: "Ask Size", Path: ".option.askSize"},
			{Header: "Bid", Path: ".option.bid"},
			{Header: "Bid Size", Path: ".option.bidSize"},
			{Header: "Company Name", Path: ".option.companyName"},
			{Header: "Days To Expiration", Path: ".option.daysToExpiration"},
			{Header: "Last Trade", Path: ".option.lastTrade"},
			{Header: "Open Interest", Path: ".option.openInterest"},
			{Header: "Option Previous Bid Price", Path: ".option.optionPreviousBidPrice"},
			{Header: "Option Previous Ask Price", Path: ".option.optionPreviousAskPrice"},
			{Header: "Options Symbology Initiative Key", Path: ".option.osiKey"},
			{Header: "Intrinsic Value", Path: ".option.intrinsicValue"},
			{Header: "Time Premium", Path: ".option.timePremium"},
			{Header: "Option Multiplier", Path: ".option.optionMultiplier"},
			{Header: "Contract Size", Path: ".option.contractSize"},
			{Header: "Symbol Description", Path: ".option.symbolDescription"},
			{Header: "Rho", Path: ".option.optionGreeks.rho"},
			{Header: "Vega", Path: ".option.optionGreeks.vega"},
			{Header: "Theta", Path: ".option.optionGreeks.theta"},
			{Header: "Delta", Path: ".option.optionGreeks.delta"},
			{Header: "Gamma", Path: ".option.optionGreeks.gamma"},
			{Header: "Implied Volatility", Path: ".option.optionGreeks.iv"},
			{Header: "Current Value", Path: ".option.optionGreeks.currentValue"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}

var quoteListWeek52Descriptor = []RenderDescriptor{
	{
		ObjectPath: ".quotes",
		Values: []RenderValue{
			{Header: "Date", Path: ".dateTimeUTC", Transformer: dateTimeTransformer},
			{Header: "Quote Status", Path: ".quoteStatus"},
			{Header: "After Hours", Path: ".ahFlag"},
			{Header: "Has Mini Options", Path: ".hasMiniOptions"},
			{Header: "Symbol", Path: ".product.symbol"},
			{Header: "Security Type", Path: ".product.securityType"},
			{Header: "Company Name", Path: ".week52.companyName"},
			{Header: "52-Week High", Path: ".week52.high52"},
			{Header: "Last Trade", Path: ".week52.lastTrade"},
			{Header: "52-Week High", Path: ".week52.low52"},
			{Header: "12-month Performance", Path: ".week52.perf12Months"},
			{Header: "Previous Close", Path: ".week52.previousClose"},
			{Header: "Symbol Description", Path: ".week52.symbolDescription"},
			{Header: "Total Volume", Path: ".week52.totalVolume"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}

var quoteListMutualFundDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".quotes",
		Values: []RenderValue{
			{Header: "Date", Path: ".dateTimeUTC", Transformer: dateTimeTransformer},
			{Header: "Quote Status", Path: ".quoteStatus"},
			{Header: "After Hours", Path: ".ahFlag"},
			{Header: "Has Mini Options", Path: ".hasMiniOptions"},
			{Header: "Symbol", Path: ".product.symbol"},
			{Header: "Security Type", Path: ".product.securityType"},
			{Header: "Symbol Description", Path: ".mutualFund.symbolDescription"},
			{Header: "Identifier", Path: ".mutualFund.cusip"},
			{Header: "Change Close", Path: ".mutualFund.changeClose"},
			{Header: "Previous Close", Path: ".mutualFund.previousClose"},
			{Header: "Transaction Fee", Path: ".mutualFund.transactionFee"},
			{Header: "Early Redemption Fee", Path: ".mutualFund.earlyRedemptionFee"},
			{Header: "Availability", Path: ".mutualFund.availability"},
			{Header: "Initial Investment", Path: ".mutualFund.initialInvestment"},
			{Header: "Subsequent Investment", Path: ".mutualFund.subsequentInvestment"},
			{Header: "Fund Family", Path: ".mutualFund.fundFamily"},
			{Header: "Fund Name", Path: ".mutualFund.fundName"},
			{Header: "Change Close Percentage", Path: ".mutualFund.changeClosePercentage"},
			{Header: "Time Of Last Trade", Path: ".mutualFund.timeOfLastTrade"},
			{Header: "Net Asset Value", Path: ".mutualFund.netAssetValue"},
			{Header: "Public Offer Price", Path: ".mutualFund.publicOfferPrice"},
			{Header: "Net Expense Ratio", Path: ".mutualFund.netExpenseRatio"},
			{Header: "Gross Expense Ratio", Path: ".mutualFund.grossExpenseRatio"},
			{Header: "Order Cutoff Time", Path: ".mutualFund.orderCutoffTime", Transformer: dateTimeTransformer},
			{Header: "Sales Charge", Path: ".mutualFund.salesCharge"},
			{Header: "Initial IRA Investment", Path: ".mutualFund.initialIraInvestment"},
			{Header: "Subsequent IRA Investment", Path: ".mutualFund.subsequentIraInvestment"},
			{Header: "Fund Inception Date", Path: ".mutualFund.fundInceptionDate", Transformer: dateTimeTransformer},
			{Header: "Average Annual Returns", Path: ".mutualFund.averageAnnualReturns"},
			{Header: "Seven Day Current Yield", Path: ".mutualFund.sevenDayCurrentYield"},
			{Header: "Annual Total Return", Path: ".mutualFund.annualTotalReturn"},
			{Header: "Weighted Average Maturity", Path: ".mutualFund.weightedAverageMaturity"},
			{Header: "Average Annual Return 1Yr", Path: ".mutualFund.averageAnnualReturn1Yr"},
			{Header: "Average Annual Return 3Yr", Path: ".mutualFund.averageAnnualReturn3Yr"},
			{Header: "Average Annual Return 5Yr", Path: ".mutualFund.averageAnnualReturn5Yr"},
			{Header: "Average Annual Return 10Yr", Path: ".mutualFund.averageAnnualReturn10Yr"},
			{Header: "52-week High", Path: ".mutualFund.high52"},
			{Header: "52-week Low", Path: ".mutualFund.low52"},
			{Header: "52-week High Date", Path: ".mutualFund.week52HiDate"},
			{Header: "52-week Low Date", Path: ".mutualFund.week52LowDate"},
			{Header: "Exchange Name", Path: ".mutualFund.exchangeName"},
			{Header: "Since Inception", Path: ".mutualFund.sinceInception"},
			{Header: "Quarterly Since Inception", Path: ".mutualFund.quarterlySinceInception"},
			{Header: "Last Trade", Path: ".mutualFund.lastTrade"},
			{Header: "Annual Marketing/Distribution Fee", Path: ".mutualFund.actual12B1Fee"},
			{Header: "Performance As Of Date", Path: ".mutualFund.performanceAsOfDate"},
			{Header: "Quarterly Performance As Of Date", Path: ".mutualFund.qtrlyPerformanceAsOfDate"},
			{Header: "Morningstar Category", Path: ".mutualFund.morningStarCategory"},
			{Header: "Monthly Trailing Return 1Y", Path: ".mutualFund.monthlyTrailingReturn1Y"},
			{Header: "Monthly Trailing Return 3Y", Path: ".mutualFund.monthlyTrailingReturn3Y"},
			{Header: "Monthly Trailing Return 5Y", Path: ".mutualFund.monthlyTrailingReturn5Y"},
			{Header: "Monthly Trailing Return 10Y", Path: ".mutualFund.monthlyTrailingReturn10Y"},
			{Header: "ETrade Early Redemption Fee", Path: ".mutualFund.etradeEarlyRedemptionFee"},
			{Header: "Max Sales Load", Path: ".mutualFund.maxSalesLoad"},
			{Header: "Monthly Trailing Return YTD", Path: ".mutualFund.monthlyTrailingReturnYTD"},
			{Header: "Monthly Trailing Return 1M", Path: ".mutualFund.monthlyTrailingReturn1M"},
			{Header: "Monthly Trailing Return 3M", Path: ".mutualFund.monthlyTrailingReturn3M"},
			{Header: "Monthly Trailing Return 6M", Path: ".mutualFund.monthlyTrailingReturn6M"},
			{Header: "Quarterly Trailing Return YTD", Path: ".mutualFund.qtrlyTrailingReturnYTD"},
			{Header: "Quarterly Trailing Return 1M", Path: ".mutualFund.qtrlyTrailingReturn1M"},
			{Header: "Quarterly Trailing Return 3M", Path: ".mutualFund.qtrlyTrailingReturn3M"},
			{Header: "Quarterly Trailing Return 6M", Path: ".mutualFund.qtrlyTrailingReturn6M"},
			{Header: "Exchange Code", Path: ".mutualFund.exchangeCode"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
