package etradelib

import (
	"errors"
	"fmt"
	client2 "github.com/jerryryle/etrade-cli/pkg/etradelib/client"
)

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

type eTradeCustomer struct {
	client       client2.ETradeClient
	customerName string
}

func CreateETradeCustomer(client client2.ETradeClient, customerName string) ETradeCustomer {
	return &eTradeCustomer{
		client:       client,
		customerName: customerName,
	}
}

func (c *eTradeCustomer) GetCustomerName() string {
	return c.customerName
}

func (c *eTradeCustomer) GetAllAccounts() ([]ETradeAccount, error) {
	response, err := c.client.ListAccounts()
	if err != nil {
		return nil, err
	}
	var accounts = make([]ETradeAccount, 0)
	for _, account := range response.Accounts {
		accounts = append(
			accounts,
			CreateETradeAccount(c.client, CreateETradeAccountInfoFromResponse(account)))
	}
	return accounts, err
}

func (c *eTradeCustomer) GetAccountById(accountID string) (ETradeAccount, error) {
	accounts, err := c.GetAllAccounts()
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.GetAccountInfo().AccountId == accountID {
			return account, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no account found with the id '%s'", accountID))
}

func (c *eTradeCustomer) GetAllAlerts() ([]ETradeAlert, error) {
	response, err := c.client.ListAlerts()
	if err != nil {
		return nil, err
	}
	var alerts = make([]ETradeAlert, 0)
	for _, alert := range response.Alerts {
		alerts = append(
			alerts,
			CreateETradeAlert(c.client, CreateETradeAlertInfoFromResponse(alert)))
	}
	return alerts, err
}

func (c *eTradeCustomer) GetAlertById(alertID int64) (ETradeAlert, error) {
	alerts, err := c.GetAllAlerts()
	if err != nil {
		return nil, err
	}
	for _, alert := range alerts {
		if alert.GetAlertInfo().Id == alertID {
			return alert, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no alert found with the id '%d'", alertID))
}

func (c *eTradeCustomer) GetQuotesAll(symbols []string) ([]ETradeQuoteAllInfo, error) {
	return nil, nil
}

func (c *eTradeCustomer) GetQuotesFundamental(symbols []string) ([]ETradeQuoteFundamentalInfo, error) {
	return nil, nil
}

func (c *eTradeCustomer) GetQuotesIntraday(symbols []string) ([]ETradeQuoteIntradayInfo, error) {
	return nil, nil
}

func (c *eTradeCustomer) GetQuotesOptions(symbols []string) ([]ETradeQuoteOptionsInfo, error) {
	return nil, nil
}

func (c *eTradeCustomer) GetQuotesWeek52(symbols []string) ([]ETradeQuoteWeek52Info, error) {
	return nil, nil
}

func (c *eTradeCustomer) GetQuotesMutualFund(symbols []string) ([]ETradeQuoteMutualFundInfo, error) {
	return nil, nil
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
