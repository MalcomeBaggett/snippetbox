package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		// Check if the current request URL path exactly matches "/". If it doesn't, use
		// the http.NotFound() function to send a 404 response to the client.
		// Importantly, we then return from the handler. If we don't return the handler
		// would keep executing and also write the "Hello from SnippetBox" message.
		http.NotFound(w, r)
		return

	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.method to check whether the request is using POST or not. Note that
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write(([]byte("Create a new snippet...")))
}

func main() {
	// use the http.NewServerMux() function to initialize a new new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	// use the http.HandleFunc() function to register the home() as the handler
	// Two parameters: the tcp network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use log.fatal() function to log the error message and exit.
	// note any error returned by http.ListenAndServe() is always non-nil.
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
