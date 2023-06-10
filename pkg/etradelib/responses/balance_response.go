package responses

type BalanceResponse struct {
	AccountId          string                 `xml:"accountId"`
	InstitutionType    string                 `xml:"institutionType"`
	AsOfDate           int64                  `xml:"asOfDate"`
	AccountType        string                 `xml:"accountType"`
	OptionLevel        string                 `xml:"optionLevel"`
	AccountDescription string                 `xml:"accountDescription"`
	QuoteMode          int32                  `xml:"quoteMode"`
	DayTraderStatus    string                 `xml:"dayTraderStatus"`
	AccountMode        string                 `xml:"accountMode"`
	AccountDesc        string                 `xml:"accountDesc"`
	OpenCalls          []BalanceOpenCalls     `xml:"openCalls"`
	Cash               BalanceCash            `xml:"cash"`
	Margin             BalanceMargin          `xml:"margin"`
	Lending            BalanceLending         `xml:"lending"`
	ComputedBalance    BalanceComputedBalance `xml:"computedBalance"`
}

type BalanceCash struct {
	FundsForOpenOrdersCash float64 `xml:"fundsForOpenOrdersCash"`
	MoneyMktBalance        float64 `xml:"moneyMktBalance"`
}

type BalanceComputedBalance struct {
	CashAvailableForInvestment     float64                `xml:"cashAvailableForInvestment"`
	CashAvailableForWithdrawal     float64                `xml:"cashAvailableForWithdrawal"`
	TotalAvailableForWithdrawal    float64                `xml:"totalAvailableForWithdrawal"`
	NetCash                        float64                `xml:"netCash"`
	CashBalance                    float64                `xml:"cashBalance"`
	SettledCashForInvestment       float64                `xml:"settledCashForInvestment"`
	UnSettledCashForInvestment     float64                `xml:"unSettledCashForInvestment"`
	FundsWithheldFromPurchasePower float64                `xml:"fundsWithheldFromPurchasePower"`
	FundsWithheldFromWithdrawal    float64                `xml:"fundsWithheldFromWithdrawal"`
	MarginBuyingPower              float64                `xml:"marginBuyingPower"`
	CashBuyingPower                float64                `xml:"cashBuyingPower"`
	DtMarginBuyingPower            float64                `xml:"dtMarginBuyingPower"`
	DtCashBuyingPower              float64                `xml:"dtCashBuyingPower"`
	MarginBalance                  float64                `xml:"marginBalance"`
	ShortAdjustBalance             float64                `xml:"shortAdjustBalance"`
	RegtEquity                     float64                `xml:"regtEquity"`
	RegtEquityPercent              float64                `xml:"regtEquityPercent"`
	AccountBalance                 float64                `xml:"accountBalance"`
	OpenCalls                      BalanceOpenCalls       `xml:"openCalls"`
	RealTimeValues                 BalanceRealTimeValues  `xml:"realTimeValues"`
	PortfolioMargin                BalancePortfolioMargin `xml:"portfolioMargin"`
}

type BalanceLending struct {
	CurrentBalance          float64 `xml:"currentBalance"`
	CreditLine              float64 `xml:"creditLine"`
	OutstandingBalance      float64 `xml:"outstandingBalance"`
	MinPaymentDue           float64 `xml:"minPaymentDue"`
	AmountPastDue           float64 `xml:"amountPastDue"`
	AvailableCredit         float64 `xml:"availableCredit"`
	YtdInterestPaid         float64 `xml:"ytdInterestPaid"`
	LastYtdInterestPaid     float64 `xml:"lastYtdInterestPaid"`
	PaymentDueDate          int64   `xml:"paymentDueDate"`
	LastPaymentReceivedDate int64   `xml:"lastPaymentReceivedDate"`
	PaymentReceivedMtd      float64 `xml:"paymentReceivedMtd"`
}

type BalanceMargin struct {
	DtCashOpenOrderReserve   float64 `xml:"dtCashOpenOrderReserve"`
	DtMarginOpenOrderReserve float64 `xml:"dtMarginOpenOrderReserve"`
}

type BalanceOpenCalls struct {
	MinEquityCall float64 `xml:"minEquityCall"`
	FedCall       float64 `xml:"fedCall"`
	CashCall      float64 `xml:"cashCall"`
	HouseCall     float64 `xml:"houseCall"`
}

type BalancePortfolioMargin struct {
	DtCashOpenOrderReserve       float64 `xml:"dtCashOpenOrderReserve"`
	DtMarginOpenOrderReserve     float64 `xml:"dtMarginOpenOrderReserve"`
	LiquidatingEquity            float64 `xml:"liquidatingEquity"`
	HouseExcessEquity            float64 `xml:"houseExcessEquity"`
	TotalHouseRequirement        float64 `xml:"totalHouseRequirement"`
	ExcessEquityMinusRequirement float64 `xml:"excessEquityMinusRequirement"`
	TotalMarginRqmts             float64 `xml:"totalMarginRqmts"`
	AvailExcessEquity            float64 `xml:"availExcessEquity"`
	ExcessEquity                 float64 `xml:"excessEquity"`
	OpenOrderReserve             float64 `xml:"openOrderReserve"`
	FundsOnHold                  float64 `xml:"fundsOnHold"`
}

type BalanceRealTimeValues struct {
	TotalAccountValue float64 `xml:"totalAccountValue"`
	NetMv             float64 `xml:"netMv"`
	NetMvLong         float64 `xml:"netMvLong"`
	NetMvShort        float64 `xml:"netMvShort"`
	TotalLongValue    float64 `xml:"totalLongValue"`
}
