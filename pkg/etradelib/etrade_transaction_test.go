package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeTransaction(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeTransaction
	}{
		{
			name: "Creates Transaction",
			testJson: `
{
  "transactionId": "1234"
}`,
			expectErr: false,
			expectValue: &eTradeTransaction{
				id: "1234",
				jsonMap: jsonmap.JsonMap{
					"transactionId": "1234",
				},
			},
		},
		{
			name: "Fails If Missing Transaction ID",
			testJson: `
{
  "MISSING": "1234"
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
				actualValue, err := CreateETradeTransaction(responseMap)
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

func TestETradeTransaction_GetId(t *testing.T) {
	testObject := &eTradeTransaction{
		id: "1234",
		jsonMap: jsonmap.JsonMap{
			"transactionId": "1234",
		},
	}
	expectedValue := "1234"

	actualValue := testObject.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeTransaction_AsJsonMap(t *testing.T) {
	testObject := &eTradeTransaction{
		id: "1234",
		jsonMap: jsonmap.JsonMap{
			"transactionId": "1234",
		},
	}
	expectedValue := jsonmap.JsonMap{
		"transactionId": "1234",
	}

	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
