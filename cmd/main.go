package main

// go run ./cmd/main.go

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tane10/dexory_assignment/api/routes"
)

func main() {

	routes.UploadRoutes()
	routes.HomeRoutes()

	// http.HandleFunc("/view", viewHandler)     // For JSON uploads
	// http.HandleFunc("/report", reportHandler) // For JSON uploads
	// http.HandleFunc("/export", exportHandler) // For JSON uploads

	// Start the server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
