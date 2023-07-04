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
	// QuoteDetailFlagNil indicates no quote detail
	// (e.g. to make a query use the default value from ETrade)
	QuoteDetailFlagNil QuoteDetailFlag = iota

	// QuoteDetailFlagAll gets all available quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/AllQuoteDetails
	QuoteDetailFlagAll

	// QuoteDetailFlagFundamental gets fundamental quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/FundamentalQuoteDetails
	QuoteDetailFlagFundamental

	// QuoteDetailFlagIntraday gets intraday quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/IntradayQuoteDetails
	QuoteDetailFlagIntraday

	// QuoteDetailFlagOptions gets options quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/OptionQuoteDetails
	QuoteDetailFlagOptions

	// QuoteDetailFlagWeek52 gets 52-week quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/Week52QuoteDetails
	QuoteDetailFlagWeek52

	// QuoteDetailFlagMutualFund gets mutual fund quote detail
	// See https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/MutualFund
	QuoteDetailFlagMutualFund
)

var quoteDetailFlagToString = map[QuoteDetailFlag]string{
	QuoteDetailFlagAll:         "ALL",
	QuoteDetailFlagFundamental: "FUNDAMENTAL",
	QuoteDetailFlagIntraday:    "INTRADAY",
	QuoteDetailFlagOptions:     "OPTIONS",
	QuoteDetailFlagWeek52:      "WEEK_52",
	QuoteDetailFlagMutualFund:  "MF_DETAIL",
}

// String converts a QuoteDetailFlag to its string representation.
func (e QuoteDetailFlag) String() string {
	if s, found := quoteDetailFlagToString[e]; found {
		return s
	}
	return "UNKNOWN"
}
