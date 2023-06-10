package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestETradeClient_ListAccounts(t *testing.T) {
	testXml := `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<AccountListResponse>
    <Accounts>
        <Account>
            <accountId>82314598</accountId>
            <accountIdKey>dBZOKt9xDrtRSAOl4MSiiA</accountIdKey>
            <accountMode>IRA</accountMode>
            <accountDesc>Brokerage</accountDesc>
            <accountName>NickName-1</accountName>
            <accountType>MARGIN</accountType>
            <institutionType>BROKERAGE</institutionType>
            <accountStatus>ACTIVE</accountStatus>
            <closedDate>0</closedDate>
            <shareWorksAccount>false</shareWorksAccount>
            <fcManagedMssbClosedAccount>false</fcManagedMssbClosedAccount>
        </Account>
        <Account>
            <accountId>58315636</accountId>
            <accountIdKey>vQMsebA1H5WltUfDkJP48g</accountIdKey>
            <accountMode>CASH</accountMode>
            <accountDesc>Complete Savings</accountDesc>
            <accountName>NickName-2</accountName>
            <accountType>INDIVIDUAL</accountType>
            <institutionType>BROKERAGE</institutionType>
            <accountStatus>ACTIVE</accountStatus>
            <closedDate>0</closedDate>
            <shareWorksAccount>false</shareWorksAccount>
            <fcManagedMssbClosedAccount>false</fcManagedMssbClosedAccount>
        </Account>
    </Accounts>
</AccountListResponse>`
	expectedResponse := responses.AccountListResponse{
		Accounts: []responses.AccountListAccount{
			{
				AccountId:                  "82314598",
				AccountIdKey:               "dBZOKt9xDrtRSAOl4MSiiA",
				AccountMode:                "IRA",
				AccountDesc:                "Brokerage",
				AccountName:                "NickName-1",
				AccountType:                "MARGIN",
				InstitutionType:            "BROKERAGE",
				AccountStatus:              "ACTIVE",
				ClosedDate:                 responses.ETradeTime{Time: time.Unix(0, 0)},
				ShareWorksAccount:          false,
				ShareWorksSource:           "",
				FcManagedMssbClosedAccount: false,
			},
			{
				AccountId:                  "58315636",
				AccountIdKey:               "vQMsebA1H5WltUfDkJP48g",
				AccountMode:                "CASH",
				AccountDesc:                "Complete Savings",
				AccountName:                "NickName-2",
				AccountType:                "INDIVIDUAL",
				InstitutionType:            "BROKERAGE",
				AccountStatus:              "ACTIVE",
				ClosedDate:                 responses.ETradeTime{Time: time.Unix(0, 0)},
				ShareWorksAccount:          false,
				ShareWorksSource:           "",
				FcManagedMssbClosedAccount: false,
			},
		},
	}

	httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testXml)),
		}
	})

	client := CreateETradeClient(GetEndpointUrls(true), httpClient)
	response, err := client.ListAccounts()
	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, response)
}

func TestETradeClient_ListAlerts(t *testing.T) {
	testXml := `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<AlertsResponse>
    <totalAlerts>4</totalAlerts>
    <Alert>
        <id>1107</id>
        <createTime>1299519940</createTime>
        <subject>Bank Statement Available Feb'12</subject>
        <status>UNREAD</status>
    </Alert>
    <Alert>
        <id>1099</id>
        <createTime>1328115640</createTime>
        <subject>Bank Statement Available for Jan'12</subject>
        <status>READ</status>
    </Alert>
    <Alert>
        <id>1090</id>
        <createTime>1315230340</createTime>
        <subject>Buy 2 MSFT Cancelled</subject>
        <status>UNDELETED</status>
    </Alert>
    <Alert>
        <id>1089</id>
        <createTime>1314888340</createTime>
        <subject>Buy 4 IBM Cancelled</subject>
        <status>DELETED</status>
    </Alert>
</AlertsResponse>`
	expectedResponse := responses.AlertsResponse{
		TotalAlerts: 4,
		Alerts: []responses.AlertsAlert{
			{
				Id:         1107,
				CreateTime: responses.ETradeTime{Time: time.Unix(1299519940, 0)},
				Subject:    "Bank Statement Available Feb'12",
				Status:     "UNREAD",
			},
			{
				Id:         1099,
				CreateTime: responses.ETradeTime{Time: time.Unix(1328115640, 0)},
				Subject:    "Bank Statement Available for Jan'12",
				Status:     "READ",
			},
			{
				Id:         1090,
				CreateTime: responses.ETradeTime{Time: time.Unix(1315230340, 0)},
				Subject:    "Buy 2 MSFT Cancelled",
				Status:     "UNDELETED",
			},
			{
				Id:         1089,
				CreateTime: responses.ETradeTime{Time: time.Unix(1314888340, 0)},
				Subject:    "Buy 4 IBM Cancelled",
				Status:     "DELETED",
			},
		},
	}

	httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testXml)),
		}
	})

	client := CreateETradeClient(GetEndpointUrls(true), httpClient)
	response, err := client.ListAlerts()
	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, response)
}
