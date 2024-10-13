package routes

import (
	"net/http"

	"github.com/tane10/dexory_assignment/internal/api/handlers"
)

func HomeRoutes() {
	// Serve static files from the 'static' directory

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.HomeHandler)

}
