package main

import "log"

func main() {
	log.Println("This is a log message")
	log.Println("INFO: this is an info message")
	log.Printf("WARNING: this is a warning message")
	log.Fatalf("ERROR: this is an error message")
}
