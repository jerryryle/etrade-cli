package constants

// TransactionsMaxCount is the maximum count that can be included in a List
// Transactions request. Note that ETrade does not document this value, so I
// determined it empirically by increasing the count until I got a
// 500 Internal Server Error. The resulting value was 50, which is the
// documented default number of transactions returned if no count is specified.
const TransactionsMaxCount = 50
