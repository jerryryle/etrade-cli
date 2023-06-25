package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeOrderList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeOrderList
	}{
		{
			name: "CreateETradeOrderList Creates List With Valid Response",
			testJson: `
{
  "OrdersResponse": {
    "Order": [
      {
        "orderId": 1234
      },
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
					&eTradeOrder{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"orderId": json.Number("5678"),
						},
					},
				},
			},
		},
		{
			name: "CreateETradeOrderList Can Create Empty List",
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
			name: "CreateETradeOrderList Fails With Invalid Response",
			// The "Order" level is not an array in the following string
			testJson: `
{
  "OrdersResponse": {
    "Order": {
      "orderId": 1234
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				responseMap, err := NewNormalizedJsonMap([]byte(tt.testJson))
				require.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := CreateETradeOrderList(responseMap)
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
			name: "GetAllOrders Returns All Orders",
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
			name: "GetAllOrders Can Return Empty List",
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
			name: "GetOrderById Returns Order For Valid ID",
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
			name: "GetOrderById Returns Nil For Invalid ID",
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

func TestETradeOrderList_AddPage(t *testing.T) {
	type pageTest struct {
		testJson    string
		expectErr   bool
		expectValue ETradeOrderList
	}
	tests := []struct {
		name      string
		pageTests []pageTest
	}{
		{
			name: "AddPage Can Add Pages",
			pageTests: []pageTest{
				{
					testJson: `
{
  "OrdersResponse": {
    "marker": "2",
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
						nextPage: "2",
					},
				},
				{
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
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var orderList ETradeOrderList
				for testIndex, pt := range tt.pageTests {
					responseMap, err := NewNormalizedJsonMap([]byte(pt.testJson))
					require.Nil(t, err)

					if testIndex == 0 {
						orderList, err = CreateETradeOrderList(responseMap)
					} else {
						// Call the Method Under Test
						err = orderList.AddPage(responseMap)
					}
					if pt.expectErr {
						assert.Error(t, err)
					} else {
						assert.Nil(t, err)
					}
					assert.Equal(t, pt.expectValue, orderList)
				}
			},
		)
	}
}

func TestETradeOrderList_NextPage(t *testing.T) {
	testOrderList := &eTradeOrderList{
		orders:   []ETradeOrder{},
		nextPage: "1234",
	}
	assert.Equal(t, "1234", testOrderList.NextPage())

	testOrderList = &eTradeOrderList{
		orders:   []ETradeOrder{},
		nextPage: "",
	}
	assert.Equal(t, "", testOrderList.NextPage())
}
