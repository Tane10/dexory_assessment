package handlers

import (
	"fmt"
	"github.com/tane10/dexory_assignment/utils"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/tane10/dexory_assignment/api"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	cwd, wdErr := utils.GetWorkingDirectory(w)
	if wdErr != nil {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, api.NewCustomError("Invalid request method, only POST is allowed", ""), http.StatusMethodNotAllowed)
		return
	}

	fileType := r.URL.Query().Get("type")

	if fileType == "" {
		http.Error(w, api.NewCustomError("File type is required", ""), http.StatusBadRequest)
		return
	}

	// bitwise left shit operation => 10 * 2^20 => 10485760
	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		http.Error(w,
			api.NewCustomError("File size exceeds 10MB limit", err.Error()),
			http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")

	if contentType != "multipart/form-data" {
		http.Error(w, api.NewCustomError("Unsupported Content-Type", ""), http.StatusUnsupportedMediaType)
	}

	file, fileHeader, err := r.FormFile("file")

	if !strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".json") &&
		!strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".csv") {
		http.Error(w, api.NewCustomError("Unsupported file type, only json or csv is supported.", ""), http.StatusBadRequest)

		file.Close()
		return
	}

	if err != nil {
		http.Error(w,
			api.NewCustomError(fmt.Sprintf("Unable to retrieve %s file", strings.ToUpper(fileType)), err.Error()),
			http.StatusBadRequest)
		return
	}

	defer file.Close()

	filePath := fmt.Sprintf("%s/data/%s", cwd, fileHeader.Filename)

	//TODO: Optimise to add handing duplicated named files

	destinationFile, err := os.Create(filePath)

	if err != nil {
		http.Error(w,
			api.NewCustomError(fmt.Sprintf("Unable to create %s file", strings.ToUpper(fileType)), err.Error()),
			http.StatusInternalServerError)
		return
	}

	defer destinationFile.Close()

	if _, err := io.Copy(destinationFile, file); err != nil {
		http.Error(w, "Unable to save csv file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Files uploaded successfully!"))

}
