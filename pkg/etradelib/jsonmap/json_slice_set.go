package jsonmap

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// SetString sets a string at the specified index in the slice.
// It will not grow the slice, and it will return an error if the index is out
// of bounds for the slice.
func (s *JsonSlice) SetString(index int, value string) error {
	return s.SetValue(index, value)
}

// SetInt sets an int at the specified index in the slice.
// It will not grow the slice, and it will return an error if the index is out
// of bounds for the slice.
func (s *JsonSlice) SetInt(index int, value int64) error {
	valueNumber := json.Number(fmt.Sprintf("%d", value))
	return s.SetValue(index, valueNumber)
}

// SetFloat sets a float at the specified index in the slice.
// It will not grow the slice, and it will return an error if the index is out
// of bounds for the slice.
func (s *JsonSlice) SetFloat(index int, value float64) error {
	valueNumber := json.Number(strconv.FormatFloat(value, 'f', -1, 64))
	return s.SetValue(index, valueNumber)
}

// SetBool sets a bool at the specified index in the slice.
// It will not grow the slice, and it will return an error if the index is out
// of bounds for the slice.
func (s *JsonSlice) SetBool(index int, value bool) error {
	return s.SetValue(index, value)
}

// SetMap sets a JsonMap at the specified index in the slice.
// It will not grow the slice, and it will return an error if the index is out
// of bounds for the slice.
// Use caution with this method. If you add a map that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the slice.
func (s *JsonSlice) SetMap(index int, value JsonMap) error {
	return s.SetValue(index, value)
}

// SetSlice sets a JsonSlice at the specified index in the slice.
// It will not grow the slice, and it will return an error if the index is out
// of bounds for the slice.
// Use caution with this method. If you add a slice that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the slice.
func (s *JsonSlice) SetSlice(index int, value JsonSlice) error {
	return s.SetValue(index, value)
}

// SetValue sets an object of any type at the specified index in the slice.
// It will not grow the slice, and it will return an error if the index is out
// of bounds for the slice.
// Use caution with this method. If you add an object that contains invalid
// Json elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the slice.
func (s *JsonSlice) SetValue(index int, value interface{}) error {
	if index < 0 || index >= len(*s) {
		return fmt.Errorf("cannot set value: slice index %d out of bounds for slice of length %d", index, len(*s))
	}
	(*s)[index] = value
	return nil
}

// SetStringAtPath sets a string at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Note that slice paths should always begin with an index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) SetStringAtPath(path string, value string) error {
	return s.SetValueAtPath(path, value)
}

// SetIntAtPath sets an int at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Note that slice paths should always begin with an index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) SetIntAtPath(path string, value int64) error {
	valueNumber := json.Number(fmt.Sprintf("%d", value))
	return s.SetValueAtPath(path, valueNumber)
}

// SetFloatAtPath sets a float at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Note that slice paths should always begin with an index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) SetFloatAtPath(path string, value float64) error {
	valueNumber := json.Number(strconv.FormatFloat(value, 'f', -1, 64))
	return s.SetValueAtPath(path, valueNumber)
}

// SetBoolAtPath sets a bool at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Note that slice paths should always begin with an index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) SetBoolAtPath(path string, value bool) error {
	return s.SetValueAtPath(path, value)
}

// SetMapAtPath sets a JsonMap at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Use caution with this method. If you add a map that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the slice.
// Note that slice paths should always begin with an index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) SetMapAtPath(path string, value JsonMap) error {
	return s.SetValueAtPath(path, value)
}

// SetSliceAtPath ses a JsonSlice at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Use caution with this method. If you add a slice that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the slice.
// Note that slice paths should always begin with an index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) SetSliceAtPath(path string, value JsonSlice) error {
	return s.SetValueAtPath(path, value)
}

// SetValueAtPath sets an object of any type at the specified path.
// It will attempt to create missing intermediate map elements of the path.
// It will fail to create missing intermediate slice elements.
// Use caution with this method. If you add an object that contains invalid Json
// elements (e.g. arbitrary structs or pointers to objects), then you may
// break the ability to traverse the slice.
// Note that slice paths should always begin with an index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) SetValueAtPath(path string, value interface{}) error {
	return pathSet(*s, path, value)
}
