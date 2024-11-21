package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
