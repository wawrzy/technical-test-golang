package controller

import (
	"net/http"
	"../shared"
	"fmt"
)

type Error struct {
	Error string `json:"error"`
}

func ErrorRequest(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		shared.ResponseJSON(w, Error{Error: "404 not found"})
		return
	}
	fmt.Println(err)
	shared.ResponseJSON(w, Error{Error: err.Error()})
}
