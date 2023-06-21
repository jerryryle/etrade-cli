package jsonmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonSlice_SetValue(t *testing.T) {
	tests := []struct {
		name          string
		startingValue JsonSlice
		testIndex     int
		testValue     interface{}
		expectErr     bool
		expectValue   JsonSlice
	}{
		{
			name:          "SetValue Replaces Value At Index With Value",
			startingValue: JsonSlice{0},
			testIndex:     0,
			testValue:     "TestValue",
			expectErr:     false,
			expectValue:   JsonSlice{"TestValue"},
		},
		{
			name:          "Index Too Big Returns Error",
			startingValue: JsonSlice{0},
			testIndex:     1,
			testValue:     "TestValue",
			expectErr:     true,
			expectValue:   JsonSlice{0},
		},
		{
			name:          "Index Negative Returns Error",
			startingValue: JsonSlice{0},
			testIndex:     -1,
			testValue:     "TestValue",
			expectErr:     true,
			expectValue:   JsonSlice{0},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				actualValue := tt.startingValue
				// Call the Method Under Test
				err := actualValue.SetValue(tt.testIndex, tt.testValue)
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

func TestJsonSlice_SetType(t *testing.T) {
	type testFn func(s *JsonSlice) error
	tests := []struct {
		name          string
		startingValue JsonSlice
		testFn        testFn
		expectErr     bool
		expectValue   JsonSlice
	}{
		{
			name:          "SetString Replaces Value At Index With String",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetString(0, "StringValue")
			},
			expectErr:   false,
			expectValue: JsonSlice{"StringValue"},
		},
		{
			name:          "SetInt Replaces Value At Index With Int",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetInt(0, 1234)
			},
			expectErr:   false,
			expectValue: JsonSlice{json.Number("1234")},
		},
		{
			name:          "SetFloat Replaces Value At Index With Float",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetFloat(0, 1234.5678)
			},
			expectErr:   false,
			expectValue: JsonSlice{json.Number("1234.5678")},
		},
		{
			name:          "SetBool Replaces Value At Index With Bool",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetBool(0, true)
			},
			expectErr:   false,
			expectValue: JsonSlice{true},
		},
		{
			name:          "SetMap Replaces Value At Index With Map",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetMap(0, JsonMap{"Foo": "Bar"})
			},
			expectErr:   false,
			expectValue: JsonSlice{JsonMap{"Foo": "Bar"}},
		},
		{
			name:          "SetSlice Replaces Value At Index With Slice",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetSlice(0, JsonSlice{1, 2, 3, 4})
			},
			expectErr:   false,
			expectValue: JsonSlice{JsonSlice{1, 2, 3, 4}},
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

func TestJsonSlice_SetValueAtPath(t *testing.T) {
	tests := []struct {
		name          string
		startingValue JsonSlice
		testPath      string
		testValue     interface{}
		expectErr     bool
		expectValue   JsonSlice
	}{
		{
			name:          "Can Set Value",
			startingValue: JsonSlice{0},
			testPath:      "[0]",
			testValue:     "TestValue",
			expectErr:     false,
			expectValue:   JsonSlice{"TestValue"},
		},
		{
			name:          "Can Set Value With Intermediate Map Creation",
			startingValue: JsonSlice{0},
			testPath:      "[0].Key",
			testValue:     "TestValue",
			expectErr:     false,
			expectValue:   JsonSlice{JsonMap{"Key": "TestValue"}},
		},
		{
			name:          "Empty Path Returns Error",
			startingValue: JsonSlice{0},
			testPath:      "",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue:   JsonSlice{0},
		},
		{
			name: "Cannot Create Missing Slice In Slice",
			startingValue: JsonSlice{
				// Slice contains one int. "[0]" can be accessed because it
				// exists; however [0][0] cannot be accessed because the second
				// slice doesn't exist and cannot be automatically created
				// because the length is unknown.
				0,
			},
			testPath:  "Level1[0][0].Key",
			testValue: "TestValue",
			expectErr: true,
			expectValue: JsonSlice{
				// The starting slice should remain untouched
				0,
			},
		},
		{
			name:          "Cannot Set Value For Path That Indexes Existing Slice With Index Too Big",
			startingValue: JsonSlice{0},
			testPath:      "[1]",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue:   JsonSlice{0},
		},
		{
			name:          "Cannot Set Value For Path That Indexes Existing Slice With Index Negative",
			startingValue: JsonSlice{0},
			testPath:      "[-1]",
			testValue:     "TestValue",
			expectErr:     true,
			expectValue: JsonSlice{
				// The starting map should remain untouched
				0,
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

func TestJsonSlice_SetTypeAtPath(t *testing.T) {
	type testFn func(s *JsonSlice) error
	tests := []struct {
		name          string
		startingValue JsonSlice
		testFn        testFn
		expectErr     bool
		expectValue   JsonSlice
	}{
		{
			name:          "SetStringAtPath Sets String",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetStringAtPath("[0]", "StringValue")
			},
			expectErr:   false,
			expectValue: JsonSlice{"StringValue"},
		},
		{
			name:          "SetIntAtPath Sets Int",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetIntAtPath("[0]", 1234)
			},
			expectErr:   false,
			expectValue: JsonSlice{json.Number("1234")},
		},
		{
			name:          "SetFloatAtPath Sets Float",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetFloatAtPath("[0]", 1234.5678)
			},
			expectErr:   false,
			expectValue: JsonSlice{json.Number("1234.5678")},
		},
		{
			name:          "SetBoolAtPath Sets Bool",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetBoolAtPath("[0]", true)
			},
			expectErr:   false,
			expectValue: JsonSlice{true},
		},
		{
			name:          "SetMapAtPath Sets Map",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetMapAtPath("[0]", JsonMap{"Foo": "Bar"})
			},
			expectErr:   false,
			expectValue: JsonSlice{JsonMap{"Foo": "Bar"}},
		},
		{
			name:          "SetSliceAtPath Sets Slice",
			startingValue: JsonSlice{0},
			testFn: func(s *JsonSlice) error {
				return s.SetSliceAtPath("[0]", JsonSlice{1, 2, 3, 4})
			},
			expectErr:   false,
			expectValue: JsonSlice{JsonSlice{1, 2, 3, 4}},
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
