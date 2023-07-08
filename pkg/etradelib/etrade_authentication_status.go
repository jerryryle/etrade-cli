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
	//   "authorizationUrl": "<verify URL>" (if status=="authorize")
	// }

	// authenticationStatusStatusResponseKey is the status key
	authenticationStatusStatusResponseKey = "status"

	// authenticationStatusAuthorizationUrlResponseKey is the authorization
	// URL key
	authenticationStatusAuthorizationUrlResponseKey = "authorizationUrl"
)

func CreateETradeAuthenticationStatusFromResponse(response []byte) (ETradeAuthenticationStatus, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeAuthenticationStatus(responseMap)
}

func CreateETradeAuthenticationStatus(responseMap jsonmap.JsonMap) (
	ETradeAuthenticationStatus, error,
) {
	status, err := responseMap.GetString(authenticationStatusStatusResponseKey)
	if err != nil {
		return nil, err
	}

	authorizationUrl := ""
	if status == "authorize" {
		if authorizationUrl, err = responseMap.GetString(authenticationStatusAuthorizationUrlResponseKey); err != nil {
			return nil, err
		}
	} else if status != "success" {
		return nil, fmt.Errorf("unknown authentication status %s", status)
	}

	return &eTradeAuthenticationStatus{
		authorizationUrl: authorizationUrl, jsonMap: responseMap,
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
