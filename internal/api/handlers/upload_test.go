package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func prepareFileUpload(filename string) (*bytes.Buffer, string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}

	filePath := fmt.Sprintf("%s/test_data/%s", cwd, filename)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a file field and copy the file content into it
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, "", err
	}

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil

}

func TestUploadHandler(t *testing.T) {
	os.Setenv("TESTING", "true")
	defer os.Unsetenv("TESTING")

	scenarios := []struct {
		name         string
		method       string
		body         string
		fileType     string
		expectedCode int
		expectedBody string
		contentType  string
		uploadFile   bool
		filename     string
	}{
		{
			name:         "Should return method not allowed when not using a POST method",
			method:       http.MethodGet,
			body:         "",
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Invalid request method, only POST is allowed: \n",
		},
		{
			name:         "Should return bad request when file type isn't presented in query params",
			method:       http.MethodPost,
			body:         "",
			fileType:     "",
			expectedCode: http.StatusBadRequest,
			expectedBody: "File type is required: \n",
		},
		//{
		//	//TODO: FIXME
		//	name:         "Should return bad request when file exceeds 10mb limit",
		//	method:       http.MethodPost,
		//	body:         "",
		//	fileType:     "file",
		//	expectedCode: http.StatusBadRequest,
		//	expectedBody: "Unsupported",
		//	uploadFile:   true,
		//	filename:     "robot_scan.json",
		//},
		{
			name:         "Should return unsupported media when content type is not multipart/form-data",
			method:       http.MethodPost,
			body:         "",
			fileType:     "file",
			expectedCode: http.StatusUnsupportedMediaType,
			expectedBody: "Unsupported Content-Type: \n",
			contentType:  "application/json",
		},
		{
			name:         "Should return bad request when file type isn't json or CSV",
			method:       http.MethodPost,
			body:         "",
			fileType:     "file",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Unsupported file type, only json or csv is supported.: \n",
			uploadFile:   true,
			filename:     "robot_scan_test.txt",
		},
		{
			name:         "Should return upload a json file",
			method:       http.MethodPost,
			body:         "",
			fileType:     "file",
			expectedCode: http.StatusOK,
			expectedBody: "Files uploaded successfully!",
			uploadFile:   true,
			filename:     "robot_scan_test.json",
		},
		{
			name:         "Should return upload a CSV file",
			method:       http.MethodPost,
			body:         "",
			fileType:     "file",
			expectedCode: http.StatusOK,
			expectedBody: "Files uploaded successfully!",
			uploadFile:   true,
			filename:     "mini-cust_test.csv",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			url := fmt.Sprintf("/upload?type=%s", scenario.fileType)
			req := httptest.NewRequest(scenario.method, url, strings.NewReader(scenario.body))

			if scenario.uploadFile {
				filename := scenario.filename // Ensure this file exists
				body, contentType, err := prepareFileUpload(filename)
				if err != nil {
					t.Fatalf("Failed to create multipart form file from local file: %v", err)
				}
				req = httptest.NewRequest(scenario.method, url, body)
				req.Header.Set("Content-Type", contentType)
			} else {
				req.Header.Set("Content-Type", scenario.contentType)
			}

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()

			// Call the handler
			handler := http.HandlerFunc(UploadHandler)
			handler.ServeHTTP(rr, req)

			// Check the status code
			if rr.Code != scenario.expectedCode {
				t.Errorf("Expected status code %d, got %d", scenario.expectedCode, rr.Code)
			}

			// Check the response body
			if rr.Body.String() != scenario.expectedBody {
				t.Errorf("Expected body %q, got %q", scenario.expectedBody, rr.Body.String())
			}
		})
	}

}
