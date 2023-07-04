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
	// {
	//   "optionChainPairs": [
	//     {
	//       <optionChainPair info>
	//     }
	//   ]
	//   "timeStamp": 1234,
	//   "quoteType": "Type",
	//   "nearPrice": 123.4,
	//   "selected": {
	//       <selected info>
	//   }
	// }

	// OptionChainPairListOptionChainPairsPath is the path to a slice of optionChainPairs.
	OptionChainPairListOptionChainPairsPath = ".optionChainPairs"

	// OptionChainPairListTimeStampPath is the path to timestamp.
	OptionChainPairListTimeStampPath = ".timeStamp"

	// OptionChainPairListQuoteTypePath is the path to quote type.
	OptionChainPairListQuoteTypePath = ".quoteType"

	// OptionChainPairListNearPricePath is the path to near price.
	OptionChainPairListNearPricePath = ".nearPrice"

	// OptionChainPairListSelectedPath is the path to selected map.
	OptionChainPairListSelectedPath = ".selected"
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

	// optionChainPairListOptionChainPairsResponsePath is the path to a slice of OptionChainPairs.
	optionChainPairListOptionChainPairsResponsePath = ".optionChainResponse.optionPair"

	// optionChainPairListTimeStampResponsePath is the path to timestamp.
	optionChainPairListTimeStampResponsePath = ".optionChainResponse.timeStamp"

	// optionChainPairListQuoteTypeResponsePath is the path to quote type.
	optionChainPairListQuoteTypeResponsePath = ".optionChainResponse.quoteType"

	// optionChainPairListNearPriceResponsePath is the path to near price.
	optionChainPairListNearPriceResponsePath = ".optionChainResponse.nearPrice"

	// optionChainPairListSelectedResponsePath is the path to selected map
	optionChainPairListSelectedResponsePath = ".optionChainResponse.selectedED"
)

func CreateETradeOptionChainPairListFromResponse(response []byte) (ETradeOptionChainPairList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeOptionChainPairList(responseMap)
}

func CreateETradeOptionChainPairList(responseMap jsonmap.JsonMap) (ETradeOptionChainPairList, error) {
	optionChainPairsSlice, err := responseMap.GetSliceOfMapsAtPath(optionChainPairListOptionChainPairsResponsePath)
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

	timeStamp, err := responseMap.GetIntAtPathWithDefault(
		optionChainPairListTimeStampResponsePath, 0,
	)
	if err != nil {
		return nil, err
	}

	quoteType, err := responseMap.GetStringAtPathWithDefault(
		optionChainPairListQuoteTypeResponsePath, "",
	)
	if err != nil {
		return nil, err
	}

	nearPrice, err := responseMap.GetFloatAtPathWithDefault(
		optionChainPairListNearPriceResponsePath, 0,
	)
	if err != nil {
		return nil, err
	}

	selected, err := responseMap.GetMapAtPathWithDefault(
		optionChainPairListSelectedResponsePath, nil,
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

	optionChainPairsSlice := make(jsonmap.JsonSlice, 0, len(e.optionChainPairs))
	for _, optionChainPair := range e.optionChainPairs {
		optionChainPairsSlice = append(optionChainPairsSlice, optionChainPair.AsJsonMap())
	}
	err := optionChainPairListMap.SetSliceAtPath(
		OptionChainPairListOptionChainPairsPath, optionChainPairsSlice,
	)
	if err != nil {
		panic(err)
	}

	err = optionChainPairListMap.SetIntAtPath(OptionChainPairListTimeStampPath, e.timeStamp)
	if err != nil {
		panic(err)
	}

	err = optionChainPairListMap.SetStringAtPath(OptionChainPairListQuoteTypePath, e.quoteType)
	if err != nil {
		panic(err)
	}

	err = optionChainPairListMap.SetFloatAtPath(OptionChainPairListNearPricePath, e.nearPrice)
	if err != nil {
		panic(err)
	}

	err = optionChainPairListMap.SetMapAtPath(OptionChainPairListSelectedPath, e.selected)
	if err != nil {
		panic(err)
	}

	return optionChainPairListMap
}
