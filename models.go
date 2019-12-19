package here

type HereMap struct {
	AppID   string
	AppCode string
}

type Request struct {
	ScaleType       string        `json:"scale_type"`
	ViewType        int           `json:"view_type"`
	Point           Coordinates   `json:"point"`
	Center          Coordinates   `json:"center"`
	Address         Address       `json:"address"`
	ShowAddressInfo bool          `json:"show_address_info"`
	Zoom            int           `json:"zoom"`
	Style           string        `json:"style"`
	PIP             int           `json:"pip"`
	POI             []Coordinates `json:"poi"`
	POILabel        []Label       `json:"poi_label"`
	PPI             int           `json:"ppi"`
	TextLabel       []Label       `json:"text_label"`
	FileType        int           `json:"file_type"`
	Width           int           `json:"width"`
	Height          int           `json:"height"`
	Terrain         int           `json:"terrain"`
	HideCompass     bool          `json:"hide_compass"`
	HideCopyright   bool          `json:"hide_copyright"`
	HideCenterDot   bool          `json:"hide_center_dot"`
	HideMarkers     int           `json:"hide_markers"`
	NoCroppedLabels bool          `json:"no_cropped_labels"`
	JpegQuality     int           `json:"jpeg_quality"`
	MaxHits         int           `json:"max_hits"`
	Language        string        `json:"language"`
}

type Address struct {
	Number  string `json:"number"`
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Country string `json:"country"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Label struct {
	Point      Coordinates `json:"point"`
	Label      string      `json:"label"`
	BgColor    string      `json:"bg_color"`
	LabelColor string      `json:"label_color"`
	Size       int         `json:"size"`
}
