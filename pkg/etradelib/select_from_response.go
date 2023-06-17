package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ClientCallFn func() ([]byte, error)

func SelectMapFromResponse(path string, fn ClientCallFn) (jsonmap.JsonMap, error) {
	responseMap, err := getMapFromResponse(fn)
	if err != nil {
		return nil, err
	}
	responseValue, err := responseMap.GetMapAtPath(path)
	if err != nil {
		return nil, err
	}
	return responseValue, nil
}

func SelectSliceFromResponse(path string, fn ClientCallFn) (jsonmap.JsonSlice, error) {
	responseMap, err := getMapFromResponse(fn)
	if err != nil {
		return nil, err
	}
	responseValue, err := responseMap.GetSliceAtPath(path)
	if err != nil {
		return nil, err
	}
	return responseValue, nil
}

func SelectMapSliceFromResponse(path string, fn ClientCallFn) ([]jsonmap.JsonMap, error) {
	responseMap, err := getMapFromResponse(fn)
	if err != nil {
		return nil, err
	}
	responseValue, err := responseMap.GetMapSliceAtPath(path)
	if err != nil {
		return nil, err
	}
	return responseValue, nil
}

func SelectValueFromResponse(path string, fn ClientCallFn) (interface{}, error) {
	responseMap, err := getMapFromResponse(fn)
	if err != nil {
		return nil, err
	}
	responseValue, err := responseMap.GetValueAtPath(path)
	if err != nil {
		return nil, err
	}
	return responseValue, nil
}

func getMapFromResponse(fn ClientCallFn) (jsonmap.JsonMap, error) {
	// Call the client method
	responseBytes, err := fn()
	if err != nil {
		return nil, err
	}
	// Unmarshal into a key-normalized Json Map
	responseMap, err := NewNormalizedJsonMap(responseBytes)
	if err != nil {
		return nil, err
	}
	return responseMap, nil
}
