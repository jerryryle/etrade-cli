package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeQuote(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeQuote
	}{
		{
			name: "CreateETradeQuote Creates Quote With Valid Response",
			testJson: `
{
  "key": "value"
}`,
			expectErr: false,
			expectValue: &eTradeQuote{
				jsonMap: jsonmap.JsonMap{
					"key": "value",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				responseMap, err := NewNormalizedJsonMap([]byte(tt.testJson))
				require.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := CreateETradeQuote(responseMap)
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

func TestETradeQuote_AsJsonMap(t *testing.T) {
	testQuote := &eTradeQuote{
		jsonMap: jsonmap.JsonMap{
			"key": "value",
		},
	}
	expectedValue := jsonmap.JsonMap{
		"key": "value",
	}

	actualValue := testQuote.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
