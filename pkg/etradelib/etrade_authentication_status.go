package etradelib

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAuthenticationStatus interface {
	NeedAuthorization() bool
	GetAuthorizationUrl() string
	AsJsonMap() jsonmap.JsonMap
}

type eTradeAuthenticationStatus struct {
	authorizationUrl string
	jsonMap          jsonmap.JsonMap
}

const (
	// The authenticate response JSON looks like this:
	// {
	//   "status": "authorize" || "success"
	//   "verifyUrl": "<verify URL>" (if status=="verify")
	// }

	// authenticationStatusStatusKey is the status key
	authenticationStatusStatusKey = "status"

	// authenticationStatusAuthorizationUrlKey is the authorization URL key
	authenticationStatusAuthorizationUrlKey = "authorizationUrl"
)

func CreateETradeAuthenticationStatusFromResponse(response []byte) (ETradeAuthenticationStatus, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeAuthenticationStatus(responseMap)
}

func CreateETradeAuthenticationStatus(authenticationStatusResponseMap jsonmap.JsonMap) (
	ETradeAuthenticationStatus, error,
) {
	status, err := authenticationStatusResponseMap.GetString(authenticationStatusStatusKey)
	if err != nil {
		return nil, err
	}

	authorizationUrl := ""
	if status == "authorize" {
		if authorizationUrl, err = authenticationStatusResponseMap.GetString(authenticationStatusAuthorizationUrlKey); err != nil {
			return nil, err
		}
	} else if status != "success" {
		return nil, fmt.Errorf("unknown authentication status %s", status)
	}

	return &eTradeAuthenticationStatus{
		authorizationUrl: authorizationUrl, jsonMap: authenticationStatusResponseMap,
	}, nil
}

func (e *eTradeAuthenticationStatus) NeedAuthorization() bool {
	return e.authorizationUrl != ""
}

func (e *eTradeAuthenticationStatus) GetAuthorizationUrl() string {
	return e.authorizationUrl
}

func (e *eTradeAuthenticationStatus) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
