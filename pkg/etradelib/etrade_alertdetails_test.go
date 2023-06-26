package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeAlertDetailsFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAlertDetails
	}{
		{
			name: "Creates Alert Details",
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
			name: "Fails With Bad JSON",
			testJson: `
{
  "AlertDetailsResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails If Missing Alert ID",
			testJson: `
{
  "AlertDetailsResponse": {
    "MISSING": 1234
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails If Missing AlertDetailsResponse",
			testJson: `
{
  "MISSING": {
    "id": 1234
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
				actualValue, err := CreateETradeAlertDetailsFromResponse([]byte(tt.testJson))
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
	testObject := &eTradeAlertDetails{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"id": json.Number("1234"),
		},
	}
	expectedValue := int64(1234)

	actualValue := testObject.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeAlertDetails_AsJsonMap(t *testing.T) {
	testObject := &eTradeAlertDetails{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"id": json.Number("1234"),
		},
	}
	expectedValue := jsonmap.JsonMap{
		"id": json.Number("1234"),
	}

	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
