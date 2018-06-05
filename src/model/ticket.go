package model

import (
	"errors"
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
)

type Ticket struct {
	ID		uint	`gorm:"primary_key;AUTO_INCREMENT"`
	Author	string
	Status	string
	Title	string
}

type SingleTicket struct {
	ID			uint			`json:"id"`
	Author		string			`json:"author"`
	Status		string			`json:"status"`
	Title		string			`json:"title"`
	Messages	[]MessageJson	`json:"messages"`
}

func CreateTicket(author string, title string) error {
	ticket := Ticket{Author: author, Status: "open", Title: title }

	if err := db.Create(&ticket).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTicket(ticket_id uint, author string, status string, title string) error {
	ticket := Ticket{ID: ticket_id}

	if db.First(&ticket).RecordNotFound() {
		return errors.New(fmt.Sprintf("ticket with id %d not found\n", ticket_id))
	}

	ticket.Author = author
	ticket.Status = status
	ticket.Title = title

	if err := db.Save(&ticket).Error; err != nil {
		return err
	}

	return nil
}

func CloseTicket(ticket_id uint) error {
	ticket := Ticket{ID: ticket_id}

	if db.First(&ticket).RecordNotFound() {
		return errors.New(fmt.Sprintf("ticket with id %d not found\n", ticket_id))
	}

	ticket.Status = "close"

	if err := db.Save(&ticket).Error; err != nil {
		return err
	}

	return nil
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	jData, err := json.Marshal(data)
	if err != nil {
		panic(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
func getSingleTicket(ticketId uint) (interface{}, error) {
	ticket := Ticket{ID: ticketId}
	if db.First(&ticket).RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("ticket with id %d not found\n", ticketId))
	}
	response := SingleTicket{ID: ticket.ID, Author: ticket.Author, Title: ticket.Author, Status: ticket.Status}

	var messages []Message
	db.Find(&messages, "ticket = ?", ticketId)
	for _, message := range messages {
		response.Messages = append(
			response.Messages,
			MessageJson{Author: message.Author, Message: message.Message, ID: message.ID, Ticket: message.Ticket})
	}

	return response, nil
}

func GetTicket(r *http.Request) (interface{}, error) {
	ticketId_str := r.URL.Query().Get("ticket_id")
	//userEmail := url.Query().Get("user_email")

	if len(ticketId_str) != 0 {
		var ticketId uint64
		var err error
		if ticketId, err = strconv.ParseUint(ticketId_str, 10, 64); err != nil {
			return nil, errors.New("ticket_id query param should be an integer")
		}
		return getSingleTicket(uint(ticketId))
	}
	return nil, nil
}