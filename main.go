package main

import (
	"consume-api/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", api.GetApi)
	router.HandleFunc("/{id}", api.GetApiId)

	log.Fatal(http.ListenAndServe(":8000", router))
}
