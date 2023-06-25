package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeQuoteList interface {
	GetAllQuotes() []ETradeQuote
	AsJsonMap() jsonmap.JsonMap
}

type eTradeQuoteList struct {
	quotes []ETradeQuote
}

const (
	// The AsJsonMap() map looks like this:
	// "quotes": [
	//   {
	//     <lookup info>
	//   }
	// ]

	// QuoteListQuotesSliceJsonMapPath is the path to a slice of lookup quotes.
	QuoteListQuotesSliceJsonMapPath = ".quotes"
)

const (
	// The lookup list response JSON looks like this:
	// {
	//   "QuoteResponse": {
	//     "QuoteData": [
	//       {
	//         <quote info>
	//       }
	//     ]
	//   }
	// }

	// quoteListQuotesSliceResponsePath is the path to a slice of quotes.
	quoteListQuotesSliceResponsePath = ".quoteResponse.quoteData"
)

func CreateETradeQuoteListFromResponse(response []byte) (ETradeQuoteList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeQuoteList(responseMap)
}

func CreateETradeQuoteList(lookupListResponseMap jsonmap.JsonMap) (ETradeQuoteList, error) {
	quotesSlice, err := lookupListResponseMap.GetSliceOfMapsAtPath(quoteListQuotesSliceResponsePath)
	if err != nil {
		return nil, err
	}
	allQuotes := make([]ETradeQuote, 0, len(quotesSlice))
	for _, quoteJsonMap := range quotesSlice {
		quote, err := CreateETradeQuote(quoteJsonMap)
		if err != nil {
			return nil, err
		}
		allQuotes = append(allQuotes, quote)
	}
	return &eTradeQuoteList{quotes: allQuotes}, nil
}

func (e *eTradeQuoteList) GetAllQuotes() []ETradeQuote {
	return e.quotes
}

func (e *eTradeQuoteList) AsJsonMap() jsonmap.JsonMap {
	quotesSlice := make(jsonmap.JsonSlice, 0, len(e.quotes))
	for _, quote := range e.quotes {
		quotesSlice = append(quotesSlice, quote.AsJsonMap())
	}
	var quoteListMap = jsonmap.JsonMap{}
	err := quoteListMap.SetSliceAtPath(QuoteListQuotesSliceJsonMapPath, quotesSlice)
	if err != nil {
		panic(err)
	}
	return quoteListMap
}
