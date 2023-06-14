package client

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
