package controller

import (
	"net/http"
	"../shared"
	"../model"
)

type Credential struct {
	Username	string
	Password	string
}


type Response struct {
	Token		string `json:"token"`
}


func (u *Credential) OK() error {
	if len(u.Username) == 0 {
		return shared.ErrMissingField("username")
	} else if len(u.Password) == 0 {
		return shared.ErrMissingField("password")
	}
	return nil
}


func signinPost(w http.ResponseWriter, r *http.Request) {
	var u Credential
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	var token string
	var err error

	if token, err = model.UpdateTokenCredential(u.Username, u.Password); err != nil {
		ErrorRequest(w, r, 400, err)
	}

	shared.ResponseJSON(w, Response{Token: token})
}

func Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		signinPost(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}