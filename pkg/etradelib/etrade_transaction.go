package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeTransaction interface {
	GetId() string
	AsJsonMap() jsonmap.JsonMap
}

type eTradeTransaction struct {
	id      string
	jsonMap jsonmap.JsonMap
}

const (
	// The transaction response JSON looks like this:
	// {
	//   "transactionId": 1234,
	//   <other alert keys/values>
	// }

	// transactionTransactionIdResponseKey is the key for the transaction ID
	transactionTransactionIdResponseKey = "transactionId"
)

func CreateETradeTransaction(responseMap jsonmap.JsonMap) (ETradeTransaction, error) {
	transactionId, err := responseMap.GetString(transactionTransactionIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeTransaction{
		id:      transactionId,
		jsonMap: responseMap,
	}, nil
}

func (e *eTradeTransaction) GetId() string {
	return e.id
}

func (e *eTradeTransaction) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
