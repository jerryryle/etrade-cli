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

func CreateETradeLookupResult(lookupJsonMap jsonmap.JsonMap) (ETradeLookupResult, error) {
	return &eTradeLookupResult{
		jsonMap: lookupJsonMap,
	}, nil
}

func (e *eTradeLookupResult) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
