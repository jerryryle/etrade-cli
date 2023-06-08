package etradelib

type eTradeClientAccountListResponse struct {
	Accounts []eTradeClientAccount `json:"Accounts"`
}

type eTradeClientAccount struct {
	AccountId         string
	AccountIdKey      string
	AccountMode       string
	AccountDesc       string
	AccountName       string
	AccountType       string
	InstitutionType   string
	AccountStatus     string
	ClosedDate        int64
	ShareWorksAccount bool
	ShareWorksSource  string
}
