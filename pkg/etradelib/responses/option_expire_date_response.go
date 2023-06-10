package responses

type OptionExpireDateResponse struct {
	ExpirationDates []OptionExpireDateExpirationDate `xml:"expirationDates"`
}

type OptionExpireDateExpirationDate struct {
	Year       int32  `xml:"year"`
	Month      int32  `xml:"month"`
	Day        int32  `xml:"day"`
	ExpiryType string `xml:"expiryType"`
}
