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

	"github.com/tane10/dexory_assignment/utils"

	"github.com/google/uuid"
	"github.com/tane10/dexory_assignment/internal/api"
	"github.com/tane10/dexory_assignment/internal/api/models"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, api.NewCustomError("Invalid request method, only POST is allowed", ""), http.StatusMethodNotAllowed)
		return
	}

	var reqBody *models.ReportRequestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		http.Error(w, api.NewCustomError("Invalid body", err.Error()), http.StatusBadRequest)
		return
	}

	if len(reqBody.File) != 2 {
		http.Error(w, api.NewCustomError("Please supply 2 files, 1 JSON file and 1 CSV file.", ""), http.StatusBadRequest)
		return
	}

	var jsonFilename string
	var csvFilename string

	jsonFileCount := 0
	csvFileCount := 0

	for _, file := range reqBody.File {
		file = strings.ToLower(file)

		switch {
		case strings.HasSuffix(file, ".json"):
			jsonFilename = file
			jsonFileCount++
		case strings.HasSuffix(file, ".csv"):
			csvFilename = file
			csvFileCount++
		}
	}

	if jsonFileCount != 1 && csvFileCount != 1 {
		http.Error(w, api.NewCustomError("Please supply 2 files, 1 JSON file and 1 CSV file.", ""),
			http.StatusBadRequest)
		return

	}

	cwd, wdErr := utils.ReportDataRouteHandler(&w)
	if wdErr != nil {
		return
	}

	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", cwd, jsonFilename))
	if err != nil {
		http.Error(w,
			api.NewCustomError("Error opening JSON file", err.Error()),
			http.StatusInternalServerError)
		return
	}

	defer jsonFile.Close()

	// opens file, returns *os.File -> read with new reader, can read line-line or in chunks
	csvFile, err := os.Open(fmt.Sprintf("%s/%s", cwd, csvFilename))
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
	var jsonLocations *[]models.Locations
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

	reportFile, err := os.Create(fmt.Sprintf("%s/reports/%s", cwd, reportFilename))
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
			api.NewCustomError("Failed to format report", err.Error()),
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
	json.NewEncoder(w).Encode(&models.ReportHandlerResp{
		Report:   report,
		Filename: reportFilename,
	})

}

func genReport(csvData map[string]string, jsonLocation *[]models.Locations) *[]models.Report {
	var report []models.Report

	for _, loc := range *jsonLocation {
		expectedItem := csvData[loc.Name]
		detectedItems := strings.Join(loc.DetectedBarcodes, ", ")

		result := models.Report{
			Location:         loc.Name,
			Scanned:          loc.Scanned,
			Occupied:         loc.Occupied,
			ExpectedItems:    csvData[loc.Name],
			DetectedBarcodes: strings.Join(loc.DetectedBarcodes, ", "),
		}

		switch {
		case loc.Occupied && expectedItem == "":
			result.Outcome = "The location was occupied by an item, but should have been empty"

		case loc.Occupied && expectedItem != detectedItems && detectedItems != "":
			result.Outcome = "The location was occupied by the wrong items"

		case loc.Occupied && expectedItem == detectedItems:
			result.Outcome = "The location was occupied by the expected items"

		case !loc.Occupied && expectedItem != "":
			result.Outcome = "The location was empty, but it should have been occupied"

		case !loc.Occupied && expectedItem == "":
			result.Outcome = "The location was empty, as expected"

		case loc.Occupied && len(loc.DetectedBarcodes) == 0:
			result.Outcome = "The location was occupied, but no barcode could be identified"

		default:
			result.Outcome = "Unexpected condition"
		}

		report = append(report, result)
	}

	return &report

}
