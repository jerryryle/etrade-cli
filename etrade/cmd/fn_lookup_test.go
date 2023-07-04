package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLookup(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Lookup",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testLookupResponse := []byte(`
{
  "LookupResponse": {
    "Data": [
      {
        "testKey": "testValue"
      }
    ]
  }
}`)
				mockClient.On("LookupProduct", "test search").Return(testLookupResponse, nil)
				return Lookup(
					mockClient, "test search",
				)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"results": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"testKey": "testValue",
					},
				},
			},
		},
		{
			name: "Fails On LookupProduct Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("LookupProduct", "test search").Return([]byte{}, errors.New("test error"))
				return Lookup(
					mockClient, "test search",
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testLookupResponse := []byte(`
{
  "LookupResponse": {
}`)
				mockClient.On("LookupProduct", "test search").Return(testLookupResponse, nil)
				return Lookup(
					mockClient, "test search",
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
