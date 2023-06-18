package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAccount interface {
	GetId() string
	GetIdKey() string
	GetInfoMap() jsonmap.JsonMap
}

type eTradeAccount struct {
	id      string
	idKey   string
	infoMap jsonmap.JsonMap
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
		id:      accountId,
		idKey:   accountIdKey,
		infoMap: accountResponseMap,
	}, nil
}

func (e *eTradeAccount) GetId() string {
	return e.id
}

func (e *eTradeAccount) GetIdKey() string {
	return e.idKey
}

func (e *eTradeAccount) GetInfoMap() jsonmap.JsonMap {
	return e.infoMap
}
