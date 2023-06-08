package etradelib

type ETradeAlert interface {
	ListAlertDetails() string
	DeleteAlert() string
}
