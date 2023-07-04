package cmd

import (
	"encoding/json"
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListTransactionDetails(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Lists Transaction Details",
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
				testTransactionDetailsResponse := []byte(`
{
  "TransactionDetailsResponse": {
    "transactionId": 1234
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On("ListTransactionDetails", "test key", "1234").Return(testTransactionDetailsResponse, nil)
				return ListTransactionDetails(mockClient, "test id", "1234")
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"transactionId": json.Number("1234"),
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
				return ListTransactionDetails(mockClient, "bad id", "1234")
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails With ListTransactionDetails Error",
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
				mockClient.On("ListTransactionDetails", "test key", "1234").Return([]byte{}, errors.New("test error"))
				return ListTransactionDetails(mockClient, "test id", "1234")
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Response",
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
				testTransactionDetailsResponse := []byte(`
{
  "TransactionDetailsResponse": {
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On("ListTransactionDetails", "test key", "1234").Return(testTransactionDetailsResponse, nil)
				return ListTransactionDetails(mockClient, "test id", "1234")
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
