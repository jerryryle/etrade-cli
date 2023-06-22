package jsonmap

import "fmt"

// GetString retrieves, from the slice, a string at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a string.
func (s *JsonSlice) GetString(index int) (string, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return "", err
	}
	return valueToString(value)
}

// GetInt retrieves, from the slice, an int at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not an int.
func (s *JsonSlice) GetInt(index int) (int64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return 0, err
	}
	return valueToInt(value)
}

// GetFloat retrieves, from the slice, a float at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a float (if the value at the index is an int,
// it will be quietly converted to a float).
func (s *JsonSlice) GetFloat(index int) (float64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return 0, err
	}
	return valueToFloat(value)
}

// GetBool retrieves, from the slice, a bool at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a bool.
func (s *JsonSlice) GetBool(index int) (bool, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return false, err
	}
	return valueToBool(value)
}

// GetMap retrieves, from the slice, a JsonMap at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a JsonMap.
func (s *JsonSlice) GetMap(index int) (JsonMap, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToMap(value)
}

// GetSlice retrieves, from the slice, a JsonSlice at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a JsonSlice.
func (s *JsonSlice) GetSlice(index int) (JsonSlice, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToSlice(value)
}

// GetSliceOfStrings retrieves, from the slice, a []string at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a string.
func (s *JsonSlice) GetSliceOfStrings(index int) ([]string, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToStringSlice(value)
}

// GetSliceOfInts retrieves, from the slice, a []int64 at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not an int64.
func (s *JsonSlice) GetSliceOfInts(index int) ([]int64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToIntSlice(value)
}

// GetSliceOfFloats retrieves, from the slice, a []float64 at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a float64 (or an int64, which will be quietly converted to a float64).
func (s *JsonSlice) GetSliceOfFloats(index int) ([]float64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToFloatSlice(value)
}

// GetSliceOfBools retrieves, from the slice, a []bool at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a bool.
func (s *JsonSlice) GetSliceOfBools(index int) ([]bool, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToBoolSlice(value)
}

// GetSliceOfMaps retrieves, from the slice, a []JsonMap at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a JsonMap.
func (s *JsonSlice) GetSliceOfMaps(index int) ([]JsonMap, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToMapSlice(value)
}

// GetSliceOfSlices retrieves, from the slice, a []JsonSlice at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a JsonSlice.
func (s *JsonSlice) GetSliceOfSlices(index int) ([]JsonSlice, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return nil, err
	}
	return valueToSliceSlice(value)
}

// GetValue retrieves, from the slice, an object of any type at the specified
// index. It will return an error if the index is out of bounds for the slice.
func (s *JsonSlice) GetValue(index int) (interface{}, error) {
	if index < 0 || index >= len(*s) {
		return nil, fmt.Errorf("cannot set value: slice index %d out of bounds for slice of length %d", index, len(*s))
	}
	value := (*s)[index]
	return value, nil
}

// GetStringWithDefault retrieves, from the slice, a string at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a string.
func (s *JsonSlice) GetStringWithDefault(index int, defaultValue string) (string, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToString(value)
}

// GetIntWithDefault retrieves, from the slice, an int at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not an int.
func (s *JsonSlice) GetIntWithDefault(index int, defaultValue int64) (int64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToInt(value)
}

// GetFloatWithDefault retrieves, from the slice, a float at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a float (if the value at the index is an int,
// it will be quietly converted to a float).
func (s *JsonSlice) GetFloatWithDefault(index int, defaultValue float64) (float64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToFloat(value)
}

// GetBoolWithDefault retrieves, from the slice, a bool at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a bool.
func (s *JsonSlice) GetBoolWithDefault(index int, defaultValue bool) (bool, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToBool(value)
}

// GetMapWithDefault retrieves, from the slice, a JsonMap at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a JsonMap.
func (s *JsonSlice) GetMapWithDefault(index int, defaultValue JsonMap) (JsonMap, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToMap(value)
}

// GetSliceWithDefault retrieves, from the slice, a JsonSlice at the specified index.
// It will return an error if the index is out of bounds for the slice or if
// the value at the index is not a JsonSlice.
func (s *JsonSlice) GetSliceWithDefault(index int, defaultValue JsonSlice) (JsonSlice, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToSlice(value)
}

// GetSliceOfStringsWithDefault retrieves, from the slice, a []string at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a string.
func (s *JsonSlice) GetSliceOfStringsWithDefault(index int, defaultValue []string) ([]string, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToStringSlice(value)
}

// GetSliceOfIntsWithDefault retrieves, from the slice, a []int64 at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not an int64.
func (s *JsonSlice) GetSliceOfIntsWithDefault(index int, defaultValue []int64) ([]int64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToIntSlice(value)
}

// GetSliceOfFloatsWithDefault retrieves, from the slice, a []float64 at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a float64 (or an int64, which will be quietly converted to a float64).
func (s *JsonSlice) GetSliceOfFloatsWithDefault(index int, defaultValue []float64) ([]float64, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToFloatSlice(value)
}

// GetSliceOfBoolsWithDefault retrieves, from the slice, a []bool at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a bool.
func (s *JsonSlice) GetSliceOfBoolsWithDefault(index int, defaultValue []bool) ([]bool, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToBoolSlice(value)
}

// GetSliceOfMapsWithDefault retrieves, from the slice, a []JsonMap at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a JsonMap.
func (s *JsonSlice) GetSliceOfMapsWithDefault(index int, defaultValue []JsonMap) ([]JsonMap, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToMapSlice(value)
}

// GetSliceOfSlicesWithDefault retrieves, from the slice, a []JsonSlice at the specified
// index. It will return an error if the index is out of bounds for the slice
// or if the value at the index is not a JsonSlice or if any value in the slice
// is not a JsonSlice.
func (s *JsonSlice) GetSliceOfSlicesWithDefault(index int, defaultValue []JsonSlice) ([]JsonSlice, error) {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue, nil
	}
	return valueToSliceSlice(value)
}

// GetValueWithDefault retrieves, from the slice, an object of any type at the
// specified index. If the index is out of bounds, then it returns the default
// value.
func (s *JsonSlice) GetValueWithDefault(index int, defaultValue interface{}) interface{} {
	value, err := s.GetValue(index)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetStringAtPath retrieves, from the slice, a string at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a string.
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

// GetIntAtPath retrieves, from the slice, an int at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not an int.
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

// GetFloatAtPath retrieves, from the slice, a float at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a float (if the value at the index is an int, it will be quietly
// converted to a float).
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

// GetBoolAtPath retrieves, from the slice, a bool at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a bool.
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

// GetMapAtPath retrieves, from the slice, a JsonMap at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a JsonMap.
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

// GetSliceAtPath retrieves, from the slice, a JsonSlice at the specified path.
// It will return an error if the path is invalid or if the value at the path
// is not a JsonSlice.
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

// GetSliceOfStringsAtPath retrieves, from the slice, a []string at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a
// string.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfStringsAtPath(path string) ([]string, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToStringSlice(value)
}

// GetSliceOfIntsAtPath retrieves, from the slice, a []int64 at the specified
// path. It will return an error if the path is invalid, if the value at the
// path is not a JsonSlice, or if any value in the slice is not an int.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfIntsAtPath(path string) ([]int64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToIntSlice(value)
}

// GetSliceOfFloatsAtPath retrieves, from the slice, a []float64 at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a float.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfFloatsAtPath(path string) ([]float64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToFloatSlice(value)
}

// GetSliceOfBoolsAtPath retrieves, from the slice, a []bool at the specified
// path. It will return an error if the path is invalid, if the value at the
// path is not a JsonSlice, or if any value in the slice is not a bool.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfBoolsAtPath(path string) ([]bool, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToBoolSlice(value)
}

// GetSliceOfMapsAtPath retrieves, from the slice, a []JsonMap at the specified
// path. It will return an error if the path is invalid, if the value at the
// path is not a JsonSlice, or if any value in the slice is not a map.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfMapsAtPath(path string) ([]JsonMap, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToMapSlice(value)
}

// GetSliceOfSlicesAtPath retrieves, from the slice, a []JsonSlice at the
// specified path. It will return an error if the path is invalid, if the value
// at the path is not a JsonSlice, or if any value in the slice is not a slice.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfSlicesAtPath(path string) ([]JsonSlice, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return valueToSliceSlice(value)
}

// GetValueAtPath retrieves, from the slice, an object of any type at the
// specified path. It will return an error if the path is invalid.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) GetValueAtPath(path string) (interface{}, error) {
	return pathGet(*s, path)
}

// GetStringAtPathWithDefault retrieves, from the slice, a string at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value.
// It will return an error if the value at the path is not a string.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForString" (slice of maps with string value)
// or "[0][0]" (slice of slices of string values)
func (s *JsonSlice) GetStringAtPathWithDefault(path string, defaultValue string) (string, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToString(value)
}

// GetIntAtPathWithDefault retrieves, from the slice, an int at the specified
// path. If the value cannot be found for any reason (including an invalid
// path), then it returns the default value. It will return an error if the
// value at the path is not an int.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForInt" (slice of maps with int value)
// or "[0][0]" (slice of slices of int values)
func (s *JsonSlice) GetIntAtPathWithDefault(path string, defaultValue int64) (int64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToInt(value)
}

// GetFloatAtPathWithDefault retrieves, from the slice, a float at the
// specified path. It will return an error if the value at the path is not a
// float (if the value at the index is an int, it will be quietly converted to
// a float).
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForFloat" (slice of maps with float value)
// or "[0][0]" (slice of slices of float values)
func (s *JsonSlice) GetFloatAtPathWithDefault(path string, defaultValue float64) (float64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToFloat(value)
}

// GetBoolAtPathWithDefault retrieves, from the slice, a bool at the specified
// path. If the value cannot be found for any reason (including an invalid
// path), then it returns the default value. It will return an error if the
// value at the path is not a bool.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForBool" (slice of maps with bool value)
// or "[0][0]" (slice of slices of bool values)
func (s *JsonSlice) GetBoolAtPathWithDefault(path string, defaultValue bool) (bool, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToBool(value)
}

// GetMapAtPathWithDefault retrieves, from the slice, a JsonMap at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a JsonMap.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForMap" (slice of maps with map value)
// or "[0][0]" (slice of slices of map values)
func (s *JsonSlice) GetMapAtPathWithDefault(path string, defaultValue JsonMap) (JsonMap, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToMap(value)
}

// GetSliceAtPathWithDefault retrieves, from the slice, a JsonSlice at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a JsonSlice.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceAtPathWithDefault(path string, defaultValue JsonSlice) (JsonSlice, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToSlice(value)
}

// GetSliceOfStringsAtPathWithDefault retrieves, from the slice, a []string at
// the specified path. If the value cannot be found for any reason (including
// an invalid path), then it returns the default value. It will return an error
// if the value at the path is not a JsonSlice or if any value in the slice is
// not a string.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfStringsAtPathWithDefault(path string, defaultValue []string) ([]string, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToStringSlice(value)
}

// GetSliceOfIntsAtPathWithDefault retrieves, from the slice, a []int64 at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a JsonSlice or if any value in the slice is not
// an int.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfIntsAtPathWithDefault(path string, defaultValue []int64) ([]int64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToIntSlice(value)
}

// GetSliceOfFloatsAtPathWithDefault retrieves, from the slice, a []float64 at
// the specified path. If the value cannot be found for any reason (including
// an invalid path), then it returns the default value. It will return an error
// if the value at the path is not a JsonSlice or if any value in the slice is
// not a float.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfFloatsAtPathWithDefault(path string, defaultValue []float64) ([]float64, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToFloatSlice(value)
}

// GetSliceOfBoolsAtPathWithDefault retrieves, from the slice, a []bool at the
// specified path. If the value cannot be found for any reason (including an
// invalid path), then it returns the default value. It will return an error if
// the value at the path is not a JsonSlice or if any value in the slice is not
// a bool.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfBoolsAtPathWithDefault(path string, defaultValue []bool) ([]bool, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToBoolSlice(value)
}

// GetSliceOfMapsAtPathWithDefault retrieves, from the slice, a []JsonMap at
// the specified path. If the value cannot be found for any reason (including
// an invalid path), then it returns the default value. It will return an error
// if the value at the path is not a JsonSlice or if any value in the slice is
// not a map.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfMapsAtPathWithDefault(path string, defaultValue []JsonMap) ([]JsonMap, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToMapSlice(value)
}

// GetSliceOfSlicesAtPathWithDefault retrieves, from the slice, a []JsonSlice
// at the specified path. If the value cannot be found for any reason
// (including an invalid path), then it returns the default value. It will
// return an error if the value at the path is not a JsonSlice or if any value
// in the slice is not a slice.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForSlice" (slice of maps with slice value)
// or "[0][0]" (slice of slices of slice values)
func (s *JsonSlice) GetSliceOfSlicesAtPathWithDefault(path string, defaultValue []JsonSlice) ([]JsonSlice, error) {
	value, err := s.GetValueAtPath(path)
	if err != nil {
		return defaultValue, nil
	}
	return valueToSliceSlice(value)
}

// GetValueAtPathWithDefault retrieves, from the slice, an object of any type
// at the specified path. If the value cannot be found for any reason
// (including an invalid path), then it returns the default value.
// Note that slice paths should always begin with an array index.
// e.g. "[0].keyForValue" (slice of maps with value)
// or "[0][0]" (slice of slices of values)
func (s *JsonSlice) GetValueAtPathWithDefault(path string, defaultValue interface{}) interface{} {
	value, err := pathGet(*s, path)
	if err == nil {
		return value
	}
	return defaultValue
}
