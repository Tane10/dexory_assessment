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

// Post file:[file_1, file_2]
func TestReportHandler(t *testing.T) {
	os.Setenv("TESTING", "true")
	defer os.Unsetenv("TESTING")

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
			name:   "Should return bad request when 2 files supplied are not 1 JSON and 1 CSV",
			method: http.MethodPost,
			body: &models.ReportRequestBody{
				File: []string{"mini-cust_test.csv", "robot_bad_scan_test.json"},
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Please supply 2 files, 1 JSON file and 1 CSV file.: \n",
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

			// Check the response body
			if reqRecorder.Body.String() != scenario.expectedBody {
				t.Errorf("Expected body %q, got %q", scenario.expectedBody, reqRecorder.Body.String())
			}
		})
	}
}
