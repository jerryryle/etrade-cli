package constants

type SortOrder int

const (
	// SortOrderNil indicates no sort order
	// (e.g. to make a query use the default value from ETrade)
	SortOrderNil SortOrder = iota
	SortOrderAsc
	SortOrderDesc
)

// MarketSession specifies the market session for which to get results
// See the constants below for semantics.
type MarketSession int

const (
	// MarketSessionNil indicates no market session
	// (e.g. to make a query use the default value from ETrade)
	MarketSessionNil MarketSession = iota

	// MarketSessionRegular gets results for the regular market session
	MarketSessionRegular

	// MarketSessionExtended gets results for the extended market session
	MarketSessionExtended
)

var sortOrderToString = map[SortOrder]string{
	SortOrderAsc:  "ASC",
	SortOrderDesc: "DESC",
}

// String converts a SortOrder to its string representation.
func (e SortOrder) String() string {
	if s, found := sortOrderToString[e]; found {
		return s
	}
	return "UNKNOWN"
}

// SortOrderFromString returns the SortOrder for the specified string,
// or an error if the string doesn't represent a valid value.
func SortOrderFromString(s string) (SortOrder, error) {
	return getKeyForValue(sortOrderToString, s)
}

var marketSessionToString = map[MarketSession]string{
	MarketSessionRegular:  "REGULAR",
	MarketSessionExtended: "EXTENDED",
}

// String converts a MarketSession to its string representation.
func (e *MarketSession) String() string {
	if s, found := marketSessionToString[*e]; found {
		return s
	}
	return "UNKNOWN"
}

// MarketSessionFromString returns the MarketSession for the
// specified string, or an error if the string doesn't represent a valid value.
func MarketSessionFromString(s string) (MarketSession, error) {
	return getKeyForValue(marketSessionToString, s)
}
