package jsonmap

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type JsonMap map[string]interface{}

// NewMapFromIoReader creates a JsonMap from an io.Reader. It returns an error
// if valid JSON cannot be decoded from the io.Reader.
// Note: This function expects the top-level JSON object to be a map and will
// fail if it is a slice (use NewSliceFromIoReader for a top-level slice).
func NewMapFromIoReader(jsonReader io.Reader) (JsonMap, error) {
	var m map[string]interface{}
	decoder := json.NewDecoder(jsonReader)
	// Decode numbers using the json.Number type instead of float64
	decoder.UseNumber()
	err := decoder.Decode(&m)
	if err != nil {
		return nil, err
	}
	// This nil map will recursively ensure that all map[string]interface{}
	// values are replaced with JsonMap and all []interface{} values are
	// replaced with JsonSlice
	jsonMap := JsonMap(m)
	jsonMap = jsonMap.Map(nil, nil)
	return jsonMap, nil
}

// NewMapFromJsonBytes creates a JsonMap from a byte slice. It returns an error
// if valid JSON cannot be decoded from the bytes.
// Note: This function expects the top-level JSON object to be a map and will
// fail if it is a slice (use NewSliceFromJsonBytes for a top-level slice).
func NewMapFromJsonBytes(jsonBytes []byte) (JsonMap, error) {
	return NewMapFromIoReader(bytes.NewReader(jsonBytes))
}

// NewMapFromJsonString creates a JsonMap from a string. It returns an error
// if valid JSON cannot be decoded from the string.
// Note: This function expects the top-level JSON object to be a map and will
// fail if it is a slice (use NewSliceFromJsonString for a top-level slice).
func NewMapFromJsonString(jsonString string) (JsonMap, error) {
	return NewMapFromIoReader(strings.NewReader(jsonString))
}

func (m *JsonMap) ToIoWriter(jsonWriter io.Writer, pretty bool) error {
	encoder := json.NewEncoder(jsonWriter)
	if pretty {
		encoder.SetIndent("", "  ")
	}
	err := encoder.Encode(*m)
	if err != nil {
		return err
	}
	return nil
}

func (m *JsonMap) ToJsonBytes(pretty bool) ([]byte, error) {
	var byteBuffer bytes.Buffer
	err := m.ToIoWriter(&byteBuffer, pretty)
	if err != nil {
		return nil, err
	}
	return byteBuffer.Bytes(), nil
}

func (m *JsonMap) ToJsonString(pretty bool) (string, error) {
	var byteBuffer bytes.Buffer
	err := m.ToIoWriter(&byteBuffer, pretty)
	if err != nil {
		return "", err
	}
	return byteBuffer.String(), nil
}
