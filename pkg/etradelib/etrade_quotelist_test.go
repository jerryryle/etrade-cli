package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeQuoteList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeQuoteList
	}{
		{
			name: "CreateETradeQuoteList Creates List With Valid Response",
			testJson: `
{
  "QuoteResponse": {
    "QuoteData": [
      {
        "key1": "value1"
      },
      {
        "key2": "value2"
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeQuoteList{
				quotes: []ETradeQuote{
					&eTradeQuote{
						jsonMap: jsonmap.JsonMap{
							"key1": "value1",
						},
					},
					&eTradeQuote{
						jsonMap: jsonmap.JsonMap{
							"key2": "value2",
						},
					},
				},
			},
		},
		{
			name: "CreateETradeQuoteList Can Create Empty List",
			testJson: `
{
  "QuoteResponse": {
    "QuoteData": [
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeQuoteList{
				quotes: []ETradeQuote{},
			},
		},
		{
			name: "CreateETradeQuoteList Fails With Invalid Response",
			// The "Quote" level is not an array in the following string
			testJson: `
{
  "QuoteResponse": {
    "QuoteData": {
      "key": "value"
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				responseMap, err := NewNormalizedJsonMap([]byte(tt.testJson))
				require.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := CreateETradeQuoteList(responseMap)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeQuoteList_GetAllQuotes(t *testing.T) {
	tests := []struct {
		name          string
		testQuoteList ETradeQuoteList
		expectValue   []ETradeQuote
	}{
		{
			name: "GetAllQuotes Returns All Quotes",
			testQuoteList: &eTradeQuoteList{
				quotes: []ETradeQuote{
					&eTradeQuote{
						jsonMap: jsonmap.JsonMap{
							"key1": "value1",
						},
					},
					&eTradeQuote{
						jsonMap: jsonmap.JsonMap{
							"key2": "value2",
						},
					},
				},
			},
			expectValue: []ETradeQuote{
				&eTradeQuote{
					jsonMap: jsonmap.JsonMap{
						"key1": "value1",
					},
				},
				&eTradeQuote{
					jsonMap: jsonmap.JsonMap{
						"key2": "value2",
					},
				},
			},
		},
		{
			name: "GetAllQuotes Can Return Empty List",
			testQuoteList: &eTradeQuoteList{
				quotes: []ETradeQuote{},
			},
			expectValue: []ETradeQuote{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testQuoteList.GetAllQuotes()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}
