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

func HandleTestReportGen(w *http.ResponseWriter) (string, error) {
	cwd, err := GetWorkingDirectory(w)
	if err != nil {
		return "", err
	}

	if os.Getenv("TESTING") == "true" {
		steps := filepath.Join("..", "..", "..")
		cwd = filepath.Join(cwd, steps)
	}

	return cwd, nil

}
