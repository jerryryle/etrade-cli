package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeAccount(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAccount
	}{
		{
			name: "CreateETradeAccount Creates Account With Valid Response",
			testJson: `
{
  "accountId": "Account 1 ID",
  "accountIdKey": "Account 1 ID Key"
}`,
			expectErr: false,
			expectValue: &eTradeAccount{
				accountInfoMap: jsonmap.JsonMap{
					"accountId":    "Account 1 ID",
					"accountIdKey": "Account 1 ID Key",
				},
				accountId:    "Account 1 ID",
				accountIdKey: "Account 1 ID Key",
			},
		},
		{
			name: "CreateETradeAccount Fails If Missing Account ID",
			testJson: `
{
  "accountIdKey": "Account 1 ID Key"
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "CreateETradeAccount Fails If Missing Account ID Key",
			testJson: `
{
  "accountId": "Account 1 ID"
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
				actualValue, err := CreateETradeAccount(responseMap)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
					assert.Equal(t, tt.expectValue, actualValue)
				}
			},
		)
	}
}

func TestETradeAccount_GetAccountInfoMap(t *testing.T) {
	testAccount := &eTradeAccount{
		accountInfoMap: jsonmap.JsonMap{
			"accountId":    "Account 1 ID",
			"accountIdKey": "Account 1 ID Key",
		},
		accountId:    "Account 1 ID",
		accountIdKey: "Account 1 ID Key",
	}
	expectedValue := jsonmap.JsonMap{
		"accountId":    "Account 1 ID",
		"accountIdKey": "Account 1 ID Key",
	}

	actualValue := testAccount.GetAccountInfoMap()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeAccount_GetAccountId(t *testing.T) {
	testAccount := &eTradeAccount{
		accountInfoMap: jsonmap.JsonMap{
			"accountId":    "Account 1 ID",
			"accountIdKey": "Account 1 ID Key",
		},
		accountId:    "Account 1 ID",
		accountIdKey: "Account 1 ID Key",
	}
	expectedValue := "Account 1 ID"

	actualValue := testAccount.GetAccountId()
	assert.Equal(t, expectedValue, actualValue)
}

func TestETradeAccount_GetAccountIdKey(t *testing.T) {
	testAccount := &eTradeAccount{
		accountInfoMap: jsonmap.JsonMap{
			"accountId":    "Account 1 ID",
			"accountIdKey": "Account 1 ID Key",
		},
		accountId:    "Account 1 ID",
		accountIdKey: "Account 1 ID Key",
	}
	expectedValue := "Account 1 ID Key"

	actualValue := testAccount.GetAccountIdKey()
	assert.Equal(t, expectedValue, actualValue)
}
