package jsonmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMap_GetValue(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		testKey     string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name:        "Key Returns Value",
			testJson:    `{"TestKey": "TestValue"}`,
			testKey:     "TestKey",
			expectErr:   false,
			expectValue: "TestValue",
		},
		{
			name:        "Empty Key Returns Error",
			testJson:    `{"TestKey": "TestValue"}`,
			testKey:     "",
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewMapFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := testMap.GetValue(tt.testKey)
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

func TestJsonMap_GetType(t *testing.T) {
	type testFn func(m *JsonMap) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetString Gets String As String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("TestKey")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetString Gets Null As Empty String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Int As String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Float As String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("TestKey")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Bool As String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Map As String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("TestKey")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Cannot Get Slice As String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetInt Gets Int As Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetInt Cannot Get String As Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": "1234"}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Float As Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": 1234.5}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Bool As Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Null As Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Map As Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Cannot Get Slice As Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloat Gets Float As Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("TestKey")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   false,
			expectValue: 1234.5678,
		},
		{
			name: "GetFloat Gets Int As Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: float64(1234),
		},
		{
			name: "GetFloat Cannot Get String As Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("TestKey")
			},
			testJson:    `{"TestKey": "1234.5678"}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Bool As Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Null As Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Map As Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("TestKey")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Cannot Get Slice As Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBool Gets Bool As Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBool Cannot Get String As Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": "true"}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Int As Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Float As Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Null As Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Map As Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Cannot Get Slice As Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMap Gets Map As Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("TestKey")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   false,
			expectValue: JsonMap{"foo": "bar"},
		},
		{
			name: "GetMap Gets Null As Nil JsonMap",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get String As Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("TestKey")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Int As Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Float As Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("TestKey")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Bool As Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Cannot Get Slice As Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSlice Gets Slice As Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSlice Gets Null As Nil Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get String As Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("TestKey")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Int As Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Float As Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("TestKey")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Bool As Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Cannot Get Map As Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("TestKey")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStrings Gets Slice As String Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStrings("TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfStrings Gets Null Slice As Empty String Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStrings("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: []string{},
		},
		{
			name: "GetSliceOfStrings Cannot Get Mixed Slice As String Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStrings("TestKey")
			},
			testJson:    `{"TestKey": ["foo", 1]}`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfInts Gets Slice As Int Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfInts("TestKey")
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfInts Gets Null Slice As Empty Int Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfInts("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: []int64{},
		},
		{
			name: "GetSliceOfInts Cannot Get Mixed Slice As Int Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfInts("TestKey")
			},
			testJson:    `{"TestKey": [1, "foo"]}`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloats Gets Slice As Float Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloats("TestKey")
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfFloats Gets Null Slice As Empty Float Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloats("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: []float64{},
		},
		{
			name: "GetSliceOfFloats Cannot Get Mixed Slice As Float Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloats("TestKey")
			},
			testJson:    `{"TestKey": [1.1, "foo"]}`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBools Gets Slice As Bool Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBools("TestKey")
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfBools Gets Null Slice As Empty Bool Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBools("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: []bool{},
		},
		{
			name: "GetSliceOfBools Cannot Get Mixed Slice As Bool Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBools("TestKey")
			},
			testJson:    `{"TestKey": [true, "foo"]}`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMaps Gets Slice As Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapSlice("TestKey")
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMaps Gets Null Slice As Empty Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapSlice("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: []JsonMap{},
		},
		{
			name: "GetSliceOfMaps Cannot Get Mixed Slice As Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapSlice("TestKey")
			},
			testJson:    `{"TestKey": [{"A": 1}, "foo"]}`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlices Gets Slice As Slice Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlices("TestKey")
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
		{
			name: "GetSliceOfSlices Gets Null Slice As Empty Slice Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlices("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: []JsonSlice{},
		},
		{
			name: "GetSliceOfSlices Cannot Get Mixed Slice As Slice Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlices("TestKey")
			},
			testJson:    `{"TestKey": [[1], "foo"]}`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewMapFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := tt.testFn(&testMap)
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

func TestJsonMap_GetValueAtPath(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		testPath    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name:      "Empty Path Returns Root",
			testJson:  `{"TestLevel1KeyWithStringValue": "TestStringValue1"}`,
			testPath:  "",
			expectErr: false,
			expectValue: JsonMap{
				"TestLevel1KeyWithStringValue": "TestStringValue1",
			},
		},
		{
			name:      "Single Dot Path Returns Root",
			testJson:  `{"TestLevel1KeyWithStringValue": "TestStringValue1"}`,
			expectErr: false,
			expectValue: JsonMap{
				"TestLevel1KeyWithStringValue": "TestStringValue1",
			},
		},
		{
			name:        "Map Indexing Returns Value",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithStringValue": "TestStringValue1"}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithStringValue",
			expectErr:   false,
			expectValue: "TestStringValue1",
		},
		{
			name:        "Path That Has Extra Dots Still Returns Value",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithStringValue": "TestStringValue1"}}`,
			testPath:    ".TestLevel1KeyWithMapValue..TestLevel2KeyWithStringValue",
			expectErr:   false,
			expectValue: "TestStringValue1",
		},
		{
			name:        "Array Indexing Returns Value",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[0].TestString",
			expectErr:   false,
			expectValue: "TestStringValue1",
		},
		{
			name:        "Array Index Too Big Returns Error",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[1].TestString",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Array Index Negative Returns Error",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[-1].TestString",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Nested Array Indexing Returns Value",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [[{"TestString": "TestStringValue1"}]]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[0][0].TestString",
			expectErr:   false,
			expectValue: "TestStringValue1",
		},
		{
			name:        "Array Indexing on Non-Array Returns Error",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue[0]",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Map Indexing on Non-Map Returns Error",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue.TestString",
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewMapFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := testMap.GetValueAtPath(tt.testPath)
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

func TestJsonMap_GetTypeAtPath(t *testing.T) {
	type testFn func(m *JsonMap) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetStringAtPath Gets String As String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringAtPath(".TestKey")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetIntAtPath Gets Int As Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntAtPath(".TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetFloatAtPath Gets Float As Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatAtPath(".TestKey")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   false,
			expectValue: 1234.5678,
		},
		{
			name: "GetBoolAtPath Gets Bool As Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolAtPath(".TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetMapAtPath Gets Map As Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapAtPath(".TestKey")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   false,
			expectValue: JsonMap{"foo": "bar"},
		},
		{
			name: "GetSliceAtPath Gets Slice As Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceAtPath(".TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSliceOfStringsAtPath Gets Slice As String Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfIntsAtPath Gets Slice As Int Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfFloatsAtPath Gets Slice As Float Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfBoolsAtPath Gets Slice As Bool Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfMapsAtPath Gets Slice As Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfSlicesAtPath Gets Slice As Slice Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewMapFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := tt.testFn(&testMap)
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
