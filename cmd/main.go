package main

// go run ./cmd/main.go

import (
	"fmt"
	"log"
	"net/http"
)

// addd routes -> POST upload, GET export, GET view
// func handleUpload(w http.ResponseWriter, r *http.Request) {
// 	// Ensure the request method is POST
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method, only POST is allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Check Content-Type to determine the format
// 	contentType := r.Header.Get("Content-Type")
// 	switch contentType {
// 	case "application/json":
// 		handleJSON(w, r)
// 	case "text/csv":
// 		handleCSV(w, r)
// 	default:
// 		http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
// 	}
// }

// func handleJSON(w http.ResponseWriter, r *http.Request) {
// 	// Ensure the request method is POST
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method, only POST is allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// }

// func view(w http.ResponseWriter, r *http.Request) {
// 	// Ensure the request method is POST
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Invalid request method, only GET is allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// }

func main() {
	// Define routes and associate handler functions

	// DEFINE routes:
	// GET /view/{id} -> viewHandler
	// POST /upload -> uploadHandler
	// GET /report -> reportHandler
	// POST /report -> reportHandler
	// POST /export -> exportHandler
	http.HandleFunc("/view", view) // For JSON uploads
	http.HandleFunc("/view", view) // For JSON uploads

	// Start the server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
