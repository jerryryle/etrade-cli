package responses

type PlaceOrderResponse struct {
	OrderType       string              `xml:"orderType"`
	Messages        []Message           `xml:"messageList>message"`
	TotalOrderValue float64             `xml:"totalOrderValue"`
	TotalCommission float64             `xml:"totalCommission"`
	OrderId         int64               `xml:"orderId"`
	Order           []OrderDetail       `xml:"order"`
	DstFlag         bool                `xml:"dstFlag"`
	OptionLevelCd   int32               `xml:"optionLevelCd"`
	MarginLevelCd   string              `xml:"marginLevelCd"`
	IsEmployee      bool                `xml:"isEmployee"`
	CommissionMsg   string              `xml:"commissionMsg"`
	OrderIds        []PlaceOrderOrderId `xml:"orderIds"`
	PlacedTime      int64               `xml:"placedTime"`
	AccountId       string              `xml:"accountId"`
	PortfolioMargin PortfolioMargin     `xml:"portfolioMargin"`
	Disclosure      Disclosure          `xml:"disclosure"`
	ClientOrderId   string              `xml:"clientOrderId"`
}

type PlaceOrderOrderId struct {
	OrderId    int64  `xml:"orderId"`
	CashMargin string `xml:"cashMargin"`
}
