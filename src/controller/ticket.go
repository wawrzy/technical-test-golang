package controller

import (
	"net/http"
	"../shared"
	"../model"
	"errors"
	"strconv"
)

type NewTicket struct {
	Author	string
	Title	string
}

type CloseTicket struct {
	ID	uint
}

type ArchiveTicket struct {
	ID	uint
}

type UpdateTicket struct {
	Author	string
	Title	string
	Status	string
}

func (u *NewTicket) OK() error {
	if len(u.Author) == 0 {
		return shared.ErrMissingField("author")
	} else if len(u.Title) == 0 {
		return shared.ErrMissingField("title")
	}
	return nil
}

func (u *CloseTicket) OK() error {
	if u.ID == 0 {
		return shared.ErrMissingField("id")
	}
	return nil
}

func (u *ArchiveTicket) OK() error {
	if u.ID == 0 {
		return shared.ErrMissingField("id")
	}
	return nil
}

func (u *UpdateTicket) OK() error {
	if len(u.Author) == 0 {
		return shared.ErrMissingField("author")
	} else if len(u.Title) == 0 {
		return shared.ErrMissingField("title")
	} else if len(u.Status) == 0 {
		return shared.ErrMissingField("status")
	} else if u.Status != "closed" && u.Status != "pending reply" && u.Status != "open" {
		return errors.New("status must be closed / pending reply / open")
	}
	return nil
}

func ticketPost(w http.ResponseWriter, r *http.Request) {
	var u NewTicket
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	var err error
	var ticket interface{}
	if ticket, err = model.CreateTicket(u.Author, u.Title); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	shared.ResponseJSON(w, ticket)
}

func ticketClose(w http.ResponseWriter, r *http.Request) {
	var u CloseTicket
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	if err := model.CloseTicket(u.ID); err != nil {
		ErrorRequest(w, r, 400, err)
	}
}

func ticketPut(w http.ResponseWriter, r *http.Request) {
	var u UpdateTicket
	var ticket_id_str string
	var ticket_id uint64
	var err error

	if ticket_id_str = r.URL.Query().Get("ticket_id"); ticket_id_str == "" {
		ErrorRequest(w, r, 400, errors.New("ticket_id query param is missing"))

		return
	}
	if ticket_id, err = strconv.ParseUint(ticket_id_str, 10, 64); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	var ticket interface{}
	if ticket, err = model.UpdateTicket(uint(ticket_id), u.Author, u.Status, u.Title); err != nil {
		ErrorRequest(w, r, 400, err)
	}
	shared.ResponseJSON(w, ticket)
}

func ticketArchive(w http.ResponseWriter, r *http.Request) {
	var u ArchiveTicket
	if err := shared.DecodeJSON(r, &u); err != nil {
		ErrorRequest(w, r, 400, err)
		return
	}
	if err := model.ArchiveTicket(u.ID); err != nil {
		ErrorRequest(w, r, 400, err)
	}
}

func ticketGet(w http.ResponseWriter, r *http.Request) {
	var err error
	var response interface{}
	if response, err = model.GetTicket(r); err != nil {
		ErrorRequest(w, r,400, err)
	} else {
		shared.ResponseJSON(w, response)
	}
}

func Ticket(w http.ResponseWriter, r *http.Request) {
	if err := shared.CheckAuthToken(r); err != nil {
		ErrorRequest(w, r,401, err)
		return
	}
	switch r.Method {
		case "POST":
			ticketPost(w, r)
			break
		case "PUT":
			ticketPut(w, r)
			break
		case "GET":
			ticketGet(w, r)
			break
		default:
			ErrorRequest(w, r,404, nil)
	}
}

func TicketClose(w http.ResponseWriter, r *http.Request) {
	if err := shared.CheckAuthToken(r); err != nil {
		ErrorRequest(w, r,401, err)
		return
	}
	if r.Method == "POST" {
		ticketClose(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}

func TicketArchive(w http.ResponseWriter, r *http.Request) {
	if err := shared.CheckAuthToken(r); err != nil {
		ErrorRequest(w, r,401, err)
		return
	}
	if r.Method == "POST" {
		ticketArchive(w, r)
		return
	}
	ErrorRequest(w, r,404, nil)
}