package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeTransaction interface {
	GetId() int64
	GetInfoMap() jsonmap.JsonMap
}

type eTradeTransaction struct {
	id      int64
	infoMap jsonmap.JsonMap
}

const (
	// The transaction response JSON looks like this:
	// {
	//   "transactionId": 1234,
	//   <other alert keys/values>
	// }

	// transactionIdResponseKey is the key for the transaction ID
	transactionIdResponseKey = "transactionId"
)

func CreateETradeTransaction(transactionResponseMap jsonmap.JsonMap) (ETradeTransaction, error) {
	transactionId, err := transactionResponseMap.GetInt(transactionIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeTransaction{
		id:      transactionId,
		infoMap: transactionResponseMap,
	}, nil
}

func (e *eTradeTransaction) GetId() int64 {
	return e.id
}

func (e *eTradeTransaction) GetInfoMap() jsonmap.JsonMap {
	return e.infoMap
}
