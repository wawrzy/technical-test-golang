package controller

import (
	"net/http"
	"fmt"
)

func ErrorRequest(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 not found")
		return
	}
	fmt.Fprint(w, err)
}
