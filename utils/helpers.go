package utils

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/tane10/dexory_assignment/internal/api"
)

func GetWorkingDirectory(w *http.ResponseWriter) (string, error) {
	cwd, err := os.Getwd()

	if os.Getenv("TESTING") == "true" {
		steps := filepath.Join("..", "..", "..")
		cwd = filepath.Join(cwd, steps)
	}

	if err != nil {
		http.Error(*w,
			api.NewCustomError("Failed to get working directory", err.Error()),
			http.StatusInternalServerError)
		return "", err
	}
	return cwd, nil
}

func ReportDataRouteHandler(w *http.ResponseWriter) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		http.Error(*w,
			api.NewCustomError("Failed to get working directory", err.Error()),
			http.StatusInternalServerError)
		return "", err
	}

	if os.Getenv("TESTING") == "true" {
		cwd = filepath.Join(cwd, "/test_data/")
		return cwd, nil

	}

	cwd = filepath.Join(cwd, "/data/")
	return cwd, nil

}
