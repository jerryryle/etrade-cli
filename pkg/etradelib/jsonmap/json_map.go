package jsonmap

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type JsonMap map[string]interface{}

// NewJsonMapFromIoReader creates a JsonMap from an io.Reader. It returns an
// error if valid JSON cannot be decoded from the io.Reader.
// Note: This function expects the top-level JSON object to be a map and will
// fail if it is a slice (use NewJsonSliceFromIoReader for a top-level slice).
func NewJsonMapFromIoReader(jsonReader io.Reader) (JsonMap, error) {
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

// NewJsonMapFromJsonBytes creates a JsonMap from a byte slice. It returns an
// error if valid JSON cannot be decoded from the bytes.
// Note: This function expects the top-level JSON object to be a map and will
// fail if it is a slice (use NewJsonSliceFromJsonBytes for a top-level slice).
func NewJsonMapFromJsonBytes(jsonBytes []byte) (JsonMap, error) {
	return NewJsonMapFromIoReader(bytes.NewReader(jsonBytes))
}

// NewJsonMapFromJsonString creates a JsonMap from a string. It returns an
// error if valid JSON cannot be decoded from the string.
// Note: This function expects the top-level JSON object to be a map and will
// fail if it is a slice (use NewJsonSliceFromJsonString for a top-level slice).
func NewJsonMapFromJsonString(jsonString string) (JsonMap, error) {
	return NewJsonMapFromIoReader(strings.NewReader(jsonString))
}

func (m *JsonMap) ToIoWriter(jsonWriter io.Writer, pretty bool, escapeHtml bool) error {
	encoder := json.NewEncoder(jsonWriter)
	encoder.SetEscapeHTML(escapeHtml)
	if pretty {
		encoder.SetIndent("", "  ")
	}
	err := encoder.Encode(*m)
	if err != nil {
		return err
	}
	return nil
}

func (m *JsonMap) ToJsonBytes(pretty bool, escapeHtml bool) ([]byte, error) {
	var byteBuffer bytes.Buffer
	err := m.ToIoWriter(&byteBuffer, pretty, escapeHtml)
	if err != nil {
		return nil, err
	}
	return byteBuffer.Bytes(), nil
}

func (m *JsonMap) ToJsonString(pretty bool, escapeHtml bool) (string, error) {
	var byteBuffer bytes.Buffer
	err := m.ToIoWriter(&byteBuffer, pretty, escapeHtml)
	if err != nil {
		return "", err
	}
	return byteBuffer.String(), nil
}
