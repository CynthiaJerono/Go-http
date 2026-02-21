package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Response represents the API response structure
type Response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	GoVersion string    `json:"go_version"`
}

// helloHandler handles GET requests to /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Create response object
	response := Response{
		Message:   "Hello, World! Welcome to Go!",
		Timestamp: time.Now(),
		GoVersion: "1.21+",
	}

	// Send JSON response
	json.NewEncoder(w).Encode(response)
}

// rootHandler handles GET requests to /
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go API Server is running!\n")
	fmt.Fprintf(w, "Visit /hello to see the Hello World endpoint\n")
}

func main() {
	// Register handlers
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Println("Press Ctrl+C to stop the server")

	// Start listening (this blocks)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
