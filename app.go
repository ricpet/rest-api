package main

import (
	// formatted I/O
	"log"      // logging library
	"net/http" // handler for http requests
)

func apiResponse(w http.ResponseWriter, r *http.Request) {
	// Set the return Content-Type as JSON like before
	w.Header().Set("Content-Type", "application/json")

	// Change the response depending on the method being requested
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET method requested"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func main() {
	// HandleFunc expects a resource in uri path “/” which executes handler function “homePageHandler”.
	// This way if a client sends a request to /, func homePageHandler will get executed.
	http.HandleFunc("/", apiResponse)

	// starts a server listening to port 8080 and nil is a handler.
	// Since we have passed nil as the handler, go uses DefaultServeMux as the default handler.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
