package jsonmap

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type JsonSlice []interface{}

// NewJsonSliceFromIoReader creates a JsonSlice from an io.Reader. It returns
// an error if valid JSON cannot be decoded from the io.Reader.
// Note: This function expects the top-level JSON object to be a slice and will
// fail if it is a map (use NewJsonMapFromIoReader for a top-level map).
func NewJsonSliceFromIoReader(jsonReader io.Reader) (JsonSlice, error) {
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
	jsonSlice := JsonSlice(s)
	jsonSlice = jsonSlice.Map(nil, nil)
	return jsonSlice, nil
}

// NewJsonSliceFromJsonBytes creates a JsonSlice from an io.Reader. It returns
// an error if valid JSON cannot be decoded from the io.Reader.
// Note: This function expects the top-level JSON object to be a slice and will
// fail if it is a map (use NewJsonMapFromJsonBytes for a top-level map).
func NewJsonSliceFromJsonBytes(jsonBytes []byte) (JsonSlice, error) {
	return NewJsonSliceFromIoReader(bytes.NewReader(jsonBytes))
}

// NewJsonSliceFromJsonString creates a JsonSlice from an io.Reader. It returns
// an error if valid JSON cannot be decoded from the io.Reader.
// Note: This function expects the top-level JSON object to be a slice and will
// fail if it is a map (use NewJsonMapFromJsonString for a top-level map).
func NewJsonSliceFromJsonString(jsonString string) (JsonSlice, error) {
	return NewJsonSliceFromIoReader(strings.NewReader(jsonString))
}

// NewJsonSliceFromSlice creates a JsonSlice from a slice of any type
func NewJsonSliceFromSlice[T any](s []T) JsonSlice {
	newSlice := make(JsonSlice, len(s))
	for i, v := range s {
		newSlice[i] = v
	}
	return newSlice
}

func (s *JsonSlice) ToIoWriter(jsonWriter io.Writer, pretty bool, escapeHtml bool) error {
	encoder := json.NewEncoder(jsonWriter)
	encoder.SetEscapeHTML(escapeHtml)
	if pretty {
		encoder.SetIndent("", "  ")
	}
	err := encoder.Encode(*s)
	if err != nil {
		return err
	}
	return nil
}

func (s *JsonSlice) ToJsonBytes(pretty bool, escapeHtml bool) ([]byte, error) {
	var byteBuffer bytes.Buffer
	err := s.ToIoWriter(&byteBuffer, pretty, escapeHtml)
	if err != nil {
		return nil, err
	}
	return byteBuffer.Bytes(), nil
}

func (s *JsonSlice) ToJsonString(pretty bool, escapeHtml bool) (string, error) {
	var byteBuffer bytes.Buffer
	err := s.ToIoWriter(&byteBuffer, pretty, escapeHtml)
	if err != nil {
		return "", err
	}
	return byteBuffer.String(), nil
}
