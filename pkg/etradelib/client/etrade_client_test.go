package client

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/etradelibtest"
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
</AccountListResponse>
`
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
				ClosedDate:                 responses.ETradeTime{Time: time.Unix(0, 0).UTC()},
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
				ClosedDate:                 responses.ETradeTime{Time: time.Unix(0, 0).UTC()},
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

	client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
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
</AlertsResponse>
`
	expectedResponse := responses.AlertsResponse{
		TotalAlerts: 4,
		Alerts: []responses.AlertsAlert{
			{
				Id:         1107,
				CreateTime: responses.ETradeTime{Time: time.Unix(1299519940, 0).UTC()},
				Subject:    "Bank Statement Available Feb'12",
				Status:     "UNREAD",
			},
			{
				Id:         1099,
				CreateTime: responses.ETradeTime{Time: time.Unix(1328115640, 0).UTC()},
				Subject:    "Bank Statement Available for Jan'12",
				Status:     "READ",
			},
			{
				Id:         1090,
				CreateTime: responses.ETradeTime{Time: time.Unix(1315230340, 0).UTC()},
				Subject:    "Buy 2 MSFT Cancelled",
				Status:     "UNDELETED",
			},
			{
				Id:         1089,
				CreateTime: responses.ETradeTime{Time: time.Unix(1314888340, 0).UTC()},
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

	client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
	response, err := client.ListAlerts()
	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, response)
}

func TestETradeClient_GetQuotes(t *testing.T) {
	type args struct {
		symbols           []string
		detailFlag        QuoteDetailFlag
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectErr bool
		expect    *responses.QuoteResponse
	}{
		{
			name: "Valid QuoteDetailAll XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailAll,
				httpClientFakeXml: quoteDetailAllTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailAllTestResponse,
		},
		{
			name: "Valid QuoteDetailFundamental XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailFundamental,
				httpClientFakeXml: quoteDetailFundamentalTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailFundamentalTestResponse,
		},
		{
			name: "Valid QuoteDetailIntraday XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailIntraday,
				httpClientFakeXml: quoteDetailIntradayTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailIntradayTestResponse,
		},
		{
			name: "Valid QuoteDetailOptions XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailOptions,
				httpClientFakeXml: quoteDetailOptionsTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailOptionsTestResponse,
		},
		{
			name: "Valid QuoteDetailWeek52 XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailWeek52,
				httpClientFakeXml: quoteDetailWeek52TestXml,
			},
			expectErr: false,
			expect:    &quoteDetailWeek52TestResponse,
		},
		{
			name: "Valid QuoteDetailMutualFund XML",
			args: args{
				symbols:           []string{"VFIAX"},
				detailFlag:        QuoteDetailMutualFund,
				httpClientFakeXml: quoteDetailMutualFundTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailMutualFundTestResponse,
		},
		{
			name: "Too Many Symbols",
			args: args{
				symbols:           []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51"},
				detailFlag:        QuoteDetailAll,
				httpClientFakeXml: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><QuoteResponse></QuoteResponse>`,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
				}
			})

			client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
			response, err := client.GetQuotes(tt.args.symbols, tt.args.detailFlag)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, response)
		})
	}
}

func TestETradeClient_LookupProduct(t *testing.T) {
	type args struct {
		search            string
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectErr bool
		expect    *responses.LookupResponse
	}{
		{
			name: "Valid Search With Results",
			args: args{
				search:            "A",
				httpClientFakeXml: lookupProductResultsTestXml,
			},
			expectErr: false,
			expect:    &lookupProductResultsTestResponse,
		},
		{
			name: "Valid Search With No Results",
			args: args{
				search:            "A",
				httpClientFakeXml: lookupProductNoResultsTestXml,
			},
			expectErr: false,
			expect:    &lookupProductNoResultsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
				}
			})

			client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
			response, err := client.LookupProduct(tt.args.search)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, response)
		})
	}
}
