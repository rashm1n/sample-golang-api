package main

import (
	"net/http"
)

func main() {
	// Set up handlers for each resource
	http.HandleFunc("/status200", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // 200
		w.Write([]byte("OK - Status 200"))
	})

	http.HandleFunc("/status400", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte("Bad Request - Status 400"))
	})

	http.HandleFunc("/status500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte("Internal Server Error - Status 500"))
	})

	http.HandleFunc("/status201", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated) // 201
		w.Write([]byte("Created - Status 201"))
	})

	// Start the server on localhost port 8080
	http.ListenAndServe(":8080", nil)
}

