package controller

import (
	"net/http"
	"fmt"
	"../shared"
)

type NewUser struct {
	Email		string
	Password	string
	Firstname	string
	Lastname	string
	Type		string
}

func (u *NewUser) OK() error {
	if len(u.Email) == 0 {
		return shared.ErrMissingField("username")
	} else if len(u.Password) == 0 {
		return shared.ErrMissingField("password")
	} else if len(u.Firstname) == 0 {
		return shared.ErrMissingField("firstname")
	} else if len(u.Lastname) == 0 {
		return shared.ErrMissingField("lastname")
	} else if len(u.Type) == 0 {
		return shared.ErrMissingField("type")
	}
	return nil
}

func signupPost(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	fmt.Println(u.Email)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		signupPost(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}