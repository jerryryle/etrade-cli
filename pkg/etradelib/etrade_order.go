package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeOrder interface {
	GetId() int64
	GetInfoMap() jsonmap.JsonMap
}

type eTradeOrder struct {
	id      int64
	infoMap jsonmap.JsonMap
}

const (
	// The order response JSON looks like this:
	// {
	//   "orderId": 1234,
	//   <other order keys/values>
	// }

	// orderIdResponseKey is the key for the order ID
	orderIdResponseKey = "orderId"
)

func CreateETradeOrder(orderResponseMap jsonmap.JsonMap) (ETradeOrder, error) {
	orderId, err := orderResponseMap.GetInt(orderIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeOrder{
		id:      orderId,
		infoMap: orderResponseMap,
	}, nil
}

func (e *eTradeOrder) GetId() int64 {
	return e.id
}

func (e *eTradeOrder) GetInfoMap() jsonmap.JsonMap {
	return e.infoMap
}
