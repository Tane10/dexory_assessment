package routes

import (
	"net/http"

	"github.com/tane10/dexory_assignment/api/handlers"
)

func ViewRoutes() {
	// /view?file=file.json&action=download
	http.HandleFunc("/view", handlers.ViewHandler)

}
