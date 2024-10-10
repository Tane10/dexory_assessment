package handlers

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tane10/dexory_assignment/api"
	"github.com/tane10/dexory_assignment/api/models"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, api.NewCustomError("Invalid request method, only POST is allowed", ""), http.StatusMethodNotAllowed)
		return
	}

	var reqBody struct {
		File []string `json:"files"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		http.Error(w, api.NewCustomError("Invalid body", err.Error()), http.StatusBadRequest)
		return
	}

	if len(reqBody.File) != 2 {
		http.Error(w, api.NewCustomError("Only supply 2 files at a time", ""), http.StatusBadRequest)
		return
	}

	var jsonFilename string
	var csvFilename string

	for _, file := range reqBody.File {
		if strings.HasSuffix(strings.ToLower(file), ".json") {
			jsonFilename = file
		}

		if strings.HasSuffix(strings.ToLower(file), ".csv") {
			csvFilename = file
		}
	}

	cwd, wdErr := os.Getwd()

	if wdErr != nil {
		http.Error(w,
			api.NewCustomError("Failed to get working directory", wdErr.Error()),
			http.StatusInternalServerError)
		return
	}

	jsonFile, err := os.Open(fmt.Sprintf("%s/data/%s", cwd, jsonFilename))
	if err != nil {
		http.Error(w,
			api.NewCustomError("Error opening JSON file", err.Error()),
			http.StatusInternalServerError)
		return
	}

	defer jsonFile.Close()

	// opens file, returns *os.File -> read with new reader, can read line-line or in chunks
	csvFile, err := os.Open(fmt.Sprintf("%s/data/%s", cwd, csvFilename))
	if err != nil {
		http.Error(w,
			api.NewCustomError("Error opening CSV file", err.Error()),
			http.StatusInternalServerError)
		return
	}

	defer csvFile.Close()

	// Read CSV line by line
	csvReader := csv.NewReader(csvFile)

	//Parse JSON in chunks
	var jsonLocations []models.Locations
	if err := json.NewDecoder(jsonFile).Decode(&jsonLocations); err != nil {
		http.Error(w,
			api.NewCustomError("Error reading JSON file", err.Error()),
			http.StatusInternalServerError)
		return
	}

	csvData := make(map[string]string)

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			http.Error(w,
				api.NewCustomError("Error reading CSV file", err.Error()),
				http.StatusInternalServerError)
			return
		}

		// CSV File formate is:  LOCATION, ITEM
		location := strings.TrimSpace(record[0])
		item := strings.TrimSpace(record[1])

		csvData[location] = item
	}

	report := genReport(csvData, jsonLocations)

	//Save the file

	id := uuid.New()
	ts := time.Now().Format("02-01-2006_15:04:05")

	//id[:] -> converts uuid to slice
	uuidStr := base64.RawURLEncoding.EncodeToString(id[:])

	reportFilename := fmt.Sprintf("report_%s_%s.json", ts, uuidStr)

	reportFile, err := os.Create(fmt.Sprintf("%s/data/reports/%s", cwd, reportFilename))
	if err != nil {
		http.Error(w,
			api.NewCustomError("Failed to create report", err.Error()),
			http.StatusInternalServerError)
		return
	}

	defer reportFile.Close()

	// Marshal report to JSON with indentation
	reportJSON, err := json.MarshalIndent(report, "", " ")
	if err != nil {
		http.Error(w,
			api.NewCustomError("Failed to formate report", err.Error()),
			http.StatusInternalServerError)
		return
	}

	if _, err := reportFile.Write(reportJSON); err != nil {
		http.Error(w,
			api.NewCustomError("Failed to save report", err.Error()),
			http.StatusInternalServerError)
		return

	}

	// Output the report
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)

}

func genReport(csvData map[string]string, jsonLocation []models.Locations) []map[string]string {
	var report []map[string]string

	for _, loc := range jsonLocation {
		result := map[string]string{
			"Location":         loc.Name,
			"Scanned":          fmt.Sprintf("%v", loc.Scanned),
			"Occupied":         fmt.Sprintf("%v", loc.Occupied),
			"ExpectedItems":    csvData[loc.Name],
			"DetectedBarcodes": strings.Join(loc.DetectedBarcodes, ", "),
		}
		// Description of outcome
		if !loc.Occupied && csvData[loc.Name] == "" {
			result["Outcome"] = "The location was empty, as expected"
		} else if !loc.Occupied && csvData[loc.Name] != "" {
			result["Outcome"] = "The location was empty, but it should have been occupied"
		} else if loc.Occupied && csvData[loc.Name] == strings.Join(loc.DetectedBarcodes, ", ") {
			result["Outcome"] = "The location was occupied by the expected items"
		} else if loc.Occupied && csvData[loc.Name] != strings.Join(loc.DetectedBarcodes, ", ") {
			result["Outcome"] = "The location was occupied by the wrong items"
		} else if loc.Occupied && len(loc.DetectedBarcodes) == 0 {
			result["Outcome"] = "The location was occupied, but no barcode could be identified"
		}

		report = append(report, result)
	}

	return report

}
