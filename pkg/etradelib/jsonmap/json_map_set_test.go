package jsonmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMap_SetValue(t *testing.T) {
	tests := []struct {
		name          string
		startingValue JsonMap
		testKey       string
		testValue     interface{}
		expectValue   JsonMap
	}{
		{
			name:          "Sets Value For Key",
			startingValue: JsonMap{},
			testKey:       "TestKey",
			testValue:     "TestValue",
			expectValue:   JsonMap{"TestKey": "TestValue"},
		},
		{
			name:          "Sets Value For Key And Overwrites Existing",
			startingValue: JsonMap{"TestKey": "OldValue"},
			testKey:       "TestKey",
			testValue:     "TestValue",
			expectValue:   JsonMap{"TestKey": "TestValue"},
		},
		{
			name:          "Empty Key Succeeds",
			startingValue: JsonMap{},
			testKey:       "",
			testValue:     "TestValue",
			expectValue:   JsonMap{"": "TestValue"},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				actualValue := tt.startingValue
				// Call the Method Under Test
				actualValue.SetValue(tt.testKey, tt.testValue)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestJsonMap_SetType(t *testing.T) {
	type testFn func(m *JsonMap)
	tests := []struct {
		name          string
		startingValue JsonMap
		testFn        testFn
		expectValue   JsonMap
	}{
		{
			name:          "SetString Sets String",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) {
				m.SetString("TestKey", "StringValue")
			},
			expectValue: JsonMap{"TestKey": "StringValue"},
		},
		{
			name:          "SetInt Sets Int",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) {
				m.SetInt("TestKey", 1234)
			},
			expectValue: JsonMap{"TestKey": json.Number("1234")},
		},
		{
			name:          "SetFloat Sets Float",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) {
				m.SetFloat("TestKey", 1234.5678)
			},
			expectValue: JsonMap{"TestKey": json.Number("1234.5678")},
		},
		{
			name:          "SetBool Sets Bool",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) {
				m.SetBool("TestKey", true)
			},
			expectValue: JsonMap{"TestKey": true},
		},
		{
			name:          "SetMap Sets Map",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) {
				m.SetMap("TestKey", JsonMap{"Foo": "Bar"})
			},
			expectValue: JsonMap{"TestKey": JsonMap{"Foo": "Bar"}},
		},
		{
			name:          "SetSlice Sets Slice",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) {
				m.SetSlice("TestKey", JsonSlice{1, 2, 3, 4})
			},
			expectValue: JsonMap{"TestKey": JsonSlice{1, 2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				actualValue := tt.startingValue
				// Call the Method Under Test
				tt.testFn(&actualValue)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestJsonMap_SetValueAtPath(t *testing.T) {
	tests := []struct {
		name          string
		startingValue JsonMap
		testPath      string
		testValue     interface{}
		expectErr     bool
		expectValue   JsonMap
	}{
		{
			name:          "Can Set Value With Intermediate Map Creation",
			startingValue: JsonMap{},
			testPath:      "Level1.Level2.Level3",
			testValue:     "TestValue",
			expectErr:     false,
			expectValue:   JsonMap{"Level1": JsonMap{"Level2": JsonMap{"Level3": "TestValue"}}},
		},
		{
			name:          "Empty Path Fails",
			startingValue: JsonMap{},
			testPath:      "",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue:   JsonMap{},
		},
		{
			name:          "Create Missing Slice In Map Fails",
			startingValue: JsonMap{
				// Map is empty. The "Level1" key cannot be added because the
				// [0] index in the path implies that it should hold a slice.
				// But a slice cannot be automatically created because the
				// length is unknown.
			},
			testPath:    "Level1[0].Key",
			testValue:   "TestValue",
			expectErr:   true,
			expectValue: JsonMap{},
		},
		{
			name: "Create Missing Slice In Slice Fails",
			startingValue: JsonMap{
				// Map contains key "Level1" holding a slice with one int.
				// "Level1[0]" can be accessed because it exists; however
				// Level1[0][0] cannot be accessed because the second slice
				// doesn't exist and cannot be automatically created because
				// the length is unknown.
				"Level1": JsonSlice{0},
			},
			testPath:  "Level1[0][0].Key",
			testValue: "TestValue",
			expectErr: true,
			expectValue: JsonMap{
				// The starting map should remain untouched
				"Level1": JsonSlice{0},
			},
		},
		{
			name: "Can Set Value For Path That Indexes Existing Slice And Replaces Slice Value With New Map",
			startingValue: JsonMap{
				// Slice is initialized with one integer element. This will be
				// replaced with a map while traversing the path.
				"Level1": JsonSlice{0},
			},
			testPath:    "Level1[0].Key",
			testValue:   "TestValue",
			expectErr:   false,
			expectValue: JsonMap{"Level1": JsonSlice{JsonMap{"Key": "TestValue"}}},
		},
		{
			name:          "Set Value For Path That Indexes Existing Map As Slice Fails",
			startingValue: JsonMap{"Level1": JsonMap{"Level2": "Level2Value"}},
			testPath:      ".Level1[0]",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue: JsonMap{
				// The starting map should remain untouched
				"Level1": JsonMap{"Level2": "Level2Value"},
			},
		},
		{
			name:          "Set Value For Path That Indexes Existing Slice As Map Fails",
			startingValue: JsonMap{"Level1": JsonSlice{"A"}},
			testPath:      ".Level1.A",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue: JsonMap{
				// The starting map should remain untouched
				"Level1": JsonSlice{"A"},
			},
		},
		{
			name:          "Set Value For Path That Indexes Existing Slice With Index Too Big Fails",
			startingValue: JsonMap{"Level1": JsonSlice{0}},
			testPath:      "Level1[1].Key",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue: JsonMap{
				// The starting map should remain untouched
				"Level1": JsonSlice{0},
			},
		},
		{
			name:          "Set Value For Path That Indexes Slice With Non-Numeric Index Fails",
			startingValue: JsonMap{"Level1": JsonSlice{0}},
			testPath:      "Level1[A].Key",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue: JsonMap{
				// The starting map should remain untouched
				"Level1": JsonSlice{0},
			},
		},
		{
			name:          "Set Value For Path That Indexes Existing Slice With Index Negative Fails",
			startingValue: JsonMap{"Level1": JsonSlice{0}},
			testPath:      "Level1[-1].Key",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue: JsonMap{
				// The starting map should remain untouched
				"Level1": JsonSlice{0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				actualValue := tt.startingValue
				// Call the Method Under Test
				err := actualValue.SetValueAtPath(tt.testPath, tt.testValue)
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

func TestJsonMap_SetTypeAtPath(t *testing.T) {
	type testFn func(m *JsonMap) error
	tests := []struct {
		name          string
		startingValue JsonMap
		testFn        testFn
		expectErr     bool
		expectValue   JsonMap
	}{
		{
			name:          "SetStringAtPath Sets String",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) error {
				return m.SetStringAtPath(".TestKey", "StringValue")
			},
			expectErr:   false,
			expectValue: JsonMap{"TestKey": "StringValue"},
		},
		{
			name:          "SetIntAtPath Sets Int",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) error {
				return m.SetIntAtPath(".TestKey", 1234)
			},
			expectErr:   false,
			expectValue: JsonMap{"TestKey": json.Number("1234")},
		},
		{
			name:          "SetFloatAtPath Sets Float",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) error {
				return m.SetFloatAtPath(".TestKey", 1234.5678)
			},
			expectErr:   false,
			expectValue: JsonMap{"TestKey": json.Number("1234.5678")},
		},
		{
			name:          "SetBoolAtPath Sets Bool",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) error {
				return m.SetBoolAtPath(".TestKey", true)
			},
			expectErr:   false,
			expectValue: JsonMap{"TestKey": true},
		},
		{
			name:          "SetMapAtPath Sets Map",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) error {
				return m.SetMapAtPath(".TestKey", JsonMap{"Foo": "Bar"})
			},
			expectErr:   false,
			expectValue: JsonMap{"TestKey": JsonMap{"Foo": "Bar"}},
		},
		{
			name:          "SetSliceAtPath Sets Slice",
			startingValue: JsonMap{},
			testFn: func(m *JsonMap) error {
				return m.SetSliceAtPath(".TestKey", JsonSlice{1, 2, 3, 4})
			},
			expectErr:   false,
			expectValue: JsonMap{"TestKey": JsonSlice{1, 2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				actualValue := tt.startingValue
				// Call the Method Under Test
				err := tt.testFn(&actualValue)
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
