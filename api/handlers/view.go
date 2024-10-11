package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/tane10/dexory_assignment/api/models"
	"github.com/tane10/dexory_assignment/utils"

	"github.com/tane10/dexory_assignment/api"
)

func actionHandler(action string, w http.ResponseWriter, r *http.Request, cwd string, fileName *string, fileData *os.File) {
	fAction := strings.ToLower(action)

	var reportData *[]models.Report

	if fAction != "download" {
		http.Error(w,
			api.NewCustomError("download is the only action allowed ", ""),
			http.StatusInternalServerError)
		return
	}

	noPrefixFileName := strings.TrimPrefix(*fileName, "reports/")

	tmpCsvFileName := fmt.Sprintf("%s.csv", strings.TrimSuffix(noPrefixFileName, ".json"))

	tmpFile, err := os.CreateTemp("", tmpCsvFileName)
	if err != nil {
		http.Error(w,
			api.NewCustomError("Unable to create temporary file", err.Error()),
			http.StatusInternalServerError)
		return
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temp file after sending

	if err := json.NewDecoder(fileData).Decode(&reportData); err != nil {
		http.Error(w,
			api.NewCustomError("Failed to decode JSON", err.Error()),
			http.StatusInternalServerError)
		return
	}

	writer := csv.NewWriter(tmpFile)
	// Flush the writer
	defer writer.Flush()

	header := []string{"Location", "Scanned", "Occupied", "ExpectedItems", "DetectedBarcodes", "Outcome"}

	// Write header
	if err := writer.Write(header); err != nil {
		http.Error(w, api.NewCustomError("Failed to write to CSV", err.Error()), http.StatusInternalServerError)
		return
	}

	// Write each report to the CSV file
	for _, report := range *reportData {
		record := []string{
			report.Location,
			fmt.Sprintf("%v", report.Scanned),
			fmt.Sprintf("%v", report.Occupied),
			report.ExpectedItems,
			report.DetectedBarcodes,
			report.Outcome,
		}
		if err := writer.Write(record); err != nil {
			return
		}
	}

	// TODO: We now have a tmp file downloaded with no data

	// Serve the CSV file for download
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", tmpCsvFileName))
	http.ServeFile(w, r, tmpFile.Name())

}

// /view?file=file.json&action=download
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, api.NewCustomError("Invalid request method, only GET is allowed", ""), http.StatusMethodNotAllowed)
		return
	}
	fileName := r.URL.Query().Get("file")

	action := r.URL.Query().Get("action")

	if fileName == "" {
		http.Error(w, api.NewCustomError("File is required", ""), http.StatusBadRequest)
		return
	}

	cwd, wdErr := utils.GetWorkingDirectory(w)
	if wdErr != nil {
		return
	}

	dir := fmt.Sprintf("%s/data/%s", cwd, fileName)

	file, err := os.Open(dir)
	if err != nil {
		http.Error(w,
			api.NewCustomError("Failed to open file", err.Error()),
			http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if action == "" {
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send file contents: %v", err), http.StatusInternalServerError)
			return
		}
	} else {
		actionHandler(action, w, r, cwd, &fileName, file)

	}

}
