package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeAlertDetails interface {
	GetId() int64
	AsJsonMap() jsonmap.JsonMap
}

type eTradeAlertDetails struct {
	id      int64
	jsonMap jsonmap.JsonMap
}

const (
	// The alert detail response JSON looks like this:
	// {
	//   "AlertDetailsResponse": {
	//     "id": 1234
	//   }
	// }

	// alertDetailsAlertDetailsResponseKey is the key for the alert details map
	alertDetailsAlertDetailsResponseKey = "alertDetailsResponse"

	// alertDetailsAlertIdResponseKey is the key for the alert ID
	alertDetailsAlertIdResponseKey = "id"
)

func CreateETradeAlertDetailsFromResponse(response []byte) (ETradeAlertDetails, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeAlertDetails(responseMap)
}

func CreateETradeAlertDetails(responseMap jsonmap.JsonMap) (ETradeAlertDetails, error) {
	var err error
	// Flatten the response by removing the "alertDetailsResponse" level
	responseMap, err = responseMap.GetMap(alertDetailsAlertDetailsResponseKey)
	if err != nil {
		return nil, err
	}

	alertId, err := responseMap.GetInt(alertDetailsAlertIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeAlertDetails{
		id:      alertId,
		jsonMap: responseMap,
	}, nil
}

func (e *eTradeAlertDetails) GetId() int64 {
	return e.id
}

func (e *eTradeAlertDetails) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
