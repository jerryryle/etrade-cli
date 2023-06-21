package jsonmap

import "fmt"

// GetString attempts to retrieve a string from the slice at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a string.
func (s *JsonSlice) GetString(index int) (string, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return "", err
	}
	return valueToString(value)
}

// GetInt attempts to retrieve an int from the slice at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not an int.
func (s *JsonSlice) GetInt(index int) (int64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return 0, err
	}
	return valueToInt(value)
}

// GetFloat attempts to retrieve a float from the slice at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a float (if the value at the index is an
// int, it will be quietly converted to a float).
func (s *JsonSlice) GetFloat(index int) (float64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return 0, err
	}
	return valueToFloat(value)
}

// GetBool attempts to retrieve a bool from the slice at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a bool.
func (s *JsonSlice) GetBool(index int) (bool, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return false, err
	}
	return valueToBool(value)
}

// GetMap attempts to retrieve a JsonMap from the slice at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonMap.
func (s *JsonSlice) GetMap(index int) (JsonMap, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToMap(value)
}

// GetSlice attempts to retrieve a JsonSlice from the slice at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice.
func (s *JsonSlice) GetSlice(index int) (JsonSlice, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToSlice(value)
}

// GetStringSlice attempts to retrieve a []string with the specified key from
// the slice. It will return an error if the index is out of bounds for the
// slice or if the value at the index is not a JsonSlice or if any value in the
// slice is not a string.
func (s *JsonSlice) GetStringSlice(index int) ([]string, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToStringSlice(value)
}

// GetIntSlice attempts to retrieve a []int64 with the specified key from
// the slice. It will return an error if the index is out of bounds for the
// slice or if the value at the index is not a JsonSlice or if any value in the
// slice is not an int.
func (s *JsonSlice) GetIntSlice(index int) ([]int64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToIntSlice(value)
}

// GetFloatSlice attempts to retrieve a []float64 with the specified key from
// the slice. It will return an error if the index is out of bounds for the
// slice or if the value at the index is not a JsonSlice or if any value in the
// slice is not a float.
func (s *JsonSlice) GetFloatSlice(index int) ([]float64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToFloatSlice(value)
}

// GetBoolSlice attempts to retrieve a []bool64 with the specified key from
// the slice. It will return an error if the index is out of bounds for the
// slice or if the value at the index is not a JsonSlice or if any value in the
// slice is not a bool.
func (s *JsonSlice) GetBoolSlice(index int) ([]bool, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToBoolSlice(value)
}

// GetMapSlice attempts to retrieve a []JsonMap with the specified key from
// the slice. It will return an error if the index is out of bounds for the
// slice or if the value at the index is not a JsonSlice or if any value in the
// slice is not a map.
func (s *JsonSlice) GetMapSlice(index int) ([]JsonMap, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToMapSlice(value)
}

// GetSliceSlice attempts to retrieve a []JsonSlice with the specified key from
// the slice. It will return an error if the index is out of bounds for the
// slice or if the value at the index is not a JsonSlice or if any value in the
// slice is not a slice.
func (s *JsonSlice) GetSliceSlice(index int) ([]JsonSlice, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToSliceSlice(value)
}

// GetValue attempts to retrieve a value of any type from the slice at the
// specified index. It will return an error if the index is out of bounds for
// the slice.
func (s *JsonSlice) GetValue(index int) (interface{}, error) {
	if index >= len(*s) {
		return nil, fmt.Errorf("index %d is greater than slice len %d", index, len(*s))
	}
	if index < 0 {
		return nil, fmt.Errorf("index %d is negative", index)
	}
	value := (*s)[index]
	return value, nil
}

// GetStringAtPath attempts to retrieve a string from the slice at the
// specified path. It will return an error if the path is invalid or if the
// value at the path is not a string.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForString" (slice of maps with string value)
// or "[0][0]" (slice of slices of string values)
func (s *JsonSlice) GetStringAtPath(path string) (string, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return "", err
	}
	return valueToString(value)
}

// GetIntAtPath attempts to retrieve an int from the slice at the
// specified path. It will return an error if the path is invalid or if the
// value at the path is not an int.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForInt" (slice of maps with int value)
// or "[0][0]" (slice of slices of int values)
func (s *JsonSlice) GetIntAtPath(path string) (int64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return 0, err
	}
	return valueToInt(value)
}

// GetFloatAtPath attempts to retrieve a float from the slice at the
// specified path. It will return an error if the path is invalid or if the
// value at the path is not a float (if the value at the index is an
// int, it will be quietly converted to a float).
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForFloat" (slice of maps with float value)
// or "[0][0]" (slice of slices of float values)
func (s *JsonSlice) GetFloatAtPath(path string) (float64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return 0, err
	}
	return valueToFloat(value)
}

// GetBoolAtPath attempts to retrieve a bool from the slice at the
// specified path. It will return an error if the path is invalid or if the
// value at the path is not a bool.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForBool" (slice of maps with bool value)
// or "[0][0]" (slice of slices of bool values)
func (s *JsonSlice) GetBoolAtPath(path string) (bool, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return false, err
	}
	return valueToBool(value)
}

// GetMapAtPath attempts to retrieve a JsonMap from the slice at the
// specified path. It will return an error if the path is invalid or if the
// value at the path is not a JsonMap.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForMap" (slice of maps with map value)
// or "[0][0]" (slice of slices of map values)
func (s *JsonSlice) GetMapAtPath(path string) (JsonMap, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToMap(value)
}

// GetSliceAtPath attempts to retrieve a JsonSlice from the slice at the
// specified path. It will return an error if the path is invalid or if the
// value at the path is not a JsonSlice.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceAtPath(path string) (JsonSlice, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToSlice(value)
}

// GetStringSliceAtPath attempts to retrieve a []string from the slice at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// string.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetStringSliceAtPath(path string) ([]string, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToStringSlice(value)
}

// GetIntSliceAtPath attempts to retrieve a []int64 from the slice at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not an
// int.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetIntSliceAtPath(path string) ([]int64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToIntSlice(value)
}

// GetFloatSliceAtPath attempts to retrieve a []float64 from the slice at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// float.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetFloatSliceAtPath(path string) ([]float64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToFloatSlice(value)
}

// GetBoolSliceAtPath attempts to retrieve a []bool from the slice at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// bool.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetBoolSliceAtPath(path string) ([]bool, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToBoolSlice(value)
}

// GetMapSliceAtPath attempts to retrieve a []JsonMap from the slice at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// map.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetMapSliceAtPath(path string) ([]JsonMap, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToMapSlice(value)
}

// GetSliceSliceAtPath attempts to retrieve a []JsonSlice from the slice at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// slice.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceSliceAtPath(path string) ([]JsonSlice, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToSliceSlice(value)
}

// GetValueAtPath attempts to retrieve a value of any type from the slice at
// the specified path. It will return an error if the path is invalid.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) GetValueAtPath(path string) (interface{}, error) {
	return pathGet(*s, path)
}
