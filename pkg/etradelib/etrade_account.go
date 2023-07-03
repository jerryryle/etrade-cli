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

	// accountAccountIdResponseKey is the key for the account ID
	accountAccountIdResponseKey = "accountId"

	// accountAccountIdResponseKey is the key for the account ID Key
	accountAccountIdKeyResponseKey = "accountIdKey"
)

func CreateETradeAccount(responseMap jsonmap.JsonMap) (ETradeAccount, error) {
	accountId, err := responseMap.GetString(accountAccountIdResponseKey)
	if err != nil {
		return nil, err
	}
	accountIdKey, err := responseMap.GetString(accountAccountIdKeyResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeAccount{
		id:      accountId,
		idKey:   accountIdKey,
		jsonMap: responseMap,
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
