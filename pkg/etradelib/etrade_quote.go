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

func CreateETradeQuote(responseMap jsonmap.JsonMap) (ETradeQuote, error) {
	return &eTradeQuote{
		jsonMap: responseMap,
	}, nil
}

func (e *eTradeQuote) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
