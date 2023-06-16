package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"time"
)

type ETradeAccount interface {
	GetAccountInfo() ETradeAccountInfo
	GetAccountBalances() (string, error)
	ListTransactions() (string, error)
	ViewPortfolio() (string, error)
	ListOrders() (string, error)
	CreateOrder() (string, error)
}

type ETradeAccountInfo struct {
	AccountId                  string
	AccountIdKey               string
	AccountMode                string
	AccountDesc                string
	AccountName                string
	AccountType                string
	InstitutionType            string
	AccountStatus              string
	ClosedDate                 time.Time
	ShareWorksAccount          bool
	ShareWorksSource           string
	FcManagedMssbClosedAccount bool
}

type eTradeAccount struct {
	eTradeClient client.ETradeClient
	accountInfo  ETradeAccountInfo
}

func CreateETradeAccount(client client.ETradeClient, accountInfo *ETradeAccountInfo) ETradeAccount {
	return &eTradeAccount{
		eTradeClient: client,
		accountInfo:  *accountInfo,
	}
}

func (a *eTradeAccount) GetAccountInfo() ETradeAccountInfo {
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
