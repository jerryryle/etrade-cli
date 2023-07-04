package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeOptionChainPair interface {
	AsJsonMap() jsonmap.JsonMap
}

type eTradeOptionChainPair struct {
	jsonMap jsonmap.JsonMap
}

func CreateETradeOptionChainPair(responseMap jsonmap.JsonMap) (ETradeOptionChainPair, error) {
	return &eTradeOptionChainPair{
		jsonMap: responseMap,
	}, nil
}

func (e *eTradeOptionChainPair) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
