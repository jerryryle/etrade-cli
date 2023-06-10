package responses

type AlertDetailsResponse struct {
	Id         int64  `xml:"id"`
	CreateTime int64  `xml:"createTime"`
	Subject    string `xml:"subject"`
	MsgText    string `xml:"msgText"`
	ReadTime   int64  `xml:"readTime"`
	DeleteTime int64  `xml:"deleteTime"`
	Symbol     string `xml:"symbol"`
	Next       string `xml:"next"`
	Prev       string `xml:"prev"`
}
