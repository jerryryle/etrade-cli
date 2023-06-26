package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeAccountListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAccountList
	}{
		{
			name: "Creates Account List",
			testJson: `
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "Account 1 ID",
          "accountIdKey": "Account 1 ID Key"
        }
      ]
    }
  }
}`,
			expectErr: false,
			expectValue: &eTradeAccountList{
				accounts: []ETradeAccount{
					&eTradeAccount{
						id:    "Account 1 ID",
						idKey: "Account 1 ID Key",
						jsonMap: jsonmap.JsonMap{
							"accountId":    "Account 1 ID",
							"accountIdKey": "Account 1 ID Key",
						},
					},
				},
			},
		},
		{
			name: "Creates Empty Account List",
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
				accounts: []ETradeAccount{},
			},
		},
		{
			name: "Fails With Invalid JSON",
			testJson: `
{
  "AccountListResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails Without Account Key",
			testJson: `
{
  "AccountListResponse": {
    "Accounts": {
      "MISSING": [
        {
          "accountId": "Account 1 ID",
          "accountIdKey": "Account 1 ID Key"
        }
      ]
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails Without Account Id",
			// The "accountId" key is missing from the following string
			testJson: `
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "MISSING": "Account 1 ID",
          "accountIdKey": "Account 1 ID Key"
        }
      ]
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := CreateETradeAccountListFromResponse([]byte(tt.testJson))
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

func TestETradeAccountList_GetAllAccounts(t *testing.T) {
	tests := []struct {
		name        string
		testObject  ETradeAccountList
		expectValue []ETradeAccount
	}{
		{
			name: "Returns All Accounts",
			testObject: &eTradeAccountList{
				accounts: []ETradeAccount{
					&eTradeAccount{
						id:    "Account 1 ID",
						idKey: "Account 1 ID Key",
						jsonMap: jsonmap.JsonMap{
							"accountId":    "Account 1 ID",
							"accountIdKey": "Account 1 ID Key",
						},
					},
					&eTradeAccount{
						id:    "Account 2 ID",
						idKey: "Account 2 ID Key",
						jsonMap: jsonmap.JsonMap{
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
					jsonMap: jsonmap.JsonMap{
						"accountId":    "Account 1 ID",
						"accountIdKey": "Account 1 ID Key",
					},
				},
				&eTradeAccount{
					id:    "Account 2 ID",
					idKey: "Account 2 ID Key",
					jsonMap: jsonmap.JsonMap{
						"accountId":    "Account 2 ID",
						"accountIdKey": "Account 2 ID Key",
					},
				},
			},
		},
		{
			name: "Can Return Empty List",
			testObject: &eTradeAccountList{
				accounts: []ETradeAccount{},
			},
			expectValue: []ETradeAccount{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAllAccounts()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeAccountList_GetAccountById(t *testing.T) {
	tests := []struct {
		name        string
		testObject  ETradeAccountList
		testId      string
		expectValue ETradeAccount
	}{
		{
			name: "Returns Account For Valid ID",
			testObject: &eTradeAccountList{
				accounts: []ETradeAccount{
					&eTradeAccount{
						id:    "Account 1 ID",
						idKey: "Account 1 ID Key",
						jsonMap: jsonmap.JsonMap{
							"accountId":    "Account 1 ID",
							"accountIdKey": "Account 1 ID Key",
						},
					},
				},
			},
			testId: "Account 1 ID",
			expectValue: &eTradeAccount{
				id:    "Account 1 ID",
				idKey: "Account 1 ID Key",
				jsonMap: jsonmap.JsonMap{
					"accountId":    "Account 1 ID",
					"accountIdKey": "Account 1 ID Key",
				},
			},
		},
		{
			name: "Returns Nil For Invalid ID",
			testObject: &eTradeAccountList{
				accounts: []ETradeAccount{
					&eTradeAccount{
						id:    "Account 1 ID",
						idKey: "Account 1 ID Key",
						jsonMap: jsonmap.JsonMap{
							"accountId":    "Account 1 ID",
							"accountIdKey": "Account 1 ID Key",
						},
					},
				},
			},
			testId:      "Account 2 ID",
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAccountById(tt.testId)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeAccountList_AsJsonMap(t *testing.T) {
	testObject := &eTradeAccountList{
		accounts: []ETradeAccount{
			&eTradeAccount{
				id:    "Account 1 ID",
				idKey: "Account 1 ID Key",
				jsonMap: jsonmap.JsonMap{
					"accountId":    "Account 1 ID",
					"accountIdKey": "Account 1 ID Key",
				},
			},
		},
	}

	expectValue := jsonmap.JsonMap{
		"accounts": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"accountId":    "Account 1 ID",
				"accountIdKey": "Account 1 ID Key",
			},
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
