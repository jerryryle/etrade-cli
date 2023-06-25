package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeQuoteList interface {
	GetAllQuotes() []ETradeQuote
	AsJsonMap() jsonmap.JsonMap
}

type eTradeQuoteList struct {
	quotes   []ETradeQuote
	messages jsonmap.JsonSlice
}

const (
	// The AsJsonMap() map looks like this:
	// "quotes": [
	//   {
	//     <quote info>
	//   }
	// ]
	// "messages": [
	//   {
	//     <message info>
	//   }
	// ]

	// QuoteListQuotesSliceJsonMapPath is the path to a slice of quotes.
	QuoteListQuotesSliceJsonMapPath = ".quotes"

	// QuoteListMessagesSliceJsonMapPath is the path to a slice of messages.
	QuoteListMessagesSliceJsonMapPath = ".messages"
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
	//     "Messages": {
	//       "message": [
	//         {
	//           <message info>
	//         }
	//       ]
	//     }
	//   }
	// }

	// quoteListQuotesSliceResponsePath is the path to a slice of quotes.
	quoteListQuotesSliceResponsePath = ".quoteResponse.quoteData"

	// quoteListQuotesSliceResponsePath is the path to a slice of messages.
	quoteListMessagesSliceResponsePath = ".quoteResponse.messages.message"
)

func CreateETradeQuoteListFromResponse(response []byte) (ETradeQuoteList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeQuoteList(responseMap)
}

func CreateETradeQuoteList(lookupListResponseMap jsonmap.JsonMap) (ETradeQuoteList, error) {
	quotesSlice, err := lookupListResponseMap.GetSliceOfMapsAtPathWithDefault(quoteListQuotesSliceResponsePath, nil)
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

	messagesSlice, err := lookupListResponseMap.GetSliceAtPathWithDefault(quoteListMessagesSliceResponsePath, nil)
	if err != nil {
		return nil, err
	}

	return &eTradeQuoteList{quotes: allQuotes, messages: messagesSlice}, nil
}

func (e *eTradeQuoteList) GetAllQuotes() []ETradeQuote {
	return e.quotes
}

func (e *eTradeQuoteList) AsJsonMap() jsonmap.JsonMap {
	var quoteListMap = jsonmap.JsonMap{}

	if len(e.quotes) > 0 {
		quotesSlice := make(jsonmap.JsonSlice, 0, len(e.quotes))
		for _, quote := range e.quotes {
			quotesSlice = append(quotesSlice, quote.AsJsonMap())
		}
		err := quoteListMap.SetSliceAtPath(QuoteListQuotesSliceJsonMapPath, quotesSlice)
		if err != nil {
			panic(err)
		}
	}
	if e.messages != nil {
		err := quoteListMap.SetSliceAtPath(QuoteListMessagesSliceJsonMapPath, e.messages)
		if err != nil {
			panic(err)
		}
	}
	return quoteListMap
}
