package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccountBalances(t *testing.T) {
	testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "TestId",
          "accountIdKey": "TestKey"
        }
      ]
    }
  }
}`)

	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Gets Account Balances",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testBalances := []byte(`
{
  "BalanceResponse": {
    "testBalanceKey": "testBalanceValue"
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On("GetAccountBalances", "TestKey", true).Return(testBalances, nil)
				return GetAccountBalances(mockClient, "TestId", true)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"testBalanceKey": "testBalanceValue",
			},
		},
		{
			name: "Fails On Bad Account Id",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				return GetAccountBalances(mockClient, "BadId", true)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On GetAccountBalances Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On("GetAccountBalances", "TestKey", true).Return([]byte{}, errors.New("test error"))
				return GetAccountBalances(mockClient, "TestId", true)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Balance Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testBalances := []byte(`
{
  "MISSING": {
    "testBalanceKey": "testBalanceValue"
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On("GetAccountBalances", "TestKey", true).Return(testBalances, nil)
				return GetAccountBalances(mockClient, "TestId", true)
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
