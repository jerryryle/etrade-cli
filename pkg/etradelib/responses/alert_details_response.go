package responses

type AlertDetailsResponse struct {
	Id         int64      `xml:"id"`
	CreateTime ETradeTime `xml:"createTime"`
	Subject    string     `xml:"subject"`
	MsgText    string     `xml:"msgText"`
	ReadTime   ETradeTime `xml:"readTime"`
	DeleteTime ETradeTime `xml:"deleteTime"`
	Symbol     string     `xml:"symbol"`
	Next       string     `xml:"next"`
	Prev       string     `xml:"prev"`
}
