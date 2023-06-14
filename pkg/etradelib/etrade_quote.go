package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
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

func CreateETradeQuoteInfoInfoFromResponse(response responses.QuoteData) *ETradeQuoteInfo {
	return &ETradeQuoteInfo{
		DateTime:       response.DateTimeUTC.GetTime(),
		QuoteStatus:    response.QuoteStatus,
		AhFlag:         response.AhFlag,
		TimeZone:       response.TimeZone,
		DstFlag:        response.DstFlag,
		HasMiniOptions: response.HasMiniOptions,
	}
}

func CreateETradeQuoteAllInfoFromResponse(response responses.QuoteData) *ETradeQuoteAllInfo {
	return &ETradeQuoteAllInfo{
		ETradeQuoteInfo:             *CreateETradeQuoteInfoInfoFromResponse(response),
		AdjustedFlag:                response.All.AdjustedFlag,
		Ask:                         response.All.Ask,
		AskSize:                     response.All.AskSize,
		AskTime:                     response.All.AskTime.GetTime(),
		Bid:                         response.All.Bid,
		BidExchange:                 response.All.BidExchange,
		BidSize:                     response.All.BidSize,
		BidTime:                     response.All.BidTime.GetTime(),
		ChangeClose:                 response.All.ChangeClose,
		ChangeClosePercentage:       response.All.ChangeClosePercentage,
		CompanyName:                 response.All.CompanyName,
		DaysToExpiration:            response.All.DaysToExpiration,
		DirLast:                     response.All.DirLast,
		Dividend:                    response.All.Dividend,
		Eps:                         response.All.Eps,
		EstEarnings:                 response.All.EstEarnings,
		ExDividendDate:              response.All.ExDividendDate.GetTime(),
		High:                        response.All.High,
		High52:                      response.All.High52,
		LastTrade:                   response.All.LastTrade,
		Low:                         response.All.Low,
		Low52:                       response.All.Low52,
		Open:                        response.All.Open,
		OpenInterest:                response.All.OpenInterest,
		OptionStyle:                 response.All.OptionStyle,
		OptionUnderlier:             response.All.OptionUnderlier,
		OptionUnderlierExchange:     response.All.OptionUnderlierExchange,
		PreviousClose:               response.All.PreviousClose,
		PreviousDayVolume:           response.All.PreviousDayVolume,
		PrimaryExchange:             response.All.PrimaryExchange,
		SymbolDescription:           response.All.SymbolDescription,
		TotalVolume:                 response.All.TotalVolume,
		Upc:                         response.All.Upc,
		OptionDeliverableList:       CreateETradeQuoteOptionDeliverableListFromResponse(response.All.OptionDeliverableList),
		CashDeliverable:             response.All.CashDeliverable,
		MarketCap:                   response.All.MarketCap,
		SharesOutstanding:           response.All.SharesOutstanding,
		NextEarningDate:             response.All.NextEarningDate.GetTime(),
		Beta:                        response.All.Beta,
		Yield:                       response.All.Yield,
		DeclaredDividend:            response.All.DeclaredDividend,
		DividendPayableDate:         response.All.DividendPayableDate.GetTime(),
		Pe:                          response.All.Pe,
		Week52LowDate:               response.All.Week52LowDate.GetTime(),
		Week52HiDate:                response.All.Week52HiDate.GetTime(),
		IntrinsicValue:              response.All.IntrinsicValue,
		TimePremium:                 response.All.TimePremium,
		OptionMultiplier:            response.All.OptionMultiplier,
		ContractSize:                response.All.ContractSize,
		ExpirationDate:              response.All.ExpirationDate.GetTime(),
		ExtendedHourLastPrice:       response.All.EhQuote.LastPrice,
		ExtendedHourChange:          response.All.EhQuote.Change,
		ExtendedHourPercentChange:   response.All.EhQuote.PercentChange,
		ExtendedHourBid:             response.All.EhQuote.Bid,
		ExtendedHourBidSize:         response.All.EhQuote.BidSize,
		ExtendedHourAsk:             response.All.EhQuote.Ask,
		ExtendedHourAskSize:         response.All.EhQuote.AskSize,
		ExtendedHourVolume:          response.All.EhQuote.Volume,
		ExtendedHourTimeOfLastTrade: response.All.EhQuote.TimeOfLastTrade.GetTime(),
		ExtendedHourTimeZone:        response.All.EhQuote.TimeZone,
		ExtendedHourQuoteStatus:     response.All.EhQuote.QuoteStatus,
		OptionPreviousBidPrice:      response.All.OptionPreviousBidPrice,
		OptionPreviousAskPrice:      response.All.OptionPreviousAskPrice,
		OsiKey:                      response.All.OsiKey,
		TimeOfLastTrade:             response.All.TimeOfLastTrade.GetTime(),
		AverageVolume:               response.All.AverageVolume,
	}
}

func CreateETradeQuoteFundamentalInfoFromResponse(response responses.QuoteData) *ETradeQuoteFundamentalInfo {
	return &ETradeQuoteFundamentalInfo{
		ETradeQuoteInfo:   *CreateETradeQuoteInfoInfoFromResponse(response),
		CompanyName:       response.Fundamental.CompanyName,
		Eps:               response.Fundamental.Eps,
		EstEarnings:       response.Fundamental.EstEarnings,
		High52:            response.Fundamental.High52,
		LastTrade:         response.Fundamental.LastTrade,
		Low52:             response.Fundamental.Low52,
		SymbolDescription: response.Fundamental.SymbolDescription,
	}
}

func CreateETradeQuoteIntradayInfoFromResponse(response responses.QuoteData) *ETradeQuoteIntradayInfo {
	return &ETradeQuoteIntradayInfo{
		ETradeQuoteInfo:       *CreateETradeQuoteInfoInfoFromResponse(response),
		Ask:                   response.Intraday.Ask,
		Bid:                   response.Intraday.Bid,
		ChangeClose:           response.Intraday.ChangeClose,
		ChangeClosePercentage: response.Intraday.ChangeClosePercentage,
		CompanyName:           response.Intraday.CompanyName,
		High:                  response.Intraday.High,
		LastTrade:             response.Intraday.LastTrade,
		Low:                   response.Intraday.Low,
		TotalVolume:           response.Intraday.TotalVolume,
	}
}

func CreateETradeQuoteOptionsInfoFromResponse(response responses.QuoteData) *ETradeQuoteOptionsInfo {
	return &ETradeQuoteOptionsInfo{
		ETradeQuoteInfo:          *CreateETradeQuoteInfoInfoFromResponse(response),
		Ask:                      response.Option.Ask,
		AskSize:                  response.Option.AskSize,
		Bid:                      response.Option.Bid,
		BidSize:                  response.Option.BidSize,
		CompanyName:              response.Option.CompanyName,
		DaysToExpiration:         response.Option.DaysToExpiration,
		LastTrade:                response.Option.LastTrade,
		OpenInterest:             response.Option.OpenInterest,
		OptionPreviousBidPrice:   response.Option.OptionPreviousBidPrice,
		OptionPreviousAskPrice:   response.Option.OptionPreviousAskPrice,
		OsiKey:                   response.Option.OsiKey,
		IntrinsicValue:           response.Option.IntrinsicValue,
		TimePremium:              response.Option.TimePremium,
		OptionMultiplier:         response.Option.OptionMultiplier,
		ContractSize:             response.Option.ContractSize,
		SymbolDescription:        response.Option.SymbolDescription,
		OptionGreeksRho:          response.Option.OptionGreeks.Rho,
		OptionGreeksVega:         response.Option.OptionGreeks.Vega,
		OptionGreeksTheta:        response.Option.OptionGreeks.Theta,
		OptionGreeksDelta:        response.Option.OptionGreeks.Delta,
		OptionGreeksGamma:        response.Option.OptionGreeks.Gamma,
		OptionGreeksIv:           response.Option.OptionGreeks.Iv,
		OptionGreeksCurrentValue: response.Option.OptionGreeks.CurrentValue,
	}
}

func CreateETradeQuoteWeek52InfoFromResponse(response responses.QuoteData) *ETradeQuoteWeek52Info {
	return &ETradeQuoteWeek52Info{
		ETradeQuoteInfo:   *CreateETradeQuoteInfoInfoFromResponse(response),
		CompanyName:       response.Week52.CompanyName,
		High52:            response.Week52.High52,
		LastTrade:         response.Week52.LastTrade,
		Low52:             response.Week52.Low52,
		Perf12Months:      response.Week52.Perf12Months,
		PreviousClose:     response.Week52.PreviousClose,
		SymbolDescription: response.Week52.SymbolDescription,
		TotalVolume:       response.Week52.TotalVolume,
	}
}

func CreateETradeQuoteMutualFundInfoFromResponse(response responses.QuoteData) *ETradeQuoteMutualFundInfo {
	return &ETradeQuoteMutualFundInfo{
		ETradeQuoteInfo:                  *CreateETradeQuoteInfoInfoFromResponse(response),
		SymbolDescription:                response.MutualFund.SymbolDescription,
		Cusip:                            response.MutualFund.Cusip,
		ChangeClose:                      response.MutualFund.ChangeClose,
		PreviousClose:                    response.MutualFund.PreviousClose,
		TransactionFee:                   response.MutualFund.TransactionFee,
		EarlyRedemptionFee:               response.MutualFund.EarlyRedemptionFee,
		Availability:                     response.MutualFund.Availability,
		InitialInvestment:                response.MutualFund.InitialInvestment,
		SubsequentInvestment:             response.MutualFund.SubsequentInvestment,
		FundFamily:                       response.MutualFund.FundFamily,
		FundName:                         response.MutualFund.FundName,
		ChangeClosePercentage:            response.MutualFund.ChangeClosePercentage,
		TimeOfLastTrade:                  response.MutualFund.TimeOfLastTrade.GetTime(),
		NetAssetValue:                    response.MutualFund.NetAssetValue,
		PublicOfferPrice:                 response.MutualFund.PublicOfferPrice,
		NetExpenseRatio:                  response.MutualFund.NetExpenseRatio,
		GrossExpenseRatio:                response.MutualFund.GrossExpenseRatio,
		OrderCutoffTime:                  response.MutualFund.OrderCutoffTime,
		SalesCharge:                      response.MutualFund.SalesCharge,
		InitialIraInvestment:             response.MutualFund.InitialIraInvestment,
		SubsequentIraInvestment:          response.MutualFund.SubsequentIraInvestment,
		NetAssetTotalValue:               response.MutualFund.NetAssets.Value,
		NetAssetTotalAsOfDate:            response.MutualFund.NetAssets.AsOfDate.GetTime(),
		FundInceptionDate:                response.MutualFund.FundInceptionDate.GetTime(),
		AverageAnnualReturns:             response.MutualFund.AverageAnnualReturns,
		SevenDayCurrentYield:             response.MutualFund.SevenDayCurrentYield,
		AnnualTotalReturn:                response.MutualFund.AnnualTotalReturn,
		WeightedAverageMaturity:          response.MutualFund.WeightedAverageMaturity,
		AverageAnnualReturn1Yr:           response.MutualFund.AverageAnnualReturn1Yr,
		AverageAnnualReturn3Yr:           response.MutualFund.AverageAnnualReturn3Yr,
		AverageAnnualReturn5Yr:           response.MutualFund.AverageAnnualReturn5Yr,
		AverageAnnualReturn10Yr:          response.MutualFund.AverageAnnualReturn10Yr,
		High52:                           response.MutualFund.High52,
		Low52:                            response.MutualFund.Low52,
		Week52LowDate:                    response.MutualFund.Week52LowDate.GetTime(),
		Week52HiDate:                     response.MutualFund.Week52HiDate.GetTime(),
		ExchangeName:                     response.MutualFund.ExchangeName,
		SinceInception:                   response.MutualFund.SinceInception,
		QuarterlySinceInception:          response.MutualFund.QuarterlySinceInception,
		LastTrade:                        response.MutualFund.LastTrade,
		Actual12B1Fee:                    response.MutualFund.Actual12B1Fee,
		PerformanceAsOfDate:              response.MutualFund.PerformanceAsOfDate.GetTime(),
		QtrlyPerformanceAsOfDate:         response.MutualFund.QtrlyPerformanceAsOfDate.GetTime(),
		RedemptionMinMonth:               response.MutualFund.Redemption.MinMonth,
		RedemptionFeePercent:             response.MutualFund.Redemption.FeePercent,
		RedemptionIsFrontEnd:             response.MutualFund.Redemption.IsFrontEnd,
		RedemptionFrontEndValues:         CreateETradeQuoteValuesListFromResponse(response.MutualFund.Redemption.FrontEndValues),
		RedemptionRedemptionDurationType: response.MutualFund.Redemption.RedemptionDurationType,
		RedemptionIsSales:                response.MutualFund.Redemption.IsSales,
		RedemptionSalesDurationType:      response.MutualFund.Redemption.SalesDurationType,
		RedemptionSalesValues:            CreateETradeQuoteValuesListFromResponse(response.MutualFund.Redemption.SalesValues),
		MorningStarCategory:              response.MutualFund.MorningStarCategory,
		MonthlyTrailingReturn1Y:          response.MutualFund.MonthlyTrailingReturn1Y,
		MonthlyTrailingReturn3Y:          response.MutualFund.MonthlyTrailingReturn3Y,
		MonthlyTrailingReturn5Y:          response.MutualFund.MonthlyTrailingReturn5Y,
		MonthlyTrailingReturn10Y:         response.MutualFund.MonthlyTrailingReturn10Y,
		EtradeEarlyRedemptionFee:         response.MutualFund.EtradeEarlyRedemptionFee,
		MaxSalesLoad:                     response.MutualFund.MaxSalesLoad,
		MonthlyTrailingReturnYTD:         response.MutualFund.MonthlyTrailingReturnYTD,
		MonthlyTrailingReturn1M:          response.MutualFund.MonthlyTrailingReturn1M,
		MonthlyTrailingReturn3M:          response.MutualFund.MonthlyTrailingReturn3M,
		MonthlyTrailingReturn6M:          response.MutualFund.MonthlyTrailingReturn6M,
		QtrlyTrailingReturnYTD:           response.MutualFund.QtrlyTrailingReturnYTD,
		QtrlyTrailingReturn1M:            response.MutualFund.QtrlyTrailingReturn1M,
		QtrlyTrailingReturn3M:            response.MutualFund.QtrlyTrailingReturn3M,
		QtrlyTrailingReturn6M:            response.MutualFund.QtrlyTrailingReturn6M,
		DeferredSalesCharges:             CreateETradeQuoteSaleChargeValuesListFromResponse(response.MutualFund.DeferredSalesCharges),
		FrontEndSalesCharges:             CreateETradeQuoteSaleChargeValuesListFromResponse(response.MutualFund.FrontEndSalesCharges),
		ExchangeCode:                     response.MutualFund.ExchangeCode,
	}
}

func CreateETradeQuoteOptionDeliverableListFromResponse(response []responses.QuoteOptionDeliverable) []ETradeQuoteOptionDeliverable {
	optionDeliverableList := make([]ETradeQuoteOptionDeliverable, 0)
	for _, respOptionDeliverable := range response {
		optionDeliverableList = append(
			optionDeliverableList, ETradeQuoteOptionDeliverable{
				RootSymbol:               respOptionDeliverable.RootSymbol,
				DeliverableSymbol:        respOptionDeliverable.DeliverableSymbol,
				DeliverableTypeCode:      respOptionDeliverable.DeliverableTypeCode,
				DeliverableExchangeCode:  respOptionDeliverable.DeliverableExchangeCode,
				DeliverableStrikePercent: respOptionDeliverable.DeliverableStrikePercent,
				DeliverableCILShares:     respOptionDeliverable.DeliverableCILShares,
				DeliverableWholeShares:   respOptionDeliverable.DeliverableWholeShares,
			},
		)
	}
	return optionDeliverableList
}

func CreateETradeQuoteValuesListFromResponse(response []responses.QuoteValues) []ETradeQuoteValues {
	quoteValuesList := make([]ETradeQuoteValues, 0)
	for _, respQuoteValues := range response {
		quoteValuesList = append(
			quoteValuesList, ETradeQuoteValues{
				Low:     respQuoteValues.Low,
				High:    respQuoteValues.High,
				Percent: respQuoteValues.Percent,
			},
		)
	}
	return quoteValuesList
}

func CreateETradeQuoteSaleChargeValuesListFromResponse(response []responses.QuoteSaleChargeValues) []ETradeQuoteSaleChargeValues {
	quoteValuesList := make([]ETradeQuoteSaleChargeValues, 0)
	for _, respQuoteSaleChargeValues := range response {
		quoteValuesList = append(
			quoteValuesList, ETradeQuoteSaleChargeValues{
				LowHigh: respQuoteSaleChargeValues.LowHigh,
				Percent: respQuoteSaleChargeValues.Percent,
			},
		)
	}
	return quoteValuesList
}
