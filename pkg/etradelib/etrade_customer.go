package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
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
	eTradeClient client.ETradeClient
	customerName string
}

func CreateETradeCustomer(eTradeClient client.ETradeClient, customerName string) ETradeCustomer {
	return &eTradeCustomer{
		eTradeClient: eTradeClient,
		customerName: customerName,
	}
}

func (c *eTradeCustomer) GetCustomerName() string {
	return c.customerName
}

func (c *eTradeCustomer) GetAllAccounts() ([]ETradeAccount, error) {
	jsonMap, err := ExecuteClientCallAndWrapResponse(func() ([]byte, error) { return c.eTradeClient.ListAccounts() })
	if err != nil {
		return nil, err
	}
	accountsSlice, err := jsonMap.GetSliceAtPath("accountListResponse.accounts.account")
	if err != nil {
		return nil, err
	}
	allAccounts := make([]ETradeAccount, 0, len(accountsSlice))
	for _, accountInfo := range accountsSlice {
		accountInfoMap, err := jsonmap.FromInterface(accountInfo)
		if err != nil {
			return nil, err
		}
		account, err := CreateETradeAccount(c.eTradeClient, accountInfoMap)
		if err != nil {
			return nil, err
		}
		allAccounts = append(allAccounts, account)
	}
	return allAccounts, nil
}

func (c *eTradeCustomer) GetAccountById(accountID string) (ETradeAccount, error) {
	return nil, nil
}

func (c *eTradeCustomer) GetAllAlerts() ([]ETradeAlert, error) {
	return nil, nil
}

func (c *eTradeCustomer) GetAlertById(alertID int64) (ETradeAlert, error) {
	return nil, nil
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
