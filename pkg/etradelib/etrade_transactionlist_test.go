package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeTransactionListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeTransactionList
	}{
		{
			name: "Creates List",
			testJson: `
{
  "TransactionListResponse": {
    "Transaction": [
      {
        "transactionId": "1234"
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: "1234",
						jsonMap: jsonmap.JsonMap{
							"transactionId": "1234",
						},
					},
				},
			},
		},
		{
			name: "Can Create Empty List",
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
			name: "Fails With Invalid JSON",
			testJson: `
{
  "TransactionListResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing Transaction Key",
			testJson: `
{
  "TransactionListResponse": {
    "MISSING": [
      {
        "transactionId": "1234"
      }
    ]
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing Transaction ID",
			testJson: `
{
  "TransactionListResponse": {
    "Transaction": [
      {
        "MISSING": "1234"
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
				actualValue, err := CreateETradeTransactionListFromResponse([]byte(tt.testJson))
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
		name        string
		testObject  ETradeTransactionList
		expectValue []ETradeTransaction
	}{
		{
			name: "GetAllTransactions Returns All Transactions",
			testObject: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: "1234",
						jsonMap: jsonmap.JsonMap{
							"transactionId": "1234",
						},
					},
					&eTradeTransaction{
						id: "5678",
						jsonMap: jsonmap.JsonMap{
							"transactionId": "5678",
						},
					},
				},
			},
			expectValue: []ETradeTransaction{
				&eTradeTransaction{
					id: "1234",
					jsonMap: jsonmap.JsonMap{
						"transactionId": "1234",
					},
				},
				&eTradeTransaction{
					id: "5678",
					jsonMap: jsonmap.JsonMap{
						"transactionId": "5678",
					},
				},
			},
		},
		{
			name: "GetAllTransactions Can Return Empty List",
			testObject: &eTradeTransactionList{
				transactions: []ETradeTransaction{},
			},
			expectValue: []ETradeTransaction{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAllTransactions()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeTransactionList_GetTransactionById(t *testing.T) {
	tests := []struct {
		name        string
		testObject  ETradeTransactionList
		testId      string
		expectValue ETradeTransaction
	}{
		{
			name: "GetTransactionById Returns Transaction For Valid ID",
			testObject: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: "1234",
						jsonMap: jsonmap.JsonMap{
							"transactionId": "1234",
						},
					},
				},
			},
			testId: "1234",
			expectValue: &eTradeTransaction{
				id: "1234",
				jsonMap: jsonmap.JsonMap{
					"transactionId": "1234",
				},
			},
		},
		{
			name: "GetTransactionById Returns Nil For Invalid ID",
			testObject: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: "1234",
						jsonMap: jsonmap.JsonMap{
							"transactionId": "1234",
						},
					},
				},
			},
			testId:      "5678",
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetTransactionById(tt.testId)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeTransactionList_AddPageFromResponse(t *testing.T) {
	startingObject := &eTradeTransactionList{
		transactions: []ETradeTransaction{
			&eTradeTransaction{
				id: "1234",
				jsonMap: jsonmap.JsonMap{
					"transactionId": "1234",
				},
			},
		},
		nextPage: "2",
	}

	tests := []struct {
		name        string
		startValue  ETradeTransactionList
		testJson    string
		expectErr   bool
		expectValue ETradeTransactionList
	}{
		{
			name:       "Can Add Pages",
			startValue: startingObject,
			testJson: `
{
  "TransactionListResponse": {
    "Transaction": [
      {
        "transactionId": "5678"
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeTransactionList{
				transactions: []ETradeTransaction{
					&eTradeTransaction{
						id: "1234",
						jsonMap: jsonmap.JsonMap{
							"transactionId": "1234",
						},
					},
					// Transactions in subsequent pages are appended to
					// the transaction list.
					&eTradeTransaction{
						id: "5678",
						jsonMap: jsonmap.JsonMap{
							"transactionId": "5678",
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
  "TransactionListResponse": {
}`,
			expectErr:   true,
			expectValue: startingObject,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				transactionList := tt.startValue
				// Call the Method Under Test
				err := transactionList.AddPageFromResponse([]byte(tt.testJson))
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, transactionList)
			},
		)
	}
}

func TestETradeTransactionList_NextPage(t *testing.T) {
	testObject := &eTradeTransactionList{
		transactions: []ETradeTransaction{},
		nextPage:     "1234",
	}
	assert.Equal(t, "1234", testObject.NextPage())

	testObject = &eTradeTransactionList{
		transactions: []ETradeTransaction{},
		nextPage:     "",
	}
	assert.Equal(t, "", testObject.NextPage())
}

func TestETradeTransactionList_AsJsonMap(t *testing.T) {
	testObject := &eTradeTransactionList{
		transactions: []ETradeTransaction{
			&eTradeTransaction{
				id: "1234",
				jsonMap: jsonmap.JsonMap{
					"transactionId": "1234",
				},
			},
		},
	}

	expectValue := jsonmap.JsonMap{
		"transactions": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"transactionId": "1234",
			},
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
