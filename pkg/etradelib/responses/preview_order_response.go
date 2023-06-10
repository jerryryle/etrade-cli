package responses

type PreviewOrderResponse struct {
	OrderType       string                               `xml:"orderType"`
	Message         []Message                            `xml:"messageList>message"`
	TotalOrderValue float64                              `xml:"totalOrderValue"`
	TotalCommission float64                              `xml:"totalCommission"`
	Order           []OrderDetail                        `xml:"order"`
	PreviewIds      []PreviewOrderPreviewId              `xml:"previewIds"`
	PreviewTime     int64                                `xml:"previewTime"`
	DstFlag         bool                                 `xml:"dstFlag"`
	AccountId       string                               `xml:"accountId"`
	OptionLevelCd   int32                                `xml:"optionLevelCd"`
	MarginLevelCd   string                               `xml:"marginLevelCd"`
	PortfolioMargin PortfolioMargin                      `xml:"portfolioMargin"`
	IsEmployee      bool                                 `xml:"isEmployee"`
	CommissionMsg   string                               `xml:"commissionMessage"`
	Disclosure      Disclosure                           `xml:"disclosure"`
	ClientOrderId   string                               `xml:"clientOrderId"`
	MarginBpDetails PreviewOrderMarginBuyingPowerDetails `xml:"marginBpDetails"`
	CashBpDetails   PreviewOrderCashBuyingPowerDetails   `xml:"cashBpDetails"`
	DtBpDetails     PreviewOrderDtBuyingPowerDetails     `xml:"dtBpDetails"`
}

type PreviewOrderCashBuyingPowerDetails struct {
	Settled          OrderBuyPowerEffect `xml:"settled"`
	SettledUnsettled OrderBuyPowerEffect `xml:"settledUnsettled"`
}

type PreviewOrderDtBuyingPowerDetails struct {
	NonMarginable OrderBuyPowerEffect `xml:"nonMarginable"`
	Marginable    OrderBuyPowerEffect `xml:"marginable"`
}

type PreviewOrderMarginBuyingPowerDetails struct {
	NonMarginable OrderBuyPowerEffect `xml:"nonMarginable"`
	Marginable    OrderBuyPowerEffect `xml:"marginable"`
}

type PreviewOrderPreviewId struct {
	PreviewId  int64  `xml:"previewId"`
	CashMargin string `xml:"cashMargin"`
}
