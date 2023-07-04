package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccountById(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Gets Account By Id",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "1234",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				return GetAccountById(mockClient, "1234")
			},
			expectErr: false,
			expectValue: etradelib.CreateETradeAccount(
				"1234", "test key", jsonmap.JsonMap{
					"accountId":    "1234",
					"accountIdKey": "test key",
				},
			),
		},
		{
			name: "Fails On ListAccounts Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return([]byte{}, errors.New("test error"))
				return GetAccountById(mockClient, "1234")
			},
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails On Bad Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				return GetAccountById(mockClient, "1234")
			},
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails On Account Not Found",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "1234",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				return GetAccountById(mockClient, "1111")
			},
			expectErr:   true,
			expectValue: nil,
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
