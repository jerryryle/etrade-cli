package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradePortfolioPosition(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradePosition
	}{
		{
			name: "Creates PortfolioPosition",
			testJson: `
{
  "positionId": 1234
}`,
			expectErr: false,
			expectValue: &eTradePosition{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"positionId": json.Number("1234"),
				},
			},
		},
		{
			name: "Fails If Missing positionId",
			testJson: `
{
  "MISSING": 1234
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
				actualValue, err := CreateETradePosition(responseMap)
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

func TestETradePortfolioPosition_GetId(t *testing.T) {
	testObject := &eTradePosition{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"positionId": json.Number("1234"),
		},
	}
	expectedValue := int64(1234)

	actualValue := testObject.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradePortfolioPosition_AddLotsFromResponse(t *testing.T) {
	startingObject := &eTradePosition{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"positionId": json.Number("1234"),
		},
	}

	tests := []struct {
		name        string
		startValue  ETradePosition
		testJson    string
		expectErr   bool
		expectValue ETradePosition
	}{
		{
			name:       "Can Add Lots",
			startValue: startingObject,
			testJson: `
{
  "PositionLotsResponse": {
    "PositionLot": [
      {
        "Key": "Value"
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradePosition{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"positionId": json.Number("1234"),
					"lots": jsonmap.JsonSlice{
						jsonmap.JsonMap{
							"key": "Value",
						},
					},
				},
			},
		},
		{
			name:       "Fails With Invalid JSON",
			startValue: startingObject,
			testJson: `
{
  "PositionLotsResponse": {
}`,
			expectErr:   true,
			expectValue: startingObject,
		},
		{
			name:       "Fails With Missing PositionLot",
			startValue: startingObject,
			testJson: `
{
  "PositionLotsResponse": {
    "MISSING": [
      {
        "Key": "Value"
      }
    ]
  }
}`,
			expectErr:   true,
			expectValue: startingObject,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var position = tt.startValue
				// Call the Method Under Test
				err := position.AddLotsFromResponse([]byte(tt.testJson))
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, position)
			},
		)
	}
}

func TestETradePortfolioPosition_AsJsonMap(t *testing.T) {
	testObject := &eTradePosition{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"positionId": json.Number("1234"),
		},
	}
	expectedValue := jsonmap.JsonMap{
		"positionId": json.Number("1234"),
	}

	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
