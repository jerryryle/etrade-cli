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
}`

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
				return NewJsonMapFromJsonString(testValidJsonString)
			},
			expectErr:   false,
			expectValue: testValidJsonExpectedMap,
		},
		{
			name: "New Map From Invalid String Fails",
			testFn: func() (JsonMap, error) {
				return NewJsonMapFromJsonString(`{"TestMap": {}`)
			},
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "New Map From Top-Level Slice String Fails",
			testFn: func() (JsonMap, error) {
				// This string represents a top-level slice, so it would need
				// NewJsonSliceFromJsonString() instead of NewJsonMapFromJsonString()
				return NewJsonMapFromJsonString(`[{"TestKey":"TestValue"}]`)
			},
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "New Map From Bytes",
			testFn: func() (JsonMap, error) {
				return NewJsonMapFromJsonBytes([]byte(testValidJsonString))
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
						"TestUrl":    "https://moo.com?foo=1&loo=2",
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
          "TestString": "TestStringValue",
          "TestUrl": "https://moo.com?foo=1&loo=2"
        }
      ]
    }
  }
}` + "\n"

	const expectedJsonStringUgly = `{"TestMap":{"TestMap":{"TestSlice":[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue","TestUrl":"https://moo.com?foo=1&loo=2"}]}}}` + "\n"

	const expectedJsonStringEscapeHtml = `{"TestMap":{"TestMap":{"TestSlice":[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue","TestUrl":"https://moo.com?foo=1\u0026loo=2"}]}}}` + "\n"

	actualValue, err := testJsonMap.ToJsonString(true, false)
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonStringPretty, actualValue)

	actualValue, err = testJsonMap.ToJsonString(false, false)
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonStringUgly, actualValue)

	actualValue, err = testJsonMap.ToJsonString(false, true)
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonStringEscapeHtml, actualValue)
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
						"TestUrl":    "https://moo.com?foo=1&loo=2",
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
          "TestString": "TestStringValue",
          "TestUrl": "https://moo.com?foo=1&loo=2"
        }
      ]
    }
  }
}` + "\n"

	const expectedJsonStringUgly = `{"TestMap":{"TestMap":{"TestSlice":[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue","TestUrl":"https://moo.com?foo=1&loo=2"}]}}}` + "\n"

	const expectedJsonStringEscapeHtml = `{"TestMap":{"TestMap":{"TestSlice":[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue","TestUrl":"https://moo.com?foo=1\u0026loo=2"}]}}}` + "\n"

	actualValue, err := testJsonMap.ToJsonBytes(true, false)
	assert.Nil(t, err)
	assert.Equal(t, []byte(expectedJsonStringPretty), actualValue)

	actualValue, err = testJsonMap.ToJsonBytes(false, false)
	assert.Nil(t, err)
	assert.Equal(t, []byte(expectedJsonStringUgly), actualValue)

	actualValue, err = testJsonMap.ToJsonBytes(false, true)
	assert.Nil(t, err)
	assert.Equal(t, []byte(expectedJsonStringEscapeHtml), actualValue)
}
