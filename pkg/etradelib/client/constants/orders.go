package constants

// ListOrdersMaxSymbols is the maximum number of symbols that can be included
// in a List Orders request
const ListOrdersMaxSymbols = 25

// OrderStatus specifies the status of orders to retrieve.
// See the constants below for semantics.
type OrderStatus int

const (
	// OrderStatusNil indicates no order status
	// (e.g. to make a query use the default value from ETrade)
	OrderStatusNil OrderStatus = iota

	// OrderStatusOpen gets orders that are open
	OrderStatusOpen

	// OrderStatusExecuted gets orders that have been executed
	OrderStatusExecuted

	// OrderStatusCanceled gets orders that have been canceled
	OrderStatusCanceled

	// OrderStatusIndividualFills gets orders with individual fills
	OrderStatusIndividualFills

	// OrderStatusCancelRequested gets orders with pending cancel requests
	OrderStatusCancelRequested

	// OrderStatusExpired gets orders that have expired
	OrderStatusExpired

	// OrderStatusRejected gets orders that have been rejected
	OrderStatusRejected
)

// OrderSecurityType specifies the security type of orders to retrieve.
// See the constants below for semantics.
type OrderSecurityType int

const (
	// OrderSecurityTypeNil indicates no order security type
	// (e.g. to make a query use the default value from ETrade)
	OrderSecurityTypeNil OrderSecurityType = iota

	// OrderSecurityTypeEquity gets equity orders
	OrderSecurityTypeEquity

	// OrderSecurityTypeOption gets option orders
	OrderSecurityTypeOption

	// OrderSecurityTypeMutualFund gets mutual fund orders
	OrderSecurityTypeMutualFund

	// OrderSecurityTypeMoneyMarketFund gets money market mutual fund orders
	OrderSecurityTypeMoneyMarketFund
)

// OrderTransactionType specifies the transaction type of orders to retrieve.
// See the constants below for semantics.
type OrderTransactionType int

const (
	// OrderTransactionTypeNil indicates no order transaction type
	// (e.g. to make a query use the default value from ETrade)
	OrderTransactionTypeNil OrderTransactionType = iota

	// OrderTransactionTypeExtendedHours gets extended hours orders
	OrderTransactionTypeExtendedHours

	// OrderTransactionTypeBuy gets buy orders
	OrderTransactionTypeBuy

	// OrderTransactionTypeSell gets sell orders
	OrderTransactionTypeSell

	// OrderTransactionTypeShort gets short orders
	OrderTransactionTypeShort

	// OrderTransactionTypeBuyToCover gets buy to cover orders
	OrderTransactionTypeBuyToCover

	// OrderTransactionTypeMutualFundExchange gets mutual fund exchange orders
	OrderTransactionTypeMutualFundExchange
)

var orderStatusToString = map[OrderStatus]string{
	OrderStatusOpen:            "OPEN",
	OrderStatusExecuted:        "EXECUTED",
	OrderStatusCanceled:        "CANCELLED",
	OrderStatusIndividualFills: "INDIVIDUAL_FILLS",
	OrderStatusCancelRequested: "CANCEL_REQUESTED",
	OrderStatusExpired:         "EXPIRED",
	OrderStatusRejected:        "REJECTED",
}

// String converts an OrderStatus to its string representation.
func (e OrderStatus) String() string {
	if s, found := orderStatusToString[e]; found {
		return s
	}
	return "UNKNOWN"
}

var orderSecurityTypeToString = map[OrderSecurityType]string{
	OrderSecurityTypeEquity:          "EQ",
	OrderSecurityTypeOption:          "OPTN",
	OrderSecurityTypeMutualFund:      "MF",
	OrderSecurityTypeMoneyMarketFund: "MMF",
}

// String converts an OrderSecurityType to its string representation.
func (e OrderSecurityType) String() string {
	if s, found := orderSecurityTypeToString[e]; found {
		return s
	}
	return "UNKNOWN"
}

var orderTransactionTypeToString = map[OrderTransactionType]string{
	OrderTransactionTypeExtendedHours:      "ATNM",
	OrderTransactionTypeBuy:                "BUY",
	OrderTransactionTypeSell:               "SELL",
	OrderTransactionTypeShort:              "SELL_SHORT",
	OrderTransactionTypeBuyToCover:         "BUY_TO_COVER",
	OrderTransactionTypeMutualFundExchange: "MF_EXCHANGE",
}

// String converts an OrderTransactionType to its string representation.
func (e OrderTransactionType) String() string {
	if s, found := orderTransactionTypeToString[e]; found {
		return s
	}
	return "UNKNOWN"
}
