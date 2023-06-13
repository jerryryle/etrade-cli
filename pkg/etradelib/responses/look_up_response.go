package responses

type LookupResponse struct {
	Data []LookupData `xml:"Data"`
}

type LookupData struct {
	Symbol      string `xml:"symbol"`
	Description string `xml:"description"`
	Type        string `xml:"type"`
}
