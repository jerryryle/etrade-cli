package responses

type CancelOrderResponse struct {
	AccountId  string     `xml:"accountId"`
	OrderId    int64      `xml:"orderId"`
	CancelTime ETradeTime `xml:"cancelTime"`
	Messages   []Message  `xml:"messages>message"`
}
