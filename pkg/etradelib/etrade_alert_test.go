package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeAlert(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAlert
	}{
		{
			name: "CreateETradeAlert Creates Alert With Valid Response",
			testJson: `
{
  "id": 1234
}`,
			expectErr: false,
			expectValue: &eTradeAlert{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"id": json.Number("1234"),
				},
			},
		},
		{
			name: "CreateETradeAlert Fails If Missing Alert ID",
			testJson: `
{
  "someOtherKey": "test"
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
				actualValue, err := CreateETradeAlert(responseMap)
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

func TestETradeAlert_GetId(t *testing.T) {
	testAlert := &eTradeAlert{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"id": json.Number("1234"),
		},
	}
	expectedValue := int64(1234)

	actualValue := testAlert.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeAlert_AsJsonMap(t *testing.T) {
	testAlert := &eTradeAlert{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"id": json.Number("1234"),
		},
	}
	expectedValue := jsonmap.JsonMap{
		"id": json.Number("1234"),
	}

	actualValue := testAlert.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
