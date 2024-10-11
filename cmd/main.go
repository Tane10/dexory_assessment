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
	routes.ViewRoutes()
	routes.ReportRoutes()

	// Start the server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
