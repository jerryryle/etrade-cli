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
	expectedJsonMap := jsonmap.JsonMap{
		"testMap": jsonmap.JsonMap{
			"testMap": jsonmap.JsonMap{
				"testSlice": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"testValue": "TestStringValue",
					},
				},
			},
		},
	}

	resultJsonMap, err := NewNormalizedJsonMap([]byte(testJson))
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonMap, resultJsonMap)
}
