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
]
`

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
		expectSlice JsonSlice
	}{
		{
			name: "New Slice From String",
			testFn: func() (JsonSlice, error) {
				return NewSliceFromJsonString(testValidJsonString)
			},
			expectErr:   false,
			expectSlice: testValidJsonExpectedSlice,
		},
		{
			name: "New Slice From Invalid String Returns Error",
			testFn: func() (JsonSlice, error) {
				return NewSliceFromJsonString(`[{"TestString": TestStringValue1]`)
			},
			expectErr:   true,
			expectSlice: nil,
		},
		{
			name: "New Slice From Top-Level Map String Returns Error",
			testFn: func() (JsonSlice, error) {
				// This string represents a top-level map, so it would need
				// NewMapFromJsonString() instead of NewSliceFromJsonString()
				return NewSliceFromJsonString(`{"TestKey":["TestValue"]}`)
			},
			expectErr:   true,
			expectSlice: nil,
		},
		{
			name: "New Slice From Bytes",
			testFn: func() (JsonSlice, error) {
				return NewSliceFromJsonBytes([]byte(testValidJsonString))
			},
			expectErr:   false,
			expectSlice: testValidJsonExpectedSlice,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				testResultSlice, err := tt.testFn()
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectSlice, testResultSlice)
			},
		)
	}
}
