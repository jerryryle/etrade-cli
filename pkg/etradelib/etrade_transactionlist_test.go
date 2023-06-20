package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeTransactionList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeTransactionList
	}{
		{
			name: "CreateETradeTransactionList Creates List With Valid Response",
			testJson: `
{
  "TransactionListResponse": {
    "Transaction": [
      {
        "transactionId": 1234
      },
      {
        "transactionId": 5678
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"transactionId": json.Number("1234"),
						},
					},
					&eTradeTransaction{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"transactionId": json.Number("5678"),
						},
					},
				},
			},
		},
		{
			name: "CreateETradeTransactionList Can Create Empty List",
			testJson: `
{
  "TransactionListResponse": {
    "Transaction": [
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeTransactionList{
				transactions: []ETradeTransaction{},
			},
		},
		{
			name: "CreateETradeTransactionList Fails With Invalid Response",
			// The "Transaction" level is not an array in the following string
			testJson: `
{
  "TransactionListResponse": {
    "Transaction": {
      "transactionId": 1234
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
				actualValue, err := CreateETradeTransactionList(responseMap)
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

func TestETradeTransactionList_GetAllTransactions(t *testing.T) {
	tests := []struct {
		name                string
		testTransactionList ETradeTransactionList
		expectValue         []ETradeTransaction
	}{
		{
			name: "GetAllTransactions Returns All Transactions",
			testTransactionList: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"transactionId": json.Number("1234"),
						},
					},
					&eTradeTransaction{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"transactionId": json.Number("5678"),
						},
					},
				},
			},
			expectValue: []ETradeTransaction{
				&eTradeTransaction{
					id: 1234,
					jsonMap: jsonmap.JsonMap{
						"transactionId": json.Number("1234"),
					},
				},
				&eTradeTransaction{
					id: 5678,
					jsonMap: jsonmap.JsonMap{
						"transactionId": json.Number("5678"),
					},
				},
			},
		},
		{
			name: "GetAllAccounts Can Return Empty List",
			testTransactionList: &eTradeTransactionList{
				transactions: []ETradeTransaction{},
			},
			expectValue: []ETradeTransaction{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testTransactionList.GetAllTransactions()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeTransactionList_GetTransactionById(t *testing.T) {
	tests := []struct {
		name                string
		testTransactionList ETradeTransactionList
		testTransactionID   int64
		expectValue         ETradeTransaction
	}{
		{
			name: "GetAccountById Returns Account For Valid ID",
			testTransactionList: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"transactionId": json.Number("1234"),
						},
					},
				},
			},
			testTransactionID: 1234,
			expectValue: &eTradeTransaction{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"transactionId": json.Number("1234"),
				},
			},
		},
		{
			name: "GetAccountById Returns Nil For Invalid ID",
			testTransactionList: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"transactionId": json.Number("1234"),
						},
					},
				},
			},
			testTransactionID: 5678,
			expectValue:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testTransactionList.GetTransactionById(tt.testTransactionID)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}
