package constants

// PortfolioMaxCount is the maximum count that can be included in a View
// Portfolio request. Note that ETrade does not document this value, so I
// determined it empirically by increasing the count until I got a
// 500 Internal Server Error. I do not have a portfolio with this many
// positions in it, so I'm not sure if there's a practical lower limit.
const PortfolioMaxCount = 65535

// PortfolioSortBy specifies the column by which the portfolio results will be sorted.
// See the constants below for semantics.
type PortfolioSortBy int

const (
	// PortfolioSortByNil indicates no sorting
	// (e.g. to make a query use the default value from ETrade)
	PortfolioSortByNil PortfolioSortBy = iota

	// PortfolioSortBySymbol sorts results by the symbol
	PortfolioSortBySymbol

	// PortfolioSortByTypeName sorts results by the type name
	PortfolioSortByTypeName

	// PortfolioSortByExchangeName sorts results by the exchange name
	PortfolioSortByExchangeName

	// PortfolioSortByCurrency sorts results by the currency
	PortfolioSortByCurrency

	// PortfolioSortByQuantity sorts results by the quantity
	PortfolioSortByQuantity

	// PortfolioSortByLongOrShort sorts results by the position type (long or short)
	PortfolioSortByLongOrShort

	// PortfolioSortByDateAcquired sorts results by the date acquired
	PortfolioSortByDateAcquired

	// PortfolioSortByPricePaid sorts results by the price paid
	PortfolioSortByPricePaid

	// PortfolioSortByTotalGain sorts results by the total gain
	PortfolioSortByTotalGain

	// PortfolioSortByTotalGainPct sorts results by the total gain percent
	PortfolioSortByTotalGainPct

	// PortfolioSortByMarketValue sorts results by the market value
	PortfolioSortByMarketValue

	// PortfolioSortByBid sorts results by the bid
	// TODO: Confirm whether this is supposed to be "Bid" (ETrade documentation says "bi")
	PortfolioSortByBid

	// PortfolioSortByAsk sorts results by the ask
	PortfolioSortByAsk

	// PortfolioSortByPriceChange sorts results by the price change
	PortfolioSortByPriceChange

	// PortfolioSortByPriceChangePct sorts results by the price change percent
	PortfolioSortByPriceChangePct

	// PortfolioSortByVolume sorts results by the volume
	PortfolioSortByVolume

	// PortfolioSortByWeek52High sorts results by the 52-week high
	PortfolioSortByWeek52High

	// PortfolioSortByWeek52Low sorts results by the 52-week low
	PortfolioSortByWeek52Low

	// PortfolioSortByEps sorts results by the earnings per share
	PortfolioSortByEps

	// PortfolioSortByPeRatio sorts results by the price to earnings (P/E) ratio
	PortfolioSortByPeRatio

	// PortfolioSortByOptionType sorts results by the option type
	PortfolioSortByOptionType

	// PortfolioSortByStrikePrice sorts results by the strike price
	PortfolioSortByStrikePrice

	// PortfolioSortByPremium sorts results by the premium
	PortfolioSortByPremium

	// PortfolioSortByExpiration sorts results by the expiration
	PortfolioSortByExpiration

	// PortfolioSortByDaysGain sorts results by the day's gain
	PortfolioSortByDaysGain

	// PortfolioSortByCommission sorts results by the commission
	PortfolioSortByCommission

	// PortfolioSortByMarketCap sorts results by the market cap
	PortfolioSortByMarketCap

	// PortfolioSortByPrevClose sorts results by the previous close
	PortfolioSortByPrevClose

	// PortfolioSortByOpen sorts results by the open
	PortfolioSortByOpen

	// PortfolioSortByDaysRange sorts results by the day's range
	PortfolioSortByDaysRange

	// PortfolioSortByTotalCost sorts results by the total cost
	PortfolioSortByTotalCost

	// PortfolioSortByDaysGainPct sorts results by the day's gain percentage
	PortfolioSortByDaysGainPct

	// PortfolioSortByPctOfPortfolio sorts results by the percent of portfolio
	PortfolioSortByPctOfPortfolio

	// PortfolioSortByLastTradeTime sorts results by the last trade time
	PortfolioSortByLastTradeTime

	// PortfolioSortByBaseSymbolPrice sorts results by the base symbol price
	PortfolioSortByBaseSymbolPrice

	// PortfolioSortByWeek52Range sorts results by the 52-week range
	PortfolioSortByWeek52Range

	// PortfolioSortByLastTrade sorts results by the last trade
	PortfolioSortByLastTrade

	// PortfolioSortBySymbolDesc sorts results by the symbol description
	PortfolioSortBySymbolDesc

	// PortfolioSortByBidSize sorts results by the bid size
	PortfolioSortByBidSize

	// PortfolioSortByAskSize sorts results by the ask size
	PortfolioSortByAskSize

	// PortfolioSortByOtherFees sorts results by other fees
	PortfolioSortByOtherFees

	// PortfolioSortByHeldAs sorts results by held as
	PortfolioSortByHeldAs

	// PortfolioSortByOptionMultiplier sorts results by the option multiplier
	PortfolioSortByOptionMultiplier

	// PortfolioSortByDeliverables sorts results by the deliverables
	PortfolioSortByDeliverables

	// PortfolioSortByCostPerShare sorts results by the cost per share
	PortfolioSortByCostPerShare

	// PortfolioSortByDividend sorts results by the dividend
	PortfolioSortByDividend

	// PortfolioSortByDivYield sorts results by the dividend yield
	PortfolioSortByDivYield

	// PortfolioSortByDivPayDate sorts results by the dividend pay date
	PortfolioSortByDivPayDate

	// PortfolioSortByEstEarn sorts results by the estimated earnings
	PortfolioSortByEstEarn

	// PortfolioSortByExDivDate sorts results by the extended dividend date
	PortfolioSortByExDivDate

	// PortfolioSortByTenDayAvgVol sorts results by the 10-day average volume
	PortfolioSortByTenDayAvgVol

	// PortfolioSortByBeta sorts results by the beta
	PortfolioSortByBeta

	// PortfolioSortByBidAskSpread sorts results by the bid ask spread
	PortfolioSortByBidAskSpread

	// PortfolioSortByMarginable sorts results by the sum available for margin
	PortfolioSortByMarginable

	// PortfolioSortByDelta52wkHi sorts results by the high for the 52-week high/low delta calculation
	PortfolioSortByDelta52wkHi

	// PortfolioSortByDelta52WkLow sorts results by the low for the 52-week high/low delta calculation
	PortfolioSortByDelta52WkLow

	// PortfolioSortByPerf1Mon sorts results by the 1-month performance
	PortfolioSortByPerf1Mon

	// PortfolioSortByAnnualDiv sorts results by the annual dividend
	PortfolioSortByAnnualDiv

	// PortfolioSortByPerf12Mon sorts results by the 12-month performance
	PortfolioSortByPerf12Mon

	// PortfolioSortByPerf3Mon sorts results by the 3-month performance
	PortfolioSortByPerf3Mon

	// PortfolioSortByPerf6Mon sorts results by the 6-month performance
	PortfolioSortByPerf6Mon

	// PortfolioSortByPreDayVol sorts results by the previous day's volume
	PortfolioSortByPreDayVol

	// PortfolioSortBySv1MonAvg sorts results by the 1-month average stochastic volatility
	PortfolioSortBySv1MonAvg

	// PortfolioSortBySv10DayAvg sorts results by the 10-day average stochastic volatility
	PortfolioSortBySv10DayAvg

	// PortfolioSortBySv20DayAvg sorts results by the 20-day average stochastic volatility
	PortfolioSortBySv20DayAvg

	// PortfolioSortBySv2MonAvg sorts results by the 2-month average stochastic volatility
	PortfolioSortBySv2MonAvg

	// PortfolioSortBySv3MonAvg sorts results by the 3-month average stochastic volatility
	PortfolioSortBySv3MonAvg

	// PortfolioSortBySv4MonAvg sorts results by the 4-month average stochastic volatility
	PortfolioSortBySv4MonAvg

	// PortfolioSortBySv6MonAvg sorts results by the 6-month average stochastic volatility
	PortfolioSortBySv6MonAvg

	// PortfolioSortByDelta sorts results by the delta
	PortfolioSortByDelta

	// PortfolioSortByGamma sorts results by the gamma
	PortfolioSortByGamma

	// PortfolioSortByIvPct sorts results by the implied volatility (IV) percentage
	PortfolioSortByIvPct

	// PortfolioSortByTheta sorts results by the theta
	PortfolioSortByTheta

	// PortfolioSortByVega sorts results by the vega
	PortfolioSortByVega

	// PortfolioSortByAdjNonadjFlag sorts results by the adjusted/non-adjusted flag
	PortfolioSortByAdjNonadjFlag

	// PortfolioSortByDaysExpiration sorts results by the days remaining until expiration
	PortfolioSortByDaysExpiration

	// PortfolioSortByOpenInterest sorts results by the open interest
	PortfolioSortByOpenInterest

	// PortfolioSortByIntrinsicValue sorts results by the intrinsic value
	PortfolioSortByIntrinsicValue

	// PortfolioSortByRho sorts results by the rho
	PortfolioSortByRho

	// PortfolioSortByTypeCode sorts results by the type code
	PortfolioSortByTypeCode

	// PortfolioSortByDisplaySymbol sorts results by the display symbol
	PortfolioSortByDisplaySymbol

	// PortfolioSortByAfterHoursPctChange sorts results by the after-hours percentage change
	PortfolioSortByAfterHoursPctChange

	// PortfolioSortByPreMarketPctChange sorts results by the pre-market percentage change
	PortfolioSortByPreMarketPctChange

	// PortfolioSortByExpandCollapseFlag sorts results by the expand/collapse flag
	PortfolioSortByExpandCollapseFlag
)

// PortfolioView specifies the type of portfolio view to retrieve
// See the constants below for semantics.
type PortfolioView int

const (
	// PortfolioViewNil indicates no portfolio view
	// (e.g. to make a query use the default value from ETrade)
	PortfolioViewNil PortfolioView = iota

	// PortfolioViewPerformance gets the performance view
	// See https://apisb.etrade.com/docs/api/account/api-portfolio-v1.html#/definitions/PerformanceView
	PortfolioViewPerformance

	// PortfolioViewFundamental gets the fundamental view
	// See https://apisb.etrade.com/docs/api/account/api-portfolio-v1.html#/definitions/FundamentalView
	PortfolioViewFundamental

	// PortfolioViewOptionsWatch gets the options watch view
	// See https://apisb.etrade.com/docs/api/account/api-portfolio-v1.html#/definitions/OptionsWatchView
	PortfolioViewOptionsWatch

	// PortfolioViewQuick gets the quick view
	// See https://apisb.etrade.com/docs/api/account/api-portfolio-v1.html#/definitions/QuickView
	PortfolioViewQuick

	// PortfolioViewComplete gets the complete view
	// See https://apisb.etrade.com/docs/api/account/api-portfolio-v1.html#/definitions/CompleteView
	PortfolioViewComplete
)

var portfolioSortByToString = map[PortfolioSortBy]string{
	PortfolioSortBySymbol:              "SYMBOL",
	PortfolioSortByTypeName:            "TYPE_NAME",
	PortfolioSortByExchangeName:        "EXCHANGE_NAME",
	PortfolioSortByCurrency:            "CURRENCY",
	PortfolioSortByQuantity:            "QUANTITY",
	PortfolioSortByLongOrShort:         "LONG_OR_SHORT",
	PortfolioSortByDateAcquired:        "DATE_ACQUIRED",
	PortfolioSortByPricePaid:           "PRICEPAID",
	PortfolioSortByTotalGain:           "TOTAL_GAIN",
	PortfolioSortByTotalGainPct:        "TOTAL_GAIN_PCT",
	PortfolioSortByMarketValue:         "MARKET_VALUE",
	PortfolioSortByBid:                 "BI", // TODO: Confirm whether this is supposed to be "BID" (ETrade documentation says "BI")
	PortfolioSortByAsk:                 "ASK",
	PortfolioSortByPriceChange:         "PRICE_CHANGE",
	PortfolioSortByPriceChangePct:      "PRICE_CHANGE_PCT",
	PortfolioSortByVolume:              "VOLUME",
	PortfolioSortByWeek52High:          "WEEK_52_HIGH",
	PortfolioSortByWeek52Low:           "WEEK_52_LOW",
	PortfolioSortByEps:                 "EPS",
	PortfolioSortByPeRatio:             "PE_RATIO",
	PortfolioSortByOptionType:          "OPTION_TYPE",
	PortfolioSortByStrikePrice:         "STRIKE_PRICE",
	PortfolioSortByPremium:             "PREMIUM",
	PortfolioSortByExpiration:          "EXPIRATION",
	PortfolioSortByDaysGain:            "DAYS_GAIN",
	PortfolioSortByCommission:          "COMMISSION",
	PortfolioSortByMarketCap:           "MARKETCAP",
	PortfolioSortByPrevClose:           "PREV_CLOSE",
	PortfolioSortByOpen:                "OPEN",
	PortfolioSortByDaysRange:           "DAYS_RANGE",
	PortfolioSortByTotalCost:           "TOTAL_COST",
	PortfolioSortByDaysGainPct:         "DAYS_GAIN_PCT",
	PortfolioSortByPctOfPortfolio:      "PCT_OF_PORTFOLIO",
	PortfolioSortByLastTradeTime:       "LAST_TRADE_TIME",
	PortfolioSortByBaseSymbolPrice:     "BASE_SYMBOL_PRICE",
	PortfolioSortByWeek52Range:         "WEEK_52_RANGE",
	PortfolioSortByLastTrade:           "LAST_TRADE",
	PortfolioSortBySymbolDesc:          "SYMBOL_DESC",
	PortfolioSortByBidSize:             "BID_SIZE",
	PortfolioSortByAskSize:             "ASK_SIZE",
	PortfolioSortByOtherFees:           "OTHER_FEES",
	PortfolioSortByHeldAs:              "HELD_AS",
	PortfolioSortByOptionMultiplier:    "OPTION_MULTIPLIER",
	PortfolioSortByDeliverables:        "DELIVERABLES",
	PortfolioSortByCostPerShare:        "COST_PERSHARE",
	PortfolioSortByDividend:            "DIVIDEND",
	PortfolioSortByDivYield:            "DIV_YIELD",
	PortfolioSortByDivPayDate:          "DIV_PAY_DATE",
	PortfolioSortByEstEarn:             "EST_EARN",
	PortfolioSortByExDivDate:           "EX_DIV_DATE",
	PortfolioSortByTenDayAvgVol:        "TEN_DAY_AVG_VOL",
	PortfolioSortByBeta:                "BETA",
	PortfolioSortByBidAskSpread:        "BID_ASK_SPREAD",
	PortfolioSortByMarginable:          "MARGINABLE",
	PortfolioSortByDelta52wkHi:         "DELTA_52WK_HI",
	PortfolioSortByDelta52WkLow:        "DELTA_52WK_LOW",
	PortfolioSortByPerf1Mon:            "PERF_1MON",
	PortfolioSortByAnnualDiv:           "ANNUAL_DIV",
	PortfolioSortByPerf12Mon:           "PERF_12MON",
	PortfolioSortByPerf3Mon:            "PERF_3MON",
	PortfolioSortByPerf6Mon:            "PERF_6MON",
	PortfolioSortByPreDayVol:           "PRE_DAY_VOL",
	PortfolioSortBySv1MonAvg:           "SV_1MON_AVG",
	PortfolioSortBySv10DayAvg:          "SV_10DAY_AVG",
	PortfolioSortBySv20DayAvg:          "SV_20DAY_AVG",
	PortfolioSortBySv2MonAvg:           "SV_2MON_AVG",
	PortfolioSortBySv3MonAvg:           "SV_3MON_AVG",
	PortfolioSortBySv4MonAvg:           "SV_4MON_AVG",
	PortfolioSortBySv6MonAvg:           "SV_6MON_AVG",
	PortfolioSortByDelta:               "DELTA",
	PortfolioSortByGamma:               "GAMMA",
	PortfolioSortByIvPct:               "IV_PCT",
	PortfolioSortByTheta:               "THETA",
	PortfolioSortByVega:                "VEGA",
	PortfolioSortByAdjNonadjFlag:       "ADJ_NONADJ_FLAG",
	PortfolioSortByDaysExpiration:      "DAYS_EXPIRATION",
	PortfolioSortByOpenInterest:        "OPEN_INTEREST",
	PortfolioSortByIntrinsicValue:      "INSTRINIC_VALUE", //TODO: Confirm whether this is supposed to be "INTRINSIC" (ETrade documentation says "INSTRINIC")
	PortfolioSortByRho:                 "RHO",
	PortfolioSortByTypeCode:            "TYPE_CODE",
	PortfolioSortByDisplaySymbol:       "DISPLAY_SYMBOL",
	PortfolioSortByAfterHoursPctChange: "AFTER_HOURS_PCTCHANGE",
	PortfolioSortByPreMarketPctChange:  "PRE_MARKET_PCTCHANGE",
	PortfolioSortByExpandCollapseFlag:  "EXPAND_COLLAPSE_FLAG",
}

// String converts a PortfolioSortBy to its string representation.
func (e *PortfolioSortBy) String() string {
	if s, found := portfolioSortByToString[*e]; found {
		return s
	}
	return "UNKNOWN"
}

// PortfolioSortByFromString returns the PortfolioSortBy for the specified
// string, or an error if the string doesn't represent a valid value.
func PortfolioSortByFromString(s string) (PortfolioSortBy, error) {
	return getKeyForValue(portfolioSortByToString, s)
}

var portfolioViewToString = map[PortfolioView]string{
	PortfolioViewPerformance:  "PERFORMANCE",
	PortfolioViewFundamental:  "FUNDAMENTAL",
	PortfolioViewOptionsWatch: "OPTIONSWATCH",
	PortfolioViewQuick:        "QUICK",
	PortfolioViewComplete:     "COMPLETE",
}

// String converts a PortfolioView to its string representation.
func (e *PortfolioView) String() string {
	if s, found := portfolioViewToString[*e]; found {
		return s
	}
	return "UNKNOWN"
}

// PortfolioViewFromString returns the PortfolioView for the specified string,
// or an error if the string doesn't represent a valid value.
func PortfolioViewFromString(s string) (PortfolioView, error) {
	return getKeyForValue(portfolioViewToString, s)
}
