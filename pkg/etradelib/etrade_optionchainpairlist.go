package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeOptionChainPairList interface {
	GetAllOptionChainPairs() []ETradeOptionChainPair
	AsJsonMap() jsonmap.JsonMap
}

type eTradeOptionChainPairList struct {
	optionChainPairs []ETradeOptionChainPair
	timeStamp        int64
	quoteType        string
	nearPrice        float64
	selected         jsonmap.JsonMap
}

const (
	// The AsJsonMap() map looks like this:
	// "optionChainPairs": [
	//   {
	//     <optionChainPair info>
	//   }
	// ]
	// "timeStamp": 1234
	// "quoteType": "Type"
	// "nearPrice": 123.4
	// "selected": {
	//     <selected info>
	// }

	// OptionChainPairListOptionChainPairsSliceJsonMapPath is the path to a slice of optionChainPairs.
	OptionChainPairListOptionChainPairsSliceJsonMapPath = ".optionChainPairs"

	// OptionChainPairTimeStampJsonMapPath is the path to timestamp.
	OptionChainPairTimeStampJsonMapPath = ".timeStamp"

	// OptionChainPairQuoteTypeJsonMapPath is the path to quote type.
	OptionChainPairQuoteTypeJsonMapPath = ".quoteType"

	// OptionChainPairNearPriceJsonMapPath is the path to near price.
	OptionChainPairNearPriceJsonMapPath = ".nearPrice"

	// OptionChainPairSelectedJsonMapPath is the path to selected map.
	OptionChainPairSelectedJsonMapPath = ".selected"
)

const (
	// The lookup list response JSON looks like this:
	// {
	//   "OptionChainResponse": {
	//     "OptionPair": [
	//       {
	//         <optionChainPair info>
	//       }
	//     ]
	//     "timeStamp": 1234
	//     "quoteType": "Type"
	//     "nearPrice": 123.4
	//     "selected": {
	//         <selected info>
	//     }
	//   }
	// }

	// optionChainPairListOptionChainPairsSliceResponsePath is the path to a slice of OptionChainPairs.
	optionChainPairListOptionChainPairsSliceResponsePath = ".optionChainResponse.optionPair"

	// optionChainPairListTimeStampResponsePath is the path to timestamp.
	optionChainPairListTimeStampResponsePath = ".optionChainResponse.timeStamp"

	// optionChainPairListQuoteTypeResponsePath is the path to quote type.
	optionChainPairListQuoteTypeResponsePath = ".optionChainResponse.quoteType"

	// optionChainPairListNearPriceResponsePath is the path to near price.
	optionChainPairListNearPriceResponsePath = ".optionChainResponse.nearPrice"

	// optionChainPairListSelectedMapResponsePath is the path to selected map
	optionChainPairListSelectedMapResponsePath = ".optionChainResponse.selectedED"
)

func CreateETradeOptionChainPairListFromResponse(response []byte) (ETradeOptionChainPairList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeOptionChainPairList(responseMap)
}

func CreateETradeOptionChainPairList(lookupListResponseMap jsonmap.JsonMap) (ETradeOptionChainPairList, error) {
	optionChainPairsSlice, err := lookupListResponseMap.GetSliceOfMapsAtPathWithDefault(
		optionChainPairListOptionChainPairsSliceResponsePath, nil,
	)
	if err != nil {
		return nil, err
	}
	allOptionChainPairs := make([]ETradeOptionChainPair, 0, len(optionChainPairsSlice))
	for _, optionChainPairJsonMap := range optionChainPairsSlice {
		optionChainPair, err := CreateETradeOptionChainPair(optionChainPairJsonMap)
		if err != nil {
			return nil, err
		}
		allOptionChainPairs = append(allOptionChainPairs, optionChainPair)
	}

	timeStamp, err := lookupListResponseMap.GetIntAtPathWithDefault(
		optionChainPairListTimeStampResponsePath, 0,
	)
	if err != nil {
		return nil, err
	}

	quoteType, err := lookupListResponseMap.GetStringAtPathWithDefault(
		optionChainPairListQuoteTypeResponsePath, "",
	)
	if err != nil {
		return nil, err
	}

	nearPrice, err := lookupListResponseMap.GetFloatAtPathWithDefault(
		optionChainPairListNearPriceResponsePath, 0,
	)
	if err != nil {
		return nil, err
	}

	selected, err := lookupListResponseMap.GetMapAtPathWithDefault(
		optionChainPairListSelectedMapResponsePath, nil,
	)
	if err != nil {
		return nil, err
	}

	return &eTradeOptionChainPairList{
		optionChainPairs: allOptionChainPairs,
		timeStamp:        timeStamp,
		quoteType:        quoteType,
		nearPrice:        nearPrice,
		selected:         selected,
	}, nil
}

func (e *eTradeOptionChainPairList) GetAllOptionChainPairs() []ETradeOptionChainPair {
	return e.optionChainPairs
}

func (e *eTradeOptionChainPairList) AsJsonMap() jsonmap.JsonMap {
	var optionChainPairListMap = jsonmap.JsonMap{}

	if len(e.optionChainPairs) > 0 {
		optionChainPairsSlice := make(jsonmap.JsonSlice, 0, len(e.optionChainPairs))
		for _, optionChainPair := range e.optionChainPairs {
			optionChainPairsSlice = append(optionChainPairsSlice, optionChainPair.AsJsonMap())
		}
		err := optionChainPairListMap.SetSliceAtPath(
			OptionChainPairListOptionChainPairsSliceJsonMapPath, optionChainPairsSlice,
		)
		if err != nil {
			panic(err)
		}
	}

	err := optionChainPairListMap.SetIntAtPath(OptionChainPairTimeStampJsonMapPath, e.timeStamp)
	if err != nil {
		panic(err)
	}

	err = optionChainPairListMap.SetStringAtPath(OptionChainPairQuoteTypeJsonMapPath, e.quoteType)
	if err != nil {
		panic(err)
	}

	err = optionChainPairListMap.SetFloatAtPath(OptionChainPairNearPriceJsonMapPath, e.nearPrice)
	if err != nil {
		panic(err)
	}

	err = optionChainPairListMap.SetMapAtPath(OptionChainPairSelectedJsonMapPath, e.selected)
	if err != nil {
		panic(err)
	}

	return optionChainPairListMap
}
