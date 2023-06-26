package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeTransactionDetailsFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeTransactionDetails
	}{
		{
			name: "Creates Transaction Details",
			testJson: `
{
  "TransactionDetailsResponse": {
    "transactionId": 1234
  }
}`,
			expectErr: false,
			expectValue: &eTradeTransactionDetails{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"transactionId": json.Number("1234"),
				},
			},
		},
		{
			name: "Fails With Invalid JSON",
			testJson: `
{
  "TransactionDetailsResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing Transaction ID",
			testJson: `
{
  "TransactionDetailsResponse": {
    "MISSING": 1234
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
				actualValue, err := CreateETradeTransactionDetailsFromResponse([]byte(tt.testJson))
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

func TestETradeTransactionDetails_GetId(t *testing.T) {
	testObject := &eTradeTransactionDetails{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"transactionId": json.Number("1234"),
		},
	}
	expectedValue := int64(1234)

	actualValue := testObject.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeTransactionDetails_AsJsonMap(t *testing.T) {
	testObject := &eTradeTransactionDetails{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"transactionId": json.Number("1234"),
		},
	}
	expectedValue := jsonmap.JsonMap{
		"transactionId": json.Number("1234"),
	}

	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
