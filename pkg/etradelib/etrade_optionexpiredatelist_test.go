package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeOptionExpireDateList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeOptionExpireDateList
	}{
		{
			name: "CreateETradeOptionExpireDateList Creates List With Valid Response",
			testJson: `
{
  "OptionExpireDateResponse": {
    "ExpirationDate": [
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
			expectValue: &eTradeOptionExpireDateList{
				optionExpireDates: []ETradeOptionExpireDate{
					&eTradeOptionExpireDate{
						jsonMap: jsonmap.JsonMap{
							"key1": "value1",
						},
					},
					&eTradeOptionExpireDate{
						jsonMap: jsonmap.JsonMap{
							"key2": "value2",
						},
					},
				},
			},
		},
		{
			name: "CreateETradeOptionExpireDateList Can Create Empty List",
			testJson: `
{
  "OptionExpireDateResponse": {
    "ExpirationDate": [
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeOptionExpireDateList{
				optionExpireDates: []ETradeOptionExpireDate{},
			},
		},
		{
			name: "CreateETradeOptionExpireDateList Fails With Invalid Response",
			// The "OptionExpireDate" level is not an array in the following string
			testJson: `
{
  "OptionExpireDateResponse": {
    "ExpirationDate": {
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
				actualValue, err := CreateETradeOptionExpireDateList(responseMap)
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

func TestETradeOptionExpireDateList_GetAllOptionExpireDates(t *testing.T) {
	tests := []struct {
		name                     string
		testOptionExpireDateList ETradeOptionExpireDateList
		expectValue              []ETradeOptionExpireDate
	}{
		{
			name: "GetAllOptionExpireDates Returns All OptionExpireDates",
			testOptionExpireDateList: &eTradeOptionExpireDateList{
				optionExpireDates: []ETradeOptionExpireDate{
					&eTradeOptionExpireDate{
						jsonMap: jsonmap.JsonMap{
							"key1": "value1",
						},
					},
					&eTradeOptionExpireDate{
						jsonMap: jsonmap.JsonMap{
							"key2": "value2",
						},
					},
				},
			},
			expectValue: []ETradeOptionExpireDate{
				&eTradeOptionExpireDate{
					jsonMap: jsonmap.JsonMap{
						"key1": "value1",
					},
				},
				&eTradeOptionExpireDate{
					jsonMap: jsonmap.JsonMap{
						"key2": "value2",
					},
				},
			},
		},
		{
			name: "GetAllOptionExpireDates Can Return Empty List",
			testOptionExpireDateList: &eTradeOptionExpireDateList{
				optionExpireDates: []ETradeOptionExpireDate{},
			},
			expectValue: []ETradeOptionExpireDate{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testOptionExpireDateList.GetAllOptionExpireDates()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}
