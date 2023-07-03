package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeLookupResult interface {
	AsJsonMap() jsonmap.JsonMap
}

type eTradeLookupResult struct {
	jsonMap jsonmap.JsonMap
}

func CreateETradeLookupResult(responseMap jsonmap.JsonMap) (ETradeLookupResult, error) {
	return &eTradeLookupResult{
		jsonMap: responseMap,
	}, nil
}

func (e *eTradeLookupResult) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
