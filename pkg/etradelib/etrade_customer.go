package etradelib

type ETradeCustomer interface {
	ListAccounts() string
	ListAlerts() string
	GetQuotes(symbols string) string
	LookUpProduct(search string) string
	GetOptionChains() string
	GetOptionExpireDates() string
}

type eTradeCustomer struct {
	client eTradeClient
}

func (e *eTradeCustomer) ListAccounts() string {
	return ""
}

func (e *eTradeCustomer) ListAlerts() string {
	return ""
}

func (e *eTradeCustomer) GetQuotes(symbols string) string {
	return ""
}

func (e *eTradeCustomer) LookUpProduct(search string) string {
	return ""
}

func (e *eTradeCustomer) GetOptionChains() string {
	return ""
}

func (e *eTradeCustomer) GetOptionExpireDates() string {
	return ""
}
