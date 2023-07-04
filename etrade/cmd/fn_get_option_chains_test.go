package cmd

import (
	"encoding/json"
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOptionChains(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Gets Option Chains",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testGetOptionChainsResponse := []byte(`
{
  "OptionChainResponse": {
	"OptionPair": [
	  {
		"testKey": "testValue"
	  }
	],
	"timeStamp": 1234,
	"quoteType": "Type",
	"nearPrice": 123.4,
	"SelectedED": {
		"testKey2": "testValue2"
	}
  }
}
`)
				mockClient.On(
					"GetOptionChains", "TestSymbol", 1, 2, 3, 4, 5, true, true, constants.OptionCategoryNil,
					constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				).Return(testGetOptionChainsResponse, nil)

				return GetOptionChains(
					mockClient, "TestSymbol", 1, 2, 3,
					4, 5, true, true,
					constants.OptionCategoryNil, constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"optionChainPairs": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"testKey": "testValue",
					},
				},
				"timeStamp": json.Number("1234"),
				"quoteType": "Type",
				"nearPrice": json.Number("123.4"),
				"selectedED": jsonmap.JsonMap{
					"testKey2": "testValue2",
				},
			},
		},
		{
			name: "Fails On GetOptionChains Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On(
					"GetOptionChains", "TestSymbol", 1, 2, 3, 4, 5, true, true, constants.OptionCategoryNil,
					constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				).Return([]byte(nil), errors.New("test error"))

				return GetOptionChains(
					mockClient, "TestSymbol", 1, 2, 3,
					4, 5, true, true,
					constants.OptionCategoryNil, constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testGetOptionChainsResponse := []byte(`
{
  "MISSING": {
	"OptionPair": [
	  {
		"testKey": "testValue"
	  }
	],
	"timeStamp": 1234,
	"quoteType": "Type",
	"nearPrice": 123.4,
	"SelectedED": {
		"testKey2": "testValue2"
	}
  }
}
`)
				mockClient.On(
					"GetOptionChains", "TestSymbol", 1, 2, 3, 4, 5, true, true, constants.OptionCategoryNil,
					constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				).Return(testGetOptionChainsResponse, nil)

				return GetOptionChains(
					mockClient, "TestSymbol", 1, 2, 3,
					4, 5, true, true,
					constants.OptionCategoryNil, constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockClient := client.ETradeClientMock{}
				// Call the Method Under Test
				actualValue, err := tt.testFn(&mockClient)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, actualValue)
				mockClient.AssertExpectations(t)
			},
		)
	}
}
