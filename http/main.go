package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher!")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello users!")
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server!")
}

func main() {

	// root url
	http.HandleFunc("/", helloHandler)

	// users url
	http.HandleFunc("/users", usersHandler)

	// server url
	http.HandleFunc("/server", serverHandler)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)

}
