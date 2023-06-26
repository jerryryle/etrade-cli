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
		name        string
		testFn      testFn
		expectErr   bool
		expectValue JsonMap
	}{
		{
			name: "New Map From String",
			testFn: func() (JsonMap, error) {
				return NewMapFromJsonString(testValidJsonString)
			},
			expectErr:   false,
			expectValue: testValidJsonExpectedMap,
		},
		{
			name: "New Map From Invalid String Fails",
			testFn: func() (JsonMap, error) {
				return NewMapFromJsonString(`{"TestMap": {}`)
			},
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "New Map From Top-Level Slice String Fails",
			testFn: func() (JsonMap, error) {
				// This string represents a top-level slice, so it would need
				// NewSliceFromJsonString() instead of NewMapFromJsonString()
				return NewMapFromJsonString(`[{"TestKey":"TestValue"}]`)
			},
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "New Map From Bytes",
			testFn: func() (JsonMap, error) {
				return NewMapFromJsonBytes([]byte(testValidJsonString))
			},
			expectErr:   false,
			expectValue: testValidJsonExpectedMap,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := tt.testFn()
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

func TestJsonMap_ToJsonString(t *testing.T) {
	var testJsonMap = JsonMap{
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

	const expectedJsonStringPretty = `{
  "TestMap": {
    "TestMap": {
      "TestSlice": [
        {
          "TestBool": true,
          "TestFloat": 123.456,
          "TestInt": 123,
          "TestString": "TestStringValue"
        }
      ]
    }
  }
}
`

	const expectedJsonStringUgly = `{"TestMap":{"TestMap":{"TestSlice":[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue"}]}}}
`

	actualValue, err := testJsonMap.ToJsonString(true)
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonStringPretty, actualValue)

	actualValue, err = testJsonMap.ToJsonString(false)
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonStringUgly, actualValue)
}

func TestJsonMap_ToJsonBytes(t *testing.T) {
	var testJsonMap = JsonMap{
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

	const expectedJsonStringPretty = `{
  "TestMap": {
    "TestMap": {
      "TestSlice": [
        {
          "TestBool": true,
          "TestFloat": 123.456,
          "TestInt": 123,
          "TestString": "TestStringValue"
        }
      ]
    }
  }
}
`

	const expectedJsonStringUgly = `{"TestMap":{"TestMap":{"TestSlice":[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue"}]}}}
`

	actualValue, err := testJsonMap.ToJsonBytes(true)
	assert.Nil(t, err)
	assert.Equal(t, []byte(expectedJsonStringPretty), actualValue)

	actualValue, err = testJsonMap.ToJsonBytes(false)
	assert.Nil(t, err)
	assert.Equal(t, []byte(expectedJsonStringUgly), actualValue)
}
