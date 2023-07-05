package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnumFlagValue_newEnumFlagValue(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	expectedValue := enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	actualValue := *newEnumFlagValue(testMap, 1)

	assert.Equal(t, expectedValue, actualValue)
}

func TestEnumFlagValue_String(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	testValue := enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	assert.Equal(t, "test enum name 1", testValue.String())
}

func TestEnumFlagValue_Set(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	testValue := enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	expectedValue := enumFlagValue[int]{
		StringValue:           "test enum name 2",
		EnumValue:             2,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	// Set with valid enum name updates the value
	err := testValue.Set("test enum name 2")
	assert.Nil(t, err)
	assert.Equal(t, expectedValue, testValue)

	testValue = enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	expectedValue = enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	// Set with invalid enum name returns error and does not update the value
	err = testValue.Set("test enum name 3")
	assert.Error(t, err)
	assert.Equal(t, expectedValue, testValue)
}

func TestEnumFlagValue_Type(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	testValue := enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	// Type is always string.
	assert.Equal(t, "string", testValue.Type())
}

func TestEnumFlagValue_AllowedValues(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	testValue := enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	assert.Equal(t, []string{"test enum name 1", "test enum name 2"}, testValue.AllowedValues())
}

func TestEnumFlagValue_JoinAllowedValues(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	testValue := enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	assert.Equal(t, "test enum name 1, test enum name 2", testValue.JoinAllowedValues(", "))
}

func TestEnumFlagValue_AllowedValuesWithHelp(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	testValue := enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	assert.Equal(t, []string{"test enum name 1\thelp string 1", "test enum name 2"}, testValue.AllowedValuesWithHelp())
}

func TestEnumFlagValue_Value(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	testValue := enumFlagValue[int]{
		StringValue:           "test enum name 1",
		EnumValue:             1,
		valueMap:              testMap,
		allowedValues:         []string{"test enum name 1", "test enum name 2"},
		allowedValuesWithHelp: []string{"test enum name 1\thelp string 1", "test enum name 2"},
	}

	assert.Equal(t, 1, testValue.Value())
}

func TestEnumValueWithHelpMap_GetEnumValue(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	actualValue, err := testMap.GetEnumValue("test enum name 1")
	assert.Nil(t, err)
	assert.Equal(t, 1, actualValue)

	actualValue, err = testMap.GetEnumValue("test enum name 3")
	assert.Error(t, err)
	assert.Equal(t, 0, actualValue)
}

func TestEnumValueWithHelpMap_GetEnumValueWithDefault(t *testing.T) {
	var testMap = enumValueWithHelpMap[int]{
		"test enum name 1": {1, "help string 1"},
		"test enum name 2": {2, ""},
	}

	actualValue := testMap.GetEnumValueWithDefault("test enum name 1", 4)
	assert.Equal(t, 1, actualValue)

	actualValue = testMap.GetEnumValueWithDefault("test enum name 3", 4)
	assert.Equal(t, 4, actualValue)
}
