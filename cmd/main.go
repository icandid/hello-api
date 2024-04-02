package main

import (
	"log"
	"net/http"

	"github.com/icandid/hello-api/handlers/rest"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslationHandler)

	log.Printf("listening on %s/n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
