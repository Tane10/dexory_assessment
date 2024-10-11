package handlers

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/tane10/dexory_assignment/api"
	"github.com/tane10/dexory_assignment/api/models"
	"github.com/tane10/dexory_assignment/utils"
)

var templates = template.Must(template.ParseFiles("web/templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var files []models.FileData
	var reports []models.FileData

	cwd, wdErr := utils.GetWorkingDirectory(w)
	if wdErr != nil {
		// Error handling is already done in the function.
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
				files = append(files, models.FileData{Name: d.Name(), Dir: path}) // store the file path
			} else {
				reports = append(reports, models.FileData{Name: d.Name(), Dir: path}) // store the file path
			}

		}
		return nil

	})

	// Render template
	data := struct {
		Files   []models.FileData
		Reports []models.FileData
	}{
		Files:   files,
		Reports: reports,
	}

	templateErr := templates.Execute(w, data)

	if templateErr != nil {
		http.Error(w,
			api.NewCustomError("Unable to load index.html", err.Error()),
			http.StatusBadRequest)
		return
	}
}
