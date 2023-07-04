package cmd

import (
	"encoding/json"
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestListOrders(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Lists Orders With Pagination",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "test id",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
				testOrdersResponse1 := []byte(`
{
  "OrdersResponse": {
    "marker": "test marker",
    "Order": [
      {
        "orderId": 1234
      }
    ]
  }
}`)

				testOrdersResponse2 := []byte(`
{
  "OrdersResponse": {
    "Order": [
      {
        "orderId": 5678
      }
    ]
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListOrders", "test key", "", 100, constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				).Return(testOrdersResponse1, nil)
				mockClient.On(
					"ListOrders", "test key", "test marker", 100, constants.OrderStatusNil, (*time.Time)(nil),
					(*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				).Return(testOrdersResponse2, nil)

				return ListOrders(
					mockClient, "test id", constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"orders": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"orderId": json.Number("1234"),
					},
					jsonmap.JsonMap{
						"orderId": json.Number("5678"),
					},
				},
			},
		},
		{
			name: "Fails With Bad Account ID",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "test id",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)

				return ListOrders(
					mockClient, "bad id", constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On First Page ListOrders Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "test id",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListOrders", "test key", "", 100, constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				).Return([]byte{}, errors.New("test error"))

				return ListOrders(
					mockClient, "test id", constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Subsequent Page ListOrders Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "test id",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
				testOrdersResponse1 := []byte(`
{
  "OrdersResponse": {
    "marker": "test marker",
    "Order": [
      {
        "orderId": 1234
      }
    ]
  }
}`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListOrders", "test key", "", 100, constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				).Return(testOrdersResponse1, nil)
				mockClient.On(
					"ListOrders", "test key", "test marker", 100, constants.OrderStatusNil, (*time.Time)(nil),
					(*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				).Return([]byte{}, errors.New("test error"))

				return ListOrders(
					mockClient, "test id", constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad First Page Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "test id",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
				testOrdersResponse1 := []byte(`
{
  "OrdersResponse": {
}`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListOrders", "test key", "", 100, constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				).Return(testOrdersResponse1, nil)

				return ListOrders(
					mockClient, "test id", constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Subsequent Page Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "test id",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
				testOrdersResponse1 := []byte(`
{
  "OrdersResponse": {
    "marker": "test marker",
    "Order": [
      {
        "orderId": 1234
      }
    ]
  }
}`)
				testOrdersResponse2 := []byte(`
{
  "OrdersResponse": {
}`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListOrders", "test key", "", 100, constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				).Return(testOrdersResponse1, nil)
				mockClient.On(
					"ListOrders", "test key", "test marker", 100, constants.OrderStatusNil, (*time.Time)(nil),
					(*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				).Return(testOrdersResponse2, nil)

				return ListOrders(
					mockClient, "test id", constants.OrderStatusNil, (*time.Time)(nil), (*time.Time)(nil),
					[]string{"TestSymbol"}, constants.OrderSecurityTypeNil, constants.OrderTransactionTypeNil,
					constants.MarketSessionNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockClient := client.ETradeClientMock{}
				// Call the Method Under Test
				actualValue, err := tt.testFn(&mockClient)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, actualValue)
				mockClient.AssertExpectations(t)
			},
		)
	}
}
