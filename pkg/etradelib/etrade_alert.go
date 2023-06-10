package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
	"time"
)

type ETradeAlert interface {
	GetAlertInfo() ETradeAlertInfo
	GetAlertDetails() (string, error)
	DeleteAlert() (string, error)
}

const (
	ETradeAlertStatusUnread    = "UNREAD"
	ETradeAlertStatusRead      = "READ"
	ETradeAlertStatusDeleted   = "DELETED"
	ETradeAlertStatusUndeleted = "UNDELETED"
)

type ETradeAlertInfo struct {
	Id         int64
	CreateTime time.Time
	Subject    string
	Status     string
}

type eTradeAlert struct {
	client    ETradeClient
	alertInfo ETradeAlertInfo
}

func CreateETradeAlertInfoFromResponse(response responses.AlertsAlert) *ETradeAlertInfo {
	return &ETradeAlertInfo{
		Id:         response.Id,
		CreateTime: response.CreateTime.GetTime(),
		Subject:    response.Subject,
		Status:     response.Status,
	}
}

func CreateETradeAlert(client ETradeClient, alertInfo *ETradeAlertInfo) ETradeAlert {
	return &eTradeAlert{
		client:    client,
		alertInfo: *alertInfo,
	}
}

func (a *eTradeAlert) GetAlertInfo() ETradeAlertInfo {
	return a.alertInfo
}

func (a *eTradeAlert) GetAlertDetails() (string, error) {
	return "", nil
}

func (a *eTradeAlert) DeleteAlert() (string, error) {
	return "", nil
}
