package handlers

import (
	"encoding/json"
	"github.com/tane10/dexory_assignment/internal/api/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var exampleReport = []models.Report{
	{
		Location:         "ZA001A",
		Scanned:          true,
		Occupied:         true,
		ExpectedItems:    "DX9850004338",
		DetectedBarcodes: "DX9850004338",
		Outcome:          "The location was occupied by the expected items",
	},
	{
		Location:         "ZA002A",
		Scanned:          true,
		Occupied:         true,
		ExpectedItems:    "",
		DetectedBarcodes: "DX9850004338, DX9850004348",
		Outcome:          "The location was occupied by an item, but should have been empty",
	},
	{
		Location:         "ZA003A",
		Scanned:          true,
		Occupied:         false,
		ExpectedItems:    "",
		DetectedBarcodes: "",
		Outcome:          "The location was empty, as expected",
	},
	{
		Location:         "ZA004A",
		Scanned:          true,
		Occupied:         true,
		ExpectedItems:    "DX9850004348",
		DetectedBarcodes: "DX9850004348",
		Outcome:          "The location was occupied by the expected items",
	},
	{
		Location:         "ZA005A",
		Scanned:          true,
		Occupied:         false,
		ExpectedItems:    "",
		DetectedBarcodes: "",
		Outcome:          "The location was empty, as expected",
	},
	{
		Location:         "ZA006A",
		Scanned:          true,
		Occupied:         false,
		ExpectedItems:    "",
		DetectedBarcodes: "",
		Outcome:          "The location was empty, as expected",
	},
	{
		Location:         "ZA007A",
		Scanned:          true,
		Occupied:         false,
		ExpectedItems:    "",
		DetectedBarcodes: "",
		Outcome:          "The location was empty, as expected",
	},
	{
		Location:         "ZA008A",
		Scanned:          true,
		Occupied:         false,
		ExpectedItems:    "",
		DetectedBarcodes: "",
		Outcome:          "The location was empty, as expected",
	},
	{
		Location:         "ZA009A",
		Scanned:          true,
		Occupied:         false,
		ExpectedItems:    "",
		DetectedBarcodes: "",
		Outcome:          "The location was empty, as expected",
	},
	{
		Location:         "ZA010A",
		Scanned:          true,
		Occupied:         false,
		ExpectedItems:    "",
		DetectedBarcodes: "",
		Outcome:          "The location was empty, as expected",
	},
	{
		Location:         "ZA011A",
		Scanned:          true,
		Occupied:         true,
		ExpectedItems:    "",
		DetectedBarcodes: "",
		Outcome:          "The location was occupied by an item, but should have been empty",
	},
}

// Post file:[file_1, file_2]
func TestReportHandler(t *testing.T) {
	os.Setenv("TESTING", "true")
	defer os.Unsetenv("TESTING")

	jsonExampleReport, err := json.Marshal(exampleReport)
	if err != nil {
		log.Fatal("Failed to marshal example report")
	}

	scenarios := []struct {
		name         string
		method       string
		body         *models.ReportRequestBody
		expectedCode int
		expectedBody string
		contentType  string
		uploadFile   bool
		filename     string
	}{
		{
			name:         "Should return method not allowed when not using a POST method",
			method:       http.MethodGet,
			body:         &models.ReportRequestBody{},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Invalid request method, only POST is allowed: \n",
		},
		{
			name:         "Should return bad request when 2 files are not sent in request",
			method:       http.MethodPost,
			body:         &models.ReportRequestBody{},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Please supply 2 files, 1 JSON file and 1 CSV file.: \n",
		},
		{
			name:   "Should return bad request when 2 files supplied are not 1 JSON and 1 CSV",
			method: http.MethodPost,
			body: &models.ReportRequestBody{
				File: []string{"test_1.json", "test_2.json"},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Please supply 2 files, 1 JSON file and 1 CSV file.: \n",
		},
		{
			name:   "Should return a valid report",
			method: http.MethodPost,
			body: &models.ReportRequestBody{
				File: []string{"mini-cust_test.csv", "robot_scan_test.json"},
			},
			expectedCode: http.StatusOK,
			expectedBody: string(jsonExampleReport),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(scenario.body)

			if err != nil {
				log.Fatalln(err.Error())
				return
			}

			req := httptest.NewRequest(scenario.method, "/report", strings.NewReader(string(jsonBody)))

			reqRecorder := httptest.NewRecorder()

			handler := http.HandlerFunc(ReportHandler)
			handler.ServeHTTP(reqRecorder, req)

			// Check the status code
			if reqRecorder.Code != scenario.expectedCode {
				t.Errorf("Expected status code %d, got %d", scenario.expectedCode, reqRecorder.Code)
			}

			if scenario.name != "Should return a valid report" {
				if reqRecorder.Body.String() != scenario.expectedBody {
					t.Errorf("Expected body %q, got %q", scenario.expectedBody, reqRecorder.Body.String())
				}
			} else {
				var expectedRespBody []models.Report
				err = json.Unmarshal([]byte(scenario.expectedBody), &expectedRespBody)
				if err != nil {
					t.Fatalf("Failed to unmarshal expected JSON string: %v", err)
				}

				var actualRespBody models.ReportHandlerResp

				err := json.NewDecoder(reqRecorder.Body).Decode(&actualRespBody)
				if err != nil {
					t.Fatalf("Failed to decode JSON: %v", err)
				}

				actualReports := *actualRespBody.Report

				for i := range expectedRespBody {
					if actualReports[i] != expectedRespBody[i] {
						t.Errorf("Expected report %v, got %v", expectedRespBody[i], actualReports[i])
					}
				}

			}

		})
	}
}
