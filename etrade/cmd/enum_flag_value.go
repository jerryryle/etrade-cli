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

type EnumFlagValue interface {
	String() string
	Set(value string) error
	Type() string
	AllowedValues() []string
	JoinAllowedValues(separator string) string
	AllowedValuesWithHelp() []string
}

type enumFlagValue[T comparable] struct {
	StringValue           string
	EnumValue             T
	valueMap              map[string]enumValueWithHelp[T]
	allowedValues         []string
	allowedValuesWithHelp []string
}

// newEnumFlagValue creates a new enumFlagValue with the default value set to defaultValue
func newEnumFlagValue[T comparable](
	valueMap map[string]enumValueWithHelp[T], defaultEnumValue T,
) *enumFlagValue[T] {
	defaultStringValue := ""
	allowedValues := make([]string, 0, len(valueMap))
	allowedValuesWithHelp := make([]string, 0, len(valueMap))

	// TODO: Don't build help in this function. Do it dynamically when needed.
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
	enumValue, ok := m.valueMap[value]
	if !ok {
		return fmt.Errorf("%s is not one of the allowed values: [%s]", value, strings.Join(m.allowedValues, ", "))
	}
	m.StringValue = value
	m.EnumValue = enumValue.Value
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
