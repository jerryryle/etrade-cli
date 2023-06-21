package jsonmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// SetString sets a string with the specified key in the map.
// It will only return an error if an empty key is provided.
func (m *JsonMap) SetString(key string, value string) error {
	return m.SetValue(key, value)
}

// SetInt sets an int with the specified key in the map.
// It will only return an error if an empty key is provided.
func (m *JsonMap) SetInt(key string, value int64) error {
	valueNumber := json.Number(fmt.Sprintf("%d", value))
	return m.SetValue(key, valueNumber)
}

// SetFloat sets a float with the specified key in the map.
// It will only return an error if an empty key is provided.
func (m *JsonMap) SetFloat(key string, value float64) error {
	valueNumber := json.Number(strconv.FormatFloat(value, 'f', -1, 64))
	return m.SetValue(key, valueNumber)
}

// SetBool sets a bool with the specified key in the map.
// It will only return an error if an empty key is provided.
func (m *JsonMap) SetBool(key string, value bool) error {
	return m.SetValue(key, value)
}

// SetMap sets a map with the specified key in the map.
// It will only return an error if an empty key is provided.
// Use caution with this method. If you add a map that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the map.
func (m *JsonMap) SetMap(key string, value JsonMap) error {
	return m.SetValue(key, value)
}

// SetSlice sets a slice with the specified key in the map.
// It will only return an error if an empty key is provided.
// Use caution with this method. If you add a slice that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the map.
func (m *JsonMap) SetSlice(key string, value JsonSlice) error {
	return m.SetValue(key, value)
}

// SetValue sets an object of any type with the specified key in the map.
// It will only return an error if an empty key is provided.
// Use caution with this method. If you add an object that contains invalid
// Json elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the map.
func (m *JsonMap) SetValue(key string, value interface{}) error {
	if key == "" {
		return errors.New("no key provided")
	}
	(*m)[key] = value
	return nil
}

// SetStringAtPath sets a string value in the map at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) SetStringAtPath(path string, value string) error {
	return m.SetValueAtPath(path, value)
}

// SetIntAtPath sets an int value in the map at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) SetIntAtPath(path string, value int64) error {
	valueNumber := json.Number(fmt.Sprintf("%d", value))
	return m.SetValueAtPath(path, valueNumber)
}

// SetFloatAtPath sets a float value in the map at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) SetFloatAtPath(path string, value float64) error {
	valueNumber := json.Number(strconv.FormatFloat(value, 'f', -1, 64))
	return m.SetValueAtPath(path, valueNumber)
}

// SetBoolAtPath sets a bool value in the map at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) SetBoolAtPath(path string, value bool) error {
	return m.SetValueAtPath(path, value)
}

// SetMapAtPath sets a JsonMap value in the map at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Use caution with this method. If you add a map that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the map.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) SetMapAtPath(path string, value JsonMap) error {
	return m.SetValueAtPath(path, value)
}

// SetSliceAtPath sets a JsonSlice value in the map at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Use caution with this method. If you add a slice that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the map.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) SetSliceAtPath(path string, value JsonSlice) error {
	return m.SetValueAtPath(path, value)
}

// SetValueAtPath sets an object of any type in the map at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Use caution with this method. If you add an object that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the map.
// Note that map paths should always begin with a key.
// e.g. "keyForMap.keyForValue" (map with a map with a value)
// or "keyForValueSlice[0]" (map with a slice of values)
// or "keyForMap.keyForValueSlice[0]" (map with a map with a slice of values)
func (m *JsonMap) SetValueAtPath(path string, value interface{}) error {
	return pathSet(*m, path, value)
}
