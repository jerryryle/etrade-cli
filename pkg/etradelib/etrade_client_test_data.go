package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
	"time"
)

const quoteDetailAllTestXml = `
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
</QuoteResponse>`

var quoteDetailAllTestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All: responses.QuoteAllQuoteDetails{
				AdjustedFlag:            false,
				Ask:                     1175.79,
				AskSize:                 100,
				AskTime:                 responses.ETradeTime{Time: time.Unix(1529522220, 0).UTC()},
				Bid:                     1175.29,
				BidExchange:             "",
				BidSize:                 100,
				BidTime:                 responses.ETradeTime{Time: time.Unix(1529522220, 0).UTC()},
				ChangeClose:             7.68,
				ChangeClosePercentage:   0.66,
				CompanyName:             "ALPHABET INC CAP STK CL C",
				DaysToExpiration:        0,
				DirLast:                 "2",
				Dividend:                0.0,
				Eps:                     23.5639,
				EstEarnings:             43.981,
				ExDividendDate:          responses.ETradeTime{Time: time.Unix(1430163144, 0).UTC()},
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
				NextEarningDate:         responses.ETradeTime{Time: time.Unix(0, 0).UTC()},
				Beta:                    1.4,
				Yield:                   0.0,
				DeclaredDividend:        0.0,
				DividendPayableDate:     responses.ETradeTime{Time: time.Unix(1430767944, 0).UTC()},
				Pe:                      49.57,
				Week52LowDate:           responses.ETradeTime{Time: time.Unix(1499110344, 0).UTC()},
				Week52HiDate:            responses.ETradeTime{Time: time.Unix(1517257944, 0).UTC()},
				IntrinsicValue:          0.0,
				TimePremium:             0.0,
				OptionMultiplier:        0.0,
				ContractSize:            0.0,
				ExpirationDate:          responses.ETradeTime{Time: time.Unix(0, 0).UTC()},
				EhQuote:                 responses.QuoteExtendedHourQuoteDetail{},
				OptionPreviousBidPrice:  0,
				OptionPreviousAskPrice:  0,
				OsiKey:                  "",
				TimeOfLastTrade:         responses.ETradeTime{Time: time.Unix(1529522220, 0).UTC()},
				AverageVolume:           1451490,
			},
			DateTime:     responses.ETradeTime{Time: time.Unix(1529522220, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(1529522220, 0).UTC()},
			QuoteStatus:  "DELAYED",
			AhFlag:       false,
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
	Messages: nil,
}

const quoteDetailFundamentalTestXml = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<QuoteResponse>
    <QuoteData>
        <dateTime>16:00:00 EDT 06-20-2012</dateTime>
        <dateTimeUTC>1340222400</dateTimeUTC>
        <quoteStatus>REALTIME</quoteStatus>
        <ahFlag>false</ahFlag>
        <Fundamental>
            <companyName>GOOGLE INC CL A</companyName>
            <eps>32.99727</eps>
            <estEarnings>43.448</estEarnings>
            <high52>670.25</high52>
            <lastTrade>577.51</lastTrade>
            <low52>0.0</low52>
            <symbolDescription>GOOGLE INC CL A</symbolDescription>
        </Fundamental>
        <Product>
            <securityType>EQ</securityType>
            <symbol>GOOG</symbol>
        </Product>
    </QuoteData>
</QuoteResponse>`

var quoteDetailFundamentalTestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All:          responses.QuoteAllQuoteDetails{},
			DateTime:     responses.ETradeTime{Time: time.Unix(1340222400, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(1340222400, 0).UTC()},
			QuoteStatus:  "REALTIME",
			AhFlag:       false,
			ErrorMessage: "",
			Fundamental: responses.QuoteFundamentalQuoteDetails{
				CompanyName:       "GOOGLE INC CL A",
				Eps:               32.99727,
				EstEarnings:       43.448,
				High52:            670.25,
				LastTrade:         577.51,
				Low52:             0.0,
				SymbolDescription: "GOOGLE INC CL A",
			},
			Intraday: responses.QuoteIntradayQuoteDetails{},
			Option:   responses.QuoteOptionQuoteDetails{},
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
	Messages: nil,
}

const quoteDetailIntradayTestXml = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<QuoteResponse>
    <QuoteData>
        <dateTime>16:00:00 EDT 06-20-2012</dateTime>
        <dateTimeUTC>1340222400</dateTimeUTC>
        <quoteStatus>REALTIME</quoteStatus>
        <ahFlag>false</ahFlag>
        <Intraday>
            <ask>579.94</ask>
            <bid>574.04</bid>
            <changeClose>0.0</changeClose>
            <changeClosePercentage>0.0</changeClosePercentage>
            <companyName>GOOGLE INC CL A</companyName>
            <high>0.0</high>
            <lastTrade>577.51</lastTrade>
            <low>0.0</low>
            <totalVolume>0</totalVolume>
        </Intraday>
        <Product>
            <securityType>EQ</securityType>
            <symbol>GOOG</symbol>
        </Product>
    </QuoteData>
</QuoteResponse>`

var quoteDetailIntradayTestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All:          responses.QuoteAllQuoteDetails{},
			DateTime:     responses.ETradeTime{Time: time.Unix(1340222400, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(1340222400, 0).UTC()},
			QuoteStatus:  "REALTIME",
			AhFlag:       false,
			ErrorMessage: "",
			Fundamental:  responses.QuoteFundamentalQuoteDetails{},
			Intraday: responses.QuoteIntradayQuoteDetails{
				Ask:                   579.94,
				Bid:                   574.04,
				ChangeClose:           0.0,
				ChangeClosePercentage: 0.0,
				CompanyName:           "GOOGLE INC CL A",
				High:                  0.0,
				LastTrade:             577.51,
				Low:                   0.0,
				TotalVolume:           0,
			},
			Option: responses.QuoteOptionQuoteDetails{},
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
	Messages: nil,
}

const quoteDetailOptionsTestXml = `
<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>
<QuoteResponse>
    <QuoteData>
        <dateTime>16:00:00 EDT 06-20-2012</dateTime>
        <dateTimeUTC>1340222400</dateTimeUTC>
        <quoteStatus>REALTIME</quoteStatus>
        <ahFlag>false</ahFlag>
        <Option>
            <ask>579.65</ask>
            <askSize>100</askSize>
            <bid>574.04</bid>
            <bidSize>100</bidSize>
            <companyName>GOOGLE INC CL A</companyName>
            <daysToExpiration>0</daysToExpiration>
            <lastTrade>577.51</lastTrade>
            <openInterest>0</openInterest>
            <optionPreviousBidPrice>0</optionPreviousBidPrice>
            <optionPreviousAskPrice>0</optionPreviousAskPrice>
            <osiKey></osiKey>
            <intrinsicValue>0.0</intrinsicValue>
            <timePremium>0.0</timePremium>
            <optionMultiplier>0.0</optionMultiplier>
            <contractSize>0.0</contractSize>
            <symbolDescription></symbolDescription>
        </Option>
        <Product>
            <securityType>EQ</securityType>
            <symbol>GOOG</symbol>
        </Product>
    </QuoteData>
</QuoteResponse>`

var quoteDetailOptionsTestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All:          responses.QuoteAllQuoteDetails{},
			DateTime:     responses.ETradeTime{Time: time.Unix(1340222400, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(1340222400, 0).UTC()},
			QuoteStatus:  "REALTIME",
			AhFlag:       false,
			ErrorMessage: "",
			Fundamental:  responses.QuoteFundamentalQuoteDetails{},
			Intraday:     responses.QuoteIntradayQuoteDetails{},
			Option: responses.QuoteOptionQuoteDetails{
				Ask:                    579.65,
				AskSize:                100,
				Bid:                    574.04,
				BidSize:                100,
				CompanyName:            "GOOGLE INC CL A",
				DaysToExpiration:       0,
				LastTrade:              577.51,
				OpenInterest:           0,
				OptionPreviousBidPrice: 0,
				OptionPreviousAskPrice: 0,
				OsiKey:                 "",
				IntrinsicValue:         0.0,
				TimePremium:            0.0,
				OptionMultiplier:       0.0,
				ContractSize:           0.0,
				SymbolDescription:      "",
				OptionGreeks:           responses.QuoteOptionGreeks{},
			},
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
	Messages: nil,
}

const quoteDetailWeek52TestXml = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<QuoteResponse>
    <QuoteData>
        <dateTime>16:00:00 EDT 06-20-2012</dateTime>
        <dateTimeUTC>1340222400</dateTimeUTC>
        <quoteStatus>REALTIME</quoteStatus>
        <ahFlag>false</ahFlag>
        <Week52>
            <companyName>GOOGLE INC CL A</companyName>
            <high52>670.25</high52>
            <lastTrade>577.51</lastTrade>
            <low52>0.0</low52>
            <perf12Months>111.0</perf12Months>
            <previousClose>577.51</previousClose>
            <symbolDescription>GOOGLE INC CL A</symbolDescription>
            <totalVolume>0</totalVolume>
        </Week52>
        <Product>
            <securityType>EQ</securityType>
            <symbol>GOOG</symbol>
        </Product>
    </QuoteData>
</QuoteResponse>
`

var quoteDetailWeek52TestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All:          responses.QuoteAllQuoteDetails{},
			DateTime:     responses.ETradeTime{Time: time.Unix(1340222400, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(1340222400, 0).UTC()},
			QuoteStatus:  "REALTIME",
			AhFlag:       false,
			ErrorMessage: "",
			Fundamental:  responses.QuoteFundamentalQuoteDetails{},
			Intraday:     responses.QuoteIntradayQuoteDetails{},
			Option:       responses.QuoteOptionQuoteDetails{},
			Product: responses.Product{
				Symbol:       "GOOG",
				SecurityType: "EQ",
			},
			Week52: responses.QuoteWeek52QuoteDetails{
				CompanyName:       "GOOGLE INC CL A",
				High52:            670.25,
				LastTrade:         577.51,
				Low52:             0.0,
				Perf12Months:      111.0,
				PreviousClose:     577.51,
				SymbolDescription: "GOOGLE INC CL A",
				TotalVolume:       0,
			},
			MutualFund:     responses.QuoteMutualFund{},
			TimeZone:       "",
			DstFlag:        false,
			HasMiniOptions: false,
		},
	},
	Messages: nil,
}

const quoteDetailMutualFundTestXml = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<QuoteResponse>
    <QuoteData>
        <dateTime>1</dateTime>
        <dateTimeUTC>2</dateTimeUTC>
        <quoteStatus>3</quoteStatus>
        <ahFlag>true</ahFlag>
        <errorMessage>4</errorMessage>
        <Product>
            <symbol>5</symbol>
            <securityType>6</securityType>
        </Product>
        <MutualFund>
            <symbolDescription>7</symbolDescription>
            <cusip>8</cusip>
            <changeClose>9</changeClose>
            <previousClose>10</previousClose>
            <transactionFee>11</transactionFee>
            <earlyRedemptionFee>12</earlyRedemptionFee>
            <availability>13</availability>
            <initialInvestment>14</initialInvestment>
            <subsequentInvestment>15</subsequentInvestment>
            <fundFamily>16</fundFamily>
            <fundName>17</fundName>
            <changeClosePercentage>18</changeClosePercentage>
            <timeOfLastTrade>19</timeOfLastTrade>
            <netAssetValue>20</netAssetValue>
            <publicOfferPrice>21</publicOfferPrice>
            <netExpenseRatio>22</netExpenseRatio>
            <grossExpenseRatio>23</grossExpenseRatio>
            <orderCutoffTime>24</orderCutoffTime>
            <salesCharge>25</salesCharge>
            <initialIraInvestment>26</initialIraInvestment>
            <subsequentIraInvestment>27</subsequentIraInvestment>
            <NetAssets>
                <value>28</value>
                <asOfDate>29</asOfDate>
            </NetAssets>
            <fundInceptionDate>30</fundInceptionDate>
            <averageAnnualReturns>31</averageAnnualReturns>
            <sevenDayCurrentYield>32</sevenDayCurrentYield>
            <annualTotalReturn>33</annualTotalReturn>
            <weightedAverageMaturity>34</weightedAverageMaturity>
            <averageAnnualReturn1Yr>35</averageAnnualReturn1Yr>
            <averageAnnualReturn3Yr>36</averageAnnualReturn3Yr>
            <averageAnnualReturn5Yr>37</averageAnnualReturn5Yr>
            <averageAnnualReturn10Yr>38</averageAnnualReturn10Yr>
            <high52>39</high52>
            <low52>40</low52>
            <week52LowDate>41</week52LowDate>
            <week52HiDate>42</week52HiDate>
            <exchangeName>43</exchangeName>
            <sinceInception>44</sinceInception>
            <quarterlySinceInception>45</quarterlySinceInception>
            <lastTrade>46</lastTrade>
            <actual12B1Fee>47</actual12B1Fee>
            <performanceAsOfDate>48</performanceAsOfDate>
            <qtrlyPerformanceAsOfDate>49</qtrlyPerformanceAsOfDate>
            <Redemption>
                <minMonth>50</minMonth>
                <feePercent>51</feePercent>
                <isFrontEnd>52</isFrontEnd>
                <FrontEndValues>
                  <Values>
                    <low>53</low>
                    <high>54</high>
                    <percent>55</percent>
                  </Values>
                </FrontEndValues>
                <redemptionDurationType>56</redemptionDurationType>
                <isSales>57</isSales>
                <salesDurationType>58</salesDurationType>
                <SalesValues>
                  <Values>
                    <low>59</low>
                    <high>60</high>
                    <percent>61</percent>
                  </Values>
                </SalesValues>
            </Redemption>
            <morningStarCategory>62</morningStarCategory>
            <monthlyTrailingReturn1Y>63</monthlyTrailingReturn1Y>
            <monthlyTrailingReturn3Y>64</monthlyTrailingReturn3Y>
            <monthlyTrailingReturn5Y>65</monthlyTrailingReturn5Y>
            <monthlyTrailingReturn10Y>66</monthlyTrailingReturn10Y>
            <etradeEarlyRedemptionFee>67</etradeEarlyRedemptionFee>
            <maxSalesLoad>68</maxSalesLoad>
            <monthlyTrailingReturnYTD>69</monthlyTrailingReturnYTD>
            <monthlyTrailingReturn1M>70</monthlyTrailingReturn1M>
            <monthlyTrailingReturn3M>71</monthlyTrailingReturn3M>
            <monthlyTrailingReturn6M>72</monthlyTrailingReturn6M>
            <qtrlyTrailingReturnYTD>73</qtrlyTrailingReturnYTD>
            <qtrlyTrailingReturn1M>74</qtrlyTrailingReturn1M>
            <qtrlyTrailingReturn3M>75</qtrlyTrailingReturn3M>
            <qtrlyTrailingReturn6M>76</qtrlyTrailingReturn6M>
            <DeferredSalesCharges>
              <SaleChargeValues>
                <lowHigh>77</lowHigh>
                <percent>78</percent>
              </SaleChargeValues>
            </DeferredSalesCharges>
            <FrontEndSalesCharges>
              <SaleChargeValues>
                <lowHigh>79</lowHigh>
                <percent>80</percent>
              </SaleChargeValues>
            </FrontEndSalesCharges>
            <exchangeCode>81</exchangeCode>
        </MutualFund>
        <timeZone>82</timeZone>
        <dstFlag>true</dstFlag>
        <hasMiniOptions>true</hasMiniOptions>
    </QuoteData>
    <Messages>
      <Message>
        <description>83</description>
        <code>84</code>
        <type>85</type>
      </Message>
    </Messages>
</QuoteResponse>
`

var quoteDetailMutualFundTestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All:          responses.QuoteAllQuoteDetails{},
			DateTime:     responses.ETradeTime{Time: time.Unix(1, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(2, 0).UTC()},
			QuoteStatus:  "3",
			AhFlag:       true,
			ErrorMessage: "4",
			Fundamental:  responses.QuoteFundamentalQuoteDetails{},
			Intraday:     responses.QuoteIntradayQuoteDetails{},
			Option:       responses.QuoteOptionQuoteDetails{},
			Product: responses.Product{
				Symbol:       "5",
				SecurityType: "6",
			},
			Week52: responses.QuoteWeek52QuoteDetails{},
			MutualFund: responses.QuoteMutualFund{
				SymbolDescription:     "7",
				Cusip:                 "8",
				ChangeClose:           9,
				PreviousClose:         10,
				TransactionFee:        "11",
				EarlyRedemptionFee:    "12",
				Availability:          "13",
				InitialInvestment:     14,
				SubsequentInvestment:  15,
				FundFamily:            "16",
				FundName:              "17",
				ChangeClosePercentage: 18,
				TimeOfLastTrade: responses.ETradeTime{
					Time: time.Unix(19, 0).UTC(),
				},
				NetAssetValue:           20,
				PublicOfferPrice:        21,
				NetExpenseRatio:         22,
				GrossExpenseRatio:       23,
				OrderCutoffTime:         24,
				SalesCharge:             "25",
				InitialIraInvestment:    26,
				SubsequentIraInvestment: 27,
				NetAssets: responses.QuoteNetAsset{
					Value: 28,
					AsOfDate: responses.ETradeTime{
						Time: time.Unix(29, 0).UTC(),
					},
				},
				FundInceptionDate: responses.ETradeTime{
					Time: time.Unix(30, 0).UTC(),
				},
				AverageAnnualReturns:    31,
				SevenDayCurrentYield:    32,
				AnnualTotalReturn:       33,
				WeightedAverageMaturity: 34,
				AverageAnnualReturn1Yr:  35,
				AverageAnnualReturn3Yr:  36,
				AverageAnnualReturn5Yr:  37,
				AverageAnnualReturn10Yr: 38,
				High52:                  39,
				Low52:                   40,
				Week52LowDate: responses.ETradeTime{
					Time: time.Unix(41, 0).UTC(),
				},
				Week52HiDate: responses.ETradeTime{
					Time: time.Unix(42, 0).UTC(),
				},
				ExchangeName:            "43",
				SinceInception:          44,
				QuarterlySinceInception: 45,
				LastTrade:               46,
				Actual12B1Fee:           47,
				PerformanceAsOfDate: responses.ETradeTime{
					Time: time.Unix(48, 0).UTC(),
				},
				QtrlyPerformanceAsOfDate: responses.ETradeTime{
					Time: time.Unix(49, 0).UTC(),
				},
				Redemption: responses.QuoteRedemption{
					MinMonth:   "50",
					FeePercent: "51",
					IsFrontEnd: "52",
					FrontEndValues: []responses.QuoteValues{
						{
							Low:     "53",
							High:    "54",
							Percent: "55",
						},
					},
					RedemptionDurationType: "56",
					IsSales:                "57",
					SalesDurationType:      "58",
					SalesValues: []responses.QuoteValues{
						{
							Low:     "59",
							High:    "60",
							Percent: "61",
						},
					},
				},
				MorningStarCategory:      "62",
				MonthlyTrailingReturn1Y:  63,
				MonthlyTrailingReturn3Y:  64,
				MonthlyTrailingReturn5Y:  65,
				MonthlyTrailingReturn10Y: 66,
				EtradeEarlyRedemptionFee: "67",
				MaxSalesLoad:             68,
				MonthlyTrailingReturnYTD: 69,
				MonthlyTrailingReturn1M:  70,
				MonthlyTrailingReturn3M:  71,
				MonthlyTrailingReturn6M:  72,
				QtrlyTrailingReturnYTD:   73,
				QtrlyTrailingReturn1M:    74,
				QtrlyTrailingReturn3M:    75,
				QtrlyTrailingReturn6M:    76,
				DeferredSalesCharges: []responses.QuoteSaleChargeValues{{
					LowHigh: "77",
					Percent: "78",
				}},
				FrontEndSalesCharges: []responses.QuoteSaleChargeValues{{
					LowHigh: "79",
					Percent: "80",
				}},
				ExchangeCode: "81",
			},
			TimeZone:       "82",
			DstFlag:        true,
			HasMiniOptions: true,
		},
	},
	Messages: []responses.QuoteMessage{
		{
			Description: "83",
			Code:        84,
			Type:        "85",
		},
	},
}
