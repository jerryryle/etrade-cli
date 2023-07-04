package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeOptionExpireDate interface {
	AsJsonMap() jsonmap.JsonMap
}

type eTradeOptionExpireDate struct {
	jsonMap jsonmap.JsonMap
}

func CreateETradeOptionExpireDate(responseMap jsonmap.JsonMap) (ETradeOptionExpireDate, error) {
	return &eTradeOptionExpireDate{
		jsonMap: responseMap,
	}, nil
}

func (e *eTradeOptionExpireDate) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
