package constants

type SortOrder int

const (
	// SortOrderNil indicates no sort order
	// (e.g. to make a query use the default value from ETrade)
	SortOrderNil SortOrder = iota
	SortOrderAsc
	SortOrderDesc
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
