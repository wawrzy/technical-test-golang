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
		return shared.ErrMissingField("email")
	} else if len(u.Firstname) == 0 {
		return shared.ErrMissingField("firstname")
	} else if len(u.Lastname) == 0 {
		return shared.ErrMissingField("lastname")
	} else if len(u.Type) == 0 {
		return shared.ErrMissingField("type")
	} else if u.Type != "client" && u.Type != "support"  {
		return errors.New("type must be client or support")
	}
	return nil
}

func userPut(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	var old_email string
	if old_email = r.URL.Query().Get("old_email"); old_email == "" {
		ErrorRequest(w, r, 400, errors.New("old_email query param is missing"))
		return
	}
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	if err := model.UpdateUser(old_email, u.Email, u.Firstname, u.Lastname, u.Type); err != nil {
		ErrorRequest(w, r, 400, err)
	}
}

func userPost(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	if err := model.CreateUser(u.Email, u.Firstname, u.Lastname, u.Type); err != nil {
		ErrorRequest(w, r, 400, err)
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := shared.CheckAuthToken(r); err != nil {
			ErrorRequest(w, r,401, err)
			return
		}
		userPost(w, r)
		return
	} else if r.Method == "PUT" {
		if err := shared.CheckAuthToken(r); err != nil {
			ErrorRequest(w, r,401, err)
			return
		}
		userPut(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}