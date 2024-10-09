package handlers

import "net/http"

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method, only GET is allowed", http.StatusMethodNotAllowed)
		return
	}
}
