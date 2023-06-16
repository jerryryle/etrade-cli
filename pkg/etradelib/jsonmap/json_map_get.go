package jsonmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (m JsonMap) GetString(key string) (string, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return "", err
	}
	return valueToString(value)
}

func (m JsonMap) GetInt(key string) (int64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return 0, err
	}
	return valueToInt(value)
}

func (m JsonMap) GetFloat(key string) (float64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return 0, err
	}
	return valueToFloat(value)
}

func (m JsonMap) GetBool(key string) (bool, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return false, err
	}
	return valueToBool(value)
}

func (m JsonMap) GetMap(key string) (JsonMap, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToMap(value)
}

func (m JsonMap) GetSlice(key string) ([]interface{}, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToSlice(value)
}

func (m JsonMap) GetValue(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("no key provided")
	}

	value, found := m[key]
	if !found {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

func (m JsonMap) GetStringAtPath(path string) (string, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return "", err
	}
	return valueToString(value)
}

func (m JsonMap) GetIntAtPath(path string) (int64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return 0, err
	}
	return valueToInt(value)
}

func (m JsonMap) GetFloatAtPath(path string) (float64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return 0, err
	}
	return valueToFloat(value)
}

func (m JsonMap) GetBoolAtPath(path string) (bool, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return false, err
	}
	return valueToBool(value)
}

func (m JsonMap) GetMapAtPath(path string) (JsonMap, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToMap(value)
}

func (m JsonMap) GetSliceAtPath(path string) ([]interface{}, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToSlice(value)
}

func (m JsonMap) GetValueAtPath(path string) (interface{}, error) {
	pathElements, err := parsePath(path)
	if err != nil {
		return nil, err
	}
	// Start with the root of the map as the current value
	var currentValue interface{} = m
	// We'll build up the current path as we traverse path elements to help produce better error messages.
	var currentPath = ""

	// Iterate over the path elements.
	for _, pathElement := range pathElements {
		switch pathElementTyped := pathElement.(type) {
		// If the path element is a string, use it to index the current element as an array
		case int:
			lastPath := currentPath
			currentPath = currentPath + fmt.Sprintf("[%d]", pathElementTyped)
			switch currentValueTyped := currentValue.(type) {
			case []interface{}:
				if pathElementTyped >= len(currentValueTyped) {
					return nil, fmt.Errorf(
						"cannot access %s because array index %d is out of bounds", currentPath, pathElementTyped,
					)
				}
				currentValue = currentValueTyped[pathElementTyped]
			default:
				return nil, fmt.Errorf("cannot access %s because %s is not an array", currentPath, lastPath)
			}
		// If the path element is a string, use it to index the current element as a map
		case string:
			lastPath := currentPath
			currentPath = currentPath + "." + pathElementTyped
			switch currentValueTyped := currentValue.(type) {
			case JsonMap:
				newValue, found := currentValueTyped[pathElementTyped]
				if !found {
					return nil, fmt.Errorf(
						"cannot access %s because key %s is not found in parent map", currentPath, pathElementTyped,
					)
				}
				currentValue = newValue
			default:
				return nil, fmt.Errorf("cannot access %s because %s is not a map", currentPath, lastPath)
			}
		default:
			return nil, errors.New("internal error evaluating path elements to get value")
		}
	}
	return currentValue, nil
}

func valueToString(value interface{}) (string, error) {
	switch valueTyped := value.(type) {
	case string:
		return valueTyped, nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("type %T is not a string", valueTyped)
	}
}

func valueToInt(value interface{}) (int64, error) {
	switch valueTyped := value.(type) {
	case int64:
		return valueTyped, nil
	case int32:
		return int64(valueTyped), nil
	case int:
		return int64(valueTyped), nil
	case json.Number:
		intVal, err := valueTyped.Int64()
		if err != nil {
			return 0, fmt.Errorf("type %T is not an int: %w", valueTyped, err)
		}
		return intVal, nil
	default:
		return 0, fmt.Errorf("type %T is not an int", valueTyped)
	}
}

func valueToFloat(value interface{}) (float64, error) {
	switch valueTyped := value.(type) {
	case float64:
		return valueTyped, nil
	case float32:
		return float64(valueTyped), nil
	case json.Number:
		floatVal, err := valueTyped.Float64()
		if err != nil {
			return 0, fmt.Errorf("type %T is not a float: %w", valueTyped, err)
		}
		return floatVal, nil
	default:
		return 0, fmt.Errorf("type %T is not a float", valueTyped)
	}
}

func valueToBool(value interface{}) (bool, error) {
	switch valueTyped := value.(type) {
	case bool:
		return valueTyped, nil
	default:
		return false, fmt.Errorf("type %T is not a bool", valueTyped)
	}
}

func valueToMap(value interface{}) (JsonMap, error) {
	switch valueTyped := value.(type) {
	case JsonMap:
		return valueTyped, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("type %T is not a map", valueTyped)
	}
}

func valueToSlice(value interface{}) ([]interface{}, error) {
	switch valueTyped := value.(type) {
	case []interface{}:
		return valueTyped, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("type %T is not a slice", valueTyped)
	}
}

// parsePath attempts to parse a path into a slice consisting of any key strings
// and index integers. Leading and/or redundant path dots will be ignored.
// e.g. parsePath("foo.bar[1].moo")  -> []interface{}{"foo", "bar", 1, "moo"}
// e.g. parsePath(".foo.bar[1].moo") -> []interface{}{"foo", "bar", 1, "moo"}
// e.g. parsePath("foo..bar[1].moo") -> []interface{}{"foo", "bar", 1, "moo"}
func parsePath(path string) ([]interface{}, error) {
	pathElements := strings.Split(path, ".")
	returnPathElements := make([]interface{}, 0, len(pathElements)*2)
	for _, pathElement := range pathElements {
		if pathElement != "" {
			keyIndexPathElements, err := splitKeyIndices(pathElement)
			if err != nil {
				return nil, err
			}
			returnPathElements = append(returnPathElements, keyIndexPathElements...)
		}
	}
	return returnPathElements, nil
}

// splitKeyIndices splits a key and array indices into a slice containing the
// key string and zero or more integer indices. If the key has no indices, then
// the return slice will contain only the key.
// e.g. splitKeyIndices("test") -> []interface{}{"test"}
// e.g. splitKeyIndices("test[0]") -> []interface{}{"test", 0}
// e.g. splitKeyIndices("test[0][1]") -> []interface{}{"test", 0, 1}
// If the input string contains missing/spurious characters around the indices
// or the index values cannot be parsed as integers, then splitKeyIndices will
// return an error.
func splitKeyIndices(s string) ([]interface{}, error) {
	// Try to split the string on opening brackets.
	// e.g. s:"test[0][1]" -> keyAndIndicesSlice:[ "test", "0]", "1]" ]
	keyAndIndicesSlice := strings.Split(s, "[")
	// If the length of the resulting slice is 1 (it will never be zero--even
	// for an empty string, but we compare against <= 1 to be extra paranoid),
	// then no opening bracket was found. Just return a slice with the original
	// string since there's no index to extract.
	if len(keyAndIndicesSlice) <= 1 {
		return []interface{}{s}, nil
	}

	// Create a new slice for the key and indices, starting with the key
	// e.g. keyAndIndicesSlice:[ "test", "0]", "1]" ] -> keyAndIndicesSlice[0]:"test"
	returnKeyAndIndices := []interface{}{keyAndIndicesSlice[0]}

	// Now extract index value(s)
	// e.g. keyAndIndicesSlice:[ "test", "0]", "1]" ] -> keyAndIndicesSlice[1:]:[ "0]", "1]" ]
	for _, indexString := range keyAndIndicesSlice[1:] {
		// Try to split the string on closing brackets.
		// e.g. indexString:"0]" -> indexSlice: [ "0", "" ]
		indexSlice := strings.Split(indexString, "]")
		// If the length of the resulting slice is not exactly 2, then
		// the closing bracket is missing.
		// e.g. indexString:"0" -> indexSlice: [ "0" ]
		//
		// If the 2nd element of the slice isn't an empty string, then there
		// are spurious characters after the closing bracket.
		// indexString:"0]foo" -> indexSlice: [ "0", "foo" ]
		if len(indexSlice) != 2 || indexSlice[1] != "" {
			return nil, fmt.Errorf("invalid index in key %s", s)
		}
		// Convert the string index value to an integer
		// e.g. indexSlice[0]:"0" -> indexValue: 0
		indexValue, err := strconv.Atoi(indexSlice[0])
		if err != nil {
			return nil, fmt.Errorf("invalid index in key %s", s)
		}
		// Append the index integer to the return slice
		returnKeyAndIndices = append(returnKeyAndIndices, indexValue)
	}
	return returnKeyAndIndices, nil
}
