package constants

// GetQuotesMaxSymbolsBeforeOverride is the maximum number of symbols that can
// be included in a GetQuotes request without specifying the
// `overrideSymbolCount` flag.
const GetQuotesMaxSymbolsBeforeOverride = 25

// GetQuotesMaxSymbols is the maximum number of symbols that can be included in
// a GetQuotes request with the `overrideSymbolCount` flag set.
const GetQuotesMaxSymbols = 50

// QuoteDetailFlag specifies the quote detail to retrieve for quote queries.
// See the constants below for semantics.
type QuoteDetailFlag int

const (
	// QuoteDetailNil indicates no quote detail
	// (e.g. to make a query use the default value from ETrade)
	QuoteDetailNil QuoteDetailFlag = iota

	// QuoteDetailAll gets all available quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/AllQuoteDetails
	QuoteDetailAll

	// QuoteDetailFundamental gets fundamental quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/FundamentalQuoteDetails
	QuoteDetailFundamental

	// QuoteDetailIntraday gets intraday quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/IntradayQuoteDetails
	QuoteDetailIntraday

	// QuoteDetailOptions gets options quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/OptionQuoteDetails
	QuoteDetailOptions

	// QuoteDetailWeek52 gets 52-week quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/Week52QuoteDetails
	QuoteDetailWeek52

	// QuoteDetailMutualFund gets mutual fund quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/MutualFund
	QuoteDetailMutualFund
)

var quoteDetailFlagToString = map[QuoteDetailFlag]string{
	QuoteDetailAll:         "ALL",
	QuoteDetailFundamental: "FUNDAMENTAL",
	QuoteDetailIntraday:    "INTRADAY",
	QuoteDetailOptions:     "OPTIONS",
	QuoteDetailWeek52:      "WEEK_52",
	QuoteDetailMutualFund:  "MF_DETAIL",
}

// String converts a QuoteDetailFlag to its string representation.
func (e QuoteDetailFlag) String() string {
	if s, found := quoteDetailFlagToString[e]; found {
		return s
	}
	return "UNKNOWN"
}
