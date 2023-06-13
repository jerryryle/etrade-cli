package client

// GetQuotesMaxSymbols is the maximum number of symbols that can be included in a GetQuotes request
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

func (f QuoteDetailFlag) String() string {
	switch f {
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

func (f OptionCategory) String() string {
	switch f {
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

func (f ChainType) String() string {
	switch f {
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

func (f PriceType) String() string {
	switch f {
	case PriceTypeAtnm:
		return "ATNM"
	case PriceTypeAll:
		return "ALL"
	}
	return "UNKNOWN"
}
