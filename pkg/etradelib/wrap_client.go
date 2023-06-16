package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ClientCallFn func() ([]byte, error)

func ExecuteClientCallAndWrapResponse(fn ClientCallFn) (jsonmap.JsonMap, error) {
	responseBytes, err := fn()
	if err != nil {
		return nil, err
	}
	return NewNormalizedJsonMapFromJsonBytes(responseBytes)
}
