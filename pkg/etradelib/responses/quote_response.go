package responses

type QuoteResponse struct {
	QuoteData []QuoteData   `xml:"QuoteData"`
	Messages  QuoteMessages `xml:"Messages"`
}

type QuoteAllQuoteDetails struct {
	AdjustedFlag            bool                         `xml:"adjustedFlag"`
	Ask                     float64                      `xml:"ask"`
	AskSize                 int64                        `xml:"askSize"`
	AskTime                 string                       `xml:"askTime"`
	Bid                     float64                      `xml:"bid"`
	BidExchange             string                       `xml:"bidExchange"`
	BidSize                 int64                        `xml:"bidSize"`
	BidTime                 string                       `xml:"bidTime"`
	ChangeClose             float64                      `xml:"changeClose"`
	ChangeClosePercentage   float64                      `xml:"changeClosePercentage"`
	CompanyName             string                       `xml:"companyName"`
	DaysToExpiration        int64                        `xml:"daysToExpiration"`
	DirLast                 string                       `xml:"dirLast"`
	Dividend                float64                      `xml:"dividend"`
	Eps                     float64                      `xml:"eps"`
	EstEarnings             float64                      `xml:"estEarnings"`
	ExDividendDate          ETradeTime                   `xml:"exDividendDate"`
	High                    float64                      `xml:"high"`
	High52                  float64                      `xml:"high52"`
	LastTrade               float64                      `xml:"lastTrade"`
	Low                     float64                      `xml:"low"`
	Low52                   float64                      `xml:"low52"`
	Open                    float64                      `xml:"open"`
	OpenInterest            int64                        `xml:"openInterest"`
	OptionStyle             string                       `xml:"optionStyle"`
	OptionUnderlier         string                       `xml:"optionUnderlier"`
	OptionUnderlierExchange string                       `xml:"optionUnderlierExchange"`
	PreviousClose           float64                      `xml:"previousClose"`
	PreviousDayVolume       int64                        `xml:"previousDayVolume"`
	PrimaryExchange         string                       `xml:"primaryExchange"`
	SymbolDescription       string                       `xml:"symbolDescription"`
	TotalVolume             int64                        `xml:"totalVolume"`
	Upc                     int64                        `xml:"upc"`
	OptionDeliverableList   []QuoteOptionDeliverable     `xml:"optionDeliverableList"`
	CashDeliverable         float64                      `xml:"cashDeliverable"`
	MarketCap               float64                      `xml:"marketCap"`
	SharesOutstanding       float64                      `xml:"sharesOutstanding"`
	NextEarningDate         string                       `xml:"nextEarningDate"`
	Beta                    float64                      `xml:"beta"`
	Yield                   float64                      `xml:"yield"`
	DeclaredDividend        float64                      `xml:"declaredDividend"`
	DividendPayableDate     ETradeTime                   `xml:"dividendPayableDate"`
	Pe                      float64                      `xml:"pe"`
	Week52LowDate           ETradeTime                   `xml:"week52LowDate"`
	Week52HiDate            ETradeTime                   `xml:"week52HiDate"`
	IntrinsicValue          float64                      `xml:"intrinsicValue"`
	TimePremium             float64                      `xml:"timePremium"`
	OptionMultiplier        float64                      `xml:"optionMultiplier"`
	ContractSize            float64                      `xml:"contractSize"`
	ExpirationDate          ETradeTime                   `xml:"expirationDate"`
	EhQuote                 QuoteExtendedHourQuoteDetail `xml:"ehQuote"`
	OptionPreviousBidPrice  float64                      `xml:"optionPreviousBidPrice"`
	OptionPreviousAskPrice  float64                      `xml:"optionPreviousAskPrice"`
	OsiKey                  string                       `xml:"osiKey"`
	TimeOfLastTrade         ETradeTime                   `xml:"timeOfLastTrade"`
	AverageVolume           int64                        `xml:"averageVolume"`
}

type QuoteExtendedHourQuoteDetail struct {
	LastPrice       float64    `xml:"lastPrice"`
	Change          float64    `xml:"change"`
	PercentChange   float64    `xml:"percentChange"`
	Bid             float64    `xml:"bid"`
	BidSize         int64      `xml:"bidSize"`
	Ask             float64    `xml:"ask"`
	AskSize         int64      `xml:"askSize"`
	Volume          int64      `xml:"volume"`
	TimeOfLastTrade ETradeTime `xml:"timeOfLastTrade"`
	TimeZone        string     `xml:"timeZone"`
	QuoteStatus     string     `xml:"quoteStatus"`
}

type QuoteFundamentalQuoteDetails struct {
	CompanyName       string  `xml:"companyName"`
	Eps               float64 `xml:"eps"`
	EstEarnings       float64 `xml:"estEarnings"`
	High52            float64 `xml:"high52"`
	LastTrade         float64 `xml:"lastTrade"`
	Low52             float64 `xml:"low52"`
	SymbolDescription string  `xml:"symbolDescription"`
}

type QuoteIntradayQuoteDetails struct {
	Ask                   float64 `xml:"ask"`
	Bid                   float64 `xml:"bid"`
	ChangeClose           float64 `xml:"changeClose"`
	ChangeClosePercentage float64 `xml:"changeClosePercentage"`
	CompanyName           string  `xml:"companyName"`
	High                  float64 `xml:"high"`
	LastTrade             float64 `xml:"lastTrade"`
	Low                   float64 `xml:"low"`
	TotalVolume           int64   `xml:"totalVolume"`
}

type QuoteMessage struct {
	Description string `xml:"description"`
	Code        int32  `xml:"code"`
	Type        string `xml:"type"`
}

type QuoteMessages struct {
	Message []QuoteMessage `xml:"message"`
}

type QuoteMutualFund struct {
	SymbolDescription        string                  `xml:"symbolDescription"`
	Cusip                    string                  `xml:"cusip"`
	ChangeClose              float64                 `xml:"changeClose"`
	PreviousClose            float64                 `xml:"previousClose"`
	TransactionFee           string                  `xml:"transactionFee"`
	EarlyRedemptionFee       string                  `xml:"earlyRedemptionFee"`
	Availability             string                  `xml:"availability"`
	InitialInvestment        float64                 `xml:"initialInvestment"`
	SubsequentInvestment     float64                 `xml:"subsequentInvestment"`
	FundFamily               string                  `xml:"fundFamily"`
	FundName                 string                  `xml:"fundName"`
	ChangeClosePercentage    float64                 `xml:"changeClosePercentage"`
	TimeOfLastTrade          ETradeTime              `xml:"timeOfLastTrade"`
	NetAssetValue            float64                 `xml:"netAssetValue"`
	PublicOfferPrice         float64                 `xml:"publicOfferPrice"`
	NetExpenseRatio          float64                 `xml:"netExpenseRatio"`
	GrossExpenseRatio        float64                 `xml:"grossExpenseRatio"`
	OrderCutoffTime          ETradeTime              `xml:"orderCutoffTime"`
	SalesCharge              string                  `xml:"salesCharge"`
	InitialIraInvestment     float64                 `xml:"initialIraInvestment"`
	SubsequentIraInvestment  float64                 `xml:"subsequentIraInvestment"`
	NetAssets                QuoteNetAsset           `xml:"netAssets"`
	FundInceptionDate        ETradeTime              `xml:"fundInceptionDate"`
	AverageAnnualReturns     float64                 `xml:"averageAnnualReturns"`
	SevenDayCurrentYield     float64                 `xml:"sevenDayCurrentYield"`
	AnnualTotalReturn        float64                 `xml:"annualTotalReturn"`
	WeightedAverageMaturity  float64                 `xml:"weightedAverageMaturity"`
	AverageAnnualReturn1Yr   float64                 `xml:"averageAnnualReturn1Yr"`
	AverageAnnualReturn3Yr   float64                 `xml:"averageAnnualReturn3Yr"`
	AverageAnnualReturn5Yr   float64                 `xml:"averageAnnualReturn5Yr"`
	AverageAnnualReturn10Yr  float64                 `xml:"averageAnnualReturn10Yr"`
	High52                   float64                 `xml:"high52"`
	Low52                    float64                 `xml:"low52"`
	Week52LowDate            ETradeTime              `xml:"week52LowDate"`
	Week52HiDate             ETradeTime              `xml:"week52HiDate"`
	ExchangeName             string                  `xml:"exchangeName"`
	SinceInception           float64                 `xml:"sinceInception"`
	QuarterlySinceInception  float64                 `xml:"quarterlySinceInception"`
	LastTrade                float64                 `xml:"lastTrade"`
	Actual12B1Fee            float64                 `xml:"actual12B1Fee"`
	PerformanceAsOfDate      string                  `xml:"performanceAsOfDate"`
	QtrlyPerformanceAsOfDate string                  `xml:"qtrlyPerformanceAsOfDate"`
	Redemption               QuoteRedemption         `xml:"redemption"`
	MorningStarCategory      string                  `xml:"morningStarCategory"`
	MonthlyTrailingReturn1Y  float64                 `xml:"monthlyTrailingReturn1Y"`
	MonthlyTrailingReturn3Y  float64                 `xml:"monthlyTrailingReturn3Y"`
	MonthlyTrailingReturn5Y  float64                 `xml:"monthlyTrailingReturn5Y"`
	MonthlyTrailingReturn10Y float64                 `xml:"monthlyTrailingReturn10Y"`
	EtradeEarlyRedemptionFee string                  `xml:"etradeEarlyRedemptionFee"`
	MaxSalesLoad             float64                 `xml:"maxSalesLoad"`
	MonthlyTrailingReturnYTD float64                 `xml:"monthlyTrailingReturnYTD"`
	MonthlyTrailingReturn1M  float64                 `xml:"monthlyTrailingReturn1M"`
	MonthlyTrailingReturn3M  float64                 `xml:"monthlyTrailingReturn3M"`
	MonthlyTrailingReturn6M  float64                 `xml:"monthlyTrailingReturn6M"`
	QtrlyTrailingReturnYTD   float64                 `xml:"qtrlyTrailingReturnYTD"`
	QtrlyTrailingReturn1M    float64                 `xml:"qtrlyTrailingReturn1M"`
	QtrlyTrailingReturn3M    float64                 `xml:"qtrlyTrailingReturn3M"`
	QtrlyTrailingReturn6M    float64                 `xml:"qtrlyTrailingReturn6M"`
	DeferredSalesCharges     []QuoteSaleChargeValues `xml:"deferredSalesCharges"`
	FrontEndSalesCharges     []QuoteSaleChargeValues `xml:"frontEndSalesCharges"`
	ExchangeCode             string                  `xml:"exchangeCode"`
}

type QuoteNetAsset struct {
	Value    float64    `xml:"value"`
	AsOfDate ETradeTime `xml:"asOfDate"`
}

type QuoteOptionDeliverable struct {
	RootSymbol               string  `xml:"rootSymbol"`
	DeliverableSymbol        string  `xml:"deliverableSymbol"`
	DeliverableTypeCode      string  `xml:"deliverableTypeCode"`
	DeliverableExchangeCode  string  `xml:"deliverableExchangeCode"`
	DeliverableStrikePercent float64 `xml:"deliverableStrikePercent"`
	DeliverableCILShares     float64 `xml:"deliverableCILShares"`
	DeliverableWholeShares   int32   `xml:"deliverableWholeShares"`
}

type QuoteOptionGreeks struct {
	Rho          float64 `xml:"rho"`
	Vega         float64 `xml:"vega"`
	Theta        float64 `xml:"theta"`
	Delta        float64 `xml:"delta"`
	Gamma        float64 `xml:"gamma"`
	Iv           float64 `xml:"iv"`
	CurrentValue bool    `xml:"currentValue"`
}

type QuoteOptionQuoteDetails struct {
	Ask                    float64           `xml:"ask"`
	AskSize                int64             `xml:"askSize"`
	Bid                    float64           `xml:"bid"`
	BidSize                int64             `xml:"bidSize"`
	CompanyName            string            `xml:"companyName"`
	DaysToExpiration       int64             `xml:"daysToExpiration"`
	LastTrade              float64           `xml:"lastTrade"`
	OpenInterest           int64             `xml:"openInterest"`
	OptionPreviousBidPrice float64           `xml:"optionPreviousBidPrice"`
	OptionPreviousAskPrice float64           `xml:"optionPreviousAskPrice"`
	OsiKey                 string            `xml:"osiKey"`
	IntrinsicValue         float64           `xml:"intrinsicValue"`
	TimePremium            float64           `xml:"timePremium"`
	OptionMultiplier       float64           `xml:"optionMultiplier"`
	ContractSize           float64           `xml:"contractSize"`
	SymbolDescription      string            `xml:"symbolDescription"`
	OptionGreeks           QuoteOptionGreeks `xml:"optionGreeks"`
}

type QuoteData struct {
	All            QuoteAllQuoteDetails         `xml:"All"`
	DateTime       string                       `xml:"dateTime"`
	DateTimeUTC    ETradeTime                   `xml:"dateTimeUTC"`
	QuoteStatus    string                       `xml:"quoteStatus"`
	AhFlag         string                       `xml:"ahFlag"`
	ErrorMessage   string                       `xml:"errorMessage"`
	Fundamental    QuoteFundamentalQuoteDetails `xml:"Fundamental"`
	Intraday       QuoteIntradayQuoteDetails    `xml:"Intraday"`
	Option         QuoteOptionQuoteDetails      `xml:"Option"`
	Product        Product                      `xml:"Product"`
	Week52         QuoteWeek52QuoteDetails      `xml:"Week52"`
	MutualFund     QuoteMutualFund              `xml:"MutualFund"`
	TimeZone       string                       `xml:"timeZone"`
	DstFlag        bool                         `xml:"dstFlag"`
	HasMiniOptions bool                         `xml:"hasMiniOptions"`
}

type QuoteRedemption struct {
	MinMonth               string        `xml:"minMonth"`
	FeePercent             string        `xml:"feePercent"`
	IsFrontEnd             string        `xml:"isFrontEnd"`
	FrontEndValues         []QuoteValues `xml:"frontEndValues"`
	RedemptionDurationType string        `xml:"redemptionDurationType"`
	IsSales                string        `xml:"isSales"`
	SalesDurationType      string        `xml:"salesDurationType"`
	SalesValues            []QuoteValues `xml:"salesValues"`
}

type QuoteSaleChargeValues struct {
	LowHigh string `xml:"lowhigh"`
	Percent string `xml:"percent"`
}

type QuoteValues struct {
	Low     string `xml:"low"`
	High    string `xml:"high"`
	Percent string `xml:"percent"`
}

type QuoteWeek52QuoteDetails struct {
	CompanyName       string  `xml:"companyName"`
	High52            float64 `xml:"high52"`
	LastTrade         float64 `xml:"lastTrade"`
	Low52             float64 `xml:"low52"`
	Perf12Months      float64 `xml:"perf12Months"`
	PreviousClose     float64 `xml:"previousClose"`
	SymbolDescription string  `xml:"symbolDescription"`
	TotalVolume       int64   `xml:"totalVolume"`
}
