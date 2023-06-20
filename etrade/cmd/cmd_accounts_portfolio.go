package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type accountsPortfolioFlags struct {
	totalsRequired bool
	lotsRequired   bool
	portfolioView  enumFlagValue[constants.PortfolioView]
	sortBy         enumFlagValue[constants.PortfolioSortBy]
	sortOrder      enumFlagValue[constants.SortOrder]
	marketSession  enumFlagValue[constants.MarketSession]
}

type CommandAccountsPortfolio struct {
	Resources *CommandResources
	flags     accountsPortfolioFlags
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
	cmd.Flags().BoolVarP(&c.flags.lotsRequired, "lots-required", "l", false, "include lots in results")

	// Initialize Enum Flag Values
	c.flags.portfolioView = *newEnumFlagValue(portfolioViewMap, constants.PortfolioViewNil)
	c.flags.sortBy = *newEnumFlagValue(sortByMap, constants.PortfolioSortByNil)
	c.flags.sortOrder = *newEnumFlagValue(sortOrderMap, constants.SortOrderNil)
	c.flags.marketSession = *newEnumFlagValue(marketSessionMap, constants.MarketSessionNil)

	// Add Enum Flags
	cmd.Flags().VarP(
		&c.flags.portfolioView, "portfolio-view", "v",
		fmt.Sprintf("portfolio view (%s)", c.flags.portfolioView.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"portfolio-view",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.portfolioView.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.portfolioView, "sort-by", "s",
		fmt.Sprintf("sort results by (%s)", c.flags.sortBy.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"sort-by",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.sortBy.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.portfolioView, "sort-order", "o",
		fmt.Sprintf("sort order (%s)", c.flags.sortOrder.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"sort-order",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.sortOrder.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.portfolioView, "market-session", "m",
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

func (c *CommandAccountsPortfolio) ViewPortfolio(accountKeyId string) error {
	response, err := c.Resources.Client.ViewPortfolio(
		accountKeyId, -1, c.flags.sortBy.Value(), c.flags.sortOrder.Value(), -1, c.flags.marketSession.Value(), true,
		true, c.flags.portfolioView.Value(),
	)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(c.Resources.OFile, string(response))
	return nil
}

var portfolioViewMap = map[string]enumValueWithHelp[constants.PortfolioView]{
	"performance":  {constants.PortfolioViewPerformance, "performance view"},
	"fundamentals": {constants.PortfolioViewFundamental, "fundamentals view"},
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
