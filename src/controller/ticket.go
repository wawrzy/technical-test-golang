package controller

import (
	"net/http"
	"../shared"
	"../model"
)

type NewTicket struct {
	Author	string
	Title	string
	Status	string
}

func (u *NewTicket) OK() error {
	if len(u.Author) == 0 {
		return shared.ErrMissingField("author")
	} else if len(u.Title) == 0 {
		return shared.ErrMissingField("title")
	} else if len(u.Status) == 0 {
		return shared.ErrMissingField("status")
	}
	return nil
}

func ticketPost(w http.ResponseWriter, r *http.Request) {
	var u NewTicket
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	if err := model.CreateTicket(u.Author, u.Status, u.Title); err != nil {
		ErrorRequest(w, r, 400, err)
	}
}

func Ticket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := shared.CheckAuthToken(r); err != nil {
			ErrorRequest(w, r,401, err)
			return
		}
		ticketPost(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}
