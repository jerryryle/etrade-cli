package responses

type PortfolioResponse struct {
	Totals           PortfolioTotals             `xml:"Totals"`
	AccountPortfolio []PortfolioAccountPortfolio `xml:"AccountPortfolio"`
}

type PortfolioAccountPortfolio struct {
	AccountId  string              `xml:"accountId"`
	Next       string              `xml:"next"`
	TotalPages int32               `xml:"totalPages"`
	NextPageNo string              `xml:"nextPageNo"`
	Positions  []PortfolioPosition `xml:"Position"`
}

type PortfolioCompleteView struct {
	PriceAdjustedFlag   bool       `xml:"priceAdjustedFlag"`
	Price               float64    `xml:"price"`
	AdjPrice            float64    `xml:"adjPrice"`
	Change              float64    `xml:"change"`
	ChangePct           float64    `xml:"changePct"`
	PrevClose           float64    `xml:"prevClose"`
	AdjPrevClose        float64    `xml:"adjPrevClose"`
	Volume              float64    `xml:"volume"`
	LastTrade           float64    `xml:"lastTrade"`
	LastTradeTime       ETradeTime `xml:"lastTradeTime"`
	AdjLastTrade        float64    `xml:"adjLastTrade"`
	SymbolDescription   string     `xml:"symbolDescription"`
	Perform1Month       float64    `xml:"perform1Month"`
	Perform3Month       float64    `xml:"perform3Month"`
	Perform6Month       float64    `xml:"perform6Month"`
	Perform12Month      float64    `xml:"perform12Month"`
	PrevDayVolume       int64      `xml:"prevDayVolume"`
	TenDayVolume        int64      `xml:"tenDayVolume"`
	Beta                float64    `xml:"beta"`
	Sv10DaysAvg         float64    `xml:"sv10DaysAvg"`
	Sv20DaysAvg         float64    `xml:"sv20DaysAvg"`
	Sv1MonAvg           float64    `xml:"sv1MonAvg"`
	Sv2MonAvg           float64    `xml:"sv2MonAvg"`
	Sv3MonAvg           float64    `xml:"sv3MonAvg"`
	Sv4MonAvg           float64    `xml:"sv4MonAvg"`
	Sv6MonAvg           float64    `xml:"sv6MonAvg"`
	Week52High          float64    `xml:"week52High"`
	Week52Low           float64    `xml:"week52Low"`
	Week52Range         string     `xml:"week52Range"`
	MarketCap           float64    `xml:"marketCap"`
	DaysRange           string     `xml:"daysRange"`
	Delta52WkHigh       float64    `xml:"delta52WkHigh"`
	Delta52WkLow        float64    `xml:"delta52WkLow"`
	Currency            string     `xml:"currency"`
	Exchange            string     `xml:"exchange"`
	Marginable          bool       `xml:"marginable"`
	Bid                 float64    `xml:"bid"`
	Ask                 float64    `xml:"ask"`
	BidAskSpread        float64    `xml:"bidAskSpread"`
	BidSize             int64      `xml:"bidSize"`
	AskSize             int64      `xml:"askSize"`
	Open                float64    `xml:"open"`
	Delta               float64    `xml:"delta"`
	Gamma               float64    `xml:"gamma"`
	IvPct               float64    `xml:"ivPct"`
	Rho                 float64    `xml:"rho"`
	Theta               float64    `xml:"theta"`
	Vega                float64    `xml:"vega"`
	Premium             float64    `xml:"premium"`
	DaysToExpiration    int32      `xml:"daysToExpiration"`
	IntrinsicValue      float64    `xml:"intrinsicValue"`
	OpenInterest        float64    `xml:"openInterest"`
	OptionsAdjustedFlag bool       `xml:"optionsAdjustedFlag"`
	DeliverablesStr     string     `xml:"deliverablesStr"`
	OptionMultiplier    float64    `xml:"optionMultiplier"`
	BaseSymbolAndPrice  string     `xml:"baseSymbolAndPrice"`
	EstEarnings         float64    `xml:"estEarnings"`
	Eps                 float64    `xml:"eps"`
	PeRatio             float64    `xml:"peRatio"`
	AnnualDividend      float64    `xml:"annualDividend"`
	Dividend            float64    `xml:"dividend"`
	DivYield            float64    `xml:"divYield"`
	DivPayDate          ETradeTime `xml:"divPayDate"`
	ExDividendDate      ETradeTime `xml:"exDividendDate"`
	Cusip               string     `xml:"cusip"`
	QuoteStatus         string     `xml:"quoteStatus"`
}

type PortfolioFundamentalView struct {
	LastTrade     float64    `xml:"lastTrade"`
	LastTradeTime ETradeTime `xml:"lastTradeTime"`
	Change        float64    `xml:"change"`
	ChangePct     float64    `xml:"changePct"`
	PeRatio       float64    `xml:"peRatio"`
	Eps           float64    `xml:"eps"`
	Dividend      float64    `xml:"dividend"`
	DivYield      float64    `xml:"divYield"`
	MarketCap     float64    `xml:"marketCap"`
	Week52Range   string     `xml:"week52Range"`
	QuoteStatus   string     `xml:"quoteStatus"`
}

type PortfolioOptionsWatchView struct {
	BaseSymbolAndPrice string     `xml:"baseSymbolAndPrice"`
	Premium            float64    `xml:"premium"`
	LastTrade          float64    `xml:"lastTrade"`
	Bid                float64    `xml:"bid"`
	Ask                float64    `xml:"ask"`
	QuoteStatus        string     `xml:"quoteStatus"`
	LastTradeTime      ETradeTime `xml:"lastTradeTime"`
}

type PortfolioPerformanceView struct {
	Change        float64    `xml:"change"`
	ChangePct     float64    `xml:"changePct"`
	LastTrade     float64    `xml:"lastTrade"`
	DaysGain      float64    `xml:"daysGain"`
	TotalGain     float64    `xml:"totalGain"`
	TotalGainPct  float64    `xml:"totalGainPct"`
	MarketValue   float64    `xml:"marketValue"`
	QuoteStatus   string     `xml:"quoteStatus"`
	LastTradeTime ETradeTime `xml:"lastTradeTime"`
}

type PortfolioPosition struct {
	PositionId        int64                     `xml:"positionId"`
	AccountId         string                    `xml:"accountId"`
	Product           Product                   `xml:"Product"`
	OsiKey            string                    `xml:"osiKey"`
	SymbolDescription string                    `xml:"symbolDescription"`
	DateAcquired      ETradeTime                `xml:"dateAcquired"`
	PricePaid         float64                   `xml:"pricePaid"`
	Price             float64                   `xml:"price"`
	Commissions       float64                   `xml:"commissions"`
	OtherFees         float64                   `xml:"otherFees"`
	Quantity          float64                   `xml:"quantity"`
	PositionIndicator string                    `xml:"positionIndicator"`
	PositionType      string                    `xml:"positionType"`
	Change            float64                   `xml:"change"`
	ChangePct         float64                   `xml:"changePct"`
	DaysGain          float64                   `xml:"daysGain"`
	DaysGainPct       float64                   `xml:"daysGainPct"`
	MarketValue       float64                   `xml:"marketValue"`
	TotalCost         float64                   `xml:"totalCost"`
	TotalGain         float64                   `xml:"totalGain"`
	TotalGainPct      float64                   `xml:"totalGainPct"`
	PctOfPortfolio    float64                   `xml:"pctOfPortfolio"`
	CostPerShare      float64                   `xml:"costPerShare"`
	TodayCommissions  float64                   `xml:"todayCommissions"`
	TodayFees         float64                   `xml:"todayFees"`
	TodayPricePaid    float64                   `xml:"todayPricePaid"`
	TodayQuantity     float64                   `xml:"todayQuantity"`
	Quotestatus       string                    `xml:"quotestatus"`
	DateTimeUTC       ETradeTime                `xml:"dateTimeUTC"`
	AdjPrevClose      float64                   `xml:"adjPrevClose"`
	Performance       PortfolioPerformanceView  `xml:"Performance"`
	Fundamental       PortfolioFundamentalView  `xml:"Fundamental"`
	OptionsWatch      PortfolioOptionsWatchView `xml:"OptionsWatch"`
	Quick             PortfolioQuickView        `xml:"Quick"`
	Complete          PortfolioCompleteView     `xml:"Complete"`
	LotsDetails       string                    `xml:"lotsDetails"`
	QuoteDetails      string                    `xml:"quoteDetails"`
	PositionLot       []PortfolioPositionLot    `xml:"PositionLot"`
}

type PortfolioPositionLot struct {
	PositionId          int64      `xml:"positionId"`
	PositionLotId       int64      `xml:"positionLotId"`
	Price               float64    `xml:"price"`
	TermCode            int32      `xml:"termCode"`
	DaysGain            float64    `xml:"daysGain"`
	DaysGainPct         float64    `xml:"daysGainPct"`
	MarketValue         float64    `xml:"marketValue"`
	TotalCost           float64    `xml:"totalCost"`
	TotalCostForGainPct float64    `xml:"totalCostForGainPct"`
	TotalGain           float64    `xml:"totalGain"`
	LotSourceCode       int32      `xml:"lotSourceCode"`
	OriginalQty         float64    `xml:"originalQty"`
	RemainingQty        float64    `xml:"remainingQty"`
	AvailableQty        float64    `xml:"availableQty"`
	OrderNo             int64      `xml:"orderNo"`
	LegNo               int32      `xml:"legNo"`
	AcquiredDate        ETradeTime `xml:"acquiredDate"`
	LocationCode        int32      `xml:"locationCode"`
	ExchangeRate        float64    `xml:"exchangeRate"`
	SettlementCurrency  string     `xml:"settlementCurrency"`
	PaymentCurrency     string     `xml:"paymentCurrency"`
	AdjPrice            float64    `xml:"adjPrice"`
	CommPerShare        float64    `xml:"commPerShare"`
	FeesPerShare        float64    `xml:"feesPerShare"`
	PremiumAdj          float64    `xml:"premiumAdj"`
	ShortType           int32      `xml:"shortType"`
}

type PortfolioQuickView struct {
	LastTrade               float64    `xml:"lastTrade"`
	LastTradeTime           ETradeTime `xml:"lastTradeTime"`
	Change                  float64    `xml:"change"`
	ChangePct               float64    `xml:"changePct"`
	Volume                  ETradeTime `xml:"volume"`
	QuoteStatus             string     `xml:"quoteStatus"`
	SevenDayCurrentYield    float64    `xml:"sevenDayCurrentYield"`
	AnnualTotalReturn       float64    `xml:"annualTotalReturn"`
	WeightedAverageMaturity float64    `xml:"weightedAverageMaturity"`
}

type PortfolioTotals struct {
	TodaysGainLoss    float64 `xml:"todaysGainLoss"`
	TodaysGainLossPct float64 `xml:"todaysGainLossPct"`
	TotalMarketValue  float64 `xml:"totalMarketValue"`
	TotalGainLoss     float64 `xml:"totalGainLoss"`
	TotalGainLossPct  float64 `xml:"totalGainLossPct"`
	TotalPricePaid    float64 `xml:"totalPricePaid"`
	CashBalance       float64 `xml:"cashBalance"`
}
