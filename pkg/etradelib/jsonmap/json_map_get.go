package jsonmap

import (
	"errors"
	"fmt"
)

// GetString retrieves, from the map, a string with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a string.
func (m *JsonMap) GetString(key string) (string, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return "", err
	}
	return valueToString(value)
}

// GetInt retrieves, from the map, an int with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not an int.
func (m *JsonMap) GetInt(key string) (int64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return 0, err
	}
	return valueToInt(value)
}

// GetFloat retrieves, from the map, a float with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a float (if the value at the index is an integer, it
// will be quietly converted to a float).
func (m *JsonMap) GetFloat(key string) (float64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return 0, err
	}
	return valueToFloat(value)
}

// GetBool retrieves, from the map, a bool with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a bool.
func (m *JsonMap) GetBool(key string) (bool, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return false, err
	}
	return valueToBool(value)
}

// GetMap retrieves, from the map, a JsonMap with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a JsonMap.
func (m *JsonMap) GetMap(key string) (JsonMap, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToMap(value)
}

// GetSlice retrieves, from the map, a JsonSlice with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a JsonSlice.
func (m *JsonMap) GetSlice(key string) (JsonSlice, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToSlice(value)
}

// GetSliceOfStrings retrieves, from the map, a []string with the specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a string.
func (m *JsonMap) GetSliceOfStrings(key string) ([]string, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToStringSlice(value)
}

// GetSliceOfInts retrieves, from the map, a []int64 with the specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not an int.
func (m *JsonMap) GetSliceOfInts(key string) ([]int64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToIntSlice(value)
}

// GetSliceOfFloats retrieves, from the map, a []float64 with the specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a float.
func (m *JsonMap) GetSliceOfFloats(key string) ([]float64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToFloatSlice(value)
}

// GetSliceOfBools retrieves, from the map, a []bool64 with the specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a bool.
func (m *JsonMap) GetSliceOfBools(key string) ([]bool, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToBoolSlice(value)
}

// GetSliceOfMaps retrieves, from the map, a []JsonMap with the specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a JsonMap.
func (m *JsonMap) GetSliceOfMaps(key string) ([]JsonMap, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToMapSlice(value)
}

// GetSliceOfSlices retrieves, from the map, a []JsonSlice with the specified
// key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a JsonSlice.
func (m *JsonMap) GetSliceOfSlices(key string) ([]JsonSlice, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToSliceSlice(value)
}

// GetValue retrieves, from the map, an object of any type with the specified
// key.
// It will return an error if the key does not exist in the map.
func (m *JsonMap) GetValue(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("no key provided")
	}

	value, found := (*m)[key]
	if !found {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

// GetStringWithDefault retrieves, from the map, a string with the specified
// key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a string.
func (m *JsonMap) GetStringWithDefault(key string, defaultValue string) (string, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToString(value)
}

// GetIntWithDefault retrieves, from the map, an int with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not an int.
func (m *JsonMap) GetIntWithDefault(key string, defaultValue int64) (int64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToInt(value)
}

// GetFloatWithDefault retrieves, from the map, a float with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a float (if the value at the index is an integer, it
// will be quietly converted to a float).
func (m *JsonMap) GetFloatWithDefault(key string, defaultValue float64) (float64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToFloat(value)
}

// GetBoolWithDefault retrieves, from the map, a bool with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a bool.
func (m *JsonMap) GetBoolWithDefault(key string, defaultValue bool) (bool, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToBool(value)
}

// GetMapWithDefault retrieves, from the map, a JsonMap with the specified key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a JsonMap.
func (m *JsonMap) GetMapWithDefault(key string, defaultValue JsonMap) (JsonMap, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToMap(value)
}

// GetSliceWithDefault retrieves, from the map, a JsonSlice with the specified
// key.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a JsonSlice.
func (m *JsonMap) GetSliceWithDefault(key string, defaultValue JsonSlice) (JsonSlice, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToSlice(value)
}

// GetSliceOfStringsWithDefault retrieves, from the map, a []string with the
// specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a string.
func (m *JsonMap) GetSliceOfStringsWithDefault(key string, defaultValue []string) ([]string, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToStringSlice(value)
}

// GetSliceOfIntsWithDefault retrieves, from the map, a []int64 with the
// specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not an int.
func (m *JsonMap) GetSliceOfIntsWithDefault(key string, defaultValue []int64) ([]int64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToIntSlice(value)
}

// GetSliceOfFloatsWithDefault retrieves, from the map, a []float64 with the
// specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a float.
func (m *JsonMap) GetSliceOfFloatsWithDefault(key string, defaultValue []float64) ([]float64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToFloatSlice(value)
}

// GetSliceOfBoolsWithDefault retrieves, from the map, a []bool64 with the
// specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a bool.
func (m *JsonMap) GetSliceOfBoolsWithDefault(key string, defaultValue []bool) ([]bool, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToBoolSlice(value)
}

// GetSliceOfMapsWithDefault retrieves, from the map, a []JsonMap with the
// specified key.
// It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a JsonMap.
func (m *JsonMap) GetSliceOfMapsWithDefault(key string, defaultValue []JsonMap) ([]JsonMap, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToMapSlice(value)
}

// GetSliceOfSlicesWithDefault retrieves, from the map, a []JsonSlice with the
// specified key.
// If the key can not be found in the map (including an invalid,
// empty key, ""), then it returns the default value. It will return an error
// if the value at the index is not a JsonSlice or if any value in the slice is
// not a JsonSlice.
func (m *JsonMap) GetSliceOfSlicesWithDefault(key string, defaultValue []JsonSlice) ([]JsonSlice, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue, nil
	}
	return valueToSliceSlice(value)
}

// GetValueWithDefault retrieves, from the map, an object of any type with the
// specified key.
// If the key can not be found in the map (including an invalid,
// empty key, ""), then it returns the default value.
func (m *JsonMap) GetValueWithDefault(key string, defaultValue interface{}) interface{} {
	value, err := m.GetValue(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetStringAtPath retrieves, from the map, a string at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a string.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForString" (map with a map with a string value)
// or "keyForStringSlice[0]" (map with a slice of string values)
// or "keyForMap.keyForStringSlice[0]" (map with a map with a slice of string values)
func (m *JsonMap) GetStringAtPath(path string) (string, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return "", err
	}
	return valueToString(value)
}

// GetIntAtPath retrieves, from the map, an int at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not an int.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForInt" (map with a map with an int value)
// or "keyForIntSlice[0]" (map with a slice of int values)
// or "keyForMap.keyForIntSlice[0]" (map with a map with a slice of int values)
func (m *JsonMap) GetIntAtPath(path string) (int64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return 0, err
	}
	return valueToInt(value)
}

// GetFloatAtPath retrieves, from the map, a float at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a float (if the value at the index is an int, it will be quietly
// converted to a float).
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForFloat" (map with a map with a float value)
// or "keyForFloatSlice[0]" (map with a slice of float values)
// or "keyForMap.keyForFloatSlice[0]" (map with a map with a slice of float values)
func (m *JsonMap) GetFloatAtPath(path string) (float64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return 0, err
	}
	return valueToFloat(value)
}

// GetBoolAtPath retrieves, from the map, a bool at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a bool.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForBool" (map with a map with a bool value)
// or "keyForBoolSlice[0]" (map with a slice of bool values)
// or "keyForMap.keyForBoolSlice[0]" (map with a map with a slice of bool values)
func (m *JsonMap) GetBoolAtPath(path string) (bool, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return false, err
	}
	return valueToBool(value)
}

// GetMapAtPath retrieves, from the map, a JsonMap at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a JsonMap.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForMap" (map with a map with a map value)
// or "keyForMapSlice[0]" (map with a slice of map values)
// or "keyForMap.keyForMapSlice[0]" (map with a map with a slice of map values)
func (m *JsonMap) GetMapAtPath(path string) (JsonMap, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToMap(value)
}

// GetSliceAtPath retrieves, from the map, a JsonSlice at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a JsonSlice.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceAtPath(path string) (JsonSlice, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToSlice(value)
}

// GetSliceOfStringsAtPath retrieves, from the map, a []string at the specified
// path. It will return an error if the path is invalid, if the value at the
// path is not a JsonSlice, or if any value in the slice is not a string.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfStringsAtPath(path string) ([]string, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToStringSlice(value)
}

// GetSliceOfIntsAtPath retrieves, from the map, a []int64 at the specified
// path. It will return an error if the path is invalid, if the value at the
// path is not a JsonSlice, or if any value in the slice is not an int.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfIntsAtPath(path string) ([]int64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToIntSlice(value)
}

// GetSliceOfFloatsAtPath retrieves, from the map, a []float64 at the specified
// path. It will return an error if the path is invalid, if the value at the
// path is not a JsonSlice, or if any value in the slice is not a float.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfFloatsAtPath(path string) ([]float64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToFloatSlice(value)
}

// GetSliceOfBoolsAtPath retrieves, from the map, a []bool at the specified
// path. It will return an error if the path is invalid, if the value at the
// path is not a JsonSlice, or if any value in the slice is not a bool.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfBoolsAtPath(path string) ([]bool, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToBoolSlice(value)
}

// GetSliceOfMapsAtPath retrieves, from the map, a []JsonMap at the specified
// path. It will return an error if the path is invalid, if the value at the
// path is not a JsonSlice, or if any value in the slice is not a JsonMap.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfMapsAtPath(path string) ([]JsonMap, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToMapSlice(value)
}

// GetSliceOfSlicesAtPath retrieves, from the map, a []JsonSlice at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// JsonSlice.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfSlicesAtPath(path string) ([]JsonSlice, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToSliceSlice(value)
}

// GetValueAtPath retrieves, from the map, an object of any type at the
// specified path. It will return an error if the path is invalid.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) GetValueAtPath(path string) (interface{}, error) {
	return pathGet(*m, path)
}

// GetStringAtPathWithDefault retrieves, from the map, a string at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a string.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForString" (map with a map with a string value)
// or "keyForStringSlice[0]" (map with a slice of string values)
// or "keyForMap.keyForStringSlice[0]" (map with a map with a slice of string values)
func (m *JsonMap) GetStringAtPathWithDefault(path string, defaultValue string) (string, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToString(value)
}

// GetIntAtPathWithDefault retrieves, from the map, an int at the specified
// path. If the value cannot be found for any reason (including an invalid
// path), then it returns the default value. It will return an error if the
// value at the path is not an int.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForInt" (map with a map with an int value)
// or "keyForIntSlice[0]" (map with a slice of int values)
// or "keyForMap.keyForIntSlice[0]" (map with a map with a slice of int values)
func (m *JsonMap) GetIntAtPathWithDefault(path string, defaultValue int64) (int64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToInt(value)
}

// GetFloatAtPathWithDefault retrieves, from the map, a float at the specified
// path. If the value cannot be found for any reason (including an invalid
// path), then it returns the default value. It will return an error if the
// value at the path is not a float (if the value at the index is an int, it
// will be quietly converted to a float).
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForFloat" (map with a map with a float value)
// or "keyForFloatSlice[0]" (map with a slice of float values)
// or "keyForMap.keyForFloatSlice[0]" (map with a map with a slice of float values)
func (m *JsonMap) GetFloatAtPathWithDefault(path string, defaultValue float64) (float64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToFloat(value)
}

// GetBoolAtPathWithDefault retrieves, from the map, a bool at the specified
// path. If the value cannot be found for any reason (including an invalid
// path), then it returns the default value. It will return an error if the
// value at the path is not a bool.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForBool" (map with a map with a bool value)
// or "keyForBoolSlice[0]" (map with a slice of bool values)
// or "keyForMap.keyForBoolSlice[0]" (map with a map with a slice of bool values)
func (m *JsonMap) GetBoolAtPathWithDefault(path string, defaultValue bool) (bool, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToBool(value)
}

// GetMapAtPathWithDefault retrieves, from the map, a JsonMap at the specified
// path. If the value cannot be found for any reason (including an invalid
// path), then it returns the default value. It will return an error if the
// value at the path is not a JsonMap.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForMap" (map with a map with a map value)
// or "keyForMapSlice[0]" (map with a slice of map values)
// or "keyForMap.keyForMapSlice[0]" (map with a map with a slice of map values)
func (m *JsonMap) GetMapAtPathWithDefault(path string, defaultValue JsonMap) (JsonMap, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToMap(value)
}

// GetSliceAtPathWithDefault retrieves, from the map, a JsonSlice at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a JsonSlice.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceAtPathWithDefault(path string, defaultValue JsonSlice) (JsonSlice, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToSlice(value)
}

// GetSliceOfStringsAtPathWithDefault retrieves, from the map, a []string at
// the specified path. If the value cannot be found for any reason (including
// an invalid path), then it returns the default value. It will return an error
// if the value at the path is not a JsonSlice or if any value in the slice is
// not a string.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfStringsAtPathWithDefault(path string, defaultValue []string) ([]string, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToStringSlice(value)
}

// GetSliceOfIntsAtPathWithDefault retrieves, from the map, a []int64 at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a JsonSlice or if any value in the slice is not
// an int.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfIntsAtPathWithDefault(path string, defaultValue []int64) ([]int64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToIntSlice(value)
}

// GetSliceOfFloatsAtPathWithDefault retrieves, from the map, a []float64 at
// the specified path. If the value cannot be found for any reason (including
// an invalid path), then it returns the default value. It will return an error
// if the value at the path is not a JsonSlice or if any value in the slice is
// not a float.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfFloatsAtPathWithDefault(path string, defaultValue []float64) ([]float64, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToFloatSlice(value)
}

// GetSliceOfBoolsAtPathWithDefault retrieves, from the map, a []bool at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a JsonSlice or if any value in the slice is not
// a bool.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfBoolsAtPathWithDefault(path string, defaultValue []bool) ([]bool, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToBoolSlice(value)
}

// GetSliceOfMapsAtPathWithDefault retrieves, from the map, a []JsonMap at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a JsonSlice or if any value in the slice is not
// a JsonMap.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfMapsAtPathWithDefault(path string, defaultValue []JsonMap) ([]JsonMap, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToMapSlice(value)
}

// GetSliceOfSlicesAtPathWithDefault retrieves, from the map, a []JsonSlice at
// the specified path. If the value cannot be found for any reason (including
// an invalid path), then it returns the default value. It will return an error
// if the value at the path is not a JsonSlice or if any value in the slice is
// not a JsonSlice.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceOfSlicesAtPathWithDefault(path string, defaultValue []JsonSlice) ([]JsonSlice, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToSliceSlice(value)
}

// GetValueAtPathWithDefault retrieves, from the map, an object of any type at
// the specified path. If the value cannot be found for any reason (including
// an invalid path), then it returns the default value.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) GetValueAtPathWithDefault(path string, defaultValue interface{}) interface{} {
	value, err := pathGet(*m, path)
	if err == nil {
		return value
	}
	return defaultValue
}
