package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				responseMap, err := NewNormalizedJsonMap([]byte(tt.testJson))
				require.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := CreateETradeAuthenticationStatus(responseMap)
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
