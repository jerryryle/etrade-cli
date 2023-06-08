package etradelib

import "fmt"

type ETradeCustomer interface {
	GetCustomerName() string
	ListAccounts() (string, error)
	ListAlerts() (string, error)
	GetQuotes(symbols string) (string, error)
	LookUpProduct(search string) (string, error)
	GetOptionChains() (string, error)
	GetOptionExpireDates() (string, error)
}

type eTradeCustomer struct {
	customerName string
	client       ETradeClient
}

func (c *eTradeCustomer) GetCustomerName() string {
	return c.customerName
}

func (c *eTradeCustomer) ListAccounts() (string, error) {
	response, err := c.client.ListAccounts()
	return fmt.Sprintf("%#v", response), err
}

func (c *eTradeCustomer) ListAlerts() (string, error) {
	return "", nil
}

func (c *eTradeCustomer) GetQuotes(symbols string) (string, error) {
	return "", nil
}

func (c *eTradeCustomer) LookUpProduct(search string) (string, error) {
	return "", nil
}

func (c *eTradeCustomer) GetOptionChains() (string, error) {
	return "", nil
}

func (c *eTradeCustomer) GetOptionExpireDates() (string, error) {
	return "", nil
}
