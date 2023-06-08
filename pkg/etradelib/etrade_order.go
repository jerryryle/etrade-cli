package etradelib

type ETradeOrder interface {
	Preview() string
	Place() string
	Cancel() string
}
