package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeQuoteListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeQuoteList
	}{
		{
			name: "Creates List",
			testJson: `
{
  "QuoteResponse": {
    "QuoteData": [
      {
        "key1": "value1"
      }
    ],
    "Messages": {
      "Message": [
        {
          "messageKey": "MessageValue"
        }
      ]
    }
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
				},
				messages: jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"messageKey": "MessageValue",
					},
				},
			},
		},
		{
			name: "Can Create Empty List",
			testJson: `
{
  "QuoteResponse": {
    "QuoteData": [
    ],
    "Messages": {
      "Message": [
        {
          "messageKey": "MessageValue"
        }
      ]
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeQuoteList{
				quotes: []ETradeQuote{},
				messages: jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"messageKey": "MessageValue",
					},
				},
			},
		},
		{
			name: "Fails With Invalid JSON",
			testJson: `
{
  "QuoteResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Succeeds With Missing QuoteData And Creates Empty Slice",
			testJson: `
{
  "QuoteResponse": {
    "MISSING": [
    ],
    "Messages": {
      "Message": [
        {
          "messageKey": "MessageValue"
        }
      ]
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeQuoteList{
				quotes: []ETradeQuote{},
				messages: jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"messageKey": "MessageValue",
					},
				},
			},
		},
		{
			name: "Succeeds With Missing Messages And Creates Nil Slice",
			testJson: `
{
  "QuoteResponse": {
    "QuoteData": [
    ],
    "MISSING": {
      "Message": [
        {
          "messageKey": "MessageValue"
        }
      ]
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeQuoteList{
				quotes:   []ETradeQuote{},
				messages: jsonmap.JsonSlice(nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := CreateETradeQuoteListFromResponse([]byte(tt.testJson))
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
		name        string
		testObject  ETradeQuoteList
		expectValue []ETradeQuote
	}{
		{
			name: "Returns All Quotes",
			testObject: &eTradeQuoteList{
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
			name: "Can Return Empty List",
			testObject: &eTradeQuoteList{
				quotes: []ETradeQuote{},
			},
			expectValue: []ETradeQuote{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAllQuotes()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeQuoteList_AsJsonMap(t *testing.T) {
	testObject := &eTradeQuoteList{
		quotes: []ETradeQuote{
			&eTradeQuote{
				jsonMap: jsonmap.JsonMap{
					"key1": "value1",
				},
			},
		},
		messages: jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"messageKey": "MessageValue",
			},
		},
	}

	expectValue := jsonmap.JsonMap{
		"quotes": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"key1": "value1",
			},
		},
		"messages": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"messageKey": "MessageValue",
			},
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
