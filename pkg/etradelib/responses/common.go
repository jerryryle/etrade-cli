package responses

//TODO: Explore converting sub-fields to pointers if they may optionally appear.
// e.g. If a struct contains "Product Product" that may not appear in the XML, change it to "Product *Product" so that
// it will be nil if not in the XML. Otherwise, you get a struct full of zeroed data.

type Disclosure struct {
	EhDisclosureFlag          bool `xml:"ehDisclosureFlag"`
	AhDisclosureFlag          bool `xml:"ahDisclosureFlag"`
	ConditionalDisclosureFlag bool `xml:"conditionalDisclosureFlag"`
	AoDisclosureFlag          bool `xml:"aoDisclosureFlag"`
	MfFLConsent               bool `xml:"mfFLConsent"`
	MfEOConsent               bool `xml:"mfEOConsent"`
}

type Instrument struct {
	Product               Product    `xml:"product"`
	SymbolDescription     string     `xml:"symbolDescription"`
	OrderAction           string     `xml:"orderAction"`
	QuantityType          string     `xml:"quantityType"`
	Quantity              float64    `xml:"quantity"`
	CancelQuantity        float64    `xml:"cancelQuantity"`
	OrderedQuantity       float64    `xml:"orderedQuantity"`
	FilledQuantity        float64    `xml:"filledQuantity"`
	AverageExecutionPrice float64    `xml:"averageExecutionPrice"`
	EstimatedCommission   float64    `xml:"estimatedCommission"`
	EstimatedFees         float64    `xml:"estimatedFees"`
	Bid                   float64    `xml:"bid"`
	Ask                   float64    `xml:"ask"`
	Lastprice             float64    `xml:"lastprice"`
	Currency              string     `xml:"currency"`
	Lots                  Lots       `xml:"lots"`
	MFQuantity            MFQuantity `xml:"mfQuantity"`
	OsiKey                string     `xml:"osiKey"`
	MFTransaction         string     `xml:"mfTransaction"`
	ReserveOrder          bool       `xml:"reserveOrder"`
	ReserveQuantity       float64    `xml:"reserveQuantity"`
}

type Lot struct {
	Id   int64   `xml:"id"`
	Size float64 `xml:"size"`
}

type Lots struct {
	Lot []Lot `xml:"lot"`
}

type Message struct {
	Description string `xml:"description"`
	Code        int32  `xml:"code"`
	Type        string `xml:"type"`
}

type MFQuantity struct {
	Cash   float64 `xml:"cash"`
	Margin float64 `xml:"margin"`
	Cusip  string  `xml:"cusip"`
}

type OrderDetail struct {
	OrderNumber           int64        `xml:"orderNumber"`
	AccountId             string       `xml:"accountId"`
	PreviewTime           ETradeTime   `xml:"previewTime"`
	PlacedTime            ETradeTime   `xml:"placedTime"`
	ExecutedTime          ETradeTime   `xml:"executedTime"`
	OrderValue            float64      `xml:"orderValue"`
	Status                string       `xml:"status"`
	OrderType             string       `xml:"orderType"`
	OrderTerm             string       `xml:"orderTerm"`
	PriceType             string       `xml:"priceType"`
	PriceValue            string       `xml:"priceValue"`
	LimitPrice            float64      `xml:"limitPrice"`
	StopPrice             float64      `xml:"stopPrice"`
	StopLimitPrice        float64      `xml:"stopLimitPrice"`
	OffsetType            string       `xml:"offsetType"`
	OffsetValue           float64      `xml:"offsetValue"`
	MarketSession         string       `xml:"marketSession"`
	RoutingDestination    string       `xml:"routingDestination"`
	BracketedLimitPrice   float64      `xml:"bracketedLimitPrice"`
	InitialStopPrice      float64      `xml:"initialStopPrice"`
	TrailPrice            float64      `xml:"trailPrice"`
	TriggerPrice          float64      `xml:"triggerPrice"`
	ConditionPrice        float64      `xml:"conditionPrice"`
	ConditionSymbol       string       `xml:"conditionSymbol"`
	ConditionType         string       `xml:"conditionType"`
	ConditionFollowPrice  string       `xml:"conditionFollowPrice"`
	ConditionSecurityType string       `xml:"conditionSecurityType"`
	ReplacedByOrderId     int64        `xml:"replacedByOrderId"`
	ReplacesOrderId       int64        `xml:"replacesOrderId"`
	AllOrNone             bool         `xml:"allOrNone"`
	PreviewId             int64        `xml:"previewId"`
	Instrument            []Instrument `xml:"instrument"`
	Messages              []Message    `xml:"messages>message"`
	InvestmentAmount      float64      `xml:"investmentAmount"`
	PositionQuantity      string       `xml:"positionQuantity"`
	AipFlag               bool         `xml:"aipFlag"`
	EgQual                string       `xml:"egQual"`
	ReInvestOption        string       `xml:"reInvestOption"`
	EstimatedCommission   float64      `xml:"estimatedCommission"`
	EstimatedFees         float64      `xml:"estimatedFees"`
	EstimatedTotalAmount  float64      `xml:"estimatedTotalAmount"`
	NetPrice              float64      `xml:"netPrice"`
	NetBid                float64      `xml:"netBid"`
	NetAsk                float64      `xml:"netAsk"`
	Gcd                   int32        `xml:"gcd"`
	Ratio                 string       `xml:"ratio"`
	MfpriceType           string       `xml:"mfpriceType"`
}

type OrderBuyPowerEffect struct {
	CurrentBp          float64 `xml:"currentBp"`
	CurrentOor         float64 `xml:"currentOor"`
	CurrentNetBp       float64 `xml:"currentNetBp"`
	CurrentOrderImpact float64 `xml:"currentOrderImpact"`
	NetBp              float64 `xml:"netBp"`
}

type PortfolioMargin struct {
	HouseExcessEquityNew    float64 `xml:"houseExcessEquityNew"`
	PmEligible              bool    `xml:"pmEligible"`
	HouseExcessEquityCurr   float64 `xml:"houseExcessEquityCurr"`
	HouseExcessEquityChange float64 `xml:"houseExcessEquityChange"`
}

type Product struct {
	Symbol          string    `xml:"symbol"`
	SecurityType    string    `xml:"securityType"`
	SecuritySubType string    `xml:"securitySubType"`
	CallPut         string    `xml:"callPut"`
	ExpiryYear      int32     `xml:"expiryYear"`
	ExpiryMonth     int32     `xml:"expiryMonth"`
	ExpiryDay       int32     `xml:"expiryDay"`
	StrikePrice     float64   `xml:"strikePrice"`
	ExpiryType      string    `xml:"expiryType"`
	ProductId       ProductId `xml:"productId"`
}

type ProductId struct {
	Symbol   string `xml:"symbol"`
	TypeCode string `xml:"typeCode"`
}
