package etradelib

type ETradeAccount interface {
	GetAccountBalances() string
	ListTransactions() string
	ViewPortfolio() string
	ListOrders() string
	CreateOrder() string
}
