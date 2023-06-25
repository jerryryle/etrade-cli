package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeQuote interface {
	AsJsonMap() jsonmap.JsonMap
}

type eTradeQuote struct {
	jsonMap jsonmap.JsonMap
}

func CreateETradeQuote(lookupJsonMap jsonmap.JsonMap) (ETradeQuote, error) {
	return &eTradeQuote{
		jsonMap: lookupJsonMap,
	}, nil
}

func (e *eTradeQuote) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
