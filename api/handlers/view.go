package handlers

import (
	"fmt"
	"github.com/tane10/dexory_assignment/utils"
	"io"
	"net/http"
	"os"

	"github.com/tane10/dexory_assignment/api"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, api.NewCustomError("Invalid request method, only GET is allowed", ""), http.StatusMethodNotAllowed)
		return
	}
	fileName := r.URL.Query().Get("file")

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

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send file contents: %v", err), http.StatusInternalServerError)
		return
	}

}
