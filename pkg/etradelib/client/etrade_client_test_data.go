package client

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
	"time"
)

const quoteDetailAllTestXml = `
<?xml version="1.0" encoding="UTF-8"?>
<QuoteResponse>
   <QuoteData>
      <All>
         <adjustedFlag>true</adjustedFlag>
         <ask>1</ask>
         <askSize>2</askSize>
         <askTime>3</askTime>
         <bid>4</bid>
         <bidExchange>5</bidExchange>
         <bidSize>6</bidSize>
         <bidTime>7</bidTime>
         <changeClose>8</changeClose>
         <changeClosePercentage>9</changeClosePercentage>
         <companyName>10</companyName>
         <daysToExpiration>11</daysToExpiration>
         <dirLast>12</dirLast>
         <dividend>13</dividend>
         <eps>14</eps>
         <estEarnings>15</estEarnings>
         <exDividendDate>16</exDividendDate>
         <high>17</high>
         <high52>18</high52>
         <lastTrade>19</lastTrade>
         <low>20</low>
         <low52>21</low52>
         <open>22</open>
         <openInterest>23</openInterest>
         <optionStyle>24</optionStyle>
         <optionUnderlier>25</optionUnderlier>
         <optionUnderlierExchange>26</optionUnderlierExchange>
         <previousClose>27</previousClose>
         <previousDayVolume>28</previousDayVolume>
         <primaryExchange>29</primaryExchange>
         <symbolDescription>30</symbolDescription>
         <totalVolume>31</totalVolume>
         <upc>32</upc>
         <OptionDeliverableList>
           <OptionDeliverable>
             <rootSymbol>33</rootSymbol>
             <deliverableSymbol>34</deliverableSymbol>
             <deliverableTypeCode>35</deliverableTypeCode>
             <deliverableExchangeCode>36</deliverableExchangeCode>
             <deliverableStrikePercent>37</deliverableStrikePercent>
             <deliverableCILShares>38</deliverableCILShares>
             <deliverableWholeShares>39</deliverableWholeShares>
           </OptionDeliverable>
         </OptionDeliverableList>
         <cashDeliverable>40</cashDeliverable>
         <marketCap>41</marketCap>
         <sharesOutstanding>42</sharesOutstanding>
         <nextEarningDate>43</nextEarningDate>
         <beta>44</beta>
         <yield>45</yield>
         <declaredDividend>46</declaredDividend>
         <dividendPayableDate>47</dividendPayableDate>
         <pe>48</pe>
         <week52LowDate>49</week52LowDate>
         <week52HiDate>50</week52HiDate>
         <intrinsicValue>51</intrinsicValue>
         <timePremium>52</timePremium>
         <optionMultiplier>53</optionMultiplier>
         <contractSize>54</contractSize>
         <expirationDate>55</expirationDate>
         <EhQuote>
           <lastPrice>56</lastPrice>
           <change>57</change>
           <percentChange>58</percentChange>
           <bid>59</bid>
           <bidSize>60</bidSize>
           <ask>61</ask>
           <askSize>62</askSize>
           <volume>63</volume>
           <timeOfLastTrade>64</timeOfLastTrade>
           <timeZone>65</timeZone>
           <quoteStatus>66</quoteStatus>
         </EhQuote>
         <optionPreviousBidPrice>67</optionPreviousBidPrice>
         <optionPreviousAskPrice>68</optionPreviousAskPrice>
         <osiKey>69</osiKey>
         <timeOfLastTrade>70</timeOfLastTrade>
         <averageVolume>71</averageVolume>
      </All>
      <dateTime>72</dateTime>
      <dateTimeUTC>73</dateTimeUTC>
      <quoteStatus>74</quoteStatus>
      <ahFlag>true</ahFlag>
      <errorMessage>75</errorMessage>
      <Product>
         <symbol>76</symbol>
         <securityType>77</securityType>
      </Product>
      <timeZone>78</timeZone>
      <dstFlag>true</dstFlag>
      <hasMiniOptions>true</hasMiniOptions>
    </QuoteData>
    <Messages>
      <Message>
        <description>79</description>
        <code>80</code>
        <type>81</type>
      </Message>
    </Messages>
</QuoteResponse>`

var quoteDetailAllTestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All: responses.QuoteAllQuoteDetails{
				AdjustedFlag:            true,
				Ask:                     1,
				AskSize:                 2,
				AskTime:                 responses.ETradeTime{Time: time.Unix(3, 0).UTC()},
				Bid:                     4,
				BidExchange:             "5",
				BidSize:                 6,
				BidTime:                 responses.ETradeTime{Time: time.Unix(7, 0).UTC()},
				ChangeClose:             8,
				ChangeClosePercentage:   9,
				CompanyName:             "10",
				DaysToExpiration:        11,
				DirLast:                 "12",
				Dividend:                13,
				Eps:                     14,
				EstEarnings:             15,
				ExDividendDate:          responses.ETradeTime{Time: time.Unix(16, 0).UTC()},
				High:                    17,
				High52:                  18,
				LastTrade:               19,
				Low:                     20,
				Low52:                   21,
				Open:                    22,
				OpenInterest:            23,
				OptionStyle:             "24",
				OptionUnderlier:         "25",
				OptionUnderlierExchange: "26",
				PreviousClose:           27,
				PreviousDayVolume:       28,
				PrimaryExchange:         "29",
				SymbolDescription:       "30",
				TotalVolume:             31,
				Upc:                     32,
				OptionDeliverableList: []responses.QuoteOptionDeliverable{
					{
						RootSymbol:               "33",
						DeliverableSymbol:        "34",
						DeliverableTypeCode:      "35",
						DeliverableExchangeCode:  "36",
						DeliverableStrikePercent: 37,
						DeliverableCILShares:     38,
						DeliverableWholeShares:   39,
					},
				},
				CashDeliverable:     40,
				MarketCap:           41,
				SharesOutstanding:   42,
				NextEarningDate:     responses.ETradeTime{Time: time.Unix(43, 0).UTC()},
				Beta:                44,
				Yield:               45,
				DeclaredDividend:    46,
				DividendPayableDate: responses.ETradeTime{Time: time.Unix(47, 0).UTC()},
				Pe:                  48,
				Week52LowDate:       responses.ETradeTime{Time: time.Unix(49, 0).UTC()},
				Week52HiDate:        responses.ETradeTime{Time: time.Unix(50, 0).UTC()},
				IntrinsicValue:      51,
				TimePremium:         52,
				OptionMultiplier:    53,
				ContractSize:        54,
				ExpirationDate:      responses.ETradeTime{Time: time.Unix(55, 0).UTC()},
				EhQuote: responses.QuoteExtendedHourQuoteDetail{
					LastPrice:     56,
					Change:        57,
					PercentChange: 58,
					Bid:           59,
					BidSize:       60,
					Ask:           61,
					AskSize:       62,
					Volume:        63,
					TimeOfLastTrade: responses.ETradeTime{
						Time: time.Unix(64, 0).UTC(),
					},
					TimeZone:    "65",
					QuoteStatus: "66",
				},
				OptionPreviousBidPrice: 67,
				OptionPreviousAskPrice: 68,
				OsiKey:                 "69",
				TimeOfLastTrade:        responses.ETradeTime{Time: time.Unix(70, 0).UTC()},
				AverageVolume:          71,
			},
			DateTime:     responses.ETradeTime{Time: time.Unix(72, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(73, 0).UTC()},
			QuoteStatus:  "74",
			AhFlag:       true,
			ErrorMessage: "75",
			Fundamental:  responses.QuoteFundamentalQuoteDetails{},
			Intraday:     responses.QuoteIntradayQuoteDetails{},
			Option:       responses.QuoteOptionQuoteDetails{},
			Product: responses.Product{
				Symbol:       "76",
				SecurityType: "77",
			},
			Week52:         responses.QuoteWeek52QuoteDetails{},
			MutualFund:     responses.QuoteMutualFund{},
			TimeZone:       "78",
			DstFlag:        true,
			HasMiniOptions: true,
		},
	},
	Messages: []responses.QuoteMessage{
		{
			Description: "79",
			Code:        80,
			Type:        "81",
		},
	},
}

const quoteDetailFundamentalTestXml = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<QuoteResponse>
    <QuoteData>
        <dateTime>1</dateTime>
        <dateTimeUTC>2</dateTimeUTC>
        <quoteStatus>3</quoteStatus>
        <ahFlag>true</ahFlag>
        <errorMessage>4</errorMessage>
        <Fundamental>
            <companyName>5</companyName>
            <eps>6</eps>
            <estEarnings>7</estEarnings>
            <high52>8</high52>
            <lastTrade>9</lastTrade>
            <low52>10</low52>
            <symbolDescription>11</symbolDescription>
        </Fundamental>
        <Product>
            <symbol>12</symbol>
            <securityType>13</securityType>
        </Product>
        <timeZone>14</timeZone>
        <dstFlag>true</dstFlag>
        <hasMiniOptions>true</hasMiniOptions>
    </QuoteData>
    <Messages>
      <Message>
        <description>15</description>
        <code>16</code>
        <type>17</type>
      </Message>
    </Messages>
</QuoteResponse>`

var quoteDetailFundamentalTestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All:          responses.QuoteAllQuoteDetails{},
			DateTime:     responses.ETradeTime{Time: time.Unix(1, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(2, 0).UTC()},
			QuoteStatus:  "3",
			AhFlag:       true,
			ErrorMessage: "4",
			Fundamental: responses.QuoteFundamentalQuoteDetails{
				CompanyName:       "5",
				Eps:               6,
				EstEarnings:       7,
				High52:            8,
				LastTrade:         9,
				Low52:             10,
				SymbolDescription: "11",
			},
			Intraday: responses.QuoteIntradayQuoteDetails{},
			Option:   responses.QuoteOptionQuoteDetails{},
			Product: responses.Product{
				Symbol:       "12",
				SecurityType: "13",
			},
			Week52:         responses.QuoteWeek52QuoteDetails{},
			MutualFund:     responses.QuoteMutualFund{},
			TimeZone:       "14",
			DstFlag:        true,
			HasMiniOptions: true,
		},
	},
	Messages: []responses.QuoteMessage{
		{
			Description: "15",
			Code:        16,
			Type:        "17",
		},
	},
}

const quoteDetailIntradayTestXml = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<QuoteResponse>
    <QuoteData>
        <dateTime>1</dateTime>
        <dateTimeUTC>2</dateTimeUTC>
        <quoteStatus>3</quoteStatus>
        <ahFlag>true</ahFlag>
        <errorMessage>4</errorMessage>
        <Intraday>
            <ask>5</ask>
            <bid>6</bid>
            <changeClose>7</changeClose>
            <changeClosePercentage>8</changeClosePercentage>
            <companyName>9</companyName>
            <high>10</high>
            <lastTrade>11</lastTrade>
            <low>12</low>
            <totalVolume>13</totalVolume>
        </Intraday>
        <Product>
            <symbol>14</symbol>
            <securityType>15</securityType>
        </Product>
        <timeZone>16</timeZone>
        <dstFlag>true</dstFlag>
        <hasMiniOptions>true</hasMiniOptions>
    </QuoteData>
    <Messages>
      <Message>
        <description>17</description>
        <code>18</code>
        <type>19</type>
      </Message>
    </Messages>
</QuoteResponse>`

var quoteDetailIntradayTestResponse = responses.QuoteResponse{
	QuoteData: []responses.QuoteData{
		{
			All:          responses.QuoteAllQuoteDetails{},
			DateTime:     responses.ETradeTime{Time: time.Unix(1, 0).UTC()},
			DateTimeUTC:  responses.ETradeTime{Time: time.Unix(2, 0).UTC()},
			QuoteStatus:  "3",
			AhFlag:       true,
			ErrorMessage: "4",
			Fundamental:  responses.QuoteFundamentalQuoteDetails{},
			Intraday: responses.QuoteIntradayQuoteDetails{
				Ask:                   5,
				Bid:                   6,
				ChangeClose:           7,
				ChangeClosePercentage: 8,
				CompanyName:           "9",
				High:                  10,
				LastTrade:             11,
				Low:                   12,
				TotalVolume:           13,
			},
			Option: responses.QuoteOptionQuoteDetails{},
			Product: responses.Product{
				Symbol:       "14",
				SecurityType: "15",
			},
			Week52:         responses.QuoteWeek52QuoteDetails{},
			MutualFund:     responses.QuoteMutualFund{},
			TimeZone:       "16",
			DstFlag:        true,
			HasMiniOptions: true,
		},
	},
	Messages: []responses.QuoteMessage{
		{
			Description: "17",
			Code:        18,
			Type:        "19",
		},
	},
}

const quoteDetailOptionsTestXml = `
<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>
<QuoteResponse>
    <QuoteData>
        <dateTime>1</dateTime>
        <dateTimeUTC>2</dateTimeUTC>
        <quoteStatus>3</quoteStatus>
        <ahFlag>true</ahFlag>
        <errorMessage>4</errorMessage>
        <Option>
            <ask>5</ask>
            <askSize>6</askSize>
            <bid>7</bid>
            <bidSize>8</bidSize>
            <companyName>9</companyName>
            <daysToExpiration>10</daysToExpiration>
            <lastTrade>11</lastTrade>
            <openInterest>12</openInterest>
            <optionPreviousBidPrice>13</optionPreviousBidPrice>
            <optionPreviousAskPrice>14</optionPreviousAskPrice>
            <osiKey>15</osiKey>
            <intrinsicValue>16</intrinsicValue>
            <timePremium>17</timePremium>
            <optionMultiplier>18</optionMultiplier>
            <contractSize>19</contractSize>
            <symbolDescription>20</symbolDescription>
            <OptionGreeks>
              <rho>21</rho>
              <vega>22</vega>
              <theta>23</theta>
              <delta>24</delta>
              <gamma>25</gamma>
              <iv>26</iv>
              <currentValue>true</currentValue>
            </OptionGreeks>
        </Option>
        <Product>
            <symbol>27</symbol>
            <securityType>28</securityType>
        </Product>
        <timeZone>29</timeZone>
        <dstFlag>true</dstFlag>
        <hasMiniOptions>true</hasMiniOptions>
    </QuoteData>
    <Messages>
      <Message>
        <description>30</description>
        <code>31</code>
        <type>32</type>
      </Message>
    </Messages>
</QuoteResponse>`

var quoteDetailOptionsTestResponse = responses.QuoteResponse{
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
			Option: responses.QuoteOptionQuoteDetails{
				Ask:                    5,
				AskSize:                6,
				Bid:                    7,
				BidSize:                8,
				CompanyName:            "9",
				DaysToExpiration:       10,
				LastTrade:              11,
				OpenInterest:           12,
				OptionPreviousBidPrice: 13,
				OptionPreviousAskPrice: 14,
				OsiKey:                 "15",
				IntrinsicValue:         16,
				TimePremium:            17,
				OptionMultiplier:       18,
				ContractSize:           19,
				SymbolDescription:      "20",
				OptionGreeks: responses.QuoteOptionGreeks{
					Rho:          21,
					Vega:         22,
					Theta:        23,
					Delta:        24,
					Gamma:        25,
					Iv:           26,
					CurrentValue: true,
				},
			},
			Product: responses.Product{
				Symbol:       "27",
				SecurityType: "28",
			},
			Week52:         responses.QuoteWeek52QuoteDetails{},
			MutualFund:     responses.QuoteMutualFund{},
			TimeZone:       "29",
			DstFlag:        true,
			HasMiniOptions: true,
		},
	},
	Messages: []responses.QuoteMessage{
		{
			Description: "30",
			Code:        31,
			Type:        "32",
		},
	},
}

const quoteDetailWeek52TestXml = `
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
        <Week52>
            <companyName>7</companyName>
            <high52>8</high52>
            <lastTrade>9</lastTrade>
            <low52>10</low52>
            <perf12Months>11</perf12Months>
            <previousClose>12</previousClose>
            <symbolDescription>13</symbolDescription>
            <totalVolume>14</totalVolume>
        </Week52>
        <timeZone>15</timeZone>
        <dstFlag>true</dstFlag>
        <hasMiniOptions>true</hasMiniOptions>
    </QuoteData>
    <Messages>
      <Message>
        <description>16</description>
        <code>17</code>
        <type>18</type>
      </Message>
    </Messages>
</QuoteResponse>
`

var quoteDetailWeek52TestResponse = responses.QuoteResponse{
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
			Week52: responses.QuoteWeek52QuoteDetails{
				CompanyName:       "7",
				High52:            8,
				LastTrade:         9,
				Low52:             10,
				Perf12Months:      11,
				PreviousClose:     12,
				SymbolDescription: "13",
				TotalVolume:       14,
			},
			MutualFund:     responses.QuoteMutualFund{},
			TimeZone:       "15",
			DstFlag:        true,
			HasMiniOptions: true,
		},
	},
	Messages: []responses.QuoteMessage{
		{
			Description: "16",
			Code:        17,
			Type:        "18",
		},
	},
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
