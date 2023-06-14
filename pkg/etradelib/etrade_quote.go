package etradelib

import (
	"time"
)

type ETradeQuoteInfo struct {
	DateTime       time.Time
	QuoteStatus    string
	AhFlag         bool
	TimeZone       string
	DstFlag        bool
	HasMiniOptions bool
}

type ETradeQuoteAllInfo struct {
	ETradeQuoteInfo
	AdjustedFlag                bool
	Ask                         float64
	AskSize                     int64
	AskTime                     time.Time
	Bid                         float64
	BidExchange                 string
	BidSize                     int64
	BidTime                     time.Time
	ChangeClose                 float64
	ChangeClosePercentage       float64
	CompanyName                 string
	DaysToExpiration            int64
	DirLast                     string
	Dividend                    float64
	Eps                         float64
	EstEarnings                 float64
	ExDividendDate              time.Time
	High                        float64
	High52                      float64
	LastTrade                   float64
	Low                         float64
	Low52                       float64
	Open                        float64
	OpenInterest                int64
	OptionStyle                 string
	OptionUnderlier             string
	OptionUnderlierExchange     string
	PreviousClose               float64
	PreviousDayVolume           int64
	PrimaryExchange             string
	SymbolDescription           string
	TotalVolume                 int64
	Upc                         int64
	OptionDeliverableList       []ETradeQuoteOptionDeliverable
	CashDeliverable             float64
	MarketCap                   float64
	SharesOutstanding           float64
	NextEarningDate             time.Time
	Beta                        float64
	Yield                       float64
	DeclaredDividend            float64
	DividendPayableDate         time.Time
	Pe                          float64
	Week52LowDate               time.Time
	Week52HiDate                time.Time
	IntrinsicValue              float64
	TimePremium                 float64
	OptionMultiplier            float64
	ContractSize                float64
	ExpirationDate              time.Time
	ExtendedHourLastPrice       float64
	ExtendedHourChange          float64
	ExtendedHourPercentChange   float64
	ExtendedHourBid             float64
	ExtendedHourBidSize         int64
	ExtendedHourAsk             float64
	ExtendedHourAskSize         int64
	ExtendedHourVolume          int64
	ExtendedHourTimeOfLastTrade time.Time
	ExtendedHourTimeZone        string
	ExtendedHourQuoteStatus     string
	OptionPreviousBidPrice      float64
	OptionPreviousAskPrice      float64
	OsiKey                      string
	TimeOfLastTrade             time.Time
	AverageVolume               int64
}

type ETradeQuoteFundamentalInfo struct {
	ETradeQuoteInfo
	CompanyName       string
	Eps               float64
	EstEarnings       float64
	High52            float64
	LastTrade         float64
	Low52             float64
	SymbolDescription string
}

type ETradeQuoteIntradayInfo struct {
	ETradeQuoteInfo
	Ask                   float64
	Bid                   float64
	ChangeClose           float64
	ChangeClosePercentage float64
	CompanyName           string
	High                  float64
	LastTrade             float64
	Low                   float64
	TotalVolume           int64
}

type ETradeQuoteOptionsInfo struct {
	ETradeQuoteInfo
	Ask                      float64
	AskSize                  int64
	Bid                      float64
	BidSize                  int64
	CompanyName              string
	DaysToExpiration         int64
	LastTrade                float64
	OpenInterest             int64
	OptionPreviousBidPrice   float64
	OptionPreviousAskPrice   float64
	OsiKey                   string
	IntrinsicValue           float64
	TimePremium              float64
	OptionMultiplier         float64
	ContractSize             float64
	SymbolDescription        string
	OptionGreeksRho          float64
	OptionGreeksVega         float64
	OptionGreeksTheta        float64
	OptionGreeksDelta        float64
	OptionGreeksGamma        float64
	OptionGreeksIv           float64
	OptionGreeksCurrentValue bool
}

type ETradeQuoteWeek52Info struct {
	ETradeQuoteInfo
	CompanyName       string
	High52            float64
	LastTrade         float64
	Low52             float64
	Perf12Months      float64
	PreviousClose     float64
	SymbolDescription string
	TotalVolume       int64
}

type ETradeQuoteMutualFundInfo struct {
	ETradeQuoteInfo
	SymbolDescription                string
	Cusip                            string
	ChangeClose                      float64
	PreviousClose                    float64
	TransactionFee                   string
	EarlyRedemptionFee               string
	Availability                     string
	InitialInvestment                float64
	SubsequentInvestment             float64
	FundFamily                       string
	FundName                         string
	ChangeClosePercentage            float64
	TimeOfLastTrade                  time.Time
	NetAssetValue                    float64
	PublicOfferPrice                 float64
	NetExpenseRatio                  float64
	GrossExpenseRatio                float64
	OrderCutoffTime                  int64
	SalesCharge                      string
	InitialIraInvestment             float64
	SubsequentIraInvestment          float64
	NetAssetTotalValue               float64
	NetAssetTotalAsOfDate            time.Time
	FundInceptionDate                time.Time
	AverageAnnualReturns             float64
	SevenDayCurrentYield             float64
	AnnualTotalReturn                float64
	WeightedAverageMaturity          float64
	AverageAnnualReturn1Yr           float64
	AverageAnnualReturn3Yr           float64
	AverageAnnualReturn5Yr           float64
	AverageAnnualReturn10Yr          float64
	High52                           float64
	Low52                            float64
	Week52LowDate                    time.Time
	Week52HiDate                     time.Time
	ExchangeName                     string
	SinceInception                   float64
	QuarterlySinceInception          float64
	LastTrade                        float64
	Actual12B1Fee                    float64
	PerformanceAsOfDate              time.Time
	QtrlyPerformanceAsOfDate         time.Time
	RedemptionMinMonth               string
	RedemptionFeePercent             string
	RedemptionIsFrontEnd             string
	RedemptionFrontEndValues         []ETradeQuoteValues
	RedemptionRedemptionDurationType string
	RedemptionIsSales                string
	RedemptionSalesDurationType      string
	RedemptionSalesValues            []ETradeQuoteValues
	MorningStarCategory              string
	MonthlyTrailingReturn1Y          float64
	MonthlyTrailingReturn3Y          float64
	MonthlyTrailingReturn5Y          float64
	MonthlyTrailingReturn10Y         float64
	EtradeEarlyRedemptionFee         string
	MaxSalesLoad                     float64
	MonthlyTrailingReturnYTD         float64
	MonthlyTrailingReturn1M          float64
	MonthlyTrailingReturn3M          float64
	MonthlyTrailingReturn6M          float64
	QtrlyTrailingReturnYTD           float64
	QtrlyTrailingReturn1M            float64
	QtrlyTrailingReturn3M            float64
	QtrlyTrailingReturn6M            float64
	DeferredSalesCharges             []ETradeQuoteSaleChargeValues
	FrontEndSalesCharges             []ETradeQuoteSaleChargeValues
	ExchangeCode                     string
}

type ETradeQuoteOptionDeliverable struct {
	RootSymbol               string
	DeliverableSymbol        string
	DeliverableTypeCode      string
	DeliverableExchangeCode  string
	DeliverableStrikePercent float64
	DeliverableCILShares     float64
	DeliverableWholeShares   int32
}

type ETradeQuoteValues struct {
	Low     string
	High    string
	Percent string
}

type ETradeQuoteSaleChargeValues struct {
	LowHigh string
	Percent string
}
