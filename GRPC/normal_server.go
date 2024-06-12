package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

type Response struct {
	Id     string `json:"id"`
	Result string `json:"result"`
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	log.Printf("Received request from client IP: %s", clientIP)
	// Decode the request body into the Request struct
	var reqBody Request
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Process the request
	result := "Processed: " + reqBody.Data

	// Encode the response into JSON format
	resp := Response{Id: reqBody.Id, Result: result}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func main() {
	// Register handlers
	http.HandleFunc("/process", processHandler)

	// Start the server
	log.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
