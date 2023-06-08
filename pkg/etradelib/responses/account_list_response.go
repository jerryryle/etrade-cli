package responses

type AccountListResponse struct {
	Accounts []Account `xml:"Accounts>Account"`
}

type Account struct {
	AccountId                  string `xml:"accountId"`
	AccountIdKey               string `xml:"accountIdKey"`
	AccountMode                string `xml:"accountMode"`
	AccountDesc                string `xml:"accountDesc"`
	AccountName                string `xml:"accountName"`
	AccountType                string `xml:"accountType"`
	InstitutionType            string `xml:"institutionType"`
	AccountStatus              string `xml:"accountStatus"`
	ClosedDate                 int64  `xml:"closedDate"`
	ShareWorksAccount          bool   `xml:"shareWorksAccount"`
	ShareWorksSource           string `xml:"shareWorksSource"`
	FcManagedMssbClosedAccount bool   `xml:"fcManagedMssbClosedAccount"`
}
