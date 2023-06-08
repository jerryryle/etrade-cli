package etradelib

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestListAccounts(t *testing.T) {
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

	httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testXml)),
		}
	})

	client := CreateETradeClient(GetEndpointUrls(true), httpClient)
	response, err := client.ListAccounts()
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.Accounts, 2)
	assert.Equal(t, "82314598", response.Accounts[0].AccountId)
}
