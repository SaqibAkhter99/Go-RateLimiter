package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"serverResponse"`
}

func backed__handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	endPointName := r.RequestURI
	endPoint := strings.Split(r.RequestURI, "/")
	fmt.Printf("Received request from %s\n", endPointName)
	msg := "Request from " + endPoint[1]
	fmt.Fprint(w, msg)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("User-Agent: %s\n", r.UserAgent())
	fmt.Printf("Accept %s\n", r.Header.Get("Accept"))

}

func backed__healthCheck__handler(w http.ResponseWriter, r *http.Request) {

	response := HealthCheckResponse{
		Status:  "Healthy",
		Message: "Connection Successful",
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

	// Log the request details
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("User-Agent: %s\n", r.UserAgent())
	fmt.Printf("Accept: %s\n", r.Header.Get("Accept"))
}

func main() {
	http.HandleFunc("/limited", backed__handler)
	http.HandleFunc("/unlimited", backed__handler)
	http.HandleFunc("/hc", backed__healthCheck__handler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}

}
