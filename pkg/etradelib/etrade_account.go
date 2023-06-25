package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAccount interface {
	GetId() string
	GetIdKey() string
	AsJsonMap() jsonmap.JsonMap
}

type eTradeAccount struct {
	id      string
	idKey   string
	jsonMap jsonmap.JsonMap
}

const (
	// The account response JSON looks like this:
	// {
	//   "accountId": "12345678",
	//   "accountIdKey": "account id key",
	//   <other account keys/values>
	// }

	// accountIdResponseKey is the key for the account ID
	accountIdResponseKey = "accountId"

	// accountIdResponseKey is the key for the account ID Key
	accountIdKeyResponseKey = "accountIdKey"
)

func CreateETradeAccount(accountJsonMap jsonmap.JsonMap) (ETradeAccount, error) {
	accountId, err := accountJsonMap.GetString(accountIdResponseKey)
	if err != nil {
		return nil, err
	}
	accountIdKey, err := accountJsonMap.GetString(accountIdKeyResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeAccount{
		id:      accountId,
		idKey:   accountIdKey,
		jsonMap: accountJsonMap,
	}, nil
}

func (e *eTradeAccount) GetId() string {
	return e.id
}

func (e *eTradeAccount) GetIdKey() string {
	return e.idKey
}

func (e *eTradeAccount) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
