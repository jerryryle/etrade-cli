package jsonmap

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type JsonSlice []interface{}

// NewSliceFromIoReader creates a JsonSlice from an io.Reader. It returns an error
// if valid JSON cannot be decoded from the io.Reader.
// Note: This function expects the top-level JSON object to be a slice and will
// fail if it is a map (use NewMapFromIoReader for a top-level map).
func NewSliceFromIoReader(jsonReader io.Reader) (JsonSlice, error) {
	var s []interface{}
	decoder := json.NewDecoder(jsonReader)
	// Decode numbers using the json.Number type instead of float64
	decoder.UseNumber()
	err := decoder.Decode(&s)
	if err != nil {
		return nil, err
	}
	// This nil map will recursively ensure that all []interface{} values are
	// replaced with JsonSlice and all map[string]interface{} values are
	// replaced with JsonMap
	jsonSlice := JsonSlice(s).Map(nil, nil)
	return jsonSlice, nil
}

// NewSliceFromJsonBytes creates a JsonSlice from an io.Reader. It returns an error
// if valid JSON cannot be decoded from the io.Reader.
// Note: This function expects the top-level JSON object to be a slice and will
// fail if it is a map (use NewMapFromJsonBytes for a top-level map).
func NewSliceFromJsonBytes(jsonBytes []byte) (JsonSlice, error) {
	return NewSliceFromIoReader(bytes.NewReader(jsonBytes))
}

// NewSliceFromJsonString creates a JsonSlice from an io.Reader. It returns an error
// if valid JSON cannot be decoded from the io.Reader.
// Note: This function expects the top-level JSON object to be a slice and will
// fail if it is a map (use NewMapFromJsonString for a top-level map).
func NewSliceFromJsonString(jsonString string) (JsonSlice, error) {
	return NewSliceFromIoReader(strings.NewReader(jsonString))
}