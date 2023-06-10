package etradelib

type ETradeCustomer interface {
	GetCustomerName() string
	GetAllAccounts() ([]ETradeAccount, error)
	ListAlerts() (string, error)
	GetQuotes(symbols string) (string, error)
	LookUpProduct(search string) (string, error)
	GetOptionChains() (string, error)
	GetOptionExpireDates() (string, error)
}

type eTradeCustomer struct {
	client       ETradeClient
	customerName string
}

func CreateETradeCustomer(client ETradeClient, customerName string) ETradeCustomer {
	return &eTradeCustomer{
		client:       client,
		customerName: customerName,
	}
}

func (c *eTradeCustomer) GetCustomerName() string {
	return c.customerName
}

func (c *eTradeCustomer) GetAllAccounts() ([]ETradeAccount, error) {
	response, err := c.client.ListAccounts()
	if err != nil {
		return nil, err
	}
	var accounts = make([]ETradeAccount, 0)
	for _, account := range response.Accounts {
		accounts = append(
			accounts,
			CreateETradeAccount(c.client, CreateETradeAccountInfoFromResponse(account)))
	}
	return accounts, err
}

func (c *eTradeCustomer) ListAlerts() (string, error) {
	return "", nil
}

func (c *eTradeCustomer) GetQuotes(symbols string) (string, error) {
	return "", nil
}

func (c *eTradeCustomer) LookUpProduct(search string) (string, error) {
	return "", nil
}

func (c *eTradeCustomer) GetOptionChains() (string, error) {
	return "", nil
}

func (c *eTradeCustomer) GetOptionExpireDates() (string, error) {
	return "", nil
}
