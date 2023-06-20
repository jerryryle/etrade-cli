package jsonmap

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestJsonMap_Map(t *testing.T) {
	type testFn func() (JsonMap, error)

	const testJsonString = `
{
  "TestMap": {
    "TestNestedMap": {
      "TestKey": "TestStringValue",
      "TestNestedSlice": [
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
				upperCaseKeys := func(parentSliceIndex int, key string, value interface{}) (string, interface{}) {
					return strings.ToUpper(key), value
				}
				jmap, err := NewMapFromJsonString(testJsonString)
				jmap = jmap.Map(upperCaseKeys, nil)
				if err != nil {
					return nil, err
				}
				return jmap, nil
			},
			expectErr: false,
			expectMap: JsonMap{
				"TESTMAP": JsonMap{
					"TESTNESTEDMAP": JsonMap{
						"TESTKEY": "TestStringValue",
						"TESTNESTEDSLICE": JsonSlice{
							JsonSlice{
								JsonMap{
									"TESTSTRING": "TestStringValue1",
									"TESTFLOAT":  json.Number("123.456"),
									"TESTINT":    json.Number("123"),
									"TESTBOOL":   true,
								},
								JsonMap{
									"TESTSTRING": "TestStringValue2",
									"TESTFLOAT":  json.Number("456.789"),
									"TESTINT":    json.Number("456"),
									"TESTBOOL":   false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Map Can Replace Map Values",
			testFn: func() (JsonMap, error) {
				replaceMapStringValuesInSlice := func(parentSliceIndex int, key string, value interface{}) (
					string, interface{},
				) {
					if parentSliceIndex < 0 {
						// Only replace strings in objects within a slice.
						return key, value
					}
					switch value.(type) {
					case string:
						return key, fmt.Sprintf("New String %d", parentSliceIndex)
					default:
						return key, value
					}
				}
				jmap, err := NewMapFromJsonString(testJsonString)
				jmap = jmap.Map(replaceMapStringValuesInSlice, nil)
				if err != nil {
					return nil, err
				}
				return jmap, nil
			},
			expectErr: false,
			expectMap: JsonMap{
				"TestMap": JsonMap{
					"TestNestedMap": JsonMap{
						"TestKey": "TestStringValue",
						"TestNestedSlice": JsonSlice{
							JsonSlice{
								JsonMap{
									"TestString": "New String 0",
									"TestFloat":  json.Number("123.456"),
									"TestInt":    json.Number("123"),
									"TestBool":   true,
								},
								JsonMap{
									"TestString": "New String 1",
									"TestFloat":  json.Number("456.789"),
									"TestInt":    json.Number("456"),
									"TestBool":   false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Map Can Replace Slice Values",
			testFn: func() (JsonMap, error) {
				replaceChildSliceValuesWithInt := func(
					parentSliceIndex int, index int, value interface{},
				) interface{} {
					if parentSliceIndex >= 0 {
						// Replace the old value with an integer based on the current index
						return index
					}
					// Return the original value since we're not currently in a slice.
					return value
				}
				jmap, err := NewMapFromJsonString(testJsonString)
				jmap = jmap.Map(nil, replaceChildSliceValuesWithInt)
				if err != nil {
					return nil, err
				}
				return jmap, nil
			},
			expectErr: false,
			expectMap: JsonMap{
				"TestMap": JsonMap{
					"TestNestedMap": JsonMap{
						"TestKey": "TestStringValue",
						"TestNestedSlice": JsonSlice{
							JsonSlice{0, 1},
						},
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
				}
				assert.Equal(t, tt.expectMap, testResultMap)
			},
		)
	}
}
