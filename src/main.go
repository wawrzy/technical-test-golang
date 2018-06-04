package main

import (
	"net/http"
	"log"
	"./route"
)

func main() {

	buildRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func buildRouter() {
	for _, element := range route.Routes() {
		http.HandleFunc(element.Path, element.Callback)
	}
}