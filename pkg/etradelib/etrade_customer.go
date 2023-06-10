package etradelib

import (
	"errors"
	"fmt"
)

type ETradeCustomer interface {
	GetCustomerName() string
	GetAllAccounts() ([]ETradeAccount, error)
	GetAccountById(accountID string) (ETradeAccount, error)
	GetAllAlerts() ([]ETradeAlert, error)
	GetAlertById(alertID int64) (ETradeAlert, error)
	GetQuotes(symbols string) (string, error)
	LookUpProduct(search string) (string, error)
	GetOptionChains() (string, error)
	GetOptionExpireDates() (string, error)
}

type eTradeCustomer struct {
	client       ETradeClient
	customerName string
}

func CreateETradeCustomer(client ETradeClient, customerName string) ETradeCustomer {
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
