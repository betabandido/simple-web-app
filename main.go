package main

import (
	"fmt"
	"log"
	"net/http"
)

const version = "v1"

func main() {
	http.Handle("/",
		http.StripPrefix("/", http.FileServer(http.Dir("static"))),
	)

	http.HandleFunc("/version", versionHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func versionHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, version)
	if err != nil {
		log.Printf("error returning version: %v", err)
	}
}
