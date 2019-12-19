package here

import (
	"errors"
)

const (
	HereMapViewBaseAPIURLv1 = "https://image.maps.api.here.com/mia/1.6/mapview"
)

func NewHereMap(appId, appCode string) (*HereMap, error) {

	if appId == "" || appCode == "" {
		return nil, errors.New("appId and/or appCode is missing")
	}

	newHeremap := new(HereMap)
	newHeremap.AppID = appId
	newHeremap.AppCode = appCode

	return newHeremap, nil
}
