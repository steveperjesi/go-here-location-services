package main

import (
	"fmt"

	here "github.com/steveperjesi/go-here-map"
)

func main() {
	hereMap, err := here.NewHereMap("APP_ID", "APP_CODE")
	if err != nil {
		panic(err)
	}

	// Center dot of the map
	center := here.Coordinates{
		Latitude:  33.820015,
		Longitude: -117.922884,
	}

	request := here.Request{
		ScaleType:         "m",
		Center:            center,
		Zoom:              10,
		Style:             "mini",
		UncertaintyRadius: 50,
	}

	mapImage, contentType, requestUrl, err := hereMap.GetMap(&request)
	if err != nil {
		panic(err)
	}

	fmt.Println("Map Image Size", mapImage.Bounds().Max.X, mapImage.Bounds().Max.Y)
	fmt.Println("Map Image Content Type", contentType)
	fmt.Println("Map Image Request URL", requestUrl)
}
