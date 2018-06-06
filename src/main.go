package main

import (
	"net/http"
	"log"
	"./route"
	"./model"
	"os"
)

func main() {
	model.InitDB()
	buildRouter()
	port := os.Getenv("PORT_GO_API")
	if len(port) == 0 {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func buildRouter() {
	for _, element := range route.Routes() {
		http.HandleFunc(element.Path, element.Callback)
	}
}