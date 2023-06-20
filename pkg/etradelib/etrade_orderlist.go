package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeOrderList interface {
	GetAllOrders() []ETradeOrder
	GetOrderById(orderID int64) ETradeOrder
}

type eTradeOrderList struct {
	orders []ETradeOrder
}

const (
	// The order list response JSON looks like this:
	// {
	//   "OrdersResponse": {
	//     "Order": [
	//       {
	//         <order info>
	//       }
	//     ]
	//   }
	// }

	// ordersSliceResponsePath is the path to a slice of orders.
	ordersSliceResponsePath = "ordersResponse.order"
)

func CreateETradeOrderList(orderListResponseMap jsonmap.JsonMap) (ETradeOrderList, error) {
	ordersSlice, err := orderListResponseMap.GetSliceOfMapsAtPath(ordersSliceResponsePath)
	if err != nil {
		return nil, err
	}
	allOrders := make([]ETradeOrder, 0, len(ordersSlice))
	for _, orderInfoMap := range ordersSlice {
		order, err := CreateETradeOrder(orderInfoMap)
		if err != nil {
			return nil, err
		}
		allOrders = append(allOrders, order)
	}
	return &eTradeOrderList{orders: allOrders}, nil
}

func (e *eTradeOrderList) GetAllOrders() []ETradeOrder {
	return e.orders
}

func (e *eTradeOrderList) GetOrderById(orderID int64) ETradeOrder {
	for _, order := range e.orders {
		if order.GetId() == orderID {
			return order
		}
	}
	return nil
}
