package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAlerts(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Deletes Alerts",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testDeleteAlertsResponse := []byte(`
{
  "AlertsResponse": {
    "result": "SUCCESS"
  }
}`)
				mockClient.On("DeleteAlerts", []string{"1234"}).Return(testDeleteAlertsResponse, nil)
				return DeleteAlerts(mockClient, []string{"1234"})
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"status": "success",
			},
		},
		{
			name: "Fails On Bad Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testDeleteAlertsResponse := []byte(`
{
  "AlertsResponse": {
}`)
				mockClient.On("DeleteAlerts", []string{"1234"}).Return(testDeleteAlertsResponse, nil)
				return DeleteAlerts(mockClient, []string{"1234"})
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On DeleteAlerts Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("DeleteAlerts", []string{"1234"}).Return([]byte{}, errors.New("test error"))
				return DeleteAlerts(mockClient, []string{"1234"})
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
