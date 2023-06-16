package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAccount interface {
	GetAccountInfo() jsonmap.JsonMap
	GetAccountBalances() (string, error)
	ListTransactions() (string, error)
	ViewPortfolio() (string, error)
	ListOrders() (string, error)
	CreateOrder() (string, error)
}

type eTradeAccount struct {
	eTradeClient client.ETradeClient
	accountInfo  jsonmap.JsonMap
	accountId    string
	accountIdKey string
}

func CreateETradeAccount(client client.ETradeClient, accountInfo jsonmap.JsonMap) (ETradeAccount, error) {
	accountId, err := accountInfo.GetString("accountId")
	if err != nil {
		return nil, err
	}
	accountIdKey, err := accountInfo.GetString("accountIdKey")
	if err != nil {
		return nil, err
	}

	return &eTradeAccount{
		eTradeClient: client,
		accountInfo:  accountInfo,
		accountId:    accountId,
		accountIdKey: accountIdKey,
	}, nil
}

func (a *eTradeAccount) GetAccountInfo() jsonmap.JsonMap {
	return a.accountInfo
}

func (a *eTradeAccount) GetAccountBalances() (string, error) {
	return "", nil
}

func (a *eTradeAccount) ListTransactions() (string, error) {
	return "", nil
}

func (a *eTradeAccount) ViewPortfolio() (string, error) {
	return "", nil
}

func (a *eTradeAccount) ListOrders() (string, error) {
	return "", nil
}

func (a *eTradeAccount) CreateOrder() (string, error) {
	return "", nil
}
