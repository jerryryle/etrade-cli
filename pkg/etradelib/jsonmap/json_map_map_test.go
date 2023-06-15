package jsonmap

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestJsonMap_Map(t *testing.T) {
	type testFn func() (JsonMap, error)

	const testJsonString = `
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

	tests := []struct {
		name      string
		testFn    testFn
		expectErr bool
		expectMap JsonMap
	}{
		{
			name: "Map Recursively Applies to Maps and Slices",
			testFn: func() (JsonMap, error) {
				upperCaseKeys := func(key string, value interface{}) (string, interface{}) {
					return strings.ToUpper(key), value
				}
				jmap, err := NewFromJsonString(testJsonString)
				jmap = jmap.Map(upperCaseKeys)
				if err != nil {
					return nil, err
				}
				return jmap, nil
			},
			expectErr: false,
			expectMap: map[string]interface{}{
				"TESTMAP": map[string]interface{}{
					"TESTMAP": map[string]interface{}{
						"TESTSLICE": []interface{}{
							map[string]interface{}{
								"TESTSTRING": "TestStringValue",
								"TESTFLOAT":  123.456,
								"TESTINT":    float64(123),
								"TESTBOOL":   true,
							},
						},
					},
				},
			},
		},
		{
			name: "Map Can Replace Values",
			testFn: func() (JsonMap, error) {
				replaceSliceWithInt := func(key string, value interface{}) (string, interface{}) {
					switch value.(type) {
					case []interface{}:
						return key, 1234
					default:
						return key, value
					}
				}
				jmap, err := NewFromJsonString(testJsonString)
				jmap = jmap.Map(replaceSliceWithInt)
				if err != nil {
					return nil, err
				}
				return jmap, nil
			},
			expectErr: false,
			expectMap: map[string]interface{}{
				"TestMap": map[string]interface{}{
					"TestMap": map[string]interface{}{
						"TestSlice": 1234,
					},
				},
			},
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
