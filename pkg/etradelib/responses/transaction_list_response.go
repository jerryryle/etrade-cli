package responses

type TransactionListResponse struct {
	Transactions     []TransactionListTransaction `xml:"Transaction"`
	Next             string                       `xml:"next"`
	Marker           string                       `xml:"marker"`
	PageMarkers      string                       `xml:"pageMarkers"`
	MoreTransactions bool                         `xml:"moreTransactions"`
	TransactionCount int                          `xml:"transactionCount"`
	TotalCount       int                          `xml:"totalCount"`
}

type TransactionListTransaction struct {
	TransactionId   string                   `xml:"transactionId"`
	AccountId       string                   `xml:"accountId"`
	TransactionDate int64                    `xml:"transactionDate"`
	PostDate        int64                    `xml:"postDate"`
	Amount          float64                  `xml:"amount"`
	Description     string                   `xml:"description"`
	Description2    string                   `xml:"description2"`
	TransactionType string                   `xml:"transactionType"`
	Memo            string                   `xml:"memo"`
	ImageFlag       bool                     `xml:"imageFlag"`
	InstType        string                   `xml:"instType"`
	Brokerage       TransactionListBrokerage `xml:"brokerage"`
	DetailsURI      string                   `xml:"detailsURI"`
}

type TransactionListBrokerage struct {
	Product            Product `xml:"product"`
	Quantity           float64 `xml:"quantity"`
	Price              float64 `xml:"price"`
	SettlementCurrency string  `xml:"settlementCurrency"`
	PaymentCurrency    string  `xml:"paymentCurrency"`
	Fee                float64 `xml:"fee"`
	SettlementDate     int64   `xml:"settlementDate"`
}
