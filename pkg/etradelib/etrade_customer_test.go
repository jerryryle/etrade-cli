package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestETradeCustomer_GetAllAccounts(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue []ETradeAccount
	}{
		{
			name: "GetAllAccounts Returns All Accounts",
			testJson: `
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "1",
          "accountIdKey": "2",
          "accountMode": "3",
          "accountDesc": "4",
          "accountName": "5",
          "accountType": "6",
          "institutionType": "7",
          "accountStatus": "8",
          "closedDate": 9,
          "shareWorksAccount": true,
          "fcManagedMssbClosedAccount": true
        }
      ]
    }
  }
}`,
			expectErr: false,
			expectValue: []ETradeAccount{
				&eTradeAccount{
					accountInfo: jsonmap.JsonMap{
						"accountId":                  "1",
						"accountIdKey":               "2",
						"accountMode":                "3",
						"accountDesc":                "4",
						"accountName":                "5",
						"accountType":                "6",
						"institutionType":            "7",
						"accountStatus":              "8",
						"closedDate":                 json.Number("9"),
						"shareWorksAccount":          true,
						"fcManagedMssbClosedAccount": true,
					},
					accountId:    "1",
					accountIdKey: "2",
				},
			},
		},
		{
			name: "GetAllAccounts Fails If Account ID Is Missing",
			testJson: `
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountIdKey": "2",
          "accountMode": "3",
          "accountDesc": "4",
          "accountName": "5",
          "accountType": "6",
          "institutionType": "7",
          "accountStatus": "8",
          "closedDate": 9,
          "shareWorksAccount": true,
          "fcManagedMssbClosedAccount": true
        }
      ]
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "GetAllAccounts Fails If Account ID Key Is Missing",
			testJson: `
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "1",
          "accountMode": "3",
          "accountDesc": "4",
          "accountName": "5",
          "accountType": "6",
          "institutionType": "7",
          "accountStatus": "8",
          "closedDate": 9,
          "shareWorksAccount": true,
          "fcManagedMssbClosedAccount": true
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
				clientFake := client.NewClientFake(tt.testJson, nil)
				testCustomer := CreateETradeCustomer(clientFake, "TestCustomerName")
				// Call the Method Under Test
				testResultValue, err := testCustomer.GetAllAccounts()
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
					require.Equal(t, len(tt.expectValue), len(testResultValue))
					for i := range testResultValue {
						assert.Equal(t, tt.expectValue[i].GetAccountInfo(), testResultValue[i].GetAccountInfo())
						assert.Equal(t, tt.expectValue[i].GetAccountId(), testResultValue[i].GetAccountId())
						assert.Equal(t, tt.expectValue[i].GetAccountIdKey(), testResultValue[i].GetAccountIdKey())
					}
				}
			},
		)
	}
}

func TestETradeCustomer_GetAccountById(t *testing.T) {
	tests := []struct {
		name        string
		testId      string
		testJson    string
		expectErr   bool
		expectValue ETradeAccount
	}{
		{
			name:   "GetAccountById Returns Account For Matching ID",
			testId: "1",
			testJson: `
{
 "AccountListResponse": {
   "Accounts": {
     "Account": [
       {
         "accountId": "1",
         "accountIdKey": "2",
         "accountMode": "3",
         "accountDesc": "4",
         "accountName": "5",
         "accountType": "6",
         "institutionType": "7",
         "accountStatus": "8",
         "closedDate": 9,
         "shareWorksAccount": true,
         "fcManagedMssbClosedAccount": true
       }
     ]
   }
 }
}`,
			expectErr: false,
			expectValue: &eTradeAccount{
				accountInfo: jsonmap.JsonMap{
					"accountId":                  "1",
					"accountIdKey":               "2",
					"accountMode":                "3",
					"accountDesc":                "4",
					"accountName":                "5",
					"accountType":                "6",
					"institutionType":            "7",
					"accountStatus":              "8",
					"closedDate":                 json.Number("9"),
					"shareWorksAccount":          true,
					"fcManagedMssbClosedAccount": true,
				},
				accountId:    "1",
				accountIdKey: "2",
			},
		},
		{
			name:   "GetAccountById Returns Error If ID Not Found",
			testId: "2",
			testJson: `
{
 "AccountListResponse": {
   "Accounts": {
     "Account": [
       {
         "accountId": "1",
         "accountIdKey": "2",
         "accountMode": "3",
         "accountDesc": "4",
         "accountName": "5",
         "accountType": "6",
         "institutionType": "7",
         "accountStatus": "8",
         "closedDate": 9,
         "shareWorksAccount": true,
         "fcManagedMssbClosedAccount": true
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
				clientFake := client.NewClientFake(tt.testJson, nil)
				testCustomer := CreateETradeCustomer(clientFake, "TestCustomerName")
				// Call the Method Under Test
				testResultValue, err := testCustomer.GetAccountById(tt.testId)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
					assert.Equal(t, tt.expectValue.GetAccountInfo(), testResultValue.GetAccountInfo())
					assert.Equal(t, tt.expectValue.GetAccountId(), testResultValue.GetAccountId())
					assert.Equal(t, tt.expectValue.GetAccountIdKey(), testResultValue.GetAccountIdKey())
				}
			},
		)
	}
}
