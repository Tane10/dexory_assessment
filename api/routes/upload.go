package routes

import (
	"net/http"

	"github.com/tane10/dexory_assignment/api/handlers"
)

func UploadRoutes() {
	// /upload?type=csv || json
	http.HandleFunc("/upload", handlers.UploadHandler)
}
