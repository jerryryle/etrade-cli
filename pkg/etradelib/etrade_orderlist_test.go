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
				[]ETradeOrder{
					&eTradeOrder{
						id: 1234,
						infoMap: jsonmap.JsonMap{
							"orderId": json.Number("1234"),
						},
					},
					&eTradeOrder{
						id: 5678,
						infoMap: jsonmap.JsonMap{
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
				[]ETradeOrder{},
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
				[]ETradeOrder{
					&eTradeOrder{
						id: 1234,
						infoMap: jsonmap.JsonMap{
							"orderId": json.Number("1234"),
						},
					},
					&eTradeOrder{
						id: 5678,
						infoMap: jsonmap.JsonMap{
							"orderId": json.Number("5678"),
						},
					},
				},
			},
			expectValue: []ETradeOrder{
				&eTradeOrder{
					id: 1234,
					infoMap: jsonmap.JsonMap{
						"orderId": json.Number("1234"),
					},
				},
				&eTradeOrder{
					id: 5678,
					infoMap: jsonmap.JsonMap{
						"orderId": json.Number("5678"),
					},
				},
			},
		},
		{
			name: "GetAllAccounts Can Return Empty List",
			testOrderList: &eTradeOrderList{
				[]ETradeOrder{},
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
			name: "GetAccountById Returns Account For Valid ID",
			testOrderList: &eTradeOrderList{
				[]ETradeOrder{
					&eTradeOrder{
						id: 1234,
						infoMap: jsonmap.JsonMap{
							"orderId": json.Number("1234"),
						},
					},
				},
			},
			testOrderID: 1234,
			expectValue: &eTradeOrder{
				id: 1234,
				infoMap: jsonmap.JsonMap{
					"orderId": json.Number("1234"),
				},
			},
		},
		{
			name: "GetAccountById Returns Nil For Invalid ID",
			testOrderList: &eTradeOrderList{
				[]ETradeOrder{
					&eTradeOrder{
						id: 1234,
						infoMap: jsonmap.JsonMap{
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
