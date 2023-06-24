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
			name: "CreateETradePosition Creates PortfolioPosition With Valid Response",
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
			name: "CreateETradePosition Fails If Missing PortfolioPosition ID",
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
	testPortfolioPosition := &eTradePosition{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"positionId": json.Number("1234"),
		},
	}
	expectedValue := int64(1234)

	actualValue := testPortfolioPosition.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradePortfolioPosition_AsJsonMap(t *testing.T) {
	testPortfolioPosition := &eTradePosition{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"positionId": json.Number("1234"),
		},
	}
	expectedValue := jsonmap.JsonMap{
		"positionId": json.Number("1234"),
	}

	actualValue := testPortfolioPosition.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
