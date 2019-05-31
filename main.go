package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "v1"

func main() {
	koDataPath := os.Getenv("KO_DATA_PATH")
	if koDataPath == "" {
		log.Fatal("KO_DATA_PATH variable is not set")
	}

	http.Handle("/", http.FileServer(http.Dir(koDataPath)))

	http.HandleFunc("/version", versionHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func versionHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, version)
	if err != nil {
		log.Printf("error returning version: %v", err)
	}
}
