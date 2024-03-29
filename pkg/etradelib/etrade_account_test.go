package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeAccountFromMap(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAccount
	}{
		{
			name: "Creates Account With Valid Response",
			testJson: `
{
  "accountId": "Account 1 ID",
  "accountIdKey": "Account 1 ID Key"
}`,
			expectErr: false,
			expectValue: &eTradeAccount{
				jsonMap: jsonmap.JsonMap{
					"accountId":    "Account 1 ID",
					"accountIdKey": "Account 1 ID Key",
				},
				id:    "Account 1 ID",
				idKey: "Account 1 ID Key",
			},
		},
		{
			name: "Fails If Missing Account ID",
			testJson: `
{
  "MISSING": "Account 1 ID",
  "accountIdKey": "Account 1 ID Key"
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "CreateETradeAccountFromMap Fails If Missing Account ID Key",
			testJson: `
{
  "accountId": "Account 1 ID",
  "MISSING": "Account 1 ID Key"
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
				actualValue, err := CreateETradeAccountFromMap(responseMap)
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

func TestETradeAccount_GetId(t *testing.T) {
	testObject := &eTradeAccount{
		jsonMap: jsonmap.JsonMap{
			"accountId":    "Account 1 ID",
			"accountIdKey": "Account 1 ID Key",
		},
		id:    "Account 1 ID",
		idKey: "Account 1 ID Key",
	}
	expectedValue := "Account 1 ID"

	actualValue := testObject.GetId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeAccount_GetIdKey(t *testing.T) {
	testObject := &eTradeAccount{
		jsonMap: jsonmap.JsonMap{
			"accountId":    "Account 1 ID",
			"accountIdKey": "Account 1 ID Key",
		},
		id:    "Account 1 ID",
		idKey: "Account 1 ID Key",
	}
	expectedValue := "Account 1 ID Key"

	actualValue := testObject.GetIdKey()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeAccount_AsJsonMap(t *testing.T) {
	testObject := &eTradeAccount{
		jsonMap: jsonmap.JsonMap{
			"accountId":    "Account 1 ID",
			"accountIdKey": "Account 1 ID Key",
		},
		id:    "Account 1 ID",
		idKey: "Account 1 ID Key",
	}
	expectedValue := jsonmap.JsonMap{
		"accountId":    "Account 1 ID",
		"accountIdKey": "Account 1 ID Key",
	}

	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
