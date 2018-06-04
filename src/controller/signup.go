package controller

import (
	"net/http"
	"fmt"
	"../shared"
)

type NewUser struct {
	Username            string
}

func (u *NewUser) OK() error {
	if len(u.Username) == 0 {
		return shared.ErrMissingField("username")
	}
	return nil
}

func signupPost(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	fmt.Println(u.Username)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		signupPost(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}