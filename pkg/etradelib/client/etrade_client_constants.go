package client

import (
	"errors"
)

// GetQuotesMaxSymbolsBeforeOverride is the maximum number of symbols that can be included in a GetQuotes request without
// specifying the `overrideSymbolCount` flag.
const GetQuotesMaxSymbolsBeforeOverride = 25

// GetQuotesMaxSymbols is the maximum number of symbols that can be included in a GetQuotes request with the
// `overrideSymbolCount` flag set.
const GetQuotesMaxSymbols = 50

type QuoteDetailFlag int

const (
	QuoteDetailAll QuoteDetailFlag = iota
	QuoteDetailFundamental
	QuoteDetailIntraday
	QuoteDetailOptions
	QuoteDetailWeek52
	QuoteDetailMutualFund
)

func (e QuoteDetailFlag) String() string {
	switch e {
	case QuoteDetailAll:
		return "ALL"
	case QuoteDetailFundamental:
		return "FUNDAMENTAL"
	case QuoteDetailIntraday:
		return "INTRADAY"
	case QuoteDetailOptions:
		return "OPTIONS"
	case QuoteDetailWeek52:
		return "WEEK_52"
	case QuoteDetailMutualFund:
		return "MF_DETAIL"
	}
	return "UNKNOWN"
}

type OptionCategory int

const (
	OptionCategoryStandard OptionCategory = iota
	OptionCategoryAll
	OptionCategoryMini
)

func (e OptionCategory) String() string {
	switch e {
	case OptionCategoryStandard:
		return "STANDARD"
	case OptionCategoryAll:
		return "ALL"
	case OptionCategoryMini:
		return "MINI"
	}
	return "UNKNOWN"
}

type ChainType int

const (
	ChainTypeCall ChainType = iota
	ChainTypePut
	ChainTypeCallPut
)

func (e ChainType) String() string {
	switch e {
	case ChainTypeCall:
		return "CALL"
	case ChainTypePut:
		return "PUT"
	case ChainTypeCallPut:
		return "CALLPUT"
	}
	return "UNKNOWN"
}

type PriceType int

const (
	PriceTypeAtnm PriceType = iota
	PriceTypeAll
)

func (e PriceType) String() string {
	switch e {
	case PriceTypeAtnm:
		return "ATNM"
	case PriceTypeAll:
		return "ALL"
	}
	return "UNKNOWN"
}

type ExpiryType int

const (
	ExpiryTypeUnspecified ExpiryType = iota
	ExpiryTypeDaily
	ExpiryTypeWeekly
	ExpiryTypeMonthly
	ExpiryTypeQuarterly
	ExpiryTypeVix
	ExpiryTypeAll
	ExpiryTypeMonthEnd
)

func (e ExpiryType) String() string {
	switch e {
	case ExpiryTypeUnspecified:
		return "UNSPECIFIED"
	case ExpiryTypeDaily:
		return "DAILY"
	case ExpiryTypeWeekly:
		return "WEEKLY"
	case ExpiryTypeMonthly:
		return "MONTHLY"
	case ExpiryTypeQuarterly:
		return "QUARTERLY"
	case ExpiryTypeVix:
		return "VIX"
	case ExpiryTypeAll:
		return "ALL"
	case ExpiryTypeMonthEnd:
		return "MONTHEND"
	}
	return "UNKNOWN"
}

type SortOrder int

const (
	SortOrderAsc SortOrder = iota
	SortOrderDesc
)

func (e SortOrder) String() string {
	switch e {
	case SortOrderAsc:
		return "ASC"
	case SortOrderDesc:
		return "DESC"
	}
	return "UNKNOWN"
}

type PortfolioSortBy int

const (
	PortfolioSortBySymbol PortfolioSortBy = iota
	PortfolioSortByTypeName
	PortfolioSortByExchangeName
	PortfolioSortByCurrency
	PortfolioSortByQuantity
	PortfolioSortByLongOrShort
	PortfolioSortByDateAcquired
	PortfolioSortByPricePaid
	PortfolioSortByTotalGain
	PortfolioSortByTotalGainPct
	PortfolioSortByMarktValue
	PortfolioSortByBi
	PortfolioSortByAsk
	PortfolioSortByPriceChange
	PortfolioSortByPriceChangePct
	PortfolioSortByVolume
	PortfolioSortByWeek52High
	PortfolioSortByWeek52Low
	PortfolioSortByEps
	PortfolioSortByPeRatio
	PortfolioSortByOptionType
	PortfolioSortByStrikePrice
	PortfolioSortByPremium
	PortfolioSortByExpiration
	PortfolioSortByDaysGain
	PortfolioSortByCommission
	PortfolioSortByMarketCap
	PortfolioSortByPrevClose
	PortfolioSortByOpen
	PortfolioSortByDaysRange
	PortfolioSortByTotalCost
	PortfolioSortByDaysGainPct
	PortfolioSortByPctOfPortfolio
	PortfolioSortByLastTradeTime
	PortfolioSortByBaseSymbolPrice
	PortfolioSortByWeek52Range
	PortfolioSortByLastTrade
	PortfolioSortBySymbolDesc
	PortfolioSortByBidSize
	PortfolioSortByAskSize
	PortfolioSortByOtherFees
	PortfolioSortByHeldAs
	PortfolioSortByOptionMultiplier
	PortfolioSortByDeliverables
	PortfolioSortByCostPerShare
	PortfolioSortByDividend
	PortfolioSortByDivYield
	PortfolioSortByDivPayDate
	PortfolioSortByEstEarn
	PortfolioSortByExDivDate
	PortfolioSortByTenDayAvgVol
	PortfolioSortByBeta
	PortfolioSortByBidAskSpread
	PortfolioSortByMarginable
	PortfolioSortByDelta52wkHi
	PortfolioSortByDelta52WkLow
	PortfolioSortByPerf1Mon
	PortfolioSortByAnnualDiv
	PortfolioSortByPerf12Mon
	PortfolioSortByPerf3Mon
	PortfolioSortByPerf6Mon
	PortfolioSortByPreDayVol
	PortfolioSortBySv1MonAvg
	PortfolioSortBySv10DayAvg
	PortfolioSortBySv20DayAvg
	PortfolioSortBySv2MonAvg
	PortfolioSortBySv3MonAvg
	PortfolioSortBySv4MonAvg
	PortfolioSortBySv6MonAvg
	PortfolioSortByDelta
	PortfolioSortByGamma
	PortfolioSortByIvPct
	PortfolioSortByTheta
	PortfolioSortByVega
	PortfolioSortByAdjNonadjFlag
	PortfolioSortByDaysExpiration
	PortfolioSortByOpenInterest
	PortfolioSortByInstrinicValue
	PortfolioSortByRho
	PortfolioSortByTypeCode
	PortfolioSortByDisplaySymbol
	PortfolioSortByAfterHoursPctChange
	PortfolioSortByPreMarketPctChange
	PortfolioSortByExpandCollapseFlag
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
	PortfolioSortByMarktValue:          "MARKET_VALUE",
	PortfolioSortByBi:                  "BI",
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
	PortfolioSortByInstrinicValue:      "INSTRINIC_VALUE",
	PortfolioSortByRho:                 "RHO",
	PortfolioSortByTypeCode:            "TYPE_CODE",
	PortfolioSortByDisplaySymbol:       "DISPLAY_SYMBOL",
	PortfolioSortByAfterHoursPctChange: "AFTER_HOURS_PCTCHANGE",
	PortfolioSortByPreMarketPctChange:  "PRE_MARKET_PCTCHANGE",
	PortfolioSortByExpandCollapseFlag:  "EXPAND_COLLAPSE_FLAG",
}

var stringToPortfolioSortBy = map[string]PortfolioSortBy{
	"SYMBOL":                PortfolioSortBySymbol,
	"TYPE_NAME":             PortfolioSortByTypeName,
	"EXCHANGE_NAME":         PortfolioSortByExchangeName,
	"CURRENCY":              PortfolioSortByCurrency,
	"QUANTITY":              PortfolioSortByQuantity,
	"LONG_OR_SHORT":         PortfolioSortByLongOrShort,
	"DATE_ACQUIRED":         PortfolioSortByDateAcquired,
	"PRICEPAID":             PortfolioSortByPricePaid,
	"TOTAL_GAIN":            PortfolioSortByTotalGain,
	"TOTAL_GAIN_PCT":        PortfolioSortByTotalGainPct,
	"MARKET_VALUE":          PortfolioSortByMarktValue,
	"BI":                    PortfolioSortByBi,
	"ASK":                   PortfolioSortByAsk,
	"PRICE_CHANGE":          PortfolioSortByPriceChange,
	"PRICE_CHANGE_PCT":      PortfolioSortByPriceChangePct,
	"VOLUME":                PortfolioSortByVolume,
	"WEEK_52_HIGH":          PortfolioSortByWeek52High,
	"WEEK_52_LOW":           PortfolioSortByWeek52Low,
	"EPS":                   PortfolioSortByEps,
	"PE_RATIO":              PortfolioSortByPeRatio,
	"OPTION_TYPE":           PortfolioSortByOptionType,
	"STRIKE_PRICE":          PortfolioSortByStrikePrice,
	"PREMIUM":               PortfolioSortByPremium,
	"EXPIRATION":            PortfolioSortByExpiration,
	"DAYS_GAIN":             PortfolioSortByDaysGain,
	"COMMISSION":            PortfolioSortByCommission,
	"MARKETCAP":             PortfolioSortByMarketCap,
	"PREV_CLOSE":            PortfolioSortByPrevClose,
	"OPEN":                  PortfolioSortByOpen,
	"DAYS_RANGE":            PortfolioSortByDaysRange,
	"TOTAL_COST":            PortfolioSortByTotalCost,
	"DAYS_GAIN_PCT":         PortfolioSortByDaysGainPct,
	"PCT_OF_PORTFOLIO":      PortfolioSortByPctOfPortfolio,
	"LAST_TRADE_TIME":       PortfolioSortByLastTradeTime,
	"BASE_SYMBOL_PRICE":     PortfolioSortByBaseSymbolPrice,
	"WEEK_52_RANGE":         PortfolioSortByWeek52Range,
	"LAST_TRADE":            PortfolioSortByLastTrade,
	"SYMBOL_DESC":           PortfolioSortBySymbolDesc,
	"BID_SIZE":              PortfolioSortByBidSize,
	"ASK_SIZE":              PortfolioSortByAskSize,
	"OTHER_FEES":            PortfolioSortByOtherFees,
	"HELD_AS":               PortfolioSortByHeldAs,
	"OPTION_MULTIPLIER":     PortfolioSortByOptionMultiplier,
	"DELIVERABLES":          PortfolioSortByDeliverables,
	"COST_PERSHARE":         PortfolioSortByCostPerShare,
	"DIVIDEND":              PortfolioSortByDividend,
	"DIV_YIELD":             PortfolioSortByDivYield,
	"DIV_PAY_DATE":          PortfolioSortByDivPayDate,
	"EST_EARN":              PortfolioSortByEstEarn,
	"EX_DIV_DATE":           PortfolioSortByExDivDate,
	"TEN_DAY_AVG_VOL":       PortfolioSortByTenDayAvgVol,
	"BETA":                  PortfolioSortByBeta,
	"BID_ASK_SPREAD":        PortfolioSortByBidAskSpread,
	"MARGINABLE":            PortfolioSortByMarginable,
	"DELTA_52WK_HI":         PortfolioSortByDelta52wkHi,
	"DELTA_52WK_LOW":        PortfolioSortByDelta52WkLow,
	"PERF_1MON":             PortfolioSortByPerf1Mon,
	"ANNUAL_DIV":            PortfolioSortByAnnualDiv,
	"PERF_12MON":            PortfolioSortByPerf12Mon,
	"PERF_3MON":             PortfolioSortByPerf3Mon,
	"PERF_6MON":             PortfolioSortByPerf6Mon,
	"PRE_DAY_VOL":           PortfolioSortByPreDayVol,
	"SV_1MON_AVG":           PortfolioSortBySv1MonAvg,
	"SV_10DAY_AVG":          PortfolioSortBySv10DayAvg,
	"SV_20DAY_AVG":          PortfolioSortBySv20DayAvg,
	"SV_2MON_AVG":           PortfolioSortBySv2MonAvg,
	"SV_3MON_AVG":           PortfolioSortBySv3MonAvg,
	"SV_4MON_AVG":           PortfolioSortBySv4MonAvg,
	"SV_6MON_AVG":           PortfolioSortBySv6MonAvg,
	"DELTA":                 PortfolioSortByDelta,
	"GAMMA":                 PortfolioSortByGamma,
	"IV_PCT":                PortfolioSortByIvPct,
	"THETA":                 PortfolioSortByTheta,
	"VEGA":                  PortfolioSortByVega,
	"ADJ_NONADJ_FLAG":       PortfolioSortByAdjNonadjFlag,
	"DAYS_EXPIRATION":       PortfolioSortByDaysExpiration,
	"OPEN_INTEREST":         PortfolioSortByOpenInterest,
	"INSTRINIC_VALUE":       PortfolioSortByInstrinicValue,
	"RHO":                   PortfolioSortByRho,
	"TYPE_CODE":             PortfolioSortByTypeCode,
	"DISPLAY_SYMBOL":        PortfolioSortByDisplaySymbol,
	"AFTER_HOURS_PCTCHANGE": PortfolioSortByAfterHoursPctChange,
	"PRE_MARKET_PCTCHANGE":  PortfolioSortByPreMarketPctChange,
	"EXPAND_COLLAPSE_FLAG":  PortfolioSortByExpandCollapseFlag,
}

func (e *PortfolioSortBy) String() string {
	if s, found := portfolioSortByToString[*e]; found {
		return s
	}
	return "UNKNOWN"
}

func (e *PortfolioSortBy) FromString(s string) error {
	if newE, found := stringToPortfolioSortBy[s]; found {
		*e = newE
		return nil
	}
	return errors.New("invalid value")
}

func PortfolioSortByFromString(s string) (PortfolioSortBy, error) {
	if e, found := stringToPortfolioSortBy[s]; found {
		return e, nil
	}
	return PortfolioSortBy(0), errors.New("invalid value")
}

type PortfolioMarketSession int

const (
	PortfolioMarketSessionRegular PortfolioMarketSession = iota
	PortfolioMarketSessionExtended
)

var portfolioMarketSessionToString = map[PortfolioMarketSession]string{
	PortfolioMarketSessionRegular:  "REGULAR",
	PortfolioMarketSessionExtended: "EXTENDED",
}

func (e *PortfolioMarketSession) String() string {
	if s, found := portfolioMarketSessionToString[*e]; found {
		return s
	}
	return "UNKNOWN"
}

type PortfolioView int

const (
	PortfolioViewPerformance PortfolioView = iota
	PortfolioViewFundamental
	PortfolioViewOptionsWatch
	PortfolioViewQuick
	PortfolioViewComplete
)

var portfolioViewToString = map[PortfolioView]string{
	PortfolioViewPerformance:  "PERFORMANCE",
	PortfolioViewFundamental:  "FUNDAMENTAL",
	PortfolioViewOptionsWatch: "OPTIONSWATCH",
	PortfolioViewQuick:        "QUICK",
	PortfolioViewComplete:     "COMPLETE",
}

func (e *PortfolioView) String() string {
	if s, found := portfolioViewToString[*e]; found {
		return s
	}
	return "UNKNOWN"
}
