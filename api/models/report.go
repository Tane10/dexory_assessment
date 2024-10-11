package models

type Locations struct {
	Name             string   `json:"name"`
	Scanned          bool     `json:"scanned"`
	Occupied         bool     `json:"occupied"`
	DetectedBarcodes []string `json:"detected_barcodes"`
}

type Report struct {
	Location         string `json:"location"`
	Scanned          bool   `json:"scanned"`
	Occupied         bool   `json:"occupied"`
	ExpectedItems    string `json:"expectedItems"`
	DetectedBarcodes string `json:"detectedBarcodes"`
	Outcome          string `json:"outcome"`
}

type FileData struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}
