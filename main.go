package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	port := ":3000"
	log.Printf("http://129.154.60.150%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
