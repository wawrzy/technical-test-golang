package controller

import (
	"net/http"
	"errors"
	"strconv"
	"../shared"
	"../model"
)

type NewMessage struct {
	Author	string
	Ticket	uint
	Message	string
}

func (u *NewMessage) OK() error {
	if len(u.Author) == 0 {
		return shared.ErrMissingField("author")
	} else if u.Ticket == 0 {
		return shared.ErrMissingField("ticket")
	} else if len(u.Message) == 0 {
		return shared.ErrMissingField("message")
	}
	return nil
}

func messagePost(w http.ResponseWriter, r *http.Request) {
	var u NewMessage
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	var err error
	var message interface{}
	if message, err = model.CreateMessage(u.Ticket, u.Author, u.Message); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	shared.ResponseJSON(w, message)
}


func messagePut(w http.ResponseWriter, r *http.Request) {
	var u NewMessage
	var message_id_str string
	var message_id uint64
	var err error

	if message_id_str = r.URL.Query().Get("message_id"); message_id_str == "" {
		ErrorRequest(w, r, 400, errors.New("message_id query param is missing"))
		return
	}
	if message_id, err = strconv.ParseUint(message_id_str, 10, 64); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	var message interface{}
	if message, err = model.UpdateMessage(uint(message_id), u.Ticket, u.Author, u.Message); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	shared.ResponseJSON(w, message)
}

func messageGet(w http.ResponseWriter, r *http.Request) {
	var err error
	var response interface{}
	if response, err = model.GetMessage(r); err != nil {
		ErrorRequest(w, r,400, err)
	} else {
		shared.ResponseJSON(w, response)
	}
}

func Message(w http.ResponseWriter, r *http.Request) {
	if err := shared.CheckAuthToken(r); err != nil {
		ErrorRequest(w, r,401, err)
		return
	}
	switch r.Method {
		case "POST":
			messagePost(w, r)
			break
		case "PUT":
			messagePut(w, r)
			break
		case "GET":
			messageGet(w, r)
			break
		default:
			ErrorRequest(w, r,404, nil)
	}
}
