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
			testJson:    `["TestValue1", "TestValue2"]`,
			testIndex:   1,
			expectErr:   false,
			expectValue: "TestValue2",
		},
		{
			name:        "Index Too Big Returns Error",
			testJson:    `["TestValue1", "TestValue2"]`,
			testIndex:   2,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Index Negative Returns Error",
			testJson:    `["TestValue1", "TestValue2"]`,
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
				testResultValue, err := testSlice.GetValue(tt.testIndex)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, testResultValue)
			},
		)
	}
}

func TestJsonSlice_GetType(t *testing.T) {
	type testFn func(s JsonSlice) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetString Gets String As String",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetString(1)
			},
			testJson:    `["StringValue1", "StringValue2"]`,
			expectErr:   false,
			expectValue: "StringValue2",
		},
		{
			name: "GetString Gets Null As Empty String",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetString(1)
			},
			testJson:    `["StringValue1", null]`,
			expectErr:   false,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Int As String",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetString(1)
			},
			testJson:    `["StringValue1", 1234]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Float As String",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetString(1)
			},
			testJson:    `["StringValue1", 1234.567]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Bool As String",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetString(1)
			},
			testJson:    `["StringValue1", true]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Map As String",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetString(1)
			},
			testJson:    `["StringValue1", {"foo": "bar"}]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Slice As String",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetString(1)
			},
			testJson:    `["StringValue1", ["foo", "bar"]]`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetInt Gets Int As Int",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetInt(1)
			},
			testJson:    `[1234, 5678]`,
			expectErr:   false,
			expectValue: int64(5678),
		},
		{
			name: "GetInt Cannot Get String As Int",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetInt(1)
			},
			testJson:    `[1234, "5678"]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Float As Int",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetInt(1)
			},
			testJson:    `[1234, 567.8]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Bool As Int",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetInt(1)
			},
			testJson:    `[1234, true]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Null As Int",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetInt(1)
			},
			testJson:    `[1234, null]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Map As Int",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetInt(1)
			},
			testJson:    `[1234, {"foo": "bar"}]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Slice As Int",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetInt(1)
			},
			testJson:    `[1234, ["foo", "bar"]]`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloat Gets Float As Float",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloat(1)
			},
			testJson:    `[123.4, 567.8]`,
			expectErr:   false,
			expectValue: 567.8,
		},
		{
			name: "GetFloat Gets Int As Float",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloat(1)
			},
			testJson:    `[1234, 5678]`,
			expectErr:   false,
			expectValue: float64(5678),
		},
		{
			name: "GetFloat Cannot Get String As Float",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloat(1)
			},
			testJson:    `[123.4, "567.8"]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Bool As Float",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloat(1)
			},
			testJson:    `[123.4, true]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Null As Float",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloat(1)
			},
			testJson:    `[123.4, null]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Map As Float",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloat(1)
			},
			testJson:    `[123.4, {"foo": "bar"}]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Slice As Float",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloat(1)
			},
			testJson:    `[123.4, ["foo", "bar"]]`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBool Gets Bool As Bool",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBool(1)
			},
			testJson:    `[false, true]`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBool Cannot Get String As Bool",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBool(1)
			},
			testJson:    `[false, "true"]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Int As Bool",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBool(1)
			},
			testJson:    `[false, 1234]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Float As Bool",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBool(1)
			},
			testJson:    `[false, 1234.5678]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Null As Bool",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBool(1)
			},
			testJson:    `[false, null]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Map As Bool",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBool(1)
			},
			testJson:    `[false, {"foo": "bar"}]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Slice As Bool",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBool(1)
			},
			testJson:    `[false, ["foo", "bar"]]`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMap Gets Map As JsonMap",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMap(1)
			},
			testJson:    `[{"foo": "bar"}, {"boo": "far"}]`,
			expectErr:   false,
			expectValue: JsonMap{"boo": "far"},
		},
		{
			name: "GetMap Gets Null As Nil JsonMap",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMap(1)
			},
			testJson:    `[{"foo": "bar"}, null]`,
			expectErr:   false,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get String As JsonMap",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMap(1)
			},
			testJson:    `[{"foo": "bar"}, "moo"]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Int As JsonMap",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMap(1)
			},
			testJson:    `[{"foo": "bar"}, 1234]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Float As JsonMap",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMap(1)
			},
			testJson:    `[{"foo": "bar"}, 1234.5678]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Bool As JsonMap",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMap(1)
			},
			testJson:    `[{"foo": "bar"}, true]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Slice As JsonMap",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMap(1)
			},
			testJson:    `[{"foo": "bar"}, ["boo", "far"]]`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSlice Gets Slice As Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSlice(1)
			},
			testJson:    `[["foo", "bar"], ["boo", "far"]]`,
			expectErr:   false,
			expectValue: JsonSlice{"boo", "far"},
		},
		{
			name: "GetSlice Gets Null As Nil Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSlice(1)
			},
			testJson:    `[["foo", "bar"], null]`,
			expectErr:   false,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get String As Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSlice(1)
			},
			testJson:    `[["foo", "bar"], "moo"]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Int As Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSlice(1)
			},
			testJson:    `[["foo", "bar"], 1234]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Float As Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSlice(1)
			},
			testJson:    `[["foo", "bar"], 1234.5678]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Bool As Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSlice(1)
			},
			testJson:    `[["foo", "bar"], true]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Map As Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSlice(1)
			},
			testJson:    `[["foo", "bar"], {"boo": "far"}]`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetStringSlice Gets Slice As String Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetStringSlice(1)
			},
			testJson:    `[["foo", "bar"], ["boo", "far"]]`,
			expectErr:   false,
			expectValue: []string{"boo", "far"},
		},
		{
			name: "GetStringSlice Gets Null Slice As Empty String Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetStringSlice(1)
			},
			testJson:    `[["foo", "bar"], null]`,
			expectErr:   false,
			expectValue: []string{},
		},
		{
			name: "GetStringSlice Cannot Get Mixed Slice As String Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetStringSlice(1)
			},
			testJson:    `[["foo", "bar"], ["boo", 1]]`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetIntSlice Gets Slice As Int Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetIntSlice(1)
			},
			testJson:    `[[1, 2], [3, 4]]`,
			expectErr:   false,
			expectValue: []int64{3, 4},
		},
		{
			name: "GetIntSlice Gets Null Slice As Empty Int Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetIntSlice(1)
			},
			testJson:    `[[1, 2], null]`,
			expectErr:   false,
			expectValue: []int64{},
		},
		{
			name: "GetIntSlice Cannot Get Mixed Slice As Int Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetIntSlice(1)
			},
			testJson:    `[[1, 2], [3, "foo"]]`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetFloatSlice Gets Slice As Float Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloatSlice(1)
			},
			testJson:    `[[1.1, 2.2], [3.3, 4.4]]`,
			expectErr:   false,
			expectValue: []float64{3.3, 4.4},
		},
		{
			name: "GetFloatSlice Gets Null Slice As Empty Float Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloatSlice(1)
			},
			testJson:    `[[1.1, 2.2], null]`,
			expectErr:   false,
			expectValue: []float64{},
		},
		{
			name: "GetFloatSlice Cannot Get Mixed Slice As Float Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetFloatSlice(1)
			},
			testJson:    `[[1.1, 2.2], [3.3, "foo"]]`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetBoolSlice Gets Slice As Bool Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBoolSlice(1)
			},
			testJson:    `[[false, true], [true, false]]`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetBoolSlice Gets Null Slice As Empty Bool Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBoolSlice(1)
			},
			testJson:    `[[false, true], null]`,
			expectErr:   false,
			expectValue: []bool{},
		},
		{
			name: "GetBoolSlice Cannot Get Mixed Slice As Bool Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetBoolSlice(1)
			},
			testJson:    `[[false, true], [true, "foo"]]`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetMapSlice Gets Slice As Map Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMapSlice(1)
			},
			testJson:    `[[{"A": 1}, {"B": 2}], [{"C": 3}, {"D": 4}]]`,
			expectErr:   false,
			expectValue: []JsonMap{{"C": json.Number("3")}, {"D": json.Number("4")}},
		},
		{
			name: "GetMapSlice Gets Null Slice As Empty Map Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMapSlice(1)
			},
			testJson:    `[[{"A": 1}, {"B": 2}], null]`,
			expectErr:   false,
			expectValue: []JsonMap{},
		},
		{
			name: "GetMapSlice Cannot Get Mixed Slice As Map Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetMapSlice(1)
			},
			testJson:    `[[{"A": 1}, {"B": 2}], [{"C": 3}, "foo"]]`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceSlice Gets Slice As Slice Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSliceSlice(1)
			},
			testJson:    `[ [[1], [2]], [[3], [4]] ]`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("3")}, {json.Number("4")}},
		},
		{
			name: "GetSliceSlice Gets Null Slice As Empty Slice Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSliceSlice(1)
			},
			testJson:    `[ [[1], [2]], null ]`,
			expectErr:   false,
			expectValue: []JsonSlice{},
		},
		{
			name: "GetSliceSlice Cannot Get Mixed Slice As Slice Slice",
			testFn: func(s JsonSlice) (interface{}, error) {
				return s.GetSliceSlice(1)
			},
			testJson:    `[ [[1], [2]], [[3], "foo"] ]`,
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
				testResultValue, err := tt.testFn(testSlice)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, testResultValue)
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
			testJson:    `["TestStringValue1", "TestStringValue2"]`,
			testPath:    "",
			expectErr:   false,
			expectValue: JsonSlice{"TestStringValue1", "TestStringValue2"},
		},
		{
			name:        "Single Dot Path Returns Root",
			testJson:    `["TestStringValue1", "TestStringValue2"]`,
			expectErr:   false,
			expectValue: JsonSlice{"TestStringValue1", "TestStringValue2"},
		},
		{
			name:        "Slice Indexing Returns Value",
			testJson:    `[{"TestKey": "TestValue1"}, {"TestKey": "TestValue2"}]`,
			testPath:    "[1].TestKey",
			expectErr:   false,
			expectValue: "TestValue2",
		},
		{
			name:        "Path That Has Extra Dots Still Returns Value",
			testJson:    `[{"TestKey": "TestValue1"}, {"TestKey": "TestValue2"}]`,
			testPath:    ".[1]..TestKey.",
			expectErr:   false,
			expectValue: "TestValue2",
		},
		{
			name:        "Array Index Too Big Returns Error",
			testJson:    `[{"TestKey": "TestValue1"}, {"TestKey": "TestValue2"}]`,
			testPath:    "[2].TestKey",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Array Index Negative Returns Error",
			testJson:    `[{"TestKey": "TestValue1"}, {"TestKey": "TestValue2"}]`,
			testPath:    "[-1].TestKey",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Nested Slice Indexing Returns Value",
			testJson:    `[[{"TestKey": "TestValue1"}], [{"TestKey": "TestValue2"}]]`,
			testPath:    "[1][0].TestKey",
			expectErr:   false,
			expectValue: "TestValue2",
		},
		{
			name:        "Array Indexing on Non-Array Returns Error",
			testJson:    `[{"TestKey": "TestValue1"}, {"TestKey": "TestValue2"}]`,
			testPath:    "[0][0]",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Map Indexing on Non-Map Returns Error",
			testJson:    `[{"TestKey": "TestValue1"}, {"TestKey": "TestValue2"}]`,
			testPath:    "[1][0]",
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
				testResultValue, err := testSlice.GetValueAtPath(tt.testPath)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, testResultValue)
			},
		)
	}
}
