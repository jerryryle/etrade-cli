package etradelib

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
