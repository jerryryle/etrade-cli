package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type accountsPortfolioFlags struct {
	totalsRequired bool
	withLots       bool
	portfolioView  enumFlagValue[constants.PortfolioView]
	sortBy         enumFlagValue[constants.PortfolioSortBy]
	sortOrder      enumFlagValue[constants.SortOrder]
	marketSession  enumFlagValue[constants.MarketSession]
}

type CommandAccountsPortfolio struct {
	Context *CommandContextWithClient
	flags   accountsPortfolioFlags
}

func (c *CommandAccountsPortfolio) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "portfolio [account ID]",
		Short: "View Portfolio",
		Long:  "View Portfolio",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountId := args[0]
			if response, err := ViewPortfolio(
				c.Context.Client, accountId, c.flags.sortBy.Value(), c.flags.sortOrder.Value(),
				c.flags.marketSession.Value(),
				c.flags.totalsRequired, c.flags.portfolioView.Value(), c.flags.withLots,
			); err == nil {
				renderDescriptor := GetQuickViewRenderDescriptor(c.flags.withLots)
				switch c.flags.portfolioView.Value() {
				case constants.PortfolioViewPerformance:
					renderDescriptor = GetPerformanceViewRenderDescriptor(c.flags.withLots)
				case constants.PortfolioViewFundamental:
					renderDescriptor = GetFundamentalViewRenderDescriptor(c.flags.withLots)
				case constants.PortfolioViewOptionsWatch:
					renderDescriptor = GetOptionsWatchViewRenderDescriptor(c.flags.withLots)
				case constants.PortfolioViewComplete:
					renderDescriptor = GetCompleteViewRenderDescriptor(c.flags.withLots)
				}
				return c.Context.Renderer.Render(response, renderDescriptor)
			} else {
				return err
			}
		},
	}

	// Add Flags
	cmd.Flags().BoolVarP(&c.flags.totalsRequired, "totals-required", "t", true, "include totals in results")
	cmd.Flags().BoolVarP(&c.flags.withLots, "with-lots", "l", false, "include lots in results")

	// Initialize Enum Flag Values
	c.flags.portfolioView = *newEnumFlagValue(portfolioViewMap, constants.PortfolioViewQuick)
	c.flags.sortBy = *newEnumFlagValue(portfolioSortByMap, constants.PortfolioSortByNil)
	c.flags.sortOrder = *newEnumFlagValue(sortOrderMap, constants.SortOrderNil)
	c.flags.marketSession = *newEnumFlagValue(marketSessionMap, constants.MarketSessionNil)

	// Add Enum Flags
	cmd.Flags().VarP(
		&c.flags.portfolioView, "view", "v",
		fmt.Sprintf("portfolio view (%s)", c.flags.portfolioView.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"view",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.portfolioView.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.sortBy, "sort-by", "s",
		fmt.Sprintf("sort results by (%s)", c.flags.sortBy.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"sort-by",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.sortBy.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.sortOrder, "sort-order", "o",
		fmt.Sprintf("sort order (%s)", c.flags.sortOrder.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"sort-order",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.sortOrder.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
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

var totalsRenderDescriptor = RenderDescriptor{
	ObjectPath: ".totals",
	Values: []RenderValue{
		{Header: "Net Account Value", Path: ".totalMarketValue"},
		{Header: "Total Gain $", Path: ".totalGainLoss"},
		{Header: "Total Gain %", Path: ".totalGainLossPct"},
		{Header: "Day's Gain Unrealized $", Path: ".todaysGainLoss"},
		{Header: "Day's Gain Unrealized %", Path: ".todaysGainLossPct"},
		{Header: "Cash Balance", Path: ".cashBalance"},
	},
	DefaultValue: "",
	SpaceAfter:   false,
}

var lotsRenderDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".lots",
		Values: []RenderValue{
			{Header: "Price", Path: ".price"},
			{Header: "Term Code", Path: ".termCode"},
			{Header: "Day's Gain $", Path: ".daysGain"},
			{Header: "Day's Gain %", Path: ".daysGainPct"},
			{Header: "Market Value", Path: ".marketValue"},
			{Header: "Total Cost", Path: ".totalCost"},
			{Header: "Total Cost For Gain %", Path: ".totalCostForGainPct"},
			{Header: "Total Gain", Path: ".totalGain"},
			{Header: "Lot Source Code", Path: ".lotSourceCode"},
			{Header: "Original Qty", Path: ".originalQty"},
			{Header: "Remaining Qty", Path: ".remainingQty"},
			{Header: "Available Qty", Path: ".availableQty"},
			{Header: "Order No", Path: ".orderNo"},
			{Header: "Leg No", Path: ".legNo"},
			{Header: "Acquired Date", Path: ".acquiredDate", Transformer: dateTransformerMs},
			{Header: "Location Code", Path: ".locationCode"},
			{Header: "Exchange Rate", Path: ".exchangeRate"},
			{Header: "Settlement Currency", Path: ".settlementCurrency"},
			{Header: "Payment Currency", Path: ".paymentCurrency"},
			{Header: "Adj Price", Path: ".adjPrice"},
			{Header: "Commissions Per Share", Path: ".commPerShare"},
			{Header: "Fees Per Share", Path: ".feesPerShare"},
			{Header: "Adjusted Premium", Path: ".premiumAdj"},
			{Header: "Short Type", Path: ".shortType"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
}

func GetQuickViewRenderDescriptor(withLots bool) []RenderDescriptor {
	subObjects := []RenderDescriptor(nil)
	if withLots {
		subObjects = lotsRenderDescriptor
	}
	return []RenderDescriptor{
		{
			ObjectPath: ".positions",
			Values: []RenderValue{
				{Header: "Symbol", Path: ".product.symbol"},
				{Header: "Security Type", Path: ".product.securityType"},
				{Header: "Symbol Description", Path: ".symbolDescription"},
				{Header: "Date Acquired", Path: ".dateAcquired", Transformer: dateTransformerMs},
				{Header: "Price Paid", Path: ".pricePaid"},
				{Header: "Commissions", Path: ".commissions"},
				{Header: "Other Fees", Path: ".otherFees"},
				{Header: "Quantity", Path: ".quantity"},
				{Header: "Position Indicator", Path: ".positionIndicator"},
				{Header: "Position Type", Path: ".positionType"},
				{Header: "Day's Gain $", Path: ".daysGain"},
				{Header: "Day's Gain %", Path: ".daysGainPct"},
				{Header: "Market Value", Path: ".marketValue"},
				{Header: "Total Cost", Path: ".totalCost"},
				{Header: "Total Gain $", Path: ".totalGain"},
				{Header: "Total Gain %", Path: ".totalGainPct"},
				{Header: "% of Portfolio", Path: ".pctOfPortfolio"},
				{Header: "Cost Per Share", Path: ".costPerShare"},
				{Header: "Today's Commissions", Path: ".todayCommissions"},
				{Header: "Today's Fees", Path: ".todayFees"},
				{Header: "Today's Price Paid", Path: ".todayPricePaid"},
				{Header: "Today's Quantity", Path: ".todayQuantity"},
				{Header: "Adj Previous Close", Path: ".adjPrevClose"},

				{Header: "Last Trade", Path: ".quick.lastTrade"},
				{Header: "Last Trade Time", Path: ".quick.lastTradeTime", Transformer: dateTimeTransformer},
				{Header: "Change $", Path: ".quick.change"},
				{Header: "Change %", Path: ".quick.changePct"},
				{Header: "Volume", Path: ".quick.volume"},
				{Header: "Quote Status", Path: ".quick.quoteStatus"},
				{Header: "7-Day Current Yield", Path: ".quick.sevenDayCurrentYield"},
				{Header: "Annual Total Return", Path: ".quick.annualTotalReturn"},
				{Header: "Weighted Average Maturity", Path: ".quick.weightedAverageMaturity"},
			},
			SubObjects:   subObjects,
			DefaultValue: "",
			SpaceAfter:   true,
		},
		totalsRenderDescriptor,
	}
}

func GetPerformanceViewRenderDescriptor(withLots bool) []RenderDescriptor {
	subObjects := []RenderDescriptor(nil)
	if withLots {
		subObjects = lotsRenderDescriptor
	}
	return []RenderDescriptor{
		{
			ObjectPath: ".positions",
			Values: []RenderValue{
				{Header: "Symbol", Path: ".product.symbol"},
				{Header: "Security Type", Path: ".product.securityType"},
				{Header: "Symbol Description", Path: ".symbolDescription"},
				{Header: "Date Acquired", Path: ".dateAcquired", Transformer: dateTransformerMs},
				{Header: "Price Paid", Path: ".pricePaid"},
				{Header: "Commissions", Path: ".commissions"},
				{Header: "Other Fees", Path: ".otherFees"},
				{Header: "Quantity", Path: ".quantity"},
				{Header: "Position Indicator", Path: ".positionIndicator"},
				{Header: "Position Type", Path: ".positionType"},
				{Header: "Day's Gain $", Path: ".daysGain"},
				{Header: "Day's Gain %", Path: ".daysGainPct"},
				{Header: "Market Value", Path: ".marketValue"},
				{Header: "Total Cost", Path: ".totalCost"},
				{Header: "Total Gain $", Path: ".totalGain"},
				{Header: "Total Gain %", Path: ".totalGainPct"},
				{Header: "% of Portfolio", Path: ".pctOfPortfolio"},
				{Header: "Cost Per Share", Path: ".costPerShare"},
				{Header: "Today's Commissions", Path: ".todayCommissions"},
				{Header: "Today's Fees", Path: ".todayFees"},
				{Header: "Today's Price Paid", Path: ".todayPricePaid"},
				{Header: "Today's Quantity", Path: ".todayQuantity"},
				{Header: "Adj Previous Close", Path: ".adjPrevClose"},
				{Header: "Change $", Path: ".performance.change"},
				{Header: "Change %", Path: ".performance.changePct"},
				{Header: "Last Trade", Path: ".performance.lastTrade"},
				{Header: "Day's Gain", Path: ".performance.daysGain"},
				{Header: "Total Gain $", Path: ".performance.totalGain"},
				{Header: "Total Gain %", Path: ".performance.totalGainPct"},
				{Header: "Market Value", Path: ".performance.marketValue"},
				{Header: "Quote Status", Path: ".performance.quoteStatus"},
				{Header: "Last Trade Time", Path: ".performance.lastTradeTime", Transformer: dateTimeTransformer},
			},
			SubObjects:   subObjects,
			DefaultValue: "",
			SpaceAfter:   true,
		},
		totalsRenderDescriptor,
	}
}

func GetFundamentalViewRenderDescriptor(withLots bool) []RenderDescriptor {
	subObjects := []RenderDescriptor(nil)
	if withLots {
		subObjects = lotsRenderDescriptor
	}
	return []RenderDescriptor{
		{
			ObjectPath: ".positions",
			Values: []RenderValue{
				{Header: "Symbol", Path: ".product.symbol"},
				{Header: "Security Type", Path: ".product.securityType"},
				{Header: "Symbol Description", Path: ".symbolDescription"},
				{Header: "Date Acquired", Path: ".dateAcquired", Transformer: dateTransformerMs},
				{Header: "Price Paid", Path: ".pricePaid"},
				{Header: "Commissions", Path: ".commissions"},
				{Header: "Other Fees", Path: ".otherFees"},
				{Header: "Quantity", Path: ".quantity"},
				{Header: "Position Indicator", Path: ".positionIndicator"},
				{Header: "Position Type", Path: ".positionType"},
				{Header: "Day's Gain $", Path: ".daysGain"},
				{Header: "Day's Gain %", Path: ".daysGainPct"},
				{Header: "Market Value", Path: ".marketValue"},
				{Header: "Total Cost", Path: ".totalCost"},
				{Header: "Total Gain $", Path: ".totalGain"},
				{Header: "Total Gain %", Path: ".totalGainPct"},
				{Header: "% of Portfolio", Path: ".pctOfPortfolio"},
				{Header: "Cost Per Share", Path: ".costPerShare"},
				{Header: "Today's Commissions", Path: ".todayCommissions"},
				{Header: "Today's Fees", Path: ".todayFees"},
				{Header: "Today's Price Paid", Path: ".todayPricePaid"},
				{Header: "Today's Quantity", Path: ".todayQuantity"},
				{Header: "Adj Previous Close", Path: ".adjPrevClose"},
				{Header: "Last Trade", Path: ".fundamental.lastTrade"},
				{Header: "Last Trade Time", Path: ".fundamental.lastTradeTime", Transformer: dateTimeTransformer},
				{Header: "Change $", Path: ".fundamental.change"},
				{Header: "Change %", Path: ".fundamental.changePct"},
				{Header: "P/E Ratio", Path: ".fundamental.peRatio"},
				{Header: "EPS", Path: ".fundamental.eps"},
				{Header: "Dividend", Path: ".fundamental.dividend"},
				{Header: "Dividend Yield", Path: ".fundamental.divYield"},
				{Header: "Market Cap", Path: ".fundamental.marketCap"},
				{Header: "52-Week Range", Path: ".fundamental.week52Range"},
				{Header: "Quote Status", Path: ".fundamental.quoteStatus"},
			},
			SubObjects:   subObjects,
			DefaultValue: "",
			SpaceAfter:   true,
		},
		totalsRenderDescriptor,
	}
}

func GetOptionsWatchViewRenderDescriptor(withLots bool) []RenderDescriptor {
	subObjects := []RenderDescriptor(nil)
	if withLots {
		subObjects = lotsRenderDescriptor
	}

	return []RenderDescriptor{
		{
			ObjectPath: ".positions",
			Values: []RenderValue{
				{Header: "Symbol", Path: ".product.symbol"},
				{Header: "Security Type", Path: ".product.securityType"},
				{Header: "Symbol Description", Path: ".symbolDescription"},
				{Header: "Date Acquired", Path: ".dateAcquired", Transformer: dateTransformerMs},
				{Header: "Price Paid", Path: ".pricePaid"},
				{Header: "Commissions", Path: ".commissions"},
				{Header: "Other Fees", Path: ".otherFees"},
				{Header: "Quantity", Path: ".quantity"},
				{Header: "Position Indicator", Path: ".positionIndicator"},
				{Header: "Position Type", Path: ".positionType"},
				{Header: "Day's Gain $", Path: ".daysGain"},
				{Header: "Day's Gain %", Path: ".daysGainPct"},
				{Header: "Market Value", Path: ".marketValue"},
				{Header: "Total Cost", Path: ".totalCost"},
				{Header: "Total Gain $", Path: ".totalGain"},
				{Header: "Total Gain %", Path: ".totalGainPct"},
				{Header: "% of Portfolio", Path: ".pctOfPortfolio"},
				{Header: "Cost Per Share", Path: ".costPerShare"},
				{Header: "Today's Commissions", Path: ".todayCommissions"},
				{Header: "Today's Fees", Path: ".todayFees"},
				{Header: "Today's Price Paid", Path: ".todayPricePaid"},
				{Header: "Today's Quantity", Path: ".todayQuantity"},
				{Header: "Adj Previous Close", Path: ".adjPrevClose"},

				{Header: "Base Symbol And Price", Path: ".optionsWatch.baseSymbolAndPrice"},
				{Header: "Premium", Path: ".optionsWatch.premium"},
				{Header: "Last Trade", Path: ".optionsWatch.lastTrade"},
				{Header: "Bid", Path: ".optionsWatch.bid"},
				{Header: "Ask", Path: ".optionsWatch.ask"},
				{Header: "Quote Status", Path: ".optionsWatch.quoteStatus"},
				{Header: "Last Trade Time", Path: ".optionsWatch.lastTradeTime", Transformer: dateTimeTransformer},
			},
			SubObjects:   subObjects,
			DefaultValue: "",
			SpaceAfter:   true,
		},
		totalsRenderDescriptor,
	}
}

func GetCompleteViewRenderDescriptor(withLots bool) []RenderDescriptor {
	subObjects := []RenderDescriptor(nil)
	if withLots {
		subObjects = lotsRenderDescriptor
	}
	return []RenderDescriptor{
		{
			ObjectPath: ".positions",
			Values: []RenderValue{
				{Header: "Symbol", Path: ".product.symbol"},
				{Header: "Security Type", Path: ".product.securityType"},
				{Header: "Symbol Description", Path: ".symbolDescription"},
				{Header: "Date Acquired", Path: ".dateAcquired", Transformer: dateTransformerMs},
				{Header: "Price Paid", Path: ".pricePaid"},
				{Header: "Commissions", Path: ".commissions"},
				{Header: "Other Fees", Path: ".otherFees"},
				{Header: "Quantity", Path: ".quantity"},
				{Header: "Position Indicator", Path: ".positionIndicator"},
				{Header: "Position Type", Path: ".positionType"},
				{Header: "Day's Gain $", Path: ".daysGain"},
				{Header: "Day's Gain %", Path: ".daysGainPct"},
				{Header: "Market Value", Path: ".marketValue"},
				{Header: "Total Cost", Path: ".totalCost"},
				{Header: "Total Gain $", Path: ".totalGain"},
				{Header: "Total Gain %", Path: ".totalGainPct"},
				{Header: "% of Portfolio", Path: ".pctOfPortfolio"},
				{Header: "Cost Per Share", Path: ".costPerShare"},
				{Header: "Today's Commissions", Path: ".todayCommissions"},
				{Header: "Today's Fees", Path: ".todayFees"},
				{Header: "Today's Price Paid", Path: ".todayPricePaid"},
				{Header: "Today's Quantity", Path: ".todayQuantity"},
				{Header: "Adj Previous Close", Path: ".adjPrevClose"},

				{Header: "Price Adjusted Flag", Path: ".complete.priceAdjustedFlag"},
				{Header: "Price", Path: ".complete.price"},
				{Header: "Adj Price", Path: ".complete.adjPrice"},
				{Header: "Change $", Path: ".complete.change"},
				{Header: "Change %", Path: ".complete.changePct"},
				{Header: "Previous Close", Path: ".complete.prevClose"},
				{Header: "Adj Prev Close", Path: ".complete.adjPrevClose"},
				{Header: "Volume", Path: ".complete.volume"},
				{Header: "Last Trade", Path: ".complete.lastTrade"},
				{Header: "Last Trade Time", Path: ".complete.lastTradeTime", Transformer: dateTimeTransformer},
				{Header: "Adj Last Trade", Path: ".complete.adjLastTrade"},
				{Header: "Symbol Description", Path: ".complete.symbolDescription"},
				{Header: "1-Month Performance", Path: ".complete.perform1Month"},
				{Header: "3-Month Performance", Path: ".complete.perform3Month"},
				{Header: "6-Month Performance", Path: ".complete.perform6Month"},
				{Header: "12-Month Performance", Path: ".complete.perform12Month"},
				{Header: "Prev Day Volume", Path: ".complete.prevDayVolume"},
				{Header: "10-Day Volume", Path: ".complete.tenDayVolume"},
				{Header: "Beta", Path: ".complete.beta"},
				{Header: "10-Day Avg SV", Path: ".complete.sv10DaysAvg"},
				{Header: "20-Day Avg SV", Path: ".complete.sv20DaysAvg"},
				{Header: "1-Month Avg SV", Path: ".complete.sv1MonAvg"},
				{Header: "2-Month Avg SV", Path: ".complete.sv2MonAvg"},
				{Header: "3-Month Avg SV", Path: ".complete.sv3MonAvg"},
				{Header: "4-Month Avg SV", Path: ".complete.sv4MonAvg"},
				{Header: "6-Month Avg SV", Path: ".complete.sv6MonAvg"},
				{Header: "52-Week High", Path: ".complete.week52High"},
				{Header: "52-Week Low", Path: ".complete.week52Low"},
				{Header: "52-Week Range", Path: ".complete.week52Range"},
				{Header: "Market Cap", Path: ".complete.marketCap"},
				{Header: "Day's Range", Path: ".complete.daysRange"},
				{Header: "Delta 52-Week High", Path: ".complete.delta52WkHigh"},
				{Header: "Delta 52-Week Low", Path: ".complete.delta52WkLow"},
				{Header: "Currency", Path: ".complete.currency"},
				{Header: "Exchange", Path: ".complete.exchange"},
				{Header: "Marginable", Path: ".complete.marginable"},
				{Header: "Bid", Path: ".complete.bid"},
				{Header: "Ask", Path: ".complete.ask"},
				{Header: "Bid AskS pread", Path: ".complete.bidAskSpread"},
				{Header: "Bid Size", Path: ".complete.bidSize"},
				{Header: "Ask Size", Path: ".complete.askSize"},
				{Header: "Open", Path: ".complete.open"},
				{Header: "Delta", Path: ".complete.delta"},
				{Header: "Gamma", Path: ".complete.gamma"},
				{Header: "IV %", Path: ".complete.ivPct"},
				{Header: "Rho", Path: ".complete.rho"},
				{Header: "Theta", Path: ".complete.theta"},
				{Header: "Vega", Path: ".complete.vega"},
				{Header: "Premium", Path: ".complete.premium"},
				{Header: "Days To Expiration", Path: ".complete.daysToExpiration"},
				{Header: "Intrinsic Value", Path: ".complete.intrinsicValue"},
				{Header: "Open Interest", Path: ".complete.openInterest"},
				{Header: "Options Adjusted Flag", Path: ".complete.optionsAdjustedFlag"},
				{Header: "Deliverables", Path: ".complete.deliverablesStr"},
				{Header: "Option Multiplier", Path: ".complete.optionMultiplier"},
				{Header: "Base Symbol And Price", Path: ".complete.baseSymbolAndPrice"},
				{Header: "Est Earnings", Path: ".complete.estEarnings"},
				{Header: "EPS", Path: ".complete.eps"},
				{Header: "P/E Ratio", Path: ".complete.peRatio"},
				{Header: "Annual Dividend", Path: ".complete.annualDividend"},
				{Header: "Dividend", Path: ".complete.dividend"},
				{Header: "Div Yield", Path: ".complete.divYield"},
				{Header: "Div Pay Date", Path: ".complete.divPayDate", Transformer: dateTransformerMs},
				{
					Header: "Extended Dividend Date", Path: ".complete.exDividendDate",
					Transformer: dateTransformerMs,
				},
				{Header: "CUSIP", Path: ".complete.cusip"},
				{Header: "Quote Status", Path: ".complete.quoteStatus"},
			},
			SubObjects:   subObjects,
			DefaultValue: "",
			SpaceAfter:   true,
		},
		totalsRenderDescriptor,
	}
}
