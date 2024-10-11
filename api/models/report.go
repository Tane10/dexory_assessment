package models

type Locations struct {
	Name             string   `json:"name"`
	Scanned          bool     `json:"scanned"`
	Occupied         bool     `json:"occupied"`
	DetectedBarcodes []string `json:"detected_barcodes"`
}

type FileData struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}
