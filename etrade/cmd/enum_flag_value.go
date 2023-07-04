package cmd

import (
	"fmt"
	"sort"
	"strings"
)

type enumValueWithHelp[T comparable] struct {
	Value T
	Help  string
}

type enumValueWithHelpMap[T comparable] map[string]enumValueWithHelp[T]

type enumFlagValue[T comparable] struct {
	StringValue           string
	EnumValue             T
	valueMap              enumValueWithHelpMap[T]
	allowedValues         []string
	allowedValuesWithHelp []string
}

// newEnumFlagValue creates a new enumFlagValue with the default value set to defaultValue
func newEnumFlagValue[T comparable](valueMap enumValueWithHelpMap[T], defaultEnumValue T) *enumFlagValue[T] {
	defaultStringValue := ""
	allowedValues := make([]string, 0, len(valueMap))
	allowedValuesWithHelp := make([]string, 0, len(valueMap))

	for k, v := range valueMap {
		// Add the value to the list of allowed values
		allowedValues = append(allowedValues, k)

		// Add the value to the list of allowed values with a help string, if
		// provided. Otherwise, add without help string.
		// See: https://github.com/spf13/cobra/blob/main/shell_completions.md#completions-for-flags
		if v.Help != "" {
			allowedValuesWithHelp = append(allowedValuesWithHelp, fmt.Sprintf("%s\t%s", k, v.Help))
		} else {
			allowedValuesWithHelp = append(allowedValuesWithHelp, k)
		}

		// If the default enum has an associated string value, set it as the
		// default. Otherwise, the string value will default to ""
		if v.Value == defaultEnumValue {
			defaultStringValue = k
		}
	}

	sort.Strings(allowedValues)
	sort.Strings(allowedValuesWithHelp)

	return &enumFlagValue[T]{
		StringValue:           defaultStringValue,
		EnumValue:             defaultEnumValue,
		valueMap:              valueMap,
		allowedValues:         allowedValues,
		allowedValuesWithHelp: allowedValuesWithHelp,
	}
}

func (m *enumFlagValue[T]) String() string {
	return m.StringValue
}

func (m *enumFlagValue[T]) Set(value string) error {
	enumValue, err := m.valueMap.GetEnumValue(value)
	if err != nil {
		return fmt.Errorf("%s is not one of the allowed values: [%s]", value, strings.Join(m.allowedValues, ", "))
	}
	m.StringValue = value
	m.EnumValue = enumValue
	return nil
}

func (m *enumFlagValue[T]) Type() string {
	return "string"
}

func (m *enumFlagValue[T]) AllowedValues() []string {
	return m.allowedValues
}

func (m *enumFlagValue[T]) JoinAllowedValues(separator string) string {
	return strings.Join(m.allowedValues, separator)
}

func (m *enumFlagValue[T]) AllowedValuesWithHelp() []string {
	return m.allowedValuesWithHelp
}

func (m *enumFlagValue[T]) Value() T {
	return m.EnumValue
}

func (m enumValueWithHelpMap[T]) GetEnumValue(name string) (T, error) {
	enumValue, ok := m[name]
	if !ok {
		var emptyVal T
		return emptyVal, fmt.Errorf("invalid enum name '%s'", name)
	}
	return enumValue.Value, nil
}

func (m enumValueWithHelpMap[T]) GetEnumValueWithDefault(name string, defaultValue T) T {
	enumValue, ok := m[name]
	if !ok {
		return defaultValue
	}
	return enumValue.Value
}
