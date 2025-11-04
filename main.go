package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	log.Printf("http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
