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

	// statusStatusKey is the status key
	statusStatusKey = "status"

	// statusErrorKey is the error key
	statusErrorKey = "error"
)

func CreateETradeStatusFromResponse(response []byte) (ETradeStatus, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeStatus(responseMap)
}

func CreateETradeStatus(statusResponseMap jsonmap.JsonMap) (
	ETradeStatus, error,
) {
	status, err := statusResponseMap.GetString(statusStatusKey)
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
		if errorMessage, err = statusResponseMap.GetString(statusErrorKey); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown authentication status %s", status)
	}

	return &eTradeStatus{
		isSuccess:    isSuccess,
		errorMessage: errorMessage,
		jsonMap:      statusResponseMap,
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
