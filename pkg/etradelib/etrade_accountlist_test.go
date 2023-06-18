package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeAccountList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAccountList
	}{
		{
			name: "CreateETradeAccountList Creates List With Valid Response",
			testJson: `
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "Account 1 ID",
          "accountIdKey": "Account 1 ID Key"
        },
        {
          "accountId": "Account 2 ID",
          "accountIdKey": "Account 2 ID Key"
        }
      ]
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeAccountList{
				[]ETradeAccount{
					&eTradeAccount{
						id:    "Account 1 ID",
						idKey: "Account 1 ID Key",
						infoMap: jsonmap.JsonMap{
							"accountId":    "Account 1 ID",
							"accountIdKey": "Account 1 ID Key",
						},
					},
					&eTradeAccount{
						id:    "Account 2 ID",
						idKey: "Account 2 ID Key",
						infoMap: jsonmap.JsonMap{
							"accountId":    "Account 2 ID",
							"accountIdKey": "Account 2 ID Key",
						},
					},
				},
			},
		},
		{
			name: "CreateETradeAccountList Can Create Empty List",
			testJson: `
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
      ]
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeAccountList{
				[]ETradeAccount{},
			},
		},
		{
			name: "CreateETradeAccountList Fails With Invalid Response",
			// The "Account" level is missing from the following string
			testJson: `
{
  "AccountListResponse": {
    "Accounts": [
      {
        "accountId": "Account 1 ID",
        "accountIdKey": "Account 1 ID Key"
      },
      {
        "accountId": "Account 2 ID",
        "accountIdKey": "Account 2 ID Key"
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
				responseMap, err := NewNormalizedJsonMap([]byte(tt.testJson))
				require.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := CreateETradeAccountList(responseMap)
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

func TestETradeAccountList_GetAllAccounts(t *testing.T) {
	tests := []struct {
		name            string
		testAccountList ETradeAccountList
		expectValue     []ETradeAccount
	}{
		{
			name: "GetAllAccounts Returns All Accounts",
			testAccountList: &eTradeAccountList{
				[]ETradeAccount{
					&eTradeAccount{
						id:    "Account 1 ID",
						idKey: "Account 1 ID Key",
						infoMap: jsonmap.JsonMap{
							"accountId":    "Account 1 ID",
							"accountIdKey": "Account 1 ID Key",
						},
					},
					&eTradeAccount{
						id:    "Account 2 ID",
						idKey: "Account 2 ID Key",
						infoMap: jsonmap.JsonMap{
							"accountId":    "Account 2 ID",
							"accountIdKey": "Account 2 ID Key",
						},
					},
				},
			},
			expectValue: []ETradeAccount{
				&eTradeAccount{
					id:    "Account 1 ID",
					idKey: "Account 1 ID Key",
					infoMap: jsonmap.JsonMap{
						"accountId":    "Account 1 ID",
						"accountIdKey": "Account 1 ID Key",
					},
				},
				&eTradeAccount{
					id:    "Account 2 ID",
					idKey: "Account 2 ID Key",
					infoMap: jsonmap.JsonMap{
						"accountId":    "Account 2 ID",
						"accountIdKey": "Account 2 ID Key",
					},
				},
			},
		},
		{
			name: "GetAllAccounts Can Return Empty List",
			testAccountList: &eTradeAccountList{
				[]ETradeAccount{},
			},
			expectValue: []ETradeAccount{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testAccountList.GetAllAccounts()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeAccountList_GetAccountById(t *testing.T) {
	tests := []struct {
		name            string
		testAccountList ETradeAccountList
		testAccountID   string
		expectValue     ETradeAccount
	}{
		{
			name: "GetAccountById Returns Account For Valid ID",
			testAccountList: &eTradeAccountList{
				[]ETradeAccount{
					&eTradeAccount{
						id:    "Account 1 ID",
						idKey: "Account 1 ID Key",
						infoMap: jsonmap.JsonMap{
							"accountId":    "Account 1 ID",
							"accountIdKey": "Account 1 ID Key",
						},
					},
				},
			},
			testAccountID: "Account 1 ID",
			expectValue: &eTradeAccount{
				id:    "Account 1 ID",
				idKey: "Account 1 ID Key",
				infoMap: jsonmap.JsonMap{
					"accountId":    "Account 1 ID",
					"accountIdKey": "Account 1 ID Key",
				},
			},
		},
		{
			name: "GetAccountById Returns Nil For Invalid ID",
			testAccountList: &eTradeAccountList{
				[]ETradeAccount{
					&eTradeAccount{
						id:    "Account 1 ID",
						idKey: "Account 1 ID Key",
						infoMap: jsonmap.JsonMap{
							"accountId":    "Account 1 ID",
							"accountIdKey": "Account 1 ID Key",
						},
					},
				},
			},
			testAccountID: "Account 2 ID",
			expectValue:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testAccountList.GetAccountById(tt.testAccountID)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}
