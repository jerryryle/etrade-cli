package jsonmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonSlice_GetValue(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		testIndex   int
		expectErr   bool
		expectValue interface{}
	}{
		{
			name:        "GetValue Gets Value",
			testJson:    `["TestValue1"]`,
			testIndex:   0,
			expectErr:   false,
			expectValue: "TestValue1",
		},
		{
			name:        "GetValue Fails If Index Too Big",
			testJson:    `["TestValue1"]`,
			testIndex:   1,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "GetValue Fails If Index Negative",
			testJson:    `["TestValue1"]`,
			testIndex:   -1,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewJsonSliceFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := testSlice.GetValue(tt.testIndex)
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

func TestJsonSlice_GetType(t *testing.T) {
	type testFn func(s *JsonSlice) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetString Gets String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(0)
			},
			testJson:    `["StringValue"]`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetString Gets Null As Empty String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(0)
			},
			testJson:    `[null]`,
			expectErr:   false,
			expectValue: "",
		},
		{
			name: "GetString Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(1)
			},
			testJson:    `["StringValue"]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Fails On Non-String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(0)
			},
			testJson:    `[1234]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetInt Gets Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetInt Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(1)
			},
			testJson:    `[1234]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Fails On Non-Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `["1234"]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloat Gets Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(0)
			},
			testJson:    `[123.4]`,
			expectErr:   false,
			expectValue: 123.4,
		},
		{
			name: "GetFloat Gets Int As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(0)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: float64(1234),
		},
		{
			name: "GetFloat Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(1)
			},
			testJson:    `[123.4]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Fails On Non-Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(0)
			},
			testJson:    `["123.4"]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBool Gets Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `[true]`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBool Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(1)
			},
			testJson:    `[true]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Fails On Non-Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `["true"]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMap Gets Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(0)
			},
			testJson:    `[{"boo": "far"}]`,
			expectErr:   false,
			expectValue: JsonMap{"boo": "far"},
		},
		{
			name: "GetMap Gets Null As Nil JsonMap",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(0)
			},
			testJson:    `[null]`,
			expectErr:   false,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(1)
			},
			testJson:    `[{"boo": "far"}]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Fails On Non-Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(0)
			},
			testJson:    `["moo"]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSlice Gets Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(0)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSlice Gets Null As Nil Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(0)
			},
			testJson:    `[null]`,
			expectErr:   false,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(1)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Fails On Non-Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(0)
			},
			testJson:    `["foo"]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStrings Gets String Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStrings(0)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfStrings Gets Null Slice As Empty String Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStrings(0)
			},
			testJson:    `[null]`,
			expectErr:   false,
			expectValue: []string{},
		},
		{
			name: "GetSliceOfStrings Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStrings(1)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfStrings Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStrings(0)
			},
			testJson:    `[["foo", 1]]`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfInts Gets Int Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfInts(0)
			},
			testJson:    `[[1, 2]]`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfInts Gets Null Slice As Empty Int Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfInts(0)
			},
			testJson:    `[null]`,
			expectErr:   false,
			expectValue: []int64{},
		},
		{
			name: "GetSliceOfInts Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfInts(1)
			},
			testJson:    `[[1, 2]]`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfInts Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfInts(0)
			},
			testJson:    `[[1, "foo"]]`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloats Gets Float Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloats(0)
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfFloats Gets Null Slice As Empty Float Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloats(0)
			},
			testJson:    `[null]`,
			expectErr:   false,
			expectValue: []float64{},
		},
		{
			name: "GetSliceOfFloats Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloats(1)
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfFloats Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloats(0)
			},
			testJson:    `[[1.1, "foo"]]`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBools Gets Bool Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBools(0)
			},
			testJson:    `[[true, false]]`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfBools Gets Null Slice As Empty Bool Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBools(0)
			},
			testJson:    `[null]`,
			expectErr:   false,
			expectValue: []bool{},
		},
		{
			name: "GetSliceOfBools Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBools(1)
			},
			testJson:    `[[true, false]]`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfBools Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBools(0)
			},
			testJson:    `[[true, "foo"]]`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMaps Gets Map Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMaps(0)
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMaps Gets Null Slice As Empty Map Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMaps(0)
			},
			testJson:    `[null]`,
			expectErr:   false,
			expectValue: []JsonMap{},
		},
		{
			name: "GetSliceOfMaps Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMaps(1)
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfMaps Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMaps(0)
			},
			testJson:    `[[{"A": 1}, "foo"]]`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlices Gets Slice Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlices(0)
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
		{
			name: "GetSliceOfSlices Gets Null Slice As Empty Slice Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlices(0)
			},
			testJson:    `[ null ]`,
			expectErr:   false,
			expectValue: []JsonSlice{},
		},
		{
			name: "GetSliceOfSlices Fails On Bad Index",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlices(1)
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
		{
			name: "GetSliceOfSlices Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlices(0)
			},
			testJson:    `[ [ [1], "foo" ] ]`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewJsonSliceFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := tt.testFn(&testSlice)
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

func TestJsonSlice_GetValueWithDefault(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		testIndex   int
		testDefault interface{}
		expectValue interface{}
	}{
		{
			name:        "GetValueWithDefault Gets Value",
			testJson:    `["TestValue1"]`,
			testIndex:   0,
			testDefault: "DefaultValue",
			expectValue: "TestValue1",
		},
		{
			name:        "GetValueWithDefault Returns Default If Index Too Big",
			testJson:    `["TestValue1"]`,
			testIndex:   1,
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
		{
			name:        "GetValueWithDefault Returns Default If Index Negative",
			testJson:    `["TestValue1"]`,
			testDefault: "DefaultValue",
			testIndex:   -1,
			expectValue: "DefaultValue",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewJsonSliceFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue := testSlice.GetValueWithDefault(tt.testIndex, tt.testDefault)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestJsonSlice_GetTypeWithDefault(t *testing.T) {
	type testFn func(s *JsonSlice) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetStringWithDefault Gets String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringWithDefault(0, "DefaultValue")
			},
			testJson:    `["StringValue"]`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetStringWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringWithDefault(1, "DefaultValue")
			},
			testJson:    `["StringValue"]`,
			expectErr:   false,
			expectValue: "DefaultValue",
		},
		{
			name: "GetStringWithDefault Fails On Non-String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringWithDefault(0, "DefaultValue")
			},
			testJson:    `[1234]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetIntWithDefault Gets Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntWithDefault(0, 1)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetIntWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntWithDefault(1, 1)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: int64(1),
		},
		{
			name: "GetIntWithDefault Fails On Non-Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntWithDefault(0, 1)
			},
			testJson:    `["1234"]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloatWithDefault Gets Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatWithDefault(0, 1.1)
			},
			testJson:    `[123.4]`,
			expectErr:   false,
			expectValue: 123.4,
		},
		{
			name: "GetFloatWithDefault Gets Int As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatWithDefault(0, 1.1)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: float64(1234),
		},
		{
			name: "GetFloatWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatWithDefault(1, 1.1)
			},
			testJson:    `[123.4]`,
			expectErr:   false,
			expectValue: 1.1,
		},
		{
			name: "GetFloatWithDefault Fails On Non-Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatWithDefault(0, 1.1)
			},
			testJson:    `["123.4"]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBoolWithDefault Gets Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolWithDefault(0, false)
			},
			testJson:    `[true]`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBoolWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolWithDefault(1, false)
			},
			testJson:    `[true]`,
			expectErr:   false,
			expectValue: false,
		},
		{
			name: "GetBoolWithDefault Fails On Non-Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolWithDefault(0, false)
			},
			testJson:    `["true"]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMapWithDefault Gets Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapWithDefault(0, JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `[{"boo": "far"}]`,
			expectErr:   false,
			expectValue: JsonMap{"boo": "far"},
		},
		{
			name: "GetMapWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapWithDefault(1, JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `[{"boo": "far"}]`,
			expectErr:   false,
			expectValue: JsonMap{"DefaultKey": "DefaultValue"},
		},
		{
			name: "GetMapWithDefault Fails On Non-Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapWithDefault(0, JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `["moo"]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSliceWithDefault Gets Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceWithDefault(0, JsonSlice{"DefaultValue"})
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSliceWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceWithDefault(1, JsonSlice{"DefaultValue"})
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: JsonSlice{"DefaultValue"},
		},
		{
			name: "GetSliceWithDefault Fails On Non-Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceWithDefault(0, JsonSlice{"DefaultValue"})
			},
			testJson:    `["foo"]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStringsWithDefault Gets String Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsWithDefault(0, []string{"DefaultValue"})
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfStringsWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsWithDefault(1, []string{"DefaultValue"})
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: []string{"DefaultValue"},
		},
		{
			name: "GetSliceOfStringsWithDefault Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsWithDefault(0, []string{"DefaultValue"})
			},
			testJson:    `[["foo", 1]]`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfIntsWithDefault Gets Int Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsWithDefault(0, []int64{0})
			},
			testJson:    `[[1, 2]]`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfIntsWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsWithDefault(1, []int64{0})
			},
			testJson:    `[[1, 2]]`,
			expectErr:   false,
			expectValue: []int64{0},
		},
		{
			name: "GetSliceOfIntsWithDefault Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsWithDefault(0, []int64{0})
			},
			testJson:    `[[1, "foo"]]`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloatsWithDefault Gets Float Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsWithDefault(0, []float64{0.0})
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfFloatsWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsWithDefault(1, []float64{0.0})
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   false,
			expectValue: []float64{0.0},
		},
		{
			name: "GetSliceOfFloatsWithDefault Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsWithDefault(0, []float64{0.0})
			},
			testJson:    `[[1.1, "foo"]]`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBoolsWithDefault Gets Slice As Bool Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsWithDefault(0, []bool{false})
			},
			testJson:    `[[true, false]]`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfBoolsWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsWithDefault(1, []bool{false})
			},
			testJson:    `[[true, false]]`,
			expectErr:   false,
			expectValue: []bool{false},
		},
		{
			name: "GetSliceOfBoolsWithDefault Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsWithDefault(0, []bool{false})
			},
			testJson:    `[[true, "foo"]]`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMapsWithDefault Gets Map Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsWithDefault(0, []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMapsWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsWithDefault(1, []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   false,
			expectValue: []JsonMap{{"DefaultKey": "DefaultValue"}},
		},
		{
			name: "GetSliceOfMapsWithDefault Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsWithDefault(0, []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `[[{"A": 1}, "foo"]]`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlicesWithDefault Gets Slice Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesWithDefault(0, []JsonSlice{{json.Number("0")}})
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
		{
			name: "GetSliceOfSlicesWithDefault Returns Default If Index Out Of Bounds",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesWithDefault(1, []JsonSlice{{json.Number("0")}})
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("0")}},
		},
		{
			name: "GetSliceOfSlicesWithDefault Fails On Mixed Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesWithDefault(0, []JsonSlice{{json.Number("0")}})
			},
			testJson:    `[ [ [1], "foo" ] ]`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewJsonSliceFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := tt.testFn(&testSlice)
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

func TestJsonSlice_GetValueAtPath(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		testPath    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name:        "GetValueAtPath Returns Root For Empty Path",
			testJson:    `["TestStringValue1"]`,
			testPath:    "",
			expectErr:   false,
			expectValue: JsonSlice{"TestStringValue1"},
		},
		{
			name:        "GetValueAtPath Returns Root For Single Dot Path",
			testJson:    `["TestStringValue1"]`,
			testPath:    ".",
			expectErr:   false,
			expectValue: JsonSlice{"TestStringValue1"},
		},
		{
			name:        "GetValueAtPath Gets Value With Slice Indexing",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0].TestKey",
			expectErr:   false,
			expectValue: "TestValue",
		},
		{
			name:        "GetValueAtPath Gets Value If Path Has Extra Dots",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    ".[0]..TestKey.",
			expectErr:   false,
			expectValue: "TestValue",
		},
		{
			name:        "GetValueAtPath Fails If Slice Index Too Big",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[1].TestKey",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "GetValueAtPath Fails If Slice Index Negative",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[-1].TestKey",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "GetValueAtPath Gets Value With Nested Slice Indexing",
			testJson:    `[[{"TestKey": "TestValue"}]]`,
			testPath:    "[0][0].TestKey",
			expectErr:   false,
			expectValue: "TestValue",
		},
		{
			name:        "GetValueAtPath Fails For Slice Indexing on Non-Slice",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0][0]",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "GetValueAtPath Fails For Map Indexing on Non-Map",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0][0]",
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewJsonSliceFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := testSlice.GetValueAtPath(tt.testPath)
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

func TestJsonSlice_GetTypeAtPath(t *testing.T) {
	type testFn func(s *JsonSlice) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetStringAtPath Gets String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringAtPath("[0]")
			},
			testJson:    `["StringValue"]`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetStringAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringAtPath("[1]")
			},
			testJson:    `["StringValue"]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetIntAtPath Gets Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntAtPath("[0]")
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetIntAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntAtPath("[1]")
			},
			testJson:    `[1234]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloatAtPath Gets Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatAtPath("[0]")
			},
			testJson:    `[123.4]`,
			expectErr:   false,
			expectValue: 123.4,
		},
		{
			name: "GetFloatAtPath Gets Int As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatAtPath("[0]")
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: float64(1234),
		},
		{
			name: "GetFloatAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatAtPath("[1]")
			},
			testJson:    `[123.4]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBoolAtPath Gets Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolAtPath("[0]")
			},
			testJson:    `[true]`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBoolAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolAtPath("[1]")
			},
			testJson:    `[true]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMapAtPath Gets Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapAtPath("[0]")
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   false,
			expectValue: JsonMap{"foo": "bar"},
		},
		{
			name: "GetMapAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapAtPath("[1]")
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSliceAtPath Gets Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceAtPath("[0]")
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSliceAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceAtPath("[1]")
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStringsAtPath Gets String Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsAtPath("[0]")
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfStringsAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsAtPath("[1]")
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfIntsAtPath Gets Int Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsAtPath("[0]")
			},
			testJson:    `[[1, 2]]`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfIntsAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsAtPath("[1]")
			},
			testJson:    `[[1, 2]]`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloatsAtPath Gets Float Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsAtPath("[0]")
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfFloatsAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsAtPath("[1]")
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBoolsAtPath Gets Bool Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsAtPath("[0]")
			},
			testJson:    `[[true, false]]`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfBoolsAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsAtPath("[1]")
			},
			testJson:    `[[true, false]]`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMapsAtPath Gets Map Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsAtPath("[0]")
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMapsAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsAtPath("[1]")
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlicesAtPath Gets Slice Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesAtPath("[0]")
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
		{
			name: "GetSliceOfSlicesAtPath Fails On Bad Path Element",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesAtPath("[1]")
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewJsonSliceFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := tt.testFn(&testSlice)
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

func TestJsonSlice_GetValueAtPathWithDefault(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		testPath    string
		testDefault interface{}
		expectValue interface{}
	}{
		{
			name:        "GetValueAtPathWithDefault Returns Root For Empty Path",
			testJson:    `["TestStringValue1"]`,
			testPath:    "",
			testDefault: "DefaultValue",
			expectValue: JsonSlice{"TestStringValue1"},
		},
		{
			name:        "GetValueAtPathWithDefault Returns Root For Single Dot Path",
			testJson:    `["TestStringValue1"]`,
			testPath:    ".",
			testDefault: "DefaultValue",
			expectValue: JsonSlice{"TestStringValue1"},
		},
		{
			name:        "GetValueAtPathWithDefault Gets Value With Slice Indexing",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0].TestKey",
			testDefault: "DefaultValue",
			expectValue: "TestValue",
		},
		{
			name:        "GetValueAtPathWithDefault Gets Value If Path Has Extra Dots",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    ".[0]..TestKey.",
			testDefault: "DefaultValue",
			expectValue: "TestValue",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default If Slice Index Too Big",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[1].TestKey",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default If Slice Index Negative",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[-1].TestKey",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
		{
			name:        "GetValueAtPathWithDefault Gets Value With Nested Slice Indexing",
			testJson:    `[[{"TestKey": "TestValue"}]]`,
			testPath:    "[0][0].TestKey",
			testDefault: "DefaultValue",
			expectValue: "TestValue",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default For Slice Indexing on Non-Slice",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0][0]",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default For Map Indexing on Non-Map",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0][0]",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewJsonSliceFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue := testSlice.GetValueAtPathWithDefault(tt.testPath, tt.testDefault)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestJsonSlice_GetTypeAtPathWithDefault(t *testing.T) {
	type testFn func(s *JsonSlice) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetStringAtPathWithDefault Gets String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringAtPathWithDefault("[0]", "DefaultValue")
			},
			testJson:    `["StringValue"]`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetStringAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringAtPathWithDefault("[1]", "DefaultValue")
			},
			testJson:    `["StringValue"]`,
			expectErr:   false,
			expectValue: "DefaultValue",
		},
		{
			name: "GetStringAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringAtPathWithDefault("[0]", "DefaultValue")
			},
			testJson:    `[1]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetIntAtPathWithDefault Gets Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntAtPathWithDefault("[0]", 1)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetIntAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntAtPathWithDefault("[1]", 1)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: int64(1),
		},
		{
			name: "GetIntAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntAtPathWithDefault("[0]", 1)
			},
			testJson:    `["1234"]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloatAtPathWithDefault Gets Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatAtPathWithDefault("[0]", 1.1)
			},
			testJson:    `[123.4]`,
			expectErr:   false,
			expectValue: 123.4,
		},
		{
			name: "GetFloatAtPathWithDefault Gets Int As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatAtPathWithDefault("[0]", 1.1)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: float64(1234),
		},
		{
			name: "GetFloatAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatAtPathWithDefault("[1]", 1.1)
			},
			testJson:    `[123.4]`,
			expectErr:   false,
			expectValue: 1.1,
		},
		{
			name: "GetFloatAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloatAtPathWithDefault("[0]", 1.1)
			},
			testJson:    `["123.4"]`,
			expectErr:   true,
			expectValue: 0.0,
		},
		{
			name: "GetBoolAtPathWithDefault Gets Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolAtPathWithDefault("[0]", false)
			},
			testJson:    `[true]`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBoolAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolAtPathWithDefault("[1]", false)
			},
			testJson:    `[true]`,
			expectErr:   false,
			expectValue: false,
		},
		{
			name: "GetBoolAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolAtPathWithDefault("[0]", false)
			},
			testJson:    `["true"]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMapAtPathWithDefault Gets Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapAtPathWithDefault("[0]", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   false,
			expectValue: JsonMap{"foo": "bar"},
		},
		{
			name: "GetMapAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapAtPathWithDefault("[1]", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   false,
			expectValue: JsonMap{"DefaultKey": "DefaultValue"},
		},
		{
			name: "GetMapAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapAtPathWithDefault("[0]", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `["foobar"]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSliceAtPathWithDefault Gets Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceAtPathWithDefault("[0]", JsonSlice{"DefaultValue"})
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSliceAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceAtPathWithDefault("[1]", JsonSlice{"DefaultValue"})
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: JsonSlice{"DefaultValue"},
		},
		{
			name: "GetSliceAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceAtPathWithDefault("[0]", JsonSlice{"DefaultValue"})
			},
			testJson:    `["foobar"]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStringsAtPathWithDefault Gets String Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsAtPathWithDefault("[0]", []string{"DefaultValue"})
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfStringsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsAtPathWithDefault("[1]", []string{"DefaultValue"})
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: []string{"DefaultValue"},
		},
		{
			name: "GetSliceOfStringsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsAtPathWithDefault("[0]", []string{"DefaultValue"})
			},
			testJson:    `[[1, 2]]`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfIntsAtPathWithDefault Gets Int Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsAtPathWithDefault("[0]", []int64{0})
			},
			testJson:    `[[1, 2]]`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfIntsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsAtPathWithDefault("[1]", []int64{0})
			},
			testJson:    `[[1, 2]]`,
			expectErr:   false,
			expectValue: []int64{0},
		},
		{
			name: "GetSliceOfIntsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsAtPathWithDefault("[0]", []int64{0})
			},
			testJson:    `[["1", "2"]]`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloatsAtPathWithDefault Gets Float Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsAtPathWithDefault("[0]", []float64{0.0})
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfFloatsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsAtPathWithDefault("[1]", []float64{0.0})
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   false,
			expectValue: []float64{0.0},
		},
		{
			name: "GetSliceOfFloatsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsAtPathWithDefault("[0]", []float64{0.0})
			},
			testJson:    `[["1.1", "2.2"]]`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBoolsAtPathWithDefault Gets Bool Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsAtPathWithDefault("[0]", []bool{false})
			},
			testJson:    `[[true, false]]`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfBoolsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsAtPathWithDefault("[1]", []bool{false})
			},
			testJson:    `[[true, false]]`,
			expectErr:   false,
			expectValue: []bool{false},
		},
		{
			name: "GetSliceOfBoolsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsAtPathWithDefault("[0]", []bool{false})
			},
			testJson:    `[["true", "false"]]`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMapsAtPathWithDefault Gets Map Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsAtPathWithDefault("[0]", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMapsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsAtPathWithDefault("[1]", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   false,
			expectValue: []JsonMap{{"DefaultKey": "DefaultValue"}},
		},
		{
			name: "GetSliceOfMapsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsAtPathWithDefault("[0]", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `[["A", "B"]]`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlicesAtPathWithDefault Gets Slice Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesAtPathWithDefault("[0]", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
		{
			name: "GetSliceOfSlicesAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesAtPathWithDefault("[1]", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   false,
			expectValue: []JsonSlice{{"DefaultValue"}},
		},
		{
			name: "GetSliceOfSlicesAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesAtPathWithDefault("[0]", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `[ [ "1, 2" ] ]`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewJsonSliceFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := tt.testFn(&testSlice)
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
