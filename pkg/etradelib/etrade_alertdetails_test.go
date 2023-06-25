package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeAlertDetails(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAlertDetails
	}{
		{
			name: "CreateETradeAlertDetails Creates Alert With Valid Response",
			testJson: `
{
  "AlertDetailsResponse": {
    "id": 1234
  }
}`,
			expectErr: false,
			expectValue: &eTradeAlertDetails{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"id": json.Number("1234"),
				},
			},
		},
		{
			name: "CreateETradeAlertDetails Fails If Missing Alert ID",
			testJson: `
{
  "AlertDetailsResponse": {
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
				actualValue, err := CreateETradeAlertDetails(responseMap)
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

func TestETradeAlertDetails_GetId(t *testing.T) {
	testAlertDetails := &eTradeAlertDetails{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"id": json.Number("1234"),
		},
	}
	expectedValue := int64(1234)

	actualValue := testAlertDetails.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeAlertDetails_AsJsonMap(t *testing.T) {
	testAlertDetails := &eTradeAlertDetails{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"id": json.Number("1234"),
		},
	}
	expectedValue := jsonmap.JsonMap{
		"id": json.Number("1234"),
	}

	actualValue := testAlertDetails.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
