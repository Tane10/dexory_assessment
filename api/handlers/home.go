package handlers

import (
	"fmt"
	"github.com/tane10/dexory_assignment/api"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

var templates = template.Must(template.ParseFiles("web/templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	type FileData struct {
		Name string `json:"name"`
		Dir  string `json:"dir"`
	}

	var files []FileData

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
			files = append(files, FileData{Name: d.Name(), Dir: path}) // store the file path
		}
		return nil

	})

	// Render template
	data := struct {
		Files *[]FileData
	}{
		Files: &files,
	}

	templateErr := templates.Execute(w, data)

	if templateErr != nil {
		http.Error(w,
			api.NewCustomError("Unable to load index.html", err.Error()),
			http.StatusBadRequest)
		return
	}
}
