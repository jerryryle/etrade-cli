package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAccount interface {
	GetAccountId() string
	GetAccountIdKey() string
	GetAccountInfoMap() jsonmap.JsonMap
}

type eTradeAccount struct {
	accountInfoMap jsonmap.JsonMap
	accountId      string
	accountIdKey   string
}

const (
	// The account response JSON looks like this:
	// {
	//   "accountId": "12345678",
	//   "accountIdKey": "abcdefghijklmnop",
	//   <other account keys/values>
	// }

	// accountIdResponseKey is the key for the account ID
	accountIdResponseKey = "accountId"

	// accountIdResponseKey is the key for the account ID Key
	accountIdKeyResponseKey = "accountIdKey"
)

func CreateETradeAccount(accountResponseMap jsonmap.JsonMap) (ETradeAccount, error) {
	accountId, err := accountResponseMap.GetString(accountIdResponseKey)
	if err != nil {
		return nil, err
	}
	accountIdKey, err := accountResponseMap.GetString(accountIdKeyResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeAccount{
		accountInfoMap: accountResponseMap,
		accountId:      accountId,
		accountIdKey:   accountIdKey,
	}, nil
}

func (e *eTradeAccount) GetAccountId() string {
	return e.accountId
}

func (e *eTradeAccount) GetAccountIdKey() string {
	return e.accountIdKey
}

func (e *eTradeAccount) GetAccountInfoMap() jsonmap.JsonMap {
	return e.accountInfoMap
}
