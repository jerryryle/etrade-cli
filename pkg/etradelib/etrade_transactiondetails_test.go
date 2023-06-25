package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeTransactionDetails(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeTransactionDetails
	}{
		{
			name: "CreateETradeTransactionDetails Creates Transaction With Valid Response",
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
			name: "CreateETradeTransactionDetails Fails If Missing Transaction ID",
			testJson: `
{
  "TransactionDetailsResponse": {
    "someOtherKey": "test"
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
				actualValue, err := CreateETradeTransactionDetails(responseMap)
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
	testTransactionDetails := &eTradeTransactionDetails{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"transactionId": json.Number("1234"),
		},
	}
	expectedValue := int64(1234)

	actualValue := testTransactionDetails.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeTransactionDetails_AsJsonMap(t *testing.T) {
	testTransactionDetails := &eTradeTransactionDetails{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"transactionId": json.Number("1234"),
		},
	}
	expectedValue := jsonmap.JsonMap{
		"transactionId": json.Number("1234"),
	}

	actualValue := testTransactionDetails.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
