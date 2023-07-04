package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeOptionChainPairListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeOptionChainPairList
	}{
		{
			name: "Creates List",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "nearPrice": 123.4,
    "SelectedED": {
      "key2": "value2"
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
				},
				timeStamp: 1234,
				quoteType: "1234",
				nearPrice: 123.4,
				selected: jsonmap.JsonMap{
					"key2": "value2",
				},
			},
		},
		{
			name: "Creates Empty List",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "nearPrice": 123.4,
    "SelectedED": {
      "key1": "value1"
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
					"key1": "value1",
				},
			},
		},
		{
			name: "Fails With Invalid JSON",
			testJson: `
{
  "OptionChainResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing OptionPair",
			testJson: `
{
  "OptionChainResponse": {
    "MISSING": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "nearPrice": 123.4,
    "SelectedED": {
      "key2": "value2"
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Defaults Missing timeStamp",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "MISSING": 1234,
    "quoteType": "1234",
    "nearPrice": 123.4,
    "SelectedED": {
      "key2": "value2"
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
				},
				timeStamp: 0,
				quoteType: "1234",
				nearPrice: 123.4,
				selected: jsonmap.JsonMap{
					"key2": "value2",
				},
			},
		},
		{
			name: "Defaults Missing quoteType",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": 1234,
    "MISSING": "1234",
    "nearPrice": 123.4,
    "SelectedED": {
      "key2": "value2"
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
				},
				timeStamp: 1234,
				quoteType: "",
				nearPrice: 123.4,
				selected: jsonmap.JsonMap{
					"key2": "value2",
				},
			},
		},
		{
			name: "Defaults Missing nearPrice",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "MISSING": 123.4,
    "SelectedED": {
      "key2": "value2"
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
				},
				timeStamp: 1234,
				quoteType: "1234",
				nearPrice: 0,
				selected: jsonmap.JsonMap{
					"key2": "value2",
				},
			},
		},
		{
			name: "Defaults Missing SelectedED",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "nearPrice": 123.4,
    "MISSING": {
      "key2": "value2"
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
				},
				timeStamp: 1234,
				quoteType: "1234",
				nearPrice: 123.4,
				selected:  nil,
			},
		},
		{
			name: "Fails If Wrong Type For timeStamp",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": [1234],
    "quoteType": "1234",
    "nearPrice": 123.4,
    "SelectedED": {
      "key2": "value2"
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails If Wrong Type For quoteType",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": 1234,
    "quoteType": ["1234"],
    "nearPrice": 123.4,
    "SelectedED": {
      "key2": "value2"
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails If Wrong Type For nearPrice",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "nearPrice": [123.4],
    "SelectedED": {
      "key2": "value2"
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails If Wrong Type For SelectedED",
			testJson: `
{
  "OptionChainResponse": {
    "OptionPair": [
      {
        "key1": "value1"
      }
    ],
    "timeStamp": 1234,
    "quoteType": "1234",
    "nearPrice": 123.4,
    "SelectedED": [{
      "key2": "value2"
    }]
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := CreateETradeOptionChainPairListFromResponse([]byte(tt.testJson))
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
		name        string
		testObject  ETradeOptionChainPairList
		expectValue []ETradeOptionChainPair
	}{
		{
			name: "Returns All OptionChainPairs",
			testObject: &eTradeOptionChainPairList{
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
			name: "Can Return Empty List",
			testObject: &eTradeOptionChainPairList{
				optionChainPairs: []ETradeOptionChainPair{},
			},
			expectValue: []ETradeOptionChainPair{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAllOptionChainPairs()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeOptionChainPairList_AsJsonMap(t *testing.T) {
	testObject := &eTradeOptionChainPairList{
		optionChainPairs: []ETradeOptionChainPair{
			&eTradeOptionChainPair{
				jsonMap: jsonmap.JsonMap{
					"key1": "value1",
				},
			},
		},
		timeStamp: 1234,
		quoteType: "1234",
		nearPrice: 123.4,
		selected: jsonmap.JsonMap{
			"key2": "value2",
		},
	}

	expectValue := jsonmap.JsonMap{
		"optionChainPairs": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"key1": "value1",
			},
		},
		"timeStamp": json.Number("1234"),
		"quoteType": "1234",
		"nearPrice": json.Number("123.4"),
		"selected": jsonmap.JsonMap{
			"key2": "value2",
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
