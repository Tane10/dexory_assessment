package handlers

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tane10/dexory_assignment/api"
)

var templates = template.Must(template.ParseFiles("web/templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	type FileData struct {
		Name string `json:"name"`
		Dir  string `json:"dir"`
	}

	var files []FileData
	var reports []FileData

	cwd, wdErr := os.Getwd()

	if wdErr != nil {
		http.Error(w,
			api.NewCustomError("Failed to get working directory", wdErr.Error()),
			http.StatusInternalServerError)
		return
	}

	dir := fmt.Sprintf("%s/data", cwd)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Check if it's a file (not a directory)
		if !d.IsDir() {
			if !strings.Contains(path, "reports") {
				files = append(files, FileData{Name: d.Name(), Dir: path}) // store the file path
			} else {
				reports = append(reports, FileData{Name: d.Name(), Dir: path}) // store the file path
			}

		}
		return nil

	})

	// Render template
	data := struct {
		Files   *[]FileData
		Reports *[]FileData
	}{
		Files:   &files,
		Reports: &reports,
	}

	templateErr := templates.Execute(w, data)

	if templateErr != nil {
		http.Error(w,
			api.NewCustomError("Unable to load index.html", err.Error()),
			http.StatusBadRequest)
		return
	}
}
