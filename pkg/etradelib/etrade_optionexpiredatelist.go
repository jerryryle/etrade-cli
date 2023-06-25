package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeOptionExpireDateList interface {
	GetAllOptionExpireDates() []ETradeOptionExpireDate
	AsJsonMap() jsonmap.JsonMap
}

type eTradeOptionExpireDateList struct {
	optionExpireDates []ETradeOptionExpireDate
}

const (
	// The AsJsonMap() map looks like this:
	// "optionExpireDates": [
	//   {
	//     <optionExpireDate info>
	//   }
	// ]

	// OptionExpireDateListOptionExpireDatesSliceJsonMapPath is the path to a slice of optionExpireDates.
	OptionExpireDateListOptionExpireDatesSliceJsonMapPath = ".optionExpireDates"
)

const (
	// The lookup list response JSON looks like this:
	// {
	//   "OptionExpireDateResponse": {
	//     "ExpirationDate": [
	//       {
	//         <optionExpireDate info>
	//       }
	//     ]
	//   }
	// }

	// optionExpireDateListOptionExpireDatesSliceResponsePath is the path to a slice of optionExpireDates.
	optionExpireDateListOptionExpireDatesSliceResponsePath = ".optionExpireDateResponse.expirationDate"
)

func CreateETradeOptionExpireDateListFromResponse(response []byte) (ETradeOptionExpireDateList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeOptionExpireDateList(responseMap)
}

func CreateETradeOptionExpireDateList(lookupListResponseMap jsonmap.JsonMap) (ETradeOptionExpireDateList, error) {
	optionExpireDatesSlice, err := lookupListResponseMap.GetSliceOfMapsAtPathWithDefault(
		optionExpireDateListOptionExpireDatesSliceResponsePath, nil,
	)
	if err != nil {
		return nil, err
	}
	allOptionExpireDates := make([]ETradeOptionExpireDate, 0, len(optionExpireDatesSlice))
	for _, optionExpireDateJsonMap := range optionExpireDatesSlice {
		optionExpireDate, err := CreateETradeOptionExpireDate(optionExpireDateJsonMap)
		if err != nil {
			return nil, err
		}
		allOptionExpireDates = append(allOptionExpireDates, optionExpireDate)
	}

	return &eTradeOptionExpireDateList{optionExpireDates: allOptionExpireDates}, nil
}

func (e *eTradeOptionExpireDateList) GetAllOptionExpireDates() []ETradeOptionExpireDate {
	return e.optionExpireDates
}

func (e *eTradeOptionExpireDateList) AsJsonMap() jsonmap.JsonMap {
	var optionExpireDateListMap = jsonmap.JsonMap{}

	optionExpireDatesSlice := make(jsonmap.JsonSlice, 0, len(e.optionExpireDates))
	for _, optionExpireDate := range e.optionExpireDates {
		optionExpireDatesSlice = append(optionExpireDatesSlice, optionExpireDate.AsJsonMap())
	}
	err := optionExpireDateListMap.SetSliceAtPath(
		OptionExpireDateListOptionExpireDatesSliceJsonMapPath, optionExpireDatesSlice,
	)
	if err != nil {
		panic(err)
	}
	return optionExpireDateListMap
}
