package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeTransactionDetails interface {
	GetId() int64
	AsJsonMap() jsonmap.JsonMap
}

type eTradeTransactionDetails struct {
	id      int64
	jsonMap jsonmap.JsonMap
}

const (
	// The transaction response JSON looks like this:
	// {
	//   "TransactionDetailsResponse": {
	//     "transactionId": 1234
	//   }
	// }

	// transactionDetailsTransactionDetailsResponseKey is the key for the
	// transaction details map
	transactionDetailsTransactionDetailsResponseKey = "transactionDetailsResponse"

	// transactionDetailsTransactionIdResponseKey is the key for the
	// transaction ID
	transactionDetailsTransactionIdResponseKey = "transactionId"
)

func CreateETradeTransactionDetailsFromResponse(response []byte) (ETradeTransactionDetails, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeTransactionDetails(responseMap)
}

func CreateETradeTransactionDetails(transactionDetailsJsonMap jsonmap.JsonMap) (ETradeTransactionDetails, error) {
	var err error
	// Flatten the response by removing the "transactionDetailsResponse" level
	transactionDetailsJsonMap, err = transactionDetailsJsonMap.GetMap(transactionDetailsTransactionDetailsResponseKey)
	if err != nil {
		return nil, err
	}

	transactionId, err := transactionDetailsJsonMap.GetInt(transactionDetailsTransactionIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeTransactionDetails{
		id:      transactionId,
		jsonMap: transactionDetailsJsonMap,
	}, nil
}

func (e *eTradeTransactionDetails) GetId() int64 {
	return e.id
}

func (e *eTradeTransactionDetails) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
