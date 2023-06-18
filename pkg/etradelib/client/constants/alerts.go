package constants

// AlertCategory specifies the category of an alert for alert queries.
// See the constants below for semantics.
type AlertCategory int

const (
	// AlertCategoryNil indicates no alert status
	// (e.g. to make a query use the default value from ETrade)
	AlertCategoryNil AlertCategory = iota

	// AlertCategoryStock indicates an alert related to a stock
	AlertCategoryStock

	// AlertCategoryAccount indicates an alert related to an account
	AlertCategoryAccount
)

// AlertStatus specifies the status of an alert for alert queries.
// See the constants below for semantics.
type AlertStatus int

const (
	// AlertStatusNil indicates no alert status
	// (e.g. to make a query use the default value from ETrade)
	AlertStatusNil AlertStatus = iota

	// AlertStatusRead indicates an alert that a customer has read
	AlertStatusRead

	// AlertStatusUnread indicates an alert that a customer has not yet read
	AlertStatusUnread

	// AlertStatusDeleted indicates an alert that a customer has deleted
	AlertStatusDeleted
)

var alertCategoryToString = map[AlertCategory]string{
	AlertCategoryStock:   "STOCK",
	AlertCategoryAccount: "ACCOUNT",
}

// String converts an AlertCategory to its string representation.
func (e AlertCategory) String() string {
	if s, found := alertCategoryToString[e]; found {
		return s
	}
	return "UNKNOWN"
}

// AlertCategoryFromString returns the AlertCategory for the specified string,
// or an error if the string doesn't represent a valid value.
func AlertCategoryFromString(s string) (AlertCategory, error) {
	return getKeyForValue(alertCategoryToString, s)
}

var alertStatusToString = map[AlertStatus]string{
	AlertStatusRead:    "READ",
	AlertStatusUnread:  "UNREAD",
	AlertStatusDeleted: "DELETED",
}

// String converts an AlertStatus to its string representation.
func (e AlertStatus) String() string {
	if s, found := alertStatusToString[e]; found {
		return s
	}
	return "UNKNOWN"
}

// AlertStatusFromString returns the AlertStatus for the specified string,
// or an error if the string doesn't represent a valid value.
func AlertStatusFromString(s string) (AlertStatus, error) {
	return getKeyForValue(alertStatusToString, s)
}
