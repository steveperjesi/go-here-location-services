package here_test

import (
	"os"
	"testing"

	here "github.com/steveperjesi/go-here-map"
	"github.com/stretchr/testify/assert"
)

const (
	DefaultAppID   = "HERE_MAP_APP_ID"
	DefaultAppCode = "HERE_MAP_APP_CODE"

	DisneylandLatitude     = 33.820015
	DisneylandLongitude    = -117.922884
	ThePizzaPressLatitude  = 33.806125
	ThePizzaPressLongitude = -117.914474
)

func AppID() string {
	if DefaultAppID != "HERE_MAP_APP_ID" {
		return DefaultAppID
	}
	return os.Getenv("HERE_MAP_APP_ID")
}

func AppCode() string {
	if DefaultAppCode != "HERE_MAP_APP_CODE" {
		return DefaultAppCode
	}
	return os.Getenv("HERE_MAP_APP_CODE")
}

func TestHereIfNoAuth(t *testing.T) {
	assert := assert.New(t)

	_, err := here.NewHereMap("", "")
	assert.Error(err)
}

func TestHereIfMissing(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	mapImage, _, _, err := hereMap.GetMap(&here.Request{})
	assert.Error(err)

	assert.Nil(mapImage)
}

func TestHereLatLongWithZoom(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Point: coords,
		Zoom:  3,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereLatLongWithPIP(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Point: coords,
		PIP:   10,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereLatLongForceJpegWithQuality(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, contentType, _, err := hereMap.GetMap(&here.Request{
		Point:       coords,
		FileType:    "1",
		JpegQuality: 90,
	})
	assert.NoError(err)

	assert.Equal(contentType, "image/jpeg")

	assert.NotNil(mapImage)
}

func TestHereLatLongForcePng(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, contentType, _, err := hereMap.GetMap(&here.Request{
		Point:    coords,
		FileType: "0",
	})
	assert.NoError(err)

	assert.Equal(contentType, "image/png")

	assert.NotNil(mapImage)
}

func TestHereLatLongForceGif(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, contentType, _, err := hereMap.GetMap(&here.Request{
		Point:    coords,
		FileType: "2",
	})
	assert.NoError(err)

	assert.Equal(contentType, "image/gif")

	assert.NotNil(mapImage)
}

func TestHereLatLongForceBmp(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, contentType, _, err := hereMap.GetMap(&here.Request{
		Point:    coords,
		FileType: "3",
	})
	assert.NoError(err)

	assert.Equal(contentType, "image/bmp")

	assert.NotNil(mapImage)
}

func TestHereCenterWithStyle(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Center: coords,
		Style:  "mini",
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereTerrainWithViewType(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Center:   coords,
		Terrain:  8,
		ViewType: 1,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereHeightAndWidth(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Center: coords,
		Height: 320,
		Width:  320,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereScaleTypeWithMaxHits(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Center:    coords,
		ScaleType: "m",
		MaxHits:   1,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHerePOIWithHiddenMarkers(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords1 := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	coords2 := here.Coordinates{
		Latitude:  ThePizzaPressLatitude,
		Longitude: ThePizzaPressLongitude,
	}

	poi := []here.Coordinates{coords1, coords2}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		POI:         poi,
		HideMarkers: 1,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHerePOIWithHiddenCompass(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords1 := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	coords2 := here.Coordinates{
		Latitude:  ThePizzaPressLatitude,
		Longitude: ThePizzaPressLongitude,
	}

	poi := []here.Coordinates{coords1, coords2}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		POI:         poi,
		HideCompass: true,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHerePOIWithPPI(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords1 := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	coords2 := here.Coordinates{
		Latitude:  ThePizzaPressLatitude,
		Longitude: ThePizzaPressLongitude,
	}

	poi := []here.Coordinates{coords1, coords2}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		POI: poi,
		PPI: 500,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereCenterWithHideCenterDot(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Center:        coords,
		HideCenterDot: true,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereCenterWithNoCroppedLabels(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Center:          coords,
		NoCroppedLabels: true,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereCenterWithHideCopyright(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Center:        coords,
		HideCopyright: true,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereAddressWithShowAddressAndLanguage(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	address := here.Address{
		Number:  "3900",
		Street:  "S Las Vegas Blvd",
		City:    "Las Vegas",
		ZipCode: "89119",
		Country: "US",
	}
	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Address:         address,
		ShowAddressInfo: true,
		Language:        "spa",
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHerePointsOfInterest(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords1 := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	coords2 := here.Coordinates{
		Latitude:  ThePizzaPressLatitude,
		Longitude: ThePizzaPressLongitude,
	}

	poi := []here.Coordinates{coords1, coords2}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		POI: poi,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHerePointsOfInterestLabels(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords1 := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	coords2 := here.Coordinates{
		Latitude:  ThePizzaPressLatitude,
		Longitude: ThePizzaPressLongitude,
	}

	label1 := here.Label{
		Point:      coords1,
		Label:      "DSNY",
		BgColor:    "blue",
		LabelColor: "white",
		Size:       20,
	}

	label2 := here.Label{
		Point:      coords2,
		Label:      "PZZA",
		BgColor:    "red",
		LabelColor: "white",
		Size:       15,
	}

	labels := []here.Label{label1, label2}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		POILabel: labels,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereTextLabels(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords1 := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}
	coords2 := here.Coordinates{
		Latitude:  ThePizzaPressLatitude,
		Longitude: ThePizzaPressLongitude,
	}

	label1 := here.Label{
		Point:      coords1,
		Label:      "DSNY",
		BgColor:    "black",
		LabelColor: "white",
		Size:       20,
	}

	label2 := here.Label{
		Point:      coords2,
		Label:      "PZZA",
		BgColor:    "green",
		LabelColor: "white",
		Size:       15,
	}

	labels := []here.Label{label1, label2}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		TextLabel: labels,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}

func TestHereRequestUrl(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	_, _, requestUrl, err := hereMap.GetMap(&here.Request{})
	assert.Error(err)

	assert.NotNil(requestUrl)
}

func TestHereUncertaintyRadius(t *testing.T) {
	assert := assert.New(t)

	hereMap, err := here.NewHereMap(AppID(), AppCode())
	assert.NoError(err)

	coords := here.Coordinates{
		Latitude:  DisneylandLatitude,
		Longitude: DisneylandLongitude,
	}

	mapImage, _, _, err := hereMap.GetMap(&here.Request{
		Center:            coords,
		UncertaintyRadius: 50,
	})
	assert.NoError(err)

	assert.NotNil(mapImage)
}
