package controller

import (
	"net/http"
	"../shared"
	"../model"
	"errors"
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

func userPut(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	var old_email string
	if old_email = r.URL.Query().Get("old_email"); old_email == "" {
		ErrorRequest(w, r, 400, errors.New("old_email query param is missing"))
	}
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	if err := model.UpdateUser(old_email, u.Email, u.Firstname, u.Lastname, u.Type); err != nil {
		ErrorRequest(w, r, 400, err)
	}
}

func userPost(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	if err := model.CreateUser(u.Email, u.Firstname, u.Lastname, u.Type); err != nil {
		ErrorRequest(w, r, 400, err)
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userPost(w, r)
		return
	} else if r.Method == "PUT" {
		userPut(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}