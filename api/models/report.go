package models

type RobotScanData struct {
	Name             string   `json:"name"`
	Scanned          bool     `json:"scanned"`
	Occupied         bool     `json:"occupied"`
	DetectedBarcodes []string `json:"detected_barcodes"`
}

type CsvCustomerData struct {
	Location string `json:"location"`
	Item     string `json:"item"`
}
