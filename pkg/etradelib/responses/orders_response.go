package responses

type OrdersResponse struct {
	Marker   string        `xml:"marker"`
	Next     string        `xml:"next"`
	Order    []OrdersOrder `xml:"order"`
	Messages []Message     `xml:"messages>message"`
}

type OrdersOrder struct {
	OrderId             int64               `xml:"orderId"`
	Details             string              `xml:"details"`
	OrderType           string              `xml:"orderType"`
	TotalOrderValue     float64             `xml:"totalOrderValue"`
	TotalCommission     float64             `xml:"totalCommission"`
	OrderDetail         []OrderDetail       `xml:"orderDetail"`
	Events              []OrdersEvent       `xml:"events>event"`
	OrderBuyPowerEffect OrderBuyPowerEffect `xml:"orderBuyPowerEffect"`
}

type OrdersEvent struct {
	Name        string       `xml:"name"`
	DateTime    int64        `xml:"dateTime"`
	OrderNumber int          `xml:"orderNumber"`
	Instrument  []Instrument `xml:"instrument"` // Include when Instrument struct is defined
}
