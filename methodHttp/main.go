package main

import (
	"fmt"
	"net/http"
)

func helloHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the /hello route!")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "POST request received")
	} else {
		fmt.Fprintf(w, "Use POST method to send data.")
	}
}

//  curl -X POST http://localhost:8080/data
// Invoke-WebRequest -Uri http://localhost:8080/data -Method POST

func main() {

	http.HandleFunc("/hello", helloHander)
	http.HandleFunc("/data", dataHandler)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)

}
