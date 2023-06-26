package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeOrderListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeOrderList
	}{
		{
			name: "Creates List",
			testJson: `
{
  "OrdersResponse": {
    "Order": [
      {
        "orderId": 1234
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeOrderList{
				orders: []ETradeOrder{
					&eTradeOrder{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"orderId": json.Number("1234"),
						},
					},
				},
			},
		},
		{
			name: "Can Create Empty List",
			testJson: `
{
  "OrdersResponse": {
    "Order": [
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeOrderList{
				orders: []ETradeOrder{},
			},
		},
		{
			name: "Fails With Invalid JSON",
			testJson: `
{
  "OrdersResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing Order",
			testJson: `
{
  "OrdersResponse": {
    "MISSING": [
      {
        "orderId": 1234
      }
    ]
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := CreateETradeOrderListFromResponse([]byte(tt.testJson))
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeOrderList_GetAllOrders(t *testing.T) {
	tests := []struct {
		name          string
		testOrderList ETradeOrderList
		expectValue   []ETradeOrder
	}{
		{
			name: "Returns All Orders",
			testOrderList: &eTradeOrderList{
				orders: []ETradeOrder{
					&eTradeOrder{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"orderId": json.Number("1234"),
						},
					},
					&eTradeOrder{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"orderId": json.Number("5678"),
						},
					},
				},
			},
			expectValue: []ETradeOrder{
				&eTradeOrder{
					id: 1234,
					jsonMap: jsonmap.JsonMap{
						"orderId": json.Number("1234"),
					},
				},
				&eTradeOrder{
					id: 5678,
					jsonMap: jsonmap.JsonMap{
						"orderId": json.Number("5678"),
					},
				},
			},
		},
		{
			name: "Can Return Empty List",
			testOrderList: &eTradeOrderList{
				orders: []ETradeOrder{},
			},
			expectValue: []ETradeOrder{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testOrderList.GetAllOrders()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeOrderList_GetOrderById(t *testing.T) {
	tests := []struct {
		name          string
		testOrderList ETradeOrderList
		testOrderID   int64
		expectValue   ETradeOrder
	}{
		{
			name: "Returns Order For Valid ID",
			testOrderList: &eTradeOrderList{
				orders: []ETradeOrder{
					&eTradeOrder{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"orderId": json.Number("1234"),
						},
					},
				},
			},
			testOrderID: 1234,
			expectValue: &eTradeOrder{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"orderId": json.Number("1234"),
				},
			},
		},
		{
			name: "Returns Nil For Invalid ID",
			testOrderList: &eTradeOrderList{
				orders: []ETradeOrder{
					&eTradeOrder{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"orderId": json.Number("1234"),
						},
					},
				},
			},
			testOrderID: 5678,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testOrderList.GetOrderById(tt.testOrderID)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeOrderList_AddPageFromResponse(t *testing.T) {
	startingObject := &eTradeOrderList{
		orders: []ETradeOrder{
			&eTradeOrder{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"orderId": json.Number("1234"),
				},
			},
		},
		nextPage: "2",
	}

	tests := []struct {
		name        string
		startValue  ETradeOrderList
		testJson    string
		expectErr   bool
		expectValue ETradeOrderList
	}{
		{
			name:       "Can Add Pages",
			startValue: startingObject,
			testJson: `
{
  "OrdersResponse": {
    "Order": [
      {
        "orderId": 5678
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeOrderList{
				orders: []ETradeOrder{
					&eTradeOrder{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"orderId": json.Number("1234"),
						},
					},
					// Orders in subsequent pages are appended to
					// the order list.
					&eTradeOrder{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"orderId": json.Number("5678"),
						},
					},
				},
				nextPage: "",
			},
		},
		{
			name:       "Fails With Invalid JSON",
			startValue: startingObject,
			testJson: `
{
  "OrdersResponse": {
}`,
			expectErr:   true,
			expectValue: startingObject,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				orderList := tt.startValue
				// Call the Method Under Test
				err := orderList.AddPageFromResponse([]byte(tt.testJson))
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, orderList)
			},
		)
	}
}

func TestETradeOrderList_NextPage(t *testing.T) {
	testObject := &eTradeOrderList{
		orders:   []ETradeOrder{},
		nextPage: "1234",
	}
	assert.Equal(t, "1234", testObject.NextPage())

	testObject = &eTradeOrderList{
		orders:   []ETradeOrder{},
		nextPage: "",
	}
	assert.Equal(t, "", testObject.NextPage())
}

func TestETradeOrderList_AsJsonMap(t *testing.T) {
	testObject := &eTradeOrderList{
		orders: []ETradeOrder{
			&eTradeOrder{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"orderId": json.Number("1234"),
				},
			},
		},
	}

	expectValue := jsonmap.JsonMap{
		"orders": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"orderId": json.Number("1234"),
			},
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
