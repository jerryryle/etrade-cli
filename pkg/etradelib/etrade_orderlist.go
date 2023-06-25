package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeOrderList interface {
	GetAllOrders() []ETradeOrder
	GetOrderById(orderID int64) ETradeOrder
	NextPage() string
	AddPage(orderListResponseMap jsonmap.JsonMap) error
	AddPageFromResponse(response []byte) error
	AsJsonMap() jsonmap.JsonMap
}

type eTradeOrderList struct {
	orders   []ETradeOrder
	nextPage string
}

const (
	// The AsJsonMap() map looks like this:
	// "orders": [
	//   {
	//     <order info>
	//   }
	// ]

	// OrderListOrdersSliceJsonMapPath is the path to a slice of
	// orders.
	OrderListOrdersSliceJsonMapPath = ".orders"
)

const (
	// The order list response JSON looks like this:
	// {
	//   "OrdersResponse": {
	//     "marker: "marker info"
	//     "Order": [
	//       {
	//         <order info>
	//       }
	//     ]
	//   }
	// }

	// orderListOrdersSliceResponsePath is the path to a slice of orders.
	orderListOrdersSliceResponsePath = "ordersResponse.order"

	// orderListMarkerStringPath is the path to the next page marker string
	orderListMarkerStringPath = "ordersResponse.marker"
)

func CreateETradeOrderListFromResponse(response []byte) (
	ETradeOrderList, error,
) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeOrderList(responseMap)
}

func CreateETradeOrderList(orderListResponseMap jsonmap.JsonMap) (ETradeOrderList, error) {
	// Create a new orderList with everything initialized to its zero value.
	orderList := eTradeOrderList{
		orders:   []ETradeOrder{},
		nextPage: "",
	}
	err := orderList.AddPage(orderListResponseMap)
	if err != nil {
		return nil, err
	}
	return &orderList, nil
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

func (e *eTradeOrderList) NextPage() string {
	return e.nextPage
}

func (e *eTradeOrderList) AddPageFromResponse(response []byte) error {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return err
	}
	return e.AddPage(responseMap)
}

func (e *eTradeOrderList) AddPage(orderListResponseMap jsonmap.JsonMap) error {
	ordersSlice, err := orderListResponseMap.GetSliceOfMapsAtPath(orderListOrdersSliceResponsePath)
	if err != nil {
		return err
	}

	// the marker key only appears if there are more pages, so ignore any
	// error and accept a possibly-zero int.
	nextPage, _ := orderListResponseMap.GetStringAtPath(orderListMarkerStringPath)

	allOrders := make([]ETradeOrder, 0, len(ordersSlice))
	for _, orderJsonMap := range ordersSlice {
		order, err := CreateETradeOrder(orderJsonMap)
		if err != nil {
			return err
		}
		allOrders = append(allOrders, order)
	}
	e.orders = append(e.orders, allOrders...)
	e.nextPage = nextPage
	return nil
}

func (e *eTradeOrderList) AsJsonMap() jsonmap.JsonMap {
	ordersSlice := make(jsonmap.JsonSlice, 0, len(e.orders))
	for _, order := range e.orders {
		ordersSlice = append(ordersSlice, order.AsJsonMap())
	}
	var ordersListMap = jsonmap.JsonMap{}
	err := ordersListMap.SetSliceAtPath(OrderListOrdersSliceJsonMapPath, ordersSlice)
	if err != nil {
		panic(err)
	}
	return ordersListMap
}
