package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeAuthenticationStatus(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAuthenticationStatus
	}{
		{
			name: "Creates Authorize AuthenticationStatus",
			testJson: `
{
  "status": "authorize",
  "authorizationUrl": "test url"
}`,
			expectErr: false,
			expectValue: &eTradeAuthenticationStatus{
				authorizationUrl: "test url",
				jsonMap: jsonmap.JsonMap{
					"status":           "authorize",
					"authorizationUrl": "test url",
				},
			},
		},
		{
			name: "Creates Success Status",
			testJson: `
{
  "status": "success"
}`,
			expectErr: false,
			expectValue: &eTradeAuthenticationStatus{
				authorizationUrl: "",
				jsonMap: jsonmap.JsonMap{
					"status": "success",
				},
			},
		},
		{
			name: "Fails Without Status",
			testJson: `
{
  "authorizationUrl": "test url"
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Authorize Status Fails Without Url",
			testJson: `
{
  "status": "authorize"
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails On Bad JSON",
			testJson: `
{
  "status": 
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := CreateETradeAuthenticationStatusFromResponse([]byte(tt.testJson))
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

func TestETradeAuthenticationStatus_AsJsonMap(t *testing.T) {
	testObject := &eTradeAuthenticationStatus{
		authorizationUrl: "test url",
		jsonMap: jsonmap.JsonMap{
			"status":           "authorize",
			"authorizationUrl": "test url",
		},
	}

	expectedValue := jsonmap.JsonMap{
		"status":           "authorize",
		"authorizationUrl": "test url",
	}

	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
