package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"time"
)

//go:embed static/*
var staticFiles embed.FS

type HelloResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	// Create sub filesystem to serve static files from root
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files
	http.Handle("/", http.FileServer(http.FS(staticFS)))

	// API endpoint
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := HelloResponse{
			Message:   "Hello from nano-golang!",
			Timestamp: time.Now(),
		}
		json.NewEncoder(w).Encode(response)
	})

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := ":8080"
	log.Printf("Server starting on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}