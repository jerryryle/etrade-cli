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
	selectedED       jsonmap.JsonMap
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
	//   "selectedED": {
	//       <selectedED info>
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

	// OptionChainPairListSelectedEDPath is the path to selectedED map.
	OptionChainPairListSelectedEDPath = ".selectedED"
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
	//     "SelectedED": {
	//         <selectedED info>
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

	// optionChainPairListSelectedEDResponsePath is the path to selectedED map
	optionChainPairListSelectedEDResponsePath = ".optionChainResponse.selectedED"
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

	selectedED, err := responseMap.GetMapAtPathWithDefault(
		optionChainPairListSelectedEDResponsePath, nil,
	)
	if err != nil {
		return nil, err
	}

	return &eTradeOptionChainPairList{
		optionChainPairs: allOptionChainPairs,
		timeStamp:        timeStamp,
		quoteType:        quoteType,
		nearPrice:        nearPrice,
		selectedED:       selectedED,
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

	err = optionChainPairListMap.SetMapAtPath(OptionChainPairListSelectedEDPath, e.selectedED)
	if err != nil {
		panic(err)
	}

	return optionChainPairListMap
}
