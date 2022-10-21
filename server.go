package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed misc/html/*
var content embed.FS

func main() {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":80", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
