package jsonmap

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type JsonMap map[string]interface{}

func NewFromIoReader(jsonReader io.Reader) (JsonMap, error) {
	var m map[string]interface{}
	decoder := json.NewDecoder(jsonReader)
	// Decode numbers using the json.Number type instead of float64
	decoder.UseNumber()
	err := decoder.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func NewFromJsonBytes(jsonBytes []byte) (JsonMap, error) {
	return NewFromIoReader(bytes.NewReader(jsonBytes))
}

func NewFromJsonString(jsonString string) (JsonMap, error) {
	return NewFromIoReader(strings.NewReader(jsonString))
}
