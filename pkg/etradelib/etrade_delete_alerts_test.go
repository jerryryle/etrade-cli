package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeDeleteAlerts(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeDeleteAlerts
	}{
		{
			name: "Creates Success DeleteAlerts",
			testJson: `
{
  "AlertsResponse": {
    "result": "SUCCESS"
  }
}`,
			expectErr: false,
			expectValue: &eTradeDeleteAlerts{
				isSuccess:    true,
				failedAlerts: []int64{},
			},
		},
		{
			name: "Creates Error DeleteAlerts",
			testJson: `
{
  "AlertsResponse": {
    "result": "ERROR",
    "failedAlerts": {
      "alertId": [
        1234
      ]
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeDeleteAlerts{
				isSuccess:    false,
				failedAlerts: []int64{1234},
			},
		},
		{
			name: "Fails With Unexpected Result",
			testJson: `
{
  "AlertsResponse": {
    "result": "BOGUS RESULT"
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails Without Result",
			testJson: `
{
  "AlertsResponse": {
    "result": "BOGUS RESULT"
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Succeeds Without Failed Alerts",
			testJson: `
{
  "AlertsResponse": {
    "result": "ERROR"
  }
}`,
			expectErr: false,
			expectValue: &eTradeDeleteAlerts{
				isSuccess:    false,
				failedAlerts: []int64{},
			},
		},
		{
			name: "Fails On Bad JSON",
			testJson: `
{
  "AlertsResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := CreateETradeDeleteAlertsFromResponse([]byte(tt.testJson))
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

func TestETradeDeleteAlerts_IsSuccess(t *testing.T) {
	testSuccessDeleteAlerts := eTradeDeleteAlerts{
		isSuccess:    true,
		failedAlerts: []int64{},
	}

	testErrorDeleteAlerts := eTradeDeleteAlerts{
		isSuccess:    false,
		failedAlerts: []int64{1234},
	}

	assert.True(t, testSuccessDeleteAlerts.IsSuccess())
	assert.False(t, testErrorDeleteAlerts.IsSuccess())
}

func TestETradeDeleteAlerts_GetFailedAlerts(t *testing.T) {
	testErrorDeleteAlerts := eTradeDeleteAlerts{
		isSuccess:    false,
		failedAlerts: []int64{1234},
	}

	assert.Equal(t, []int64{1234}, testErrorDeleteAlerts.GetFailedAlerts())
}

func TestETradeDeleteAlerts_AsJsonMap(t *testing.T) {
	testObject := eTradeDeleteAlerts{
		isSuccess:    false,
		failedAlerts: []int64{1234},
	}

	expectedValue := jsonmap.JsonMap{
		"status":       "error",
		"error":        "some alerts could not be deleted",
		"failedAlerts": jsonmap.JsonSlice{int64(1234)},
	}

	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
