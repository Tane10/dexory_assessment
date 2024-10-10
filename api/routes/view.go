package routes

import (
	"github.com/tane10/dexory_assignment/api/handlers"
	"net/http"
)

func ViewRoutes() {
	// /view?file=file.json
	http.HandleFunc("/view", handlers.ViewHandler)
}
