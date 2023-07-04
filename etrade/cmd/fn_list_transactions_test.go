package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestListTransactions(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Lists Transactions With Pagination",
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
				testTransactionsResponse1 := []byte(`
{
  "TransactionListResponse": {
    "marker": "test marker",
    "Transaction": [
      {
        "transactionId": "1234"
      }
    ]
  }
}`)

				testTransactionsResponse2 := []byte(`
{
  "TransactionListResponse": {
    "Transaction": [
      {
        "transactionId": "5678"
      }
    ]
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListTransactions", "test key", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil, "",
					50,
				).Return(testTransactionsResponse1, nil)
				mockClient.On(
					"ListTransactions", "test key", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
					"test marker", 50,
				).Return(testTransactionsResponse2, nil)

				return ListTransactions(
					mockClient, "test id", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
				)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"transactions": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"transactionId": "1234",
					},
					jsonmap.JsonMap{
						"transactionId": "5678",
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

				return ListTransactions(
					mockClient, "bad id", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On First Page ListTransactions Error",
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
					"ListTransactions", "test key", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil, "",
					50,
				).Return([]byte{}, errors.New("test error"))

				return ListTransactions(
					mockClient, "test id", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Subsequent Page ListTransactions Error",
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
				testTransactionsResponse1 := []byte(`
{
  "TransactionListResponse": {
    "marker": "test marker",
    "Transaction": [
      {
        "transactionId": "1234"
      }
    ]
  }
}`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListTransactions", "test key", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil, "",
					50,
				).Return(testTransactionsResponse1, nil)
				mockClient.On(
					"ListTransactions", "test key", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
					"test marker", 50,
				).Return([]byte{}, errors.New("test error"))

				return ListTransactions(
					mockClient, "test id", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
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
				testTransactionsResponse1 := []byte(`
{
  "TransactionListResponse": {
}`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListTransactions", "test key", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil, "",
					50,
				).Return(testTransactionsResponse1, nil)

				return ListTransactions(
					mockClient, "test id", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
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
				testTransactionsResponse1 := []byte(`
{
  "TransactionListResponse": {
    "marker": "test marker",
    "Transaction": [
      {
        "transactionId": "1234"
      }
    ]
  }
}`)
				testTransactionsResponse2 := []byte(`
{
  "TransactionListResponse": {
}`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ListTransactions", "test key", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil, "",
					50,
				).Return(testTransactionsResponse1, nil)
				mockClient.On(
					"ListTransactions", "test key", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
					"test marker", 50,
				).Return(testTransactionsResponse2, nil)

				return ListTransactions(
					mockClient, "test id", (*time.Time)(nil), (*time.Time)(nil), constants.SortOrderNil,
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
