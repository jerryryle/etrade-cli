package etradelib

type ETradeCustomer interface {
	GetCustomerName() string
	ListAccounts() string
	ListAlerts() string
	GetQuotes(symbols string) string
	LookUpProduct(search string) string
	GetOptionChains() string
	GetOptionExpireDates() string
}

type eTradeCustomer struct {
	customerName string
	client       eTradeClient
}

func (c *eTradeCustomer) GetCustomerName() string {
	return c.customerName
}

func (c *eTradeCustomer) ListAccounts() string {
	return ""
}

func (c *eTradeCustomer) ListAlerts() string {
	return ""
}

func (c *eTradeCustomer) GetQuotes(symbols string) string {
	return ""
}

func (c *eTradeCustomer) LookUpProduct(search string) string {
	return ""
}

func (c *eTradeCustomer) GetOptionChains() string {
	return ""
}

func (c *eTradeCustomer) GetOptionExpireDates() string {
	return ""
}
