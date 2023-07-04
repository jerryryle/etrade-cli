package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAccounts(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Gets Quotes",
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
				return ListAccounts(mockClient)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"accounts": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"accountId":    "1234",
						"accountIdKey": "test key",
					},
				},
			},
		},
		{
			name: "Fails On ListAccounts Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return([]byte{}, errors.New("test error"))
				return ListAccounts(mockClient)
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
}`)
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				return ListAccounts(mockClient)
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
