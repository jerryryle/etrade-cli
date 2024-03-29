package constants

// OptionCategory specifies the category of options to retrieve.
// See the constants below for semantics.
type OptionCategory int

const (
	// OptionCategoryNil indicates no option category
	// (e.g. to make a query use the default value from ETrade)
	OptionCategoryNil OptionCategory = iota

	// OptionCategoryStandard gets standard options
	OptionCategoryStandard

	// OptionCategoryAll gets all options
	OptionCategoryAll

	// OptionCategoryMini gets mini options
	OptionCategoryMini
)

// OptionChainType specifies the type of option chain to retrieve.
// See the constants below for semantics.
type OptionChainType int

const (
	// OptionChainTypeNil indicates no chain type
	// (e.g. to make a query use the default value from ETrade)
	OptionChainTypeNil OptionChainType = iota

	// OptionChainTypeCall gets calls
	OptionChainTypeCall

	// OptionChainTypePut gets puts
	OptionChainTypePut

	// OptionChainTypeCallPut gets calls and puts
	OptionChainTypeCallPut
)

// OptionPriceType specifies the option price type to retrieve.
// See the constants below for semantics.
type OptionPriceType int

const (
	// OptionPriceTypeNil indicates no price type
	// (e.g. to make a query use the default value from ETrade)
	OptionPriceTypeNil OptionPriceType = iota

	// OptionPriceTypeExtendedHours gets extended hours prices
	OptionPriceTypeExtendedHours

	// OptionPriceTypeAll gets all prices
	OptionPriceTypeAll
)

// OptionExpiryType specifies the expiration type to retrieve.
// See the constants below for semantics.
type OptionExpiryType int

const (
	// OptionExpiryTypeNil indicates no expiry type
	// (e.g. to make a query use the default value from ETrade)
	OptionExpiryTypeNil OptionExpiryType = iota

	// OptionExpiryTypeUnspecified gets options with unspecified expiry type
	OptionExpiryTypeUnspecified

	// OptionExpiryTypeDaily gets options with daily expiry type
	OptionExpiryTypeDaily

	// OptionExpiryTypeWeekly gets options with weekly expiry type
	OptionExpiryTypeWeekly

	// OptionExpiryTypeMonthly gets options with monthly expiry type
	OptionExpiryTypeMonthly

	// OptionExpiryTypeQuarterly gets options with quarterly expiry type
	OptionExpiryTypeQuarterly

	// OptionExpiryTypeVix gets options with VIX expiry type
	OptionExpiryTypeVix

	// OptionExpiryTypeAll gets options with all expiry types
	OptionExpiryTypeAll

	// OptionExpiryTypeMonthEnd gets options with month end expiry type
	OptionExpiryTypeMonthEnd
)

var optionCategoryToString = map[OptionCategory]string{
	OptionCategoryStandard: "STANDARD",
	OptionCategoryAll:      "ALL",
	OptionCategoryMini:     "MINI",
}

// String converts an OptionCategory to its string representation.
func (e OptionCategory) String() string {
	if s, found := optionCategoryToString[e]; found {
		return s
	}
	return "UNKNOWN"
}

var chainTypeToString = map[OptionChainType]string{
	OptionChainTypeCall:    "CALL",
	OptionChainTypePut:     "PUT",
	OptionChainTypeCallPut: "CALLPUT",
}

// String converts a OptionChainType to its string representation.
func (e OptionChainType) String() string {
	if s, found := chainTypeToString[e]; found {
		return s
	}
	return "UNKNOWN"
}

var priceTypeToString = map[OptionPriceType]string{
	OptionPriceTypeExtendedHours: "ATNM",
	OptionPriceTypeAll:           "ALL",
}

// String converts a OptionPriceType to its string representation.
func (e OptionPriceType) String() string {
	if s, found := priceTypeToString[e]; found {
		return s
	}
	return "UNKNOWN"
}

var expiryTypeToString = map[OptionExpiryType]string{
	OptionExpiryTypeUnspecified: "UNSPECIFIED",
	OptionExpiryTypeDaily:       "DAILY",
	OptionExpiryTypeWeekly:      "WEEKLY",
	OptionExpiryTypeMonthly:     "MONTHLY",
	OptionExpiryTypeQuarterly:   "QUARTERLY",
	OptionExpiryTypeVix:         "VIX",
	OptionExpiryTypeAll:         "ALL",
	OptionExpiryTypeMonthEnd:    "MONTHEND",
}

// String converts a OptionExpiryType to its string representation.
func (e OptionExpiryType) String() string {
	if s, found := expiryTypeToString[e]; found {
		return s
	}
	return "UNKNOWN"
}
