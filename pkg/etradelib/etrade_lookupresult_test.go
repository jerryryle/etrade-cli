package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeLookup(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeLookupResult
	}{
		{
			name: "CreateETradeLookupResult Creates Lookup With Valid Response",
			testJson: `
{
  "key": "value"
}`,
			expectErr: false,
			expectValue: &eTradeLookupResult{
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
				actualValue, err := CreateETradeLookupResult(responseMap)
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

func TestETradeLookup_AsJsonMap(t *testing.T) {
	testLookup := &eTradeLookupResult{
		jsonMap: jsonmap.JsonMap{
			"key": "value",
		},
	}
	expectedValue := jsonmap.JsonMap{
		"key": "value",
	}

	actualValue := testLookup.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
