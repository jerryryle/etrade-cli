package jsonmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestJsonMap_Map(t *testing.T) {
	type testFn func() JsonMap

	var testJsonMap = JsonMap{
		"Key 1": JsonMap{
			"Key 2": JsonSlice{
				JsonSlice{
					JsonMap{
						"Key 3": "Value 3",
					},
					JsonMap{
						"Key 4": "Value 4",
					},
				},
			},
		},
	}

	tests := []struct {
		name        string
		testFn      testFn
		expectValue JsonMap
	}{
		{
			name: "Map Can Replace Map Keys",
			testFn: func() JsonMap {
				upperCaseKeys := func(
					elementPath []interface{}, ancestorSliceIndex int, key string, value interface{},
				) (string, interface{}) {
					return strings.ToUpper(key), value
				}
				return testJsonMap.Map(upperCaseKeys, nil)
			},
			expectValue: JsonMap{
				"KEY 1": JsonMap{
					"KEY 2": JsonSlice{
						JsonSlice{
							JsonMap{
								"KEY 3": "Value 3",
							},
							JsonMap{
								"KEY 4": "Value 4",
							},
						},
					},
				},
			},
		},
		{
			name: "Map Can Replace Map Values",
			testFn: func() JsonMap {
				replaceMapStringValuesInSlice := func(
					elementPath []interface{}, ancestorSliceIndex int, key string, value interface{},
				) (
					string, interface{},
				) {
					if ancestorSliceIndex < 0 {
						// Only replace strings in objects within a slice.
						return key, value
					}
					switch value.(type) {
					case string:
						return key, fmt.Sprintf("New String %d", ancestorSliceIndex)
					default:
						return key, value
					}
				}
				return testJsonMap.Map(replaceMapStringValuesInSlice, nil)
			},
			expectValue: JsonMap{
				"Key 1": JsonMap{
					"Key 2": JsonSlice{
						JsonSlice{
							JsonMap{
								"Key 3": "New String 0",
							},
							JsonMap{
								"Key 4": "New String 1",
							},
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
				actualValue := tt.testFn()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestJsonMap_ElementPath(t *testing.T) {
	var testJsonMap = JsonMap{
		"Key1": JsonMap{
			"Key2": JsonSlice{
				JsonSlice{
					JsonMap{
						"Key3": "Value3",
					},
					JsonMap{
						"Key4": "Value4",
					},
				},
			},
		},
	}
	expectedElementPath := [][]interface{}{
		{},
		{"Key1"},
		{"Key1", "Key2", 0, 0},
		{"Key1", "Key2", 0, 1},
	}

	actualElementPath := make([][]interface{}, 0)

	recordElementPath := func(
		elementPath []interface{}, ancestorSliceIndex int, key string, value interface{},
	) (string, interface{}) {
		// Need to copy the element path because map will continue to update it
		// and if we store references to it, they'll be invalid later.
		elementPathCopy := make([]interface{}, len(elementPath))
		copy(elementPathCopy, elementPath)
		actualElementPath = append(actualElementPath, elementPathCopy)
		return key, value
	}
	_ = testJsonMap.Map(recordElementPath, nil)
	assert.Equal(t, expectedElementPath, actualElementPath)
}
