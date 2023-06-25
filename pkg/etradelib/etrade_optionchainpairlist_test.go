package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeOptionChainPairList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeOptionChainPairList
	}{
		{
			name: "CreateETradeOptionChainPairList Creates List With Valid Response",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      },
      {
        "key2": "value2"
      }
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "nearPrice": 123.4,
    "SelectedED": {
      "key3": "value3"
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeOptionChainPairList{
				optionChainPairs: []ETradeOptionChainPair{
					&eTradeOptionChainPair{
						jsonMap: jsonmap.JsonMap{
							"key1": "value1",
						},
					},
					&eTradeOptionChainPair{
						jsonMap: jsonmap.JsonMap{
							"key2": "value2",
						},
					},
				},
				timeStamp: 1234,
				quoteType: "1234",
				nearPrice: 123.4,
				selected: jsonmap.JsonMap{
					"key3": "value3",
				},
			},
		},
		{
			name: "CreateETradeOptionChainPairList Can Create Empty List",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "nearPrice": 123.4,
    "SelectedED": {
      "key3": "value3"
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeOptionChainPairList{
				optionChainPairs: []ETradeOptionChainPair{},
				timeStamp:        1234,
				quoteType:        "1234",
				nearPrice:        123.4,
				selected: jsonmap.JsonMap{
					"key3": "value3",
				},
			},
		},
		{
			name: "CreateETradeOptionChainPairList Fails With Invalid Response",
			// The "OptionPair" level is not an array in the following string
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": {
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
				actualValue, err := CreateETradeOptionChainPairList(responseMap)
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

func TestETradeOptionChainPairList_GetAllOptionChainPairs(t *testing.T) {
	tests := []struct {
		name                    string
		testOptionChainPairList ETradeOptionChainPairList
		expectValue             []ETradeOptionChainPair
	}{
		{
			name: "GetAllOptionChainPairs Returns All OptionChainPairs",
			testOptionChainPairList: &eTradeOptionChainPairList{
				optionChainPairs: []ETradeOptionChainPair{
					&eTradeOptionChainPair{
						jsonMap: jsonmap.JsonMap{
							"key1": "value1",
						},
					},
					&eTradeOptionChainPair{
						jsonMap: jsonmap.JsonMap{
							"key2": "value2",
						},
					},
				},
			},
			expectValue: []ETradeOptionChainPair{
				&eTradeOptionChainPair{
					jsonMap: jsonmap.JsonMap{
						"key1": "value1",
					},
				},
				&eTradeOptionChainPair{
					jsonMap: jsonmap.JsonMap{
						"key2": "value2",
					},
				},
			},
		},
		{
			name: "GetAllOptionChainPairs Can Return Empty List",
			testOptionChainPairList: &eTradeOptionChainPairList{
				optionChainPairs: []ETradeOptionChainPair{},
			},
			expectValue: []ETradeOptionChainPair{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testOptionChainPairList.GetAllOptionChainPairs()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}
