package utils

import (
	"net/http"
	"os"

	"github.com/tane10/dexory_assignment/internal/api"
)

func GetWorkingDirectory(w http.ResponseWriter) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		http.Error(w,
			api.NewCustomError("Failed to get working directory", err.Error()),
			http.StatusInternalServerError)
		return "", err
	}
	return cwd, nil
}
