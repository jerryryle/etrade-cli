package jsonmap

import (
	"errors"
	"fmt"
)

// GetString attempts to retrieve a string with the specified key from the map.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a string.
func (m *JsonMap) GetString(key string) (string, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return "", err
	}
	return valueToString(value)
}

// GetInt attempts to retrieve an int with the specified key from the map.
// It will return an error if the key does not exist in the map or if the value
// at the index is not an int.
func (m *JsonMap) GetInt(key string) (int64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return 0, err
	}
	return valueToInt(value)
}

// GetFloat attempts to retrieve a float with the specified key from the map.
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

// GetBool attempts to retrieve a bool with the specified key from the map.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a bool.
func (m *JsonMap) GetBool(key string) (bool, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return false, err
	}
	return valueToBool(value)
}

// GetMap attempts to retrieve a JsonMap with the specified key from the map.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a JsonMap.
func (m *JsonMap) GetMap(key string) (JsonMap, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToMap(value)
}

// GetSlice attempts to retrieve a JsonSlice with the specified key from the map.
// It will return an error if the key does not exist in the map or if the value
// at the index is not a JsonSlice.
func (m *JsonMap) GetSlice(key string) (JsonSlice, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToSlice(value)
}

// GetStringSlice attempts to retrieve a []string with the specified key from
// the map. It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a string.
func (m *JsonMap) GetStringSlice(key string) ([]string, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToStringSlice(value)
}

// GetIntSlice attempts to retrieve a []int64 with the specified key from
// the map. It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not an int.
func (m *JsonMap) GetIntSlice(key string) ([]int64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToIntSlice(value)
}

// GetFloatSlice attempts to retrieve a []float64 with the specified key from
// the map. It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a float.
func (m *JsonMap) GetFloatSlice(key string) ([]float64, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToFloatSlice(value)
}

// GetBoolSlice attempts to retrieve a []bool64 with the specified key from
// the map. It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a bool.
func (m *JsonMap) GetBoolSlice(key string) ([]bool, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToBoolSlice(value)
}

// GetMapSlice attempts to retrieve a []JsonMap with the specified key from
// the map. It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a map.
func (m *JsonMap) GetMapSlice(key string) ([]JsonMap, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToMapSlice(value)
}

// GetSliceSlice attempts to retrieve a []JsonSlice with the specified key from
// the map. It will return an error if the key does not exist in the map or if
// the value at the index is not a JsonSlice or if any value in the slice is
// not a map.
func (m *JsonMap) GetSliceSlice(key string) ([]JsonSlice, error) {
	value, err := m.GetValue(key)
	if err != nil {
		return nil, err
	}
	return valueToSliceSlice(value)
}

// GetValue attempts to retrieve a value of any type with the specified key
// from the map. It will return an error if the key does not exist in the map.
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

// GetStringAtPath attempts to retrieve a string from the map at the specified
// path. It will return an error if the path is invalid or if the value at the
// path is not a string.
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

// GetIntAtPath attempts to retrieve an int from the map at the specified
// path. It will return an error if the path is invalid or if the value at the
// path is not an int.
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

// GetFloatAtPath attempts to retrieve a float from the map at the specified
// path. It will return an error if the path is invalid or if the value at the
// path is not a float (if the value at the index is an int, it
// will be quietly converted to a float).
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

// GetBoolAtPath attempts to retrieve a bool from the map at the specified
// path. It will return an error if the path is invalid or if the value at the
// path is not a bool.
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

// GetMapAtPath attempts to retrieve a JsonMap from the map at the specified
// path. It will return an error if the path is invalid or if the value at the
// path is not a JsonMap.
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

// GetSliceAtPath attempts to retrieve a JsonSlice from the map at the specified
// path. It will return an error if the path is invalid or if the value at the
// path is not a JsonSlice.
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

// GetSliceOfStringsAtPath attempts to retrieve a []string from the map at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// string.
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

// GetSliceOfIntsAtPath attempts to retrieve a []int64 from the map at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not an
// int.
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

// GetSliceOfFloatsAtPath attempts to retrieve a []float64 from the map at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// float.
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

// GetSliceOfBoolsAtPath attempts to retrieve a []bool from the map at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// bool.
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

// GetSliceOfMapsAtPath attempts to retrieve a []JsonMap from the map at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// map.
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

// GetSliceSliceAtPath attempts to retrieve a []JsonSlice from the map at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// slice.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForSlice" (map with a map with a slice value)
// or "keyForSliceSlice[0]" (map with a slice of slice values)
// or "keyForMap.keyForSliceSlice[0]" (map with a map with a slice of slice values)
func (m *JsonMap) GetSliceSliceAtPath(path string) ([]JsonSlice, error) {
	value, err := m.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToSliceSlice(value)
}

// GetValueAtPath attempts to retrieve value of any type from the map at the
// specified path. It will return an error if the path is invalid.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) GetValueAtPath(path string) (interface{}, error) {
	return pathGet(path, *m)
}
