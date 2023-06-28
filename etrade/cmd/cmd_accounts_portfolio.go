package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
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
			return c.ViewPortfolio(args[0])
		},
	}

	// Add Flags
	cmd.Flags().BoolVarP(&c.flags.totalsRequired, "totals-required", "t", true, "include totals in results")
	cmd.Flags().BoolVarP(&c.flags.withLots, "with-lots", "l", false, "include lots in results")

	// Initialize Enum Flag Values
	c.flags.portfolioView = *newEnumFlagValue(portfolioViewMap, constants.PortfolioViewQuick)
	c.flags.sortBy = *newEnumFlagValue(sortByMap, constants.PortfolioSortByNil)
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

func (c *CommandAccountsPortfolio) ViewPortfolio(accountId string) error {
	// This determines how many portfolio items will be retrieved in each
	// request. This should normally be set to the max for efficiency, but can
	// be lowered to test the pagination logic.
	const countPerRequest = constants.PortfolioMaxCount

	account, err := GetAccountById(c.Context.Client, accountId)
	if err != nil {
		return err
	}

	response, err := c.Context.Client.ViewPortfolio(
		account.GetIdKey(), countPerRequest, c.flags.sortBy.Value(), c.flags.sortOrder.Value(), "",
		c.flags.marketSession.Value(),
		c.flags.totalsRequired, true, c.flags.portfolioView.Value(),
	)
	if err != nil {
		return err
	}

	positionList, err := etradelib.CreateETradePositionListFromResponse(response)
	if err != nil {
		return err
	}

	for positionList.NextPage() != "" {
		response, err = c.Context.Client.ViewPortfolio(
			account.GetIdKey(), countPerRequest, c.flags.sortBy.Value(), c.flags.sortOrder.Value(),
			positionList.NextPage(),
			c.flags.marketSession.Value(), c.flags.totalsRequired, true, c.flags.portfolioView.Value(),
		)
		if err != nil {
			return err
		}
		err = positionList.AddPageFromResponse(response)
		if err != nil {
			return err
		}
	}

	if c.flags.withLots {
		for _, position := range positionList.GetAllPositions() {
			response, err = c.Context.Client.ListPositionLotsDetails(account.GetIdKey(), position.GetId())
			if err != nil {
				return err
			}
			err = position.AddLotsFromResponse(response)
			if err != nil {
				return err
			}
		}
	}

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

	err = c.Context.Renderer.Render(positionList.AsJsonMap(), renderDescriptor)
	if err != nil {
		return err
	}
	return nil
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

var portfolioViewMap = map[string]enumValueWithHelp[constants.PortfolioView]{
	"performance":  {constants.PortfolioViewPerformance, "performance view"},
	"fundamental":  {constants.PortfolioViewFundamental, "fundamentals view"},
	"optionsWatch": {constants.PortfolioViewOptionsWatch, "options watch view"},
	"quick":        {constants.PortfolioViewQuick, "quick view"},
	"complete":     {constants.PortfolioViewComplete, "complete view"},
}

var sortByMap = map[string]enumValueWithHelp[constants.PortfolioSortBy]{
	"symbol":           {constants.PortfolioSortBySymbol, "sort by the symbol"},
	"typeName":         {constants.PortfolioSortByTypeName, "sort by the type name"},
	"exchangeName":     {constants.PortfolioSortByExchangeName, "sort by the exchange name"},
	"currency":         {constants.PortfolioSortByCurrency, "sort by the currency"},
	"quantity":         {constants.PortfolioSortByQuantity, "sort by the quantity"},
	"longOrShort":      {constants.PortfolioSortByLongOrShort, "sort by the position type (long or short)"},
	"dateAcquired":     {constants.PortfolioSortByDateAcquired, "sort by the date acquired"},
	"pricePaid":        {constants.PortfolioSortByPricePaid, "sort by the price paid"},
	"totalGain":        {constants.PortfolioSortByTotalGain, "sort by the total gain"},
	"totalGainPct":     {constants.PortfolioSortByTotalGainPct, "sort by the total gain percent"},
	"marketValue":      {constants.PortfolioSortByMarketValue, "sort by the market value"},
	"bid":              {constants.PortfolioSortByBid, "sort by the bid"},
	"ask":              {constants.PortfolioSortByAsk, "sort by the ask"},
	"priceChange":      {constants.PortfolioSortByPriceChange, "sort by the price change"},
	"priceChangePct":   {constants.PortfolioSortByPriceChangePct, "sort by the price change percent"},
	"volume":           {constants.PortfolioSortByVolume, "sort by the volume"},
	"week52High":       {constants.PortfolioSortByWeek52High, "sort by the 52-week high"},
	"week52Low":        {constants.PortfolioSortByWeek52Low, "sort by the 52-week low"},
	"eps":              {constants.PortfolioSortByEps, "sort by the earnings per share"},
	"peRatio":          {constants.PortfolioSortByPeRatio, "sort by the price to earnings (P/E) ratio"},
	"optionType":       {constants.PortfolioSortByOptionType, "sort by the option type"},
	"strikePrice":      {constants.PortfolioSortByStrikePrice, "sort by the strike price"},
	"premium":          {constants.PortfolioSortByPremium, "sort by the premium"},
	"expiration":       {constants.PortfolioSortByExpiration, "sort by the expiration"},
	"daysGain":         {constants.PortfolioSortByDaysGain, "sort by the day's gain"},
	"commission":       {constants.PortfolioSortByCommission, "sort by the commission"},
	"marketCap":        {constants.PortfolioSortByMarketCap, "sort by the market cap"},
	"prevClose":        {constants.PortfolioSortByPrevClose, "sort by the previous close"},
	"open":             {constants.PortfolioSortByOpen, "sort by the open"},
	"daysRange":        {constants.PortfolioSortByDaysRange, "sort by the day's range"},
	"totalCost":        {constants.PortfolioSortByTotalCost, "sort by the total cost"},
	"daysGainPct":      {constants.PortfolioSortByDaysGainPct, "sort by the day's gain percentage"},
	"pctOfPortfolio":   {constants.PortfolioSortByPctOfPortfolio, "sort by the percent of portfolio"},
	"lastTradeTime":    {constants.PortfolioSortByLastTradeTime, "sort by the last trade time"},
	"baseSymbolPrice":  {constants.PortfolioSortByBaseSymbolPrice, "sort by the base symbol price"},
	"week52Range":      {constants.PortfolioSortByWeek52Range, "sort by the 52-week range"},
	"lastTrade":        {constants.PortfolioSortByLastTrade, "sort by the last trade"},
	"symbolDesc":       {constants.PortfolioSortBySymbolDesc, "sort by the symbol description"},
	"bidSize":          {constants.PortfolioSortByBidSize, "sort by the bid size"},
	"askSize":          {constants.PortfolioSortByAskSize, "sort by the ask size"},
	"otherFees":        {constants.PortfolioSortByOtherFees, "sort by other fees"},
	"heldAs":           {constants.PortfolioSortByHeldAs, "sort by held as"},
	"optionMultiplier": {constants.PortfolioSortByOptionMultiplier, "sort by the option multiplier"},
	"deliverables":     {constants.PortfolioSortByDeliverables, "sort by the deliverables"},
	"costPerShare":     {constants.PortfolioSortByCostPerShare, "sort by the cost per share"},
	"dividend":         {constants.PortfolioSortByDividend, "sort by the dividend"},
	"divYield":         {constants.PortfolioSortByDivYield, "sort by the dividend yield"},
	"divPayDate":       {constants.PortfolioSortByDivPayDate, "sort by the dividend pay date"},
	"estEarn":          {constants.PortfolioSortByEstEarn, "sort by the estimated earnings"},
	"exDivDate":        {constants.PortfolioSortByExDivDate, "sort by the extended dividend date"},
	"tenDayAvgVol":     {constants.PortfolioSortByTenDayAvgVol, "sort by the 10-day average volume"},
	"beta":             {constants.PortfolioSortByBeta, "sort by the beta"},
	"bidAskSpread":     {constants.PortfolioSortByBidAskSpread, "sort by the bid ask spread"},
	"marginable":       {constants.PortfolioSortByMarginable, "sort by the sum available for margin"},
	"delta52wkHi": {
		constants.PortfolioSortByDelta52wkHi,
		"sort by the high for the 52-week high/low delta calculation",
	},
	"delta52WkLow": {
		constants.PortfolioSortByDelta52WkLow,
		"sort by the low for the 52-week high/low delta calculation",
	},
	"perf1Mon":            {constants.PortfolioSortByPerf1Mon, "sort by the 1-month performance"},
	"annualDiv":           {constants.PortfolioSortByAnnualDiv, "sort by the annual dividend"},
	"perf12Mon":           {constants.PortfolioSortByPerf12Mon, "sort by the 12-month performance"},
	"perf3Mon":            {constants.PortfolioSortByPerf3Mon, "sort by the 3-month performance"},
	"perf6Mon":            {constants.PortfolioSortByPerf6Mon, "sort by the 6-month performance"},
	"preDayVol":           {constants.PortfolioSortByPreDayVol, "sort by the previous day's volume"},
	"sv1MonAvg":           {constants.PortfolioSortBySv1MonAvg, "sort by the 1-month average stochastic volatility"},
	"sv10DayAvg":          {constants.PortfolioSortBySv10DayAvg, "sort by the 10-day average stochastic volatility"},
	"sv20DayAvg":          {constants.PortfolioSortBySv20DayAvg, "sort by the 20-day average stochastic volatility"},
	"sv2MonAvg":           {constants.PortfolioSortBySv2MonAvg, "sort by the 2-month average stochastic volatility"},
	"sv3MonAvg":           {constants.PortfolioSortBySv3MonAvg, "sort by the 3-month average stochastic volatility"},
	"sv4MonAvg":           {constants.PortfolioSortBySv4MonAvg, "sort by the 4-month average stochastic volatility"},
	"sv6MonAvg":           {constants.PortfolioSortBySv6MonAvg, "sort by the 6-month average stochastic volatility"},
	"delta":               {constants.PortfolioSortByDelta, "sort by the delta"},
	"gamma":               {constants.PortfolioSortByGamma, "sort by the gamma"},
	"ivPct":               {constants.PortfolioSortByIvPct, "sort by the implied volatility (IV) percentage"},
	"theta":               {constants.PortfolioSortByTheta, "sort by the theta"},
	"vega":                {constants.PortfolioSortByVega, "sort by the vega"},
	"adjNonadjFlag":       {constants.PortfolioSortByAdjNonadjFlag, "sort by the adjusted/non-adjusted flag"},
	"daysExpiration":      {constants.PortfolioSortByDaysExpiration, "sort by the days remaining until expiration"},
	"openInterest":        {constants.PortfolioSortByOpenInterest, "sort by the open interest"},
	"intrinsicValue":      {constants.PortfolioSortByIntrinsicValue, "sort by the intrinsic value"},
	"rho":                 {constants.PortfolioSortByRho, "sort by the rho"},
	"typeCode":            {constants.PortfolioSortByTypeCode, "sort by the type code"},
	"displaySymbol":       {constants.PortfolioSortByDisplaySymbol, "sort by the display symbol"},
	"afterHoursPctChange": {constants.PortfolioSortByAfterHoursPctChange, "sort by the after-hours percentage change"},
	"preMarketPctChange":  {constants.PortfolioSortByPreMarketPctChange, "sort by the pre-market percentage change"},
	"expandCollapseFlag":  {constants.PortfolioSortByExpandCollapseFlag, "sort by the expand/collapse flag"},
}
