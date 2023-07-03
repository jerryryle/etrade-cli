package etradelib

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeStatus interface {
	IsSuccess() bool
	GetErrorMessage() string
	AsJsonMap() jsonmap.JsonMap
}

type eTradeStatus struct {
	isSuccess    bool
	errorMessage string
	jsonMap      jsonmap.JsonMap
}

const (
	// The status response JSON looks like this:
	// {
	//   "status": "success" || "error"
	//   "error": "<error message>" (if status=="error")
	// }

	// statusStatusResponseKey is the status key
	statusStatusResponseKey = "status"

	// statusErrorResponseKey is the error key
	statusErrorResponseKey = "error"
)

func CreateETradeStatusFromResponse(response []byte) (ETradeStatus, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeStatus(responseMap)
}

func CreateETradeStatus(responseMap jsonmap.JsonMap) (
	ETradeStatus, error,
) {
	status, err := responseMap.GetString(statusStatusResponseKey)
	if err != nil {
		return nil, err
	}

	isSuccess := false
	errorMessage := ""
	switch status {
	case "success":
		isSuccess = true
	case "error":
		isSuccess = false
		if errorMessage, err = responseMap.GetString(statusErrorResponseKey); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown authentication status %s", status)
	}

	return &eTradeStatus{
		isSuccess:    isSuccess,
		errorMessage: errorMessage,
		jsonMap:      responseMap,
	}, nil
}

func (e *eTradeStatus) IsSuccess() bool {
	return e.isSuccess
}

func (e *eTradeStatus) GetErrorMessage() string {
	return e.errorMessage
}

func (e *eTradeStatus) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
