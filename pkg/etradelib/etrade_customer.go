package etradelib

type ETradeCustomer interface {
	GetCustomerName() string
	GetAllAccounts() ([]ETradeAccount, error)
	GetAccountById(accountID string) (ETradeAccount, error)
	GetAllAlerts() ([]ETradeAlert, error)
	GetAlertById(alertID int64) (ETradeAlert, error)
	GetQuotesAll(symbols []string) ([]ETradeQuoteAllInfo, error)
	GetQuotesFundamental(symbols []string) ([]ETradeQuoteFundamentalInfo, error)
	GetQuotesIntraday(symbols []string) ([]ETradeQuoteIntradayInfo, error)
	GetQuotesOptions(symbols []string) ([]ETradeQuoteOptionsInfo, error)
	GetQuotesWeek52(symbols []string) ([]ETradeQuoteWeek52Info, error)
	GetQuotesMutualFund(symbols []string) ([]ETradeQuoteMutualFundInfo, error)
	LookUpProduct(search string) (string, error)
	GetOptionChains() (string, error)
	GetOptionExpireDates() (string, error)
}
