package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAccountList interface {
	GetAllAccounts() []ETradeAccount
	GetAccountById(accountID string) ETradeAccount
}

type eTradeAccountList struct {
	accounts []ETradeAccount
}

const (
	// The account list response JSON looks like this:
	// "AccountListResponse": {
	//   "Accounts": {
	//     "Account": [
	//       {
	//         <account info>
	//       }
	//     ]
	//   }
	// }

	// accountsSliceResponsePath is the path to a slice of accounts.
	accountsSliceResponsePath = "accountListResponse.accounts.account"
)

func CreateETradeAccountList(accountListResponseMap jsonmap.JsonMap) (ETradeAccountList, error) {
	accountsSlice, err := accountListResponseMap.GetSliceOfMapsAtPath(accountsSliceResponsePath)
	if err != nil {
		return nil, err
	}
	allAccounts := make([]ETradeAccount, 0, len(accountsSlice))
	for _, accountInfoMap := range accountsSlice {
		account, err := CreateETradeAccount(accountInfoMap)
		if err != nil {
			return nil, err
		}
		allAccounts = append(allAccounts, account)
	}
	return &eTradeAccountList{accounts: allAccounts}, nil
}

func (e *eTradeAccountList) GetAllAccounts() []ETradeAccount {
	return e.accounts
}

func (e *eTradeAccountList) GetAccountById(accountID string) ETradeAccount {
	for _, account := range e.accounts {
		if account.GetId() == accountID {
			return account
		}
	}
	return nil
}
