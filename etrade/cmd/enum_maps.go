package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"

var sortOrderMap = map[string]enumValueWithHelp[constants.SortOrder]{
	"ascending":  {constants.SortOrderAsc, "sort in ascending order"},
	"descending": {constants.SortOrderDesc, "sort in descending order"},
}

var marketSessionMap = map[string]enumValueWithHelp[constants.MarketSession]{
	"regular":  {constants.MarketSessionRegular, "regular market session"},
	"extended": {constants.MarketSessionExtended, "extended market session"},
}

var portfolioViewMap = map[string]enumValueWithHelp[constants.PortfolioView]{
	"performance":  {constants.PortfolioViewPerformance, "performance view"},
	"fundamental":  {constants.PortfolioViewFundamental, "fundamentals view"},
	"optionsWatch": {constants.PortfolioViewOptionsWatch, "options watch view"},
	"quick":        {constants.PortfolioViewQuick, "quick view"},
	"complete":     {constants.PortfolioViewComplete, "complete view"},
}

var portfolioSortByMap = map[string]enumValueWithHelp[constants.PortfolioSortBy]{
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

var alertCategoryMap = map[string]enumValueWithHelp[constants.AlertCategory]{
	"stock":   {constants.AlertCategoryStock, "only stock-related alerts"},
	"account": {constants.AlertCategoryAccount, "only account-related alerts"},
}

var alertStatusMap = map[string]enumValueWithHelp[constants.AlertStatus]{
	"read":    {constants.AlertStatusRead, "only read alerts"},
	"unread":  {constants.AlertStatusUnread, "only unread alerts"},
	"deleted": {constants.AlertStatusDeleted, "only deleted alerts"},
}

var quoteDetailMap = map[string]enumValueWithHelp[constants.QuoteDetailFlag]{
	"all":         {constants.QuoteDetailAll, "get all details"},
	"fundamental": {constants.QuoteDetailFundamental, "get fundamental details"},
	"intraday":    {constants.QuoteDetailIntraday, "get intraday details"},
	"options":     {constants.QuoteDetailOptions, "get options details"},
	"week52":      {constants.QuoteDetailWeek52, "get 52-week details"},
	"mutualFund":  {constants.QuoteDetailMutualFund, "get mutual fund details"},
}

var optionCategoryMap = map[string]enumValueWithHelp[constants.OptionCategory]{
	"standard": {constants.OptionCategoryStandard, "only standard options"},
	"all":      {constants.OptionCategoryAll, "all options"},
	"mini":     {constants.OptionCategoryMini, "only mini options"},
}

var optionChainTypeMap = map[string]enumValueWithHelp[constants.OptionChainType]{
	"call":    {constants.OptionChainTypeCall, "only call options"},
	"put":     {constants.OptionChainTypePut, "only put options"},
	"callPut": {constants.OptionChainTypeCallPut, "call and put options"},
}

var optionPriceTypeMap = map[string]enumValueWithHelp[constants.OptionPriceType]{
	"extendedHours": {constants.OptionPriceTypeExtendedHours, "only extended hours price types"},
	"all":           {constants.OptionPriceTypeAll, "all price types"},
}

var optionExpiryTypeMap = map[string]enumValueWithHelp[constants.OptionExpiryType]{
	"unspecified": {constants.OptionExpiryTypeUnspecified, "unspecified expiry type"},
	"daily":       {constants.OptionExpiryTypeDaily, "daily expiry type"},
	"weekly":      {constants.OptionExpiryTypeWeekly, "weekly expiry type"},
	"monthly":     {constants.OptionExpiryTypeMonthly, "monthly expiry type"},
	"quarterly":   {constants.OptionExpiryTypeQuarterly, "quarterly expiry type"},
	"vix":         {constants.OptionExpiryTypeVix, "VIX expiry type"},
	"all":         {constants.OptionExpiryTypeAll, "all expiry types"},
	"monthEnd":    {constants.OptionExpiryTypeMonthEnd, "month-end expiry type"},
}
