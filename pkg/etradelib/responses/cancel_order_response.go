package responses

type CancelOrderResponse struct {
	AccountId  string    `xml:"accountId"`
	OrderId    int64     `xml:"orderId"`
	CancelTime int64     `xml:"cancelTime"`
	Messages   []Message `xml:"messages>message"`
}
