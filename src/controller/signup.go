package controller

import (
	"net/http"
	"../shared"
	"../model"
	"fmt"
)

type NewCredential struct {
	Username	string
	Password	string
}

func (u *NewCredential) OK() error {
	if len(u.Username) == 0 {
		return shared.ErrMissingField("username")
	} else if len(u.Password) == 0 {
		return shared.ErrMissingField("password")
	}
	return nil
}


func signupPost(w http.ResponseWriter, r *http.Request) {
	var u NewCredential
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	if err := model.CreateCredential(u.Username, u.Password); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	fmt.Fprintln(w, "Success")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		signupPost(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}