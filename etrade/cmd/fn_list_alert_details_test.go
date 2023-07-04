package cmd

import (
	"encoding/json"
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAlertDetails(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Lists Alert Details",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAlertDetailsResponse := []byte(`
{
  "AlertDetailsResponse": {
    "id": 1234
  }
}`)
				mockClient.On("ListAlertDetails", "1234", false).Return(testAlertDetailsResponse, nil)
				return ListAlertDetails(mockClient, "1234")
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"id": json.Number("1234"),
			},
		},
		{
			name: "Fails On ListAlertDetails Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAlertDetails", "1234", false).Return([]byte{}, errors.New("test error"))
				return ListAlertDetails(mockClient, "1234")
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAlertDetailsResponse := []byte(`
{
  "AlertDetailsResponse": {
}`)
				mockClient.On("ListAlertDetails", "1234", false).Return(testAlertDetailsResponse, nil)
				return ListAlertDetails(mockClient, "1234")
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
