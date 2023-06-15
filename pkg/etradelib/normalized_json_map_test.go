package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNormalizedJsonMapFromJsonBytes(t *testing.T) {
	testJson := `
{
  "TestMap": {
    "TestMap": {
      "TestSlice": [
        {
          "TestValue": "TestStringValue"
        }
      ]
    }
  }
}
`
	expectedJsonMap := jsonmap.JsonMap(
		map[string]interface{}{
			"testMap": map[string]interface{}{
				"testMap": map[string]interface{}{
					"testSlice": []interface{}{
						map[string]interface{}{
							"testValue": "TestStringValue",
						},
					},
				},
			},
		},
	)

	resultJsonMap, err := NewNormalizedJsonMapFromJsonBytes([]byte(testJson))
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonMap, resultJsonMap)
}
