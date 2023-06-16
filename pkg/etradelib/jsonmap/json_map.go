package jsonmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type JsonMap map[string]interface{}

func FromInterface(i interface{}) (JsonMap, error) {
	switch i := i.(type) {
	case map[string]interface{}:
		return i, nil
	case JsonMap:
		return i, nil
	default:
		return nil, fmt.Errorf("cannot treat type %T as JsonMap", i)
	}
}

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
