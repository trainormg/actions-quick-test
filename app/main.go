package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type WhereAmIResponse struct {
	Name      string              `json:"name"`
	Zone      string              `json:"zone"`
	Timestamp string              `json:"timestamp"`
	Headers   http.Header         `json:"headers"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		resp := WhereAmIResponse{
			Name:      hostname,
			Zone:      os.Getenv("ZONE"),
			Timestamp: time.Now().Format(time.RFC3339),
			Headers:   r.Header,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	log.Printf("Starting whereami test server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
