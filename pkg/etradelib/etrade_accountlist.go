package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAccountList interface {
	GetAllAccounts() []ETradeAccount
	GetAccountById(accountID string) ETradeAccount
	AsJsonMap() jsonmap.JsonMap
}

type eTradeAccountList struct {
	accounts []ETradeAccount
}

const (
	// The AsJsonMap() map looks like this:
	// {
	//   "accounts": [
	//     {
	//       <account info>
	//     }
	//   ]
	// }

	// AccountsListAccountsPath is the path to a slice of accounts.
	AccountsListAccountsPath = "accounts"
)

const (
	// The account list response JSON looks like this:
	// {
	//   "AccountListResponse": {
	//     "Accounts": {
	//       "Account": [
	//         {
	//           <account info>
	//         }
	//       ]
	//     }
	//   }
	// }

	// accountsListAccountsResponsePath is the path to a slice of accounts.
	accountsListAccountsResponsePath = "accountListResponse.accounts.account"
)

func CreateETradeAccountListFromResponse(response []byte) (ETradeAccountList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeAccountList(responseMap)
}

func CreateETradeAccountList(responseMap jsonmap.JsonMap) (ETradeAccountList, error) {
	accountsSlice, err := responseMap.GetSliceOfMapsAtPath(accountsListAccountsResponsePath)
	if err != nil {
		return nil, err
	}
	allAccounts := make([]ETradeAccount, 0, len(accountsSlice))
	for _, accountJsonMap := range accountsSlice {
		account, err := CreateETradeAccountFromMap(accountJsonMap)
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

func (e *eTradeAccountList) AsJsonMap() jsonmap.JsonMap {
	accountSlice := make(jsonmap.JsonSlice, 0, len(e.accounts))
	for _, account := range e.accounts {
		accountSlice = append(accountSlice, account.AsJsonMap())
	}
	var accountListMap = jsonmap.JsonMap{}
	err := accountListMap.SetSliceAtPath(AccountsListAccountsPath, accountSlice)
	if err != nil {
		panic(err)
	}
	return accountListMap
}
