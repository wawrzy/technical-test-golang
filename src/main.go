package main

import (
	"net/http"
	"log"
	"./route"
	"./model"
)

func main() {
	model.InitDB()
	buildRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func buildRouter() {
	for _, element := range route.Routes() {
		http.HandleFunc(element.Path, element.Callback)
	}
}