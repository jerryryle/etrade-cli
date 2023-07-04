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
			name:        "GetValue Gets Value",
			testJson:    `{"TestKey": "TestValue"}`,
			testKey:     "TestKey",
			expectErr:   false,
			expectValue: "TestValue",
		},
		{
			name:        "GetValue Fails If Key Empty",
			testJson:    `{"TestKey": "TestValue"}`,
			testKey:     "",
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewJsonMapFromJsonString(tt.testJson)
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
			name: "GetString Gets String",
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
			name: "GetString Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("MISSING")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetString Fails On Non-String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetString("TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetInt Gets Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetInt Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("MISSING")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetInt Fails On Non-Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetInt("TestKey")
			},
			testJson:    `{"TestKey": "1234"}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloat Gets Float",
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
			name: "GetFloat Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("MISSING")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetFloat Fails On Non-Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloat("TestKey")
			},
			testJson:    `{"TestKey": "1234.5678"}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBool Gets Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBool Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("MISSING")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetBool Fails On Non-Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBool("TestKey")
			},
			testJson:    `{"TestKey": "true"}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMap Gets Map",
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
			name: "GetMap Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("MISSING")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetMap Fails On Non-Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMap("TestKey")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSlice Gets Slice",
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
			name: "GetSlice Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("MISSING")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSlice Fails On Non-Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSlice("TestKey")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStrings Gets String Slice",
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
			name: "GetSliceOfStrings Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStrings("MISSING")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfStrings Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStrings("TestKey")
			},
			testJson:    `{"TestKey": ["foo", 1]}`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfInts Gets Int Slice",
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
			name: "GetSliceOfInts Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfInts("MISSING")
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfInts Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfInts("TestKey")
			},
			testJson:    `{"TestKey": [1, "foo"]}`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloats Gets Float Slice",
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
			name: "GetSliceOfFloats Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloats("MISSING")
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfFloats Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloats("TestKey")
			},
			testJson:    `{"TestKey": [1.1, "foo"]}`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBools Gets Bool Slice",
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
			name: "GetSliceOfBools Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBools("MISSING")
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfBools Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBools("TestKey")
			},
			testJson:    `{"TestKey": [true, "foo"]}`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMaps Gets Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMaps("TestKey")
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMaps Gets Null Slice As Empty Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMaps("TestKey")
			},
			testJson:    `{"TestKey": null}`,
			expectErr:   false,
			expectValue: []JsonMap{},
		},
		{
			name: "GetSliceOfMaps Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMaps("MISSING")
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfMaps Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMaps("TestKey")
			},
			testJson:    `{"TestKey": [{"A": 1}, "foo"]}`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlices Gets Slice Slice",
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
			name: "GetSliceOfSlices Fails On Missing Key",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlices("MISSING")
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
		{
			name: "GetSliceOfSlices Fails On Mixed Slice",
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
				testMap, err := NewJsonMapFromJsonString(tt.testJson)
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

func TestJsonMap_GetValueWithDefault(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		testKey     string
		testDefault interface{}
		expectValue interface{}
	}{
		{
			name:        "GetValueWithDefault Gets Value",
			testJson:    `{"TestKey": "TestValue"}`,
			testKey:     "TestKey",
			testDefault: "TestDefaultValue",
			expectValue: "TestValue",
		},
		{
			name:        "GetValueWithDefault Returns Default If Key Not Found",
			testJson:    `{"TestKey": "TestValue"}`,
			testKey:     "MissingKey",
			testDefault: "TestDefaultValue",
			expectValue: "TestDefaultValue",
		},
		{
			name:        "GetValueWithDefault Returns Default If Key Empty",
			testJson:    `{"TestKey": "TestValue"}`,
			testKey:     "",
			testDefault: "TestDefaultValue",
			expectValue: "TestDefaultValue",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewJsonMapFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue := testMap.GetValueWithDefault(tt.testKey, tt.testDefault)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestJsonMap_GetTypeWithDefault(t *testing.T) {
	type testFn func(m *JsonMap) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetStringWithDefault Gets String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringWithDefault("TestKey", "DefaultValue")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetStringWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringWithDefault("MissingKey", "DefaultValue")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   false,
			expectValue: "DefaultValue",
		},
		{
			name: "GetStringWithDefault Fails On Non-String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringWithDefault("TestKey", "DefaultValue")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetIntWithDefault Gets Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntWithDefault("TestKey", 1)
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetIntWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntWithDefault("MissingKey", 1)
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: int64(1),
		},
		{
			name: "GetIntWithDefault Fails On Non-Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntWithDefault("TestKey", 1)
			},
			testJson:    `{"TestKey": "1234"}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloatWithDefault Gets Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatWithDefault("TestKey", 1.1)
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   false,
			expectValue: 1234.5678,
		},
		{
			name: "GetFloatWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatWithDefault("MissingKey", 1.1)
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   false,
			expectValue: 1.1,
		},
		{
			name: "GetFloatWithDefault Fails On Non-Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatWithDefault("TestKey", 1.1)
			},
			testJson:    `{"TestKey": "1234.5678"}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBoolWithDefault Gets Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolWithDefault("TestKey", false)
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBoolWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolWithDefault("MissingKey", false)
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   false,
			expectValue: false,
		},
		{
			name: "GetBoolWithDefault Fails On Non-Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolWithDefault("TestKey", false)
			},
			testJson:    `{"TestKey": "true"}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMapWithDefault Gets Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapWithDefault("TestKey", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   false,
			expectValue: JsonMap{"foo": "bar"},
		},
		{
			name: "GetMapWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapWithDefault("MissingKey", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   false,
			expectValue: JsonMap{"DefaultKey": "DefaultValue"},
		},
		{
			name: "GetMapWithDefault Fails On Non-Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapWithDefault("TestKey", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSliceWithDefault Gets Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceWithDefault("TestKey", JsonSlice{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSliceWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceWithDefault("MissingKey", JsonSlice{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: JsonSlice{"DefaultValue"},
		},
		{
			name: "GetSliceWithDefault Fails On Non-Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceWithDefault("TestKey", JsonSlice{"DefaultValue"})
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStringsWithDefault Gets String Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsWithDefault("TestKey", []string{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfStringsWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsWithDefault("MissingKey", []string{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: []string{"DefaultValue"},
		},
		{
			name: "GetSliceOfStringsWithDefault Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsWithDefault("TestKey", []string{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", 1]}`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfIntsWithDefault Gets Int Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsWithDefault("TestKey", []int64{0})
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfIntsWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsWithDefault("MissingKey", []int64{0})
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   false,
			expectValue: []int64{0},
		},
		{
			name: "GetSliceOfIntsWithDefault Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsWithDefault("TestKey", []int64{0})
			},
			testJson:    `{"TestKey": [1, "foo"]}`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloatsWithDefault Gets Float Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsWithDefault("TestKey", []float64{0.0})
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfFloatsWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsWithDefault("MissingKey", []float64{0.0})
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   false,
			expectValue: []float64{0.0},
		},
		{
			name: "GetSliceOfFloatsWithDefault Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsWithDefault("TestKey", []float64{0.0})
			},
			testJson:    `{"TestKey": [1.1, "foo"]}`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBoolsWithDefault Gets Bool Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsWithDefault("TestKey", []bool{false})
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfBoolsWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsWithDefault("MissingKey", []bool{false})
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   false,
			expectValue: []bool{false},
		},
		{
			name: "GetSliceOfBoolsWithDefault Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsWithDefault("TestKey", []bool{false})
			},
			testJson:    `{"TestKey": [true, "foo"]}`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMapsWithDefault Gets Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsWithDefault("TestKey", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMapsWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsWithDefault("MissingKey", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   false,
			expectValue: []JsonMap{{"DefaultKey": "DefaultValue"}},
		},
		{
			name: "GetSliceOfMapsWithDefault Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsWithDefault("TestKey", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `{"TestKey": [{"A": 1}, "foo"]}`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlicesWithDefault Gets Slice Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesWithDefault("TestKey", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
		{
			name: "GetSliceOfSlicesWithDefault Returns Default If Key Not Found",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesWithDefault("MissingKey", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   false,
			expectValue: []JsonSlice{{"DefaultValue"}},
		},
		{
			name: "GetSliceOfSlicesWithDefault Fails On Mixed Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesWithDefault("TestKey", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `{"TestKey": [[1], "foo"]}`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewJsonMapFromJsonString(tt.testJson)
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
			name:        "GetValueAtPath Returns Root For Empty Path",
			testJson:    `{"TestLevel1KeyWithStringValue": "TestStringValue1"}`,
			testPath:    "",
			expectErr:   false,
			expectValue: JsonMap{"TestLevel1KeyWithStringValue": "TestStringValue1"},
		},
		{
			name:        "GetValueAtPath Returns Root For Single Dot Path",
			testJson:    `{"TestLevel1KeyWithStringValue": "TestStringValue1"}`,
			testPath:    ".",
			expectErr:   false,
			expectValue: JsonMap{"TestLevel1KeyWithStringValue": "TestStringValue1"},
		},
		{
			name:        "GetValueAtPath Gets Value With Map Indexing",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithStringValue": "TestStringValue1"}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithStringValue",
			expectErr:   false,
			expectValue: "TestStringValue1",
		},
		{
			name:        "GetValueAtPath Gets Value If Path Has Extra Dots",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithStringValue": "TestStringValue1"}}`,
			testPath:    ".TestLevel1KeyWithMapValue..TestLevel2KeyWithStringValue",
			expectErr:   false,
			expectValue: "TestStringValue1",
		},
		{
			name:        "GetValueAtPath Gets Value With Slice Indexing",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[0].TestString",
			expectErr:   false,
			expectValue: "TestStringValue1",
		},
		{
			name:        "GetValueAtPath Fails If Slice Index Too Big",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[1].TestString",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "GetValueAtPath Fails With Non-Numeric Slice Index",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[A].TestString",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "GetValueAtPath Fails If Slice Index Negative",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[-1].TestString",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "GetValueAtPath Gets Value With Nested Slice Indexing",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [[{"TestString": "TestStringValue1"}]]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[0][0].TestString",
			expectErr:   false,
			expectValue: "TestStringValue1",
		},
		{
			name:        "GetValueAtPath Fails For Slice Indexing on Non-Slice",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue[0]",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "GetValueAtPath Fails For Map Indexing on Non-Map",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue.TestString",
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewJsonMapFromJsonString(tt.testJson)
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
			name: "GetStringAtPath Gets String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringAtPath(".TestKey")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetStringAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringAtPath(".MISSING")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetIntAtPath Gets Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntAtPath(".TestKey")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetIntAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntAtPath(".MISSING")
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloatAtPath Gets Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatAtPath(".TestKey")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   false,
			expectValue: 1234.5678,
		},
		{
			name: "GetFloatAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatAtPath(".MISSING")
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBoolAtPath Gets Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolAtPath(".TestKey")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBoolAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolAtPath(".MISSING")
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMapAtPath Gets Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapAtPath(".TestKey")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   false,
			expectValue: JsonMap{"foo": "bar"},
		},
		{
			name: "GetMapAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapAtPath(".MISSING")
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSliceAtPath Gets Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceAtPath(".TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSliceAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceAtPath(".MISSING")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStringsAtPath Gets String Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfStringsAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsAtPath(".MISSING")
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfIntsAtPath Gets Int Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfIntsAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsAtPath(".MISSING")
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloatsAtPath Gets Float Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfFloatsAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsAtPath(".MISSING")
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBoolsAtPath Gets Bool Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfBoolsAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsAtPath(".MISSING")
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMapsAtPath Gets Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMapsAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsAtPath(".MISSING")
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlicesAtPath Gets Slice Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesAtPath(".TestKey")
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
		{
			name: "GetSliceOfSlicesAtPath Fails With Missing Path Element",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesAtPath(".MISSING")
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewJsonMapFromJsonString(tt.testJson)
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

func TestJsonMap_GetValueAtPathWithDefault(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		testPath    string
		testDefault interface{}
		expectValue interface{}
	}{
		{
			name:        "GetValueAtPathWithDefault Returns Root For Empty Path",
			testJson:    `{"TestLevel1KeyWithStringValue": "TestStringValue1"}`,
			testPath:    "",
			testDefault: "DefaultValue",
			expectValue: JsonMap{"TestLevel1KeyWithStringValue": "TestStringValue1"},
		},
		{
			name:        "GetValueAtPathWithDefault Returns Root For Single Dot Path",
			testJson:    `{"TestLevel1KeyWithStringValue": "TestStringValue1"}`,
			testPath:    ".",
			testDefault: "DefaultValue",
			expectValue: JsonMap{"TestLevel1KeyWithStringValue": "TestStringValue1"},
		},
		{
			name:        "GetValueAtPathWithDefault Gets Value With Map Indexing",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithStringValue": "TestStringValue1"}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithStringValue",
			testDefault: "DefaultValue",
			expectValue: "TestStringValue1",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default With Missing Map Key",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithStringValue": "TestStringValue1"}}`,
			testPath:    "TestLevel1KeyWithMapValue.MissingKey",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
		{
			name:        "GetValueAtPathWithDefault Gets Value If Path Has Extra Dots",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithStringValue": "TestStringValue1"}}`,
			testPath:    ".TestLevel1KeyWithMapValue..TestLevel2KeyWithStringValue",
			testDefault: "DefaultValue",
			expectValue: "TestStringValue1",
		},
		{
			name:        "GetValueAtPathWithDefault Gets Value With Slice Indexing",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[0].TestString",
			testDefault: "DefaultValue",
			expectValue: "TestStringValue1",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default If Slice Index Too Big",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[1].TestString",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default If Slice Index Negative",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[-1].TestString",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
		{
			name:        "GetValueAtPathWithDefault Gets Value With Nested Slice Indexing",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [[{"TestString": "TestStringValue1"}]]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue[0][0].TestString",
			testDefault: "DefaultValue",
			expectValue: "TestStringValue1",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default For Slice Indexing on Non-Slice",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue[0]",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
		{
			name:        "GetValueAtPathWithDefault Returns Default For Map Indexing on Non-Map",
			testJson:    `{"TestLevel1KeyWithMapValue": {"TestLevel2KeyWithSliceValue": [{"TestString": "TestStringValue1"}]}}`,
			testPath:    "TestLevel1KeyWithMapValue.TestLevel2KeyWithSliceValue.TestString",
			testDefault: "DefaultValue",
			expectValue: "DefaultValue",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewJsonMapFromJsonString(tt.testJson)
				assert.Nil(t, err)
				// Call the Method Under Test
				actualValue := testMap.GetValueAtPathWithDefault(tt.testPath, tt.testDefault)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestJsonMap_GetTypeAtPathWithDefault(t *testing.T) {
	type testFn func(m *JsonMap) (interface{}, error)

	tests := []struct {
		name        string
		testFn      testFn
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "GetStringAtPathWithDefault Gets String",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringAtPathWithDefault(".TestKey", "DefaultValue")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   false,
			expectValue: "StringValue",
		},
		{
			name: "GetStringAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringAtPathWithDefault(".MissingKey", "DefaultValue")
			},
			testJson:    `{"TestKey": "StringValue"}`,
			expectErr:   false,
			expectValue: "DefaultValue",
		},
		{
			name: "GetStringAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetStringAtPathWithDefault(".TestKey", "DefaultValue")
			},
			testJson:    `{"TestKey": 1}`,
			expectErr:   true,
			expectValue: "",
		},
		{
			name: "GetIntAtPathWithDefault Gets Int",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntAtPathWithDefault(".TestKey", 1)
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: int64(1234),
		},
		{
			name: "GetIntAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntAtPathWithDefault(".MissingKey", 1)
			},
			testJson:    `{"TestKey": 1234}`,
			expectErr:   false,
			expectValue: int64(1),
		},
		{
			name: "GetIntAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetIntAtPathWithDefault(".TestKey", 1)
			},
			testJson:    `{"TestKey": "1234"}`,
			expectErr:   true,
			expectValue: int64(0),
		},
		{
			name: "GetFloatAtPathWithDefault Gets Float",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatAtPathWithDefault(".TestKey", 1.1)
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   false,
			expectValue: 1234.5678,
		},
		{
			name: "GetFloatAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatAtPathWithDefault(".MissingKey", 1.1)
			},
			testJson:    `{"TestKey": 1234.5678}`,
			expectErr:   false,
			expectValue: 1.1,
		},
		{
			name: "GetFloatAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetFloatAtPathWithDefault(".TestKey", 1.1)
			},
			testJson:    `{"TestKey": "1234.5678"}`,
			expectErr:   true,
			expectValue: float64(0),
		},
		{
			name: "GetBoolAtPathWithDefault Gets Bool",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolAtPathWithDefault(".TestKey", false)
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   false,
			expectValue: true,
		},
		{
			name: "GetBoolAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolAtPathWithDefault(".MissingKey", false)
			},
			testJson:    `{"TestKey": true}`,
			expectErr:   false,
			expectValue: false,
		},
		{
			name: "GetBoolAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetBoolAtPathWithDefault(".TestKey", false)
			},
			testJson:    `{"TestKey": "true"}`,
			expectErr:   true,
			expectValue: false,
		},
		{
			name: "GetMapAtPathWithDefault Gets Map",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapAtPathWithDefault(".TestKey", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   false,
			expectValue: JsonMap{"foo": "bar"},
		},
		{
			name: "GetMapAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapAtPathWithDefault(".MissingKey", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `{"TestKey": {"foo": "bar"}}`,
			expectErr:   false,
			expectValue: JsonMap{"DefaultKey": "DefaultValue"},
		},
		{
			name: "GetMapAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetMapAtPathWithDefault(".TestKey", JsonMap{"DefaultKey": "DefaultValue"})
			},
			testJson:    `{"TestKey": "foobar"}`,
			expectErr:   true,
			expectValue: JsonMap(nil),
		},
		{
			name: "GetSliceAtPathWithDefault Gets Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceAtPathWithDefault(".TestKey", JsonSlice{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: JsonSlice{"foo", "bar"},
		},
		{
			name: "GetSliceAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceAtPathWithDefault(".MissingKey", JsonSlice{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: JsonSlice{"DefaultValue"},
		},
		{
			name: "GetSliceAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceAtPathWithDefault(".TestKey", JsonSlice{"DefaultValue"})
			},
			testJson:    `{"TestKey": "foobar"}`,
			expectErr:   true,
			expectValue: JsonSlice(nil),
		},
		{
			name: "GetSliceOfStringsAtPathWithDefault Gets String Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsAtPathWithDefault(".TestKey", []string{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: []string{"foo", "bar"},
		},
		{
			name: "GetSliceOfStringsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsAtPathWithDefault(".MissingKey", []string{"DefaultValue"})
			},
			testJson:    `{"TestKey": ["foo", "bar"]}`,
			expectErr:   false,
			expectValue: []string{"DefaultValue"},
		},
		{
			name: "GetSliceOfStringsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfStringsAtPathWithDefault(".TestKey", []string{"DefaultValue"})
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   true,
			expectValue: []string(nil),
		},
		{
			name: "GetSliceOfIntsAtPathWithDefault Gets Int Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsAtPathWithDefault(".TestKey", []int64{0})
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   false,
			expectValue: []int64{1, 2},
		},
		{
			name: "GetSliceOfIntsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsAtPathWithDefault(".MissingKey", []int64{0})
			},
			testJson:    `{"TestKey": [1, 2]}`,
			expectErr:   false,
			expectValue: []int64{0},
		},
		{
			name: "GetSliceOfIntsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfIntsAtPathWithDefault(".TestKey", []int64{0})
			},
			testJson:    `{"TestKey": ["1", "2"]}`,
			expectErr:   true,
			expectValue: []int64(nil),
		},
		{
			name: "GetSliceOfFloatsAtPathWithDefault Gets Float Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsAtPathWithDefault(".TestKey", []float64{0.0})
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   false,
			expectValue: []float64{1.1, 2.2},
		},
		{
			name: "GetSliceOfFloatsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsAtPathWithDefault(".MissingKey", []float64{0.0})
			},
			testJson:    `{"TestKey": [1.1, 2.2]}`,
			expectErr:   false,
			expectValue: []float64{0.0},
		},
		{
			name: "GetSliceOfFloatsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfFloatsAtPathWithDefault(".TestKey", []float64{0.0})
			},
			testJson:    `{"TestKey": ["1.1", "2.2"]}`,
			expectErr:   true,
			expectValue: []float64(nil),
		},
		{
			name: "GetSliceOfBoolsAtPathWithDefault Gets Bool Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsAtPathWithDefault(".TestKey", []bool{false})
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   false,
			expectValue: []bool{true, false},
		},
		{
			name: "GetSliceOfBoolsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsAtPathWithDefault(".MissingKey", []bool{false})
			},
			testJson:    `{"TestKey": [true, false]}`,
			expectErr:   false,
			expectValue: []bool{false},
		},
		{
			name: "GetSliceOfBoolsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfBoolsAtPathWithDefault(".TestKey", []bool{false})
			},
			testJson:    `{"TestKey": ["true", "false"]}`,
			expectErr:   true,
			expectValue: []bool(nil),
		},
		{
			name: "GetSliceOfMapsAtPathWithDefault Gets Map Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsAtPathWithDefault(".TestKey", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   false,
			expectValue: []JsonMap{{"A": json.Number("1")}, {"B": json.Number("2")}},
		},
		{
			name: "GetSliceOfMapsAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsAtPathWithDefault(".MissingKey", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `{"TestKey": [{"A": 1}, {"B": 2}]}`,
			expectErr:   false,
			expectValue: []JsonMap{{"DefaultKey": "DefaultValue"}},
		},
		{
			name: "GetSliceOfMapsAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfMapsAtPathWithDefault(".TestKey", []JsonMap{{"DefaultKey": "DefaultValue"}})
			},
			testJson:    `{"TestKey": ["A", "B"]}`,
			expectErr:   true,
			expectValue: []JsonMap(nil),
		},
		{
			name: "GetSliceOfSlicesAtPathWithDefault Gets Slice Slice",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesAtPathWithDefault(".TestKey", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   false,
			expectValue: []JsonSlice{{json.Number("1")}, {json.Number("2")}},
		},
		{
			name: "GetSliceOfSlicesAtPathWithDefault Returns Default Value For Invalid Path",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesAtPathWithDefault(".MissingKey", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `{"TestKey": [[1], [2]]}`,
			expectErr:   false,
			expectValue: []JsonSlice{{"DefaultValue"}},
		},
		{
			name: "GetSliceOfSlicesAtPathWithDefault Fails For Valid Path But Wrong Type",
			testFn: func(m *JsonMap) (interface{}, error) {
				return m.GetSliceOfSlicesAtPathWithDefault(".TestKey", []JsonSlice{{"DefaultValue"}})
			},
			testJson:    `{"TestKey": ["1", "2"]}`,
			expectErr:   true,
			expectValue: []JsonSlice(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				testMap, err := NewJsonMapFromJsonString(tt.testJson)
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
