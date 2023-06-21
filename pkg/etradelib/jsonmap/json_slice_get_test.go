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
			name:        "Index Returns Value",
			testJson:    `["TestValue1"]`,
			testIndex:   0,
			expectErr:   false,
			expectValue: "TestValue1",
		},
		{
			name:        "Index Too Big Returns Error",
			testJson:    `["TestValue1"]`,
			testIndex:   1,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Index Negative Returns Error",
			testJson:    `["TestValue1"]`,
			testIndex:   -1,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewSliceFromJsonString(tt.testJson)
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
			name: "GetString Gets String As String",
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
			name: "GetString Cannot Get Int As String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(0)
			},
			testJson:    `[1234]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Float As String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(0)
			},
			testJson:    `[1234.567]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Bool As String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(0)
			},
			testJson:    `[true]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Map As String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(0)
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Slice As String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetString(0)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetInt Gets Int As Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetInt Cannot Get String As Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `["1234"]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Float As Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `[123.4]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Bool As Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `[true]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Null As Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `[null]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Map As Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Slice As Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetInt(0)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloat Gets Float As Float",
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
			name: "GetFloat Cannot Get String As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(0)
			},
			testJson:    `["123.4"]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Bool As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(0)
			},
			testJson:    `[true]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Null As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(0)
			},
			testJson:    `[null]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Map As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(0)
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Slice As Float",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetFloat(0)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBool Gets Bool As Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `[true]`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBool Cannot Get String As Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `["true"]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Int As Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `[1234]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Float As Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `[1234.5678]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Null As Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `[null]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Map As Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Slice As Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBool(0)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMap Gets Map as Map",
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
			name: "GetMap Cannot Get String as Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(0)
			},
			testJson:    `["moo"]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Int as Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(0)
			},
			testJson:    `[1234]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Float as Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(0)
			},
			testJson:    `[123.4]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Bool as Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(0)
			},
			testJson:    `[true]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Slice as Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMap(0)
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSlice Gets Slice As Slice",
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
			name: "GetSlice Cannot Get String As Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(0)
			},
			testJson:    `["foo"]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Int As Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(0)
			},
			testJson:    `[1234]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Float As Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(0)
			},
			testJson:    `[123.4]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Bool As Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(0)
			},
			testJson:    `[true]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Map As Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSlice(0)
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStrings Gets Slice As String Slice",
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
			name: "GetSliceOfStrings Cannot Get Mixed Slice As String Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStrings(0)
			},
			testJson:    `[["foo", 1]]`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfInts Gets Slice As Int Slice",
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
			name: "GetSliceOfInts Cannot Get Mixed Slice As Int Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfInts(0)
			},
			testJson:    `[[1, "foo"]]`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloats Gets Slice As Float Slice",
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
			name: "GetSliceOfFloats Cannot Get Mixed Slice As Float Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloats(0)
			},
			testJson:    `[[1.1, "foo"]]`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBools Gets Slice As Bool Slice",
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
			name: "GetSliceOfBools Cannot Get Mixed Slice As Bool Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBools(0)
			},
			testJson:    `[[true, "foo"]]`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMaps Gets Slice As Map Slice",
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
			name: "GetSliceOfMaps Cannot Get Mixed Slice As Map Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMaps(0)
			},
			testJson:    `[[{"A": 1}, "foo"]]`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlices Gets Slice As Slice Slice",
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
			name: "GetSliceOfSlices Cannot Get Mixed Slice As Slice Slice",
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
				testSlice, err := NewSliceFromJsonString(tt.testJson)
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
			name:        "Empty Path Returns Root",
			testJson:    `["TestStringValue1"]`,
			testPath:    "",
			expectErr:   false,
			expectValue: JsonSlice{"TestStringValue1"},
		},
		{
			name:        "Single Dot Path Returns Root",
			testJson:    `["TestStringValue1"]`,
			expectErr:   false,
			expectValue: JsonSlice{"TestStringValue1"},
		},
		{
			name:        "Slice Indexing Returns Value",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0].TestKey",
			expectErr:   false,
			expectValue: "TestValue",
		},
		{
			name:        "Path That Has Extra Dots Still Returns Value",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    ".[0]..TestKey.",
			expectErr:   false,
			expectValue: "TestValue",
		},
		{
			name:        "Array Index Too Big Returns Error",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[1].TestKey",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Array Index Negative Returns Error",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[-1].TestKey",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Nested Slice Indexing Returns Value",
			testJson:    `[[{"TestKey": "TestValue"}]]`,
			testPath:    "[0][0].TestKey",
			expectErr:   false,
			expectValue: "TestValue",
		},
		{
			name:        "Array Indexing on Non-Array Returns Error",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0][0]",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Map Indexing on Non-Map Returns Error",
			testJson:    `[{"TestKey": "TestValue"}]`,
			testPath:    "[0][0]",
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewSliceFromJsonString(tt.testJson)
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
			name: "GetStringAtPath Gets String As String",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetStringAtPath("[0]")
			},
			testJson:    `["StringValue"]`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetIntAtPath Gets Int As Int",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetIntAtPath("[0]")
			},
			testJson:    `[1234]`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetFloatAtPath Gets Float As Float",
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
			name: "GetBoolAtPath Gets Bool As Bool",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetBoolAtPath("[0]")
			},
			testJson:    `[true]`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetMapAtPath Gets Map As Map",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetMapAtPath("[0]")
			},
			testJson:    `[{"foo": "bar"}]`,
			expectErr:   false,
			expectValue: JsonMap{"foo": "bar"},
		},
		{
			name: "GetSliceAtPath Gets Slice As Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceAtPath("[0]")
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSliceOfStringsAtPath Gets Slice As String Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfStringsAtPath("[0]")
			},
			testJson:    `[["foo", "bar"]]`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfIntsAtPath Gets Slice As Int Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfIntsAtPath("[0]")
			},
			testJson:    `[[1, 2]]`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfFloatsAtPath Gets Slice As Float Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfFloatsAtPath("[0]")
			},
			testJson:    `[[1.1, 2.2]]`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfBoolsAtPath Gets Slice As Bool Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfBoolsAtPath("[0]")
			},
			testJson:    `[[true, false]]`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfMapsAtPath Gets Slice As Map Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfMapsAtPath("[0]")
			},
			testJson:    `[[{"A": 1}, {"B": 2}]]`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfSlicesAtPath Gets Slice As Slice Slice",
			testFn: func(s *JsonSlice) (interface{}, error) {
				return s.GetSliceOfSlicesAtPath("[0]")
			},
			testJson:    `[ [ [1], [2] ] ]`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testSlice, err := NewSliceFromJsonString(tt.testJson)
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
