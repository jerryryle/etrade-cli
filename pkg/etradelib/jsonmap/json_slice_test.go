package jsonmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonSlice_New(t *testing.T) {
	type testFn func() (JsonSlice, error)

	const testValidJsonString = `
[
  {
    "TestString": "TestStringValue1",
    "TestFloat": 123.456,
    "TestInt": 123,
    "TestBool": true
  },
  {
    "TestString": "TestStringValue2",
    "TestFloat": 456.789,
    "TestInt": 456,
    "TestBool": false
  }
]`

	var testValidJsonExpectedSlice = JsonSlice{
		JsonMap{
			"TestString": "TestStringValue1",
			"TestFloat":  json.Number("123.456"),
			"TestInt":    json.Number("123"),
			"TestBool":   true,
		},
		JsonMap{
			"TestString": "TestStringValue2",
			"TestFloat":  json.Number("456.789"),
			"TestInt":    json.Number("456"),
			"TestBool":   false,
		},
	}

	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue JsonSlice
	}{
		{
			name: "New Slice From String",
			testFn: func() (JsonSlice, error) {
				return NewSliceFromJsonString(testValidJsonString)
			},
			expectErr:   false,
			expectValue: testValidJsonExpectedSlice,
		},
		{
			name: "New Slice From Invalid String Fails",
			testFn: func() (JsonSlice, error) {
				return NewSliceFromJsonString(`[{"TestString": TestStringValue1]`)
			},
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "New Slice From Top-Level Map String Fails",
			testFn: func() (JsonSlice, error) {
				// This string represents a top-level map, so it would need
				// NewMapFromJsonString() instead of NewSliceFromJsonString()
				return NewSliceFromJsonString(`{"TestKey":["TestValue"]}`)
			},
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "New Slice From Bytes",
			testFn: func() (JsonSlice, error) {
				return NewSliceFromJsonBytes([]byte(testValidJsonString))
			},
			expectErr:   false,
			expectValue: testValidJsonExpectedSlice,
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

func TestJsonSlice_ToJsonString(t *testing.T) {
	var testJsonSlice = JsonSlice{
		JsonMap{
			"TestString": "TestStringValue1",
			"TestFloat":  json.Number("123.456"),
			"TestInt":    json.Number("123"),
			"TestBool":   true,
		},
		JsonMap{
			"TestString": "TestStringValue2",
			"TestFloat":  json.Number("456.789"),
			"TestInt":    json.Number("456"),
			"TestBool":   false,
			"TestUrl":    "https://moo.com?foo=1&loo=2",
		},
	}

	const expectedJsonStringPretty = `[
  {
    "TestBool": true,
    "TestFloat": 123.456,
    "TestInt": 123,
    "TestString": "TestStringValue1"
  },
  {
    "TestBool": false,
    "TestFloat": 456.789,
    "TestInt": 456,
    "TestString": "TestStringValue2",
    "TestUrl": "https://moo.com?foo=1&loo=2"
  }
]` + "\n"

	const expectedJsonStringUgly = `[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue1"},{"TestBool":false,"TestFloat":456.789,"TestInt":456,"TestString":"TestStringValue2","TestUrl":"https://moo.com?foo=1&loo=2"}]` + "\n"

	const expectedJsonStringEscapeHtml = `[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue1"},{"TestBool":false,"TestFloat":456.789,"TestInt":456,"TestString":"TestStringValue2","TestUrl":"https://moo.com?foo=1\u0026loo=2"}]` + "\n"

	actualValue, err := testJsonSlice.ToJsonString(true, false)
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonStringPretty, actualValue)

	actualValue, err = testJsonSlice.ToJsonString(false, false)
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonStringUgly, actualValue)

	actualValue, err = testJsonSlice.ToJsonString(false, true)
	assert.Nil(t, err)
	assert.Equal(t, expectedJsonStringEscapeHtml, actualValue)
}

func TestJsonSlice_ToJsonBytes(t *testing.T) {
	var testJsonSlice = JsonSlice{
		JsonMap{
			"TestString": "TestStringValue1",
			"TestFloat":  json.Number("123.456"),
			"TestInt":    json.Number("123"),
			"TestBool":   true,
		},
		JsonMap{
			"TestString": "TestStringValue2",
			"TestFloat":  json.Number("456.789"),
			"TestInt":    json.Number("456"),
			"TestBool":   false,
			"TestUrl":    "https://moo.com?foo=1&loo=2",
		},
	}

	const expectedJsonStringPretty = `[
  {
    "TestBool": true,
    "TestFloat": 123.456,
    "TestInt": 123,
    "TestString": "TestStringValue1"
  },
  {
    "TestBool": false,
    "TestFloat": 456.789,
    "TestInt": 456,
    "TestString": "TestStringValue2",
    "TestUrl": "https://moo.com?foo=1&loo=2"
  }
]` + "\n"

	const expectedJsonStringUgly = `[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue1"},{"TestBool":false,"TestFloat":456.789,"TestInt":456,"TestString":"TestStringValue2","TestUrl":"https://moo.com?foo=1&loo=2"}]` + "\n"

	const expectedJsonStringEscapeHtml = `[{"TestBool":true,"TestFloat":123.456,"TestInt":123,"TestString":"TestStringValue1"},{"TestBool":false,"TestFloat":456.789,"TestInt":456,"TestString":"TestStringValue2","TestUrl":"https://moo.com?foo=1\u0026loo=2"}]` + "\n"

	actualValue, err := testJsonSlice.ToJsonBytes(true, false)
	assert.Nil(t, err)
	assert.Equal(t, []byte(expectedJsonStringPretty), actualValue)

	actualValue, err = testJsonSlice.ToJsonBytes(false, false)
	assert.Nil(t, err)
	assert.Equal(t, []byte(expectedJsonStringUgly), actualValue)

	actualValue, err = testJsonSlice.ToJsonBytes(false, true)
	assert.Nil(t, err)
	assert.Equal(t, []byte(expectedJsonStringEscapeHtml), actualValue)
}
