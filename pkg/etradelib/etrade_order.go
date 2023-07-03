package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeOrder interface {
	GetId() int64
	AsJsonMap() jsonmap.JsonMap
}

type eTradeOrder struct {
	id      int64
	jsonMap jsonmap.JsonMap
}

const (
	// The order response JSON looks like this:
	// {
	//   "orderId": 1234,
	//   <other order keys/values>
	// }

	// orderOrderIdResponseKey is the key for the order ID
	orderOrderIdResponseKey = "orderId"
)

func CreateETradeOrder(responseMap jsonmap.JsonMap) (ETradeOrder, error) {
	orderId, err := responseMap.GetInt(orderOrderIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeOrder{
		id:      orderId,
		jsonMap: responseMap,
	}, nil
}

func (e *eTradeOrder) GetId() int64 {
	return e.id
}

func (e *eTradeOrder) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
