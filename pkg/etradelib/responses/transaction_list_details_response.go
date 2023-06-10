package responses

type TransactionDetailsResponse struct {
	TransactionId   int64                       `xml:"transactionId"`
	AccountId       string                      `xml:"accountId"`
	TransactionDate ETradeTime                  `xml:"transactionDate"`
	PostDate        ETradeTime                  `xml:"postDate"`
	Amount          float64                     `xml:"amount"`
	Description     string                      `xml:"description"`
	Category        TransactionDetailsCategory  `xml:"category"`
	Brokerage       TransactionDetailsBrokerage `xml:"brokerage"`
}

type TransactionDetailsBrokerage struct {
	TransactionType    string  `xml:"transactionType"`
	Product            Product `xml:"product"`
	Quantity           float64 `xml:"quantity"`
	Price              float64 `xml:"price"`
	SettlementCurrency string  `xml:"settlementCurrency"`
	PaymentCurrency    string  `xml:"paymentCurrency"`
	Fee                float64 `xml:"fee"`
	Memo               string  `xml:"memo"`
	CheckNo            string  `xml:"checkNo"`
	OrderNo            string  `xml:"orderNo"`
}

type TransactionDetailsCategory struct {
	CategoryId   string `xml:"categoryId"`
	ParentId     string `xml:"parentId"`
	CategoryName string `xml:"categoryName"`
	ParentName   string `xml:"parentName"`
}
