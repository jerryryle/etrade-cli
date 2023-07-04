package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeLookupResultList interface {
	GetAllResults() []ETradeLookupResult
	AsJsonMap() jsonmap.JsonMap
}

type eTradeLookupResultList struct {
	results []ETradeLookupResult
}

const (
	// The AsJsonMap() map looks like this:
	// "results": [
	//   {
	//     <lookup info>
	//   }
	// ]

	// LookupResultListResultsPath is the path to a slice of lookup results.
	LookupResultListResultsPath = ".results"
)

const (
	// The lookup list response JSON looks like this:
	// {
	//   "LookupResponse": {
	//     "Data": [
	//       {
	//         <lookup info>
	//       }
	//     ]
	//   }
	// }

	// lookupResultListResultsResponsePath is the path to a slice of results.
	lookupResultListResultsResponsePath = ".lookupResponse.data"
)

func CreateETradeLookupResultListFromResponse(response []byte) (ETradeLookupResultList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeLookupResultList(responseMap)
}

func CreateETradeLookupResultList(responseMap jsonmap.JsonMap) (ETradeLookupResultList, error) {
	resultsSlice, err := responseMap.GetSliceOfMapsAtPath(lookupResultListResultsResponsePath)
	if err != nil {
		return nil, err
	}
	allResults := make([]ETradeLookupResult, 0, len(resultsSlice))
	for _, resultJsonMap := range resultsSlice {
		result, err := CreateETradeLookupResult(resultJsonMap)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, result)
	}
	return &eTradeLookupResultList{results: allResults}, nil
}

func (e *eTradeLookupResultList) GetAllResults() []ETradeLookupResult {
	return e.results
}

func (e *eTradeLookupResultList) AsJsonMap() jsonmap.JsonMap {
	resultsSlice := make(jsonmap.JsonSlice, 0, len(e.results))
	for _, result := range e.results {
		resultsSlice = append(resultsSlice, result.AsJsonMap())
	}
	var lookupResultListMap = jsonmap.JsonMap{}
	err := lookupResultListMap.SetSliceAtPath(LookupResultListResultsPath, resultsSlice)
	if err != nil {
		panic(err)
	}
	return lookupResultListMap
}
