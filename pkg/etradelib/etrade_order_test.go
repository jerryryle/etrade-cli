package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeOrder(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeOrder
	}{
		{
			name: "CreateETradeOrder Creates Order With Valid Response",
			testJson: `
{
  "orderId": 1234
}`,
			expectErr: false,
			expectValue: &eTradeOrder{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"orderId": json.Number("1234"),
				},
			},
		},
		{
			name: "CreateETradeOrder Fails If Missing Order ID",
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
				actualValue, err := CreateETradeOrder(responseMap)
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

func TestETradeOrder_GetId(t *testing.T) {
	testOrder := &eTradeOrder{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"orderId": json.Number("1234"),
		},
	}
	expectedValue := int64(1234)

	actualValue := testOrder.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeOrder_GetJsonMap(t *testing.T) {
	testOrder := &eTradeOrder{
		id: 1234,
		jsonMap: jsonmap.JsonMap{
			"orderId": json.Number("1234"),
		},
	}
	expectedValue := jsonmap.JsonMap{
		"orderId": json.Number("1234"),
	}

	actualValue := testOrder.GetJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
