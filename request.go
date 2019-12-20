package here

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/image/bmp"
)

func (h *HereMap) GetMap(request *Request) (image.Image, string, error) {
	endpoint, _ := url.Parse(HereMapViewBaseAPIURLv1)
	params := endpoint.Query()

	if h.AppID != "" {
		params.Set("app_id", h.AppID)
		params.Set("app_code", h.AppCode)
	}

	if request.Terrain > 0 {
		params.Set("t", strconv.Itoa(request.Terrain))
	}

	if request.Width > 0 {
		params.Set("w", strconv.Itoa(request.Width))
	}

	if request.Height > 0 {
		params.Set("h", strconv.Itoa(request.Height))
	}

	if request.ViewType > 0 {
		params.Set("vt", strconv.Itoa(request.ViewType))
	}

	if request.Zoom > 0 {
		params.Set("z", strconv.Itoa(request.Zoom))
	}

	if request.ScaleType != "" {
		params.Set("sb", request.ScaleType)
	}

	if request.Style != "" {
		params.Set("style", request.Style)
	}

	if len(request.POI) > 0 {
		list := []string{}
		for i, _ := range request.POI {
			list = append(list, Float64ToString(request.POI[i].Latitude))
			list = append(list, Float64ToString(request.POI[i].Longitude))
		}
		poi := strings.Join(list, ",")
		params.Set("poi", poi)
	}

	// Text labels
	if len(request.TextLabel) > 0 {
		for i, _ := range request.TextLabel {
			lat := Float64ToString(request.TextLabel[i].Point.Latitude)
			lon := Float64ToString(request.TextLabel[i].Point.Longitude)
			latLong := lat + "," + lon
			labelData := latLong + ";" + request.TextLabel[i].Label + ";" + request.TextLabel[i].BgColor + ";" + request.TextLabel[i].LabelColor + ";" + strconv.Itoa(request.TextLabel[i].Size)
			field := "tx" + strconv.Itoa(i)
			params.Set(field, labelData)
		}
	}

	if len(request.POILabel) > 0 {
		for i, _ := range request.POILabel {
			lat := Float64ToString(request.POILabel[i].Point.Latitude)
			lon := Float64ToString(request.POILabel[i].Point.Longitude)
			latLong := lat + "," + lon
			labelData := latLong + ";" + request.POILabel[i].BgColor + ";" + request.POILabel[i].LabelColor + ";" + strconv.Itoa(request.POILabel[i].Size) + ";" + request.POILabel[i].Label
			field := "poix" + strconv.Itoa(i)
			params.Set(field, labelData)
		}
	}

	if request.JpegQuality > 0 {
		params.Set("q", strconv.Itoa(request.JpegQuality))
	}

	if request.PPI > 0 {
		params.Set("ppi", strconv.Itoa(request.PPI))
	}

	if request.HideMarkers > 0 {
		params.Set("nomrk", strconv.Itoa(request.HideMarkers))
	}

	if request.PIP > 0 {
		params.Set("pip", strconv.Itoa(request.PIP))
	}

	if request.HideCompass {
		params.Set("nocmp", "1")
	}

	if request.HideCenterDot {
		params.Set("nodot", "1")
	}

	if request.NoCroppedLabels {
		params.Set("nocrop", "1")
	}

	if request.HideCopyright {
		params.Set("nocp", "1")
	}

	// File type
	if request.FileType != "" {
		params.Set("f", request.FileType)
	}

	if request.MaxHits > 0 {
		params.Set("maxhits", strconv.Itoa(request.MaxHits))
	}

	if request.Language != "" {
		params.Set("ml", request.Language)
	}

	if request.Point.Latitude != 0 {
		params.Set("lat", Float64ToString(request.Point.Latitude))
	}

	if request.Point.Longitude != 0 {
		params.Set("lon", Float64ToString(request.Point.Longitude))
	}

	// Map center
	if request.Center.Latitude != 0 && request.Center.Longitude != 0 {
		center := Float64ToString(request.Center.Latitude) + "," + Float64ToString(request.Center.Longitude)
		params.Set("c", center)
	}

	// Check for address request
	if request.Address.Number != "" {
		params.Set("n", request.Address.Number)
	}

	if request.Address.Street != "" {
		params.Set("s", request.Address.Street)
	}

	if request.Address.City != "" {
		params.Set("ci", request.Address.City)
	}

	if request.Address.ZipCode != "" {
		params.Set("z", request.Address.ZipCode)
	}

	if request.Address.Country != "" {
		params.Set("co", request.Address.Country)
	}

	if request.ShowAddressInfo {
		params.Set("i", "1")
	}

	// fmt.Println("params", len(params), params)

	// Sending a valid AppID and AppCode with no other params will result in a successul map of Berlin
	if len(params) <= 2 {
		return nil, "", errors.New("Missing Parameters")
	}

	endpoint.RawQuery = params.Encode()

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	fmt.Println("REQUEST", endpoint.String())

	response, err := client.Get(endpoint.String())
	if err != nil {
		return nil, "", err
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, "", errors.New(response.Status)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}
	// fmt.Println("HERECOM MAP RESULT >>>", response.Status, response.Header.Get("Content-Length"), response.Header.Get("Content-Type"), len(data))

	contentTypes := strings.Split(response.Header.Get("Content-Type"), ";")

	for i, _ := range contentTypes {
		if contentTypes[i] == "image/jpeg" {
			resultImage, err := jpeg.Decode(bytes.NewReader(data))
			if err != nil {
				// fmt.Println("ERROR JPEG resultImage", err)
				return nil, "", err
			}
			return resultImage, contentTypes[i], err

		} else if contentTypes[i] == "image/bmp" {
			resultImage, err := bmp.Decode(bytes.NewReader(data))
			if err != nil {
				// fmt.Println("ERROR BMP resultImage", err)
				return nil, "", err
			}
			return resultImage, contentTypes[i], err

		} else if contentTypes[i] == "image/gif" {
			resultImage, err := gif.Decode(bytes.NewReader(data))
			if err != nil {
				// fmt.Println("ERROR GIF resultImage", err)
				return nil, "", err
			}
			return resultImage, contentTypes[i], err

		} else if contentTypes[i] == "image/png" {
			resultImage, err := png.Decode(bytes.NewReader(data))
			if err != nil {
				// fmt.Println("ERROR PNG resultImage", err)
				return nil, "", err
			}
			return resultImage, contentTypes[i], err
		}
	}

	return nil, "", err
}
