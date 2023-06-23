package jsonmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonSlice_Map(t *testing.T) {
	type testFn func() JsonSlice

	var testJsonSlice = JsonSlice{
		JsonSlice{
			JsonMap{
				"Key 1": JsonMap{
					"Key 2": "Value 2",
				},
			},
			JsonMap{
				"Key 3": JsonMap{
					"Key 4": "Value 4",
				},
			},
		},
	}

	tests := []struct {
		name        string
		testFn      testFn
		expectValue JsonSlice
	}{
		{
			name: "Map Can Replace Values Within Slice",
			testFn: func() JsonSlice {
				replaceChildSliceValuesWithInt := func(
					elementPath []interface{}, ancestorSliceIndex int, index int, value interface{},
				) (interface{}, bool) {
					// Only replace values within a nested slice
					if ancestorSliceIndex >= 0 {
						// Replace the old value with an integer based on the current index
						return index, true
					}
					// Return the original value since we're not currently in a nested slice.
					return value, true
				}
				return testJsonSlice.Map(nil, replaceChildSliceValuesWithInt)
			},
			expectValue: JsonSlice{
				JsonSlice{0, 1},
			},
		},
		{
			name: "Map Can Drop Values From Slice",
			testFn: func() JsonSlice {
				replaceChildSliceValuesWithInt := func(
					elementPath []interface{}, ancestorSliceIndex int, index int, value interface{},
				) (interface{}, bool) {
					// Drop the second element from the slice
					if index == 1 {
						return nil, false
					}
					return value, true
				}
				return testJsonSlice.Map(nil, replaceChildSliceValuesWithInt)
			},
			expectValue: JsonSlice{
				JsonSlice{
					JsonMap{
						"Key 1": JsonMap{
							"Key 2": "Value 2",
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

func TestJsonSlice_ElementPath(t *testing.T) {
	var testJsonSlice = JsonSlice{
		JsonSlice{
			JsonMap{
				"Key 1": JsonMap{
					"Key 2": "Value 2",
				},
			},
			JsonMap{
				"Key 3": JsonMap{
					"Key 4": "Value 4",
				},
			},
		},
	}

	expectedMapElementPath := [][]interface{}{
		{0, 0},
		{0, 0, "Key 1"},
		{0, 1},
		{0, 1, "Key 3"},
	}

	expectedSliceElementPath := [][]interface{}{
		{0},
		{0, 0},
		{0, 1},
	}

	actualMapElementPath := make([][]interface{}, 0)
	actualSliceElementPath := make([][]interface{}, 0)

	recordMapElementPath := func(
		elementPath []interface{}, ancestorSliceIndex int, key string, value interface{},
	) (string, interface{}) {
		// Need to copy the element path because map will continue to update it
		// and if we store references to it, they'll be invalid later.
		elementPathCopy := make([]interface{}, len(elementPath))
		copy(elementPathCopy, elementPath)
		actualMapElementPath = append(actualMapElementPath, elementPathCopy)
		return key, value
	}

	recordSliceElementPath := func(
		elementPath []interface{}, ancestorSliceIndex int, index int, value interface{},
	) (interface{}, bool) {
		// Need to copy the element path because map will continue to update it
		// and if we store references to it, they'll be invalid later.
		elementPathCopy := make([]interface{}, len(elementPath))
		copy(elementPathCopy, elementPath)
		actualSliceElementPath = append(actualSliceElementPath, elementPathCopy)
		return value, true
	}

	_ = testJsonSlice.Map(recordMapElementPath, recordSliceElementPath)
	assert.Equal(t, expectedSliceElementPath, actualSliceElementPath)
	assert.Equal(t, expectedMapElementPath, actualMapElementPath)
}
