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

	client := CreateETradeClient(GetEndpointUrls(true), httpClient, CreateNullLogger())
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

	client := CreateETradeClient(GetEndpointUrls(true), httpClient, CreateNullLogger())
	response, err := client.ListAlerts()
	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, response)
}

func TestETradeClient_GetQuotes(t *testing.T) {
	testXml := `
<?xml version="1.0" encoding="UTF-8"?>
<QuoteResponse>
   <QuoteData>
      <dateTime>15:17:00 EDT 06-20-2018</dateTime>
      <dateTimeUTC>1529522220</dateTimeUTC>
      <quoteStatus>DELAYED</quoteStatus>
      <ahFlag>false</ahFlag>
      <hasMiniOptions>false</hasMiniOptions>
      <All>
         <adjustedFlag>false</adjustedFlag>
         <ask>1175.79</ask>
         <askSize>100</askSize>
         <askTime>15:17:00 EDT 06-20-2018</askTime>
         <bid>1175.29</bid>
         <bidExchange />
         <bidSize>100</bidSize>
         <bidTime>15:17:00 EDT 06-20-2018</bidTime>
         <changeClose>7.68</changeClose>
         <changeClosePercentage>0.66</changeClosePercentage>
         <companyName>ALPHABET INC CAP STK CL C</companyName>
         <daysToExpiration>0</daysToExpiration>
         <dirLast>2</dirLast>
         <dividend>0.0</dividend>
         <eps>23.5639</eps>
         <estEarnings>43.981</estEarnings>
         <exDividendDate>1430163144</exDividendDate>
         <high>1186.2856</high>
         <high52>1186.89</high52>
         <lastTrade>1175.74</lastTrade>
         <low>1171.76</low>
         <low52>894.79</low52>
         <open>1175.31</open>
         <openInterest>0</openInterest>
         <optionStyle />
         <optionUnderlier />
         <previousClose>1168.06</previousClose>
         <previousDayVolume>1620909</previousDayVolume>
         <primaryExchange>NSDQ</primaryExchange>
         <symbolDescription>ALPHABET INC CAP STK CL C</symbolDescription>
         <totalVolume>1167544</totalVolume>
         <upc>0</upc>
         <cashDeliverable>0</cashDeliverable>
         <marketCap>410276824480.00</marketCap>
         <sharesOutstanding>348952000</sharesOutstanding>
         <nextEarningDate />
         <beta>1.4</beta>
         <yield>0.0</yield>
         <declaredDividend>0.0</declaredDividend>
         <dividendPayableDate>1430767944</dividendPayableDate>
         <pe>49.57</pe>
         <week52LowDate>1499110344</week52LowDate>
         <week52HiDate>1517257944</week52HiDate>
         <intrinsicValue>0.0</intrinsicValue>
         <timePremium>0.0</timePremium>
         <optionMultiplier>0.0</optionMultiplier>
         <contractSize>0.0</contractSize>
         <expirationDate>0</expirationDate>
         <timeOfLastTrade>1529522220</timeOfLastTrade>
         <averageVolume>1451490</averageVolume>
      </All>
      <Product>
         <securityType>EQ</securityType>
         <symbol>GOOG</symbol>
      </Product>
   </QuoteData>
</QuoteResponse>
`
	expectedResponse := responses.QuoteResponse{
		QuoteData: []responses.QuoteData{
			{
				All: responses.QuoteAllQuoteDetails{
					AdjustedFlag:            false,
					Ask:                     1175.79,
					AskSize:                 100,
					AskTime:                 "15:17:00 EDT 06-20-2018",
					Bid:                     1175.29,
					BidExchange:             "",
					BidSize:                 100,
					BidTime:                 "15:17:00 EDT 06-20-2018",
					ChangeClose:             7.68,
					ChangeClosePercentage:   0.66,
					CompanyName:             "ALPHABET INC CAP STK CL C",
					DaysToExpiration:        0,
					DirLast:                 "2",
					Dividend:                0.0,
					Eps:                     23.5639,
					EstEarnings:             43.981,
					ExDividendDate:          responses.ETradeTime{Time: time.Unix(1430163144, 0)},
					High:                    1186.2856,
					High52:                  1186.89,
					LastTrade:               1175.74,
					Low:                     1171.76,
					Low52:                   894.79,
					Open:                    1175.31,
					OpenInterest:            0,
					OptionStyle:             "",
					OptionUnderlier:         "",
					OptionUnderlierExchange: "",
					PreviousClose:           1168.06,
					PreviousDayVolume:       1620909,
					PrimaryExchange:         "NSDQ",
					SymbolDescription:       "ALPHABET INC CAP STK CL C",
					TotalVolume:             1167544,
					Upc:                     0,
					OptionDeliverableList:   nil,
					CashDeliverable:         0,
					MarketCap:               410276824480.00,
					SharesOutstanding:       348952000,
					NextEarningDate:         "",
					Beta:                    1.4,
					Yield:                   0.0,
					DeclaredDividend:        0.0,
					DividendPayableDate:     responses.ETradeTime{Time: time.Unix(1430767944, 0)},
					Pe:                      49.57,
					Week52LowDate:           responses.ETradeTime{Time: time.Unix(1499110344, 0)},
					Week52HiDate:            responses.ETradeTime{Time: time.Unix(1517257944, 0)},
					IntrinsicValue:          0.0,
					TimePremium:             0.0,
					OptionMultiplier:        0.0,
					ContractSize:            0.0,
					ExpirationDate:          responses.ETradeTime{Time: time.Unix(0, 0)},
					EhQuote:                 responses.QuoteExtendedHourQuoteDetail{},
					OptionPreviousBidPrice:  0,
					OptionPreviousAskPrice:  0,
					OsiKey:                  "",
					TimeOfLastTrade:         responses.ETradeTime{Time: time.Unix(1529522220, 0)},
					AverageVolume:           1451490,
				},
				DateTime:     "15:17:00 EDT 06-20-2018",
				DateTimeUTC:  responses.ETradeTime{Time: time.Unix(1529522220, 0)},
				QuoteStatus:  "DELAYED",
				AhFlag:       "false",
				ErrorMessage: "",
				Fundamental:  responses.QuoteFundamentalQuoteDetails{},
				Intraday:     responses.QuoteIntradayQuoteDetails{},
				Option:       responses.QuoteOptionQuoteDetails{},
				Product: responses.Product{
					Symbol:       "GOOG",
					SecurityType: "EQ",
				},
				Week52:         responses.QuoteWeek52QuoteDetails{},
				MutualFund:     responses.QuoteMutualFund{},
				TimeZone:       "",
				DstFlag:        false,
				HasMiniOptions: false,
			},
		},
		Messages: responses.QuoteMessages{},
	}

	httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(testXml)),
		}
	})

	client := CreateETradeClient(GetEndpointUrls(true), httpClient, CreateNullLogger())
	response, err := client.GetQuotes([]string{"GOOG"}, QuoteDetailAll)
	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, response)
}
