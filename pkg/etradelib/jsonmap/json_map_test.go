package jsonmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMap_New(t *testing.T) {
	type testFn func() (JsonMap, error)

	const testValidJsonString = `
{
  "TestMap": {
    "TestMap": {
      "TestSlice": [
        {
          "TestString": "TestStringValue",
          "TestFloat": 123.456,
          "TestInt": 123,
          "TestBool": true
        }
      ]
    }
  }
}
`

	var testValidJsonExpectedMap = JsonMap{
		"TestMap": JsonMap{
			"TestMap": JsonMap{
				"TestSlice": JsonSlice{
					JsonMap{
						"TestString": "TestStringValue",
						"TestFloat":  json.Number("123.456"),
						"TestInt":    json.Number("123"),
						"TestBool":   true,
					},
				},
			},
		},
	}

	tests := []struct {
		name      string
		testFn    testFn
		expectErr bool
		expectMap JsonMap
	}{
		{
			name: "New Map From String",
			testFn: func() (JsonMap, error) {
				return NewMapFromJsonString(testValidJsonString)
			},
			expectErr: false,
			expectMap: testValidJsonExpectedMap,
		},
		{
			name: "New Map From Invalid String Returns Error",
			testFn: func() (JsonMap, error) {
				return NewMapFromJsonString(`{"TestMap": {}`)
			},
			expectErr: true,
			expectMap: nil,
		},
		{
			name: "New Map From Top-Level Slice String Returns Error",
			testFn: func() (JsonMap, error) {
				// This string represents a top-level slice, so it would need
				// NewSliceFromJsonString() instead of NewMapFromJsonString()
				return NewMapFromJsonString(`[{"TestKey":"TestValue"}]`)
			},
			expectErr: true,
			expectMap: nil,
		},
		{
			name: "New Map From Bytes",
			testFn: func() (JsonMap, error) {
				return NewMapFromJsonBytes([]byte(testValidJsonString))
			},
			expectErr: false,
			expectMap: testValidJsonExpectedMap,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				testResultMap, err := tt.testFn()
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
					assert.Equal(t, tt.expectMap, testResultMap)
				}
			},
		)
	}
}
