package jsonmap

import (
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

	var testValidJsonExpectedMap = map[string]interface{}{
		"TestMap": map[string]interface{}{
			"TestMap": map[string]interface{}{
				"TestSlice": []interface{}{
					map[string]interface{}{
						"TestString": "TestStringValue",
						"TestFloat":  123.456,
						"TestInt":    float64(123),
						"TestBool":   true,
					},
				},
			},
		},
	}

	const testInvalidJsonString = `
{
  "TestMap": {
}
`

	tests := []struct {
		name      string
		testFn    testFn
		expectErr bool
		expectMap JsonMap
	}{
		{
			name: "New Map From String",
			testFn: func() (JsonMap, error) {
				return NewFromJsonString(testValidJsonString)
			},
			expectErr: false,
			expectMap: testValidJsonExpectedMap,
		},
		{
			name: "New Map From Invalid String Returns Error",
			testFn: func() (JsonMap, error) {
				return NewFromJsonString(testInvalidJsonString)
			},
			expectErr: true,
			expectMap: nil,
		},
		{
			name: "New Map From Bytes",
			testFn: func() (JsonMap, error) {
				return NewFromJsonBytes([]byte(testValidJsonString))
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
