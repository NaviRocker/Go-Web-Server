package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve files in the public folder
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Handle requests to the /message endpoint
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		// Set the response header
		w.Header().Set("Content-Type", "text/plain")

		// Write the response message
		w.Write([]byte("This is a message from the server!"))
	})

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
