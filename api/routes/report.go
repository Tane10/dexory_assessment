package routes

import (
	"net/http"

	"github.com/tane10/dexory_assignment/api/handlers"
)

func ReportRoutes() {

	// Post file:[file_1, file_2]
	http.HandleFunc("/report", handlers.ReportHandler)
}
