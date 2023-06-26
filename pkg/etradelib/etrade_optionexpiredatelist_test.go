package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeOptionExpireDateListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeOptionExpireDateList
	}{
		{
			name: "Creates List",
			testJson: `
{
  "OptionExpireDateResponse": {
    "ExpirationDate": [
      {
        "key1": "value1"
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
				},
			},
		},
		{
			name: "Creates Empty List",
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
			name: "Fails With Invalid JSON",
			testJson: `
{
  "OptionExpireDateResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing ExpirationDate",
			testJson: `
{
  "OptionExpireDateResponse": {
    "MISSING": [
      {
        "key1": "value1"
      }
    ]
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
				actualValue, err := CreateETradeOptionExpireDateListFromResponse([]byte(tt.testJson))
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
		name        string
		testObject  ETradeOptionExpireDateList
		expectValue []ETradeOptionExpireDate
	}{
		{
			name: "Returns All OptionExpireDates",
			testObject: &eTradeOptionExpireDateList{
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
			name: "Can Return Empty List",
			testObject: &eTradeOptionExpireDateList{
				optionExpireDates: []ETradeOptionExpireDate{},
			},
			expectValue: []ETradeOptionExpireDate{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAllOptionExpireDates()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeOptionExpireDateList_AsJsonMap(t *testing.T) {
	testObject := &eTradeOptionExpireDateList{
		optionExpireDates: []ETradeOptionExpireDate{
			&eTradeOptionExpireDate{
				jsonMap: jsonmap.JsonMap{
					"key1": "value1",
				},
			},
		},
	}

	expectValue := jsonmap.JsonMap{
		"optionExpireDates": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"key1": "value1",
			},
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
