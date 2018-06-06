package model

import (
	"errors"
	"fmt"
	"strconv"
	"net/http"
)

type Ticket struct {
	ID		uint	`gorm:"primary_key;AUTO_INCREMENT"`
	Author	string
	Status	string
	Title	string
}

type TicketArchive struct {
	ID		uint
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

func CreateTicket(author string, title string) (interface{}, error) {
	ticket := Ticket{Author: author, Status: "open", Title: title }

	if err := db.Create(&ticket).Error; err != nil {
		return nil, err
	}
	singleTicket := SingleTicket{ID: ticket.ID, Author: author, Status: "open", Title: title}
	return singleTicket, nil
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

	ticket.Status = "closed"

	if err := db.Save(&ticket).Error; err != nil {
		return err
	}

	return nil
}

func GetSingleTicket(ticketId uint) (interface{}, error) {
	ticket := Ticket{ID: ticketId}
	if db.First(&ticket).RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("ticket with id %d not found\n", ticketId))
	}
	response := SingleTicket{ID: ticket.ID, Author: ticket.Author, Title: ticket.Title, Status: ticket.Status}

	var messages []Message
	db.Find(&messages, "ticket = ?", ticketId)
	for _, message := range messages {
		response.Messages = append(
			response.Messages,
			MessageJson{Author: message.Author, Message: message.Message, ID: message.ID, Ticket: message.Ticket})
	}

	return response, nil
}

func getUserTickets(userEmail string) (interface{}, error) {
	user := User{Email: userEmail}
	if db.First(&user).RecordNotFound() {
		return nil, errors.New("user with email " + userEmail + " not found")
	}
	var tickets []Ticket
	var response []SingleTicket
	db.Find(&tickets, "author = ?", userEmail)
	for _, ticket := range tickets {
		signTicket := SingleTicket{ID: ticket.ID, Author: ticket.Author, Title: ticket.Title, Status: ticket.Status}
		var messages []Message
		db.Find(&messages, "ticket = ?", ticket.ID)
		for _, message := range messages {
			signTicket.Messages = append(
				signTicket.Messages,
				MessageJson{Author: message.Author, Message: message.Message, ID: message.ID, Ticket: message.Ticket})
		}
		response = append(response, signTicket)
	}
	return response, nil
}

func getAllTickets() (interface{}, error) {
	var tickets []Ticket
	var response []SingleTicket
	db.Find(&tickets)
	for _, ticket := range tickets {
		signTicket := SingleTicket{ID: ticket.ID, Author: ticket.Author, Title: ticket.Title, Status: ticket.Status}
		var messages []Message
		db.Find(&messages, "ticket = ?", ticket.ID)
		for _, message := range messages {
			signTicket.Messages = append(
				signTicket.Messages,
				MessageJson{Author: message.Author, Message: message.Message, ID: message.ID, Ticket: message.Ticket})
		}
		response = append(response, signTicket)
	}
	return response, nil
}

func GetTicket(r *http.Request) (interface{}, error) {
	ticketId_str := r.URL.Query().Get("ticket_id")
	userEmail := r.URL.Query().Get("user_email")

	if len(ticketId_str) != 0 {
		var ticketId uint64
		var err error
		if ticketId, err = strconv.ParseUint(ticketId_str, 10, 64); err != nil {
			return nil, errors.New("ticket_id query param should be an integer")
		}
		return GetSingleTicket(uint(ticketId))
	} else if len(userEmail) != 0 {
		return getUserTickets(userEmail)
	}
	return getAllTickets()
}

func ArchiveTicket(ticketId uint) error {
	ticket := Ticket{ID: ticketId}
	if db.First(&ticket).RecordNotFound() {
		return errors.New(fmt.Sprintf("ticket with id %d not found", ticketId))
	}
	if ticket.Status != "closed" {
		return errors.New("ticket should be closed")
	}
	archiveTicket := TicketArchive{ID: ticketId, Author: ticket.Author, Title: ticket.Title, Status: "closed"}
	if err := db.Create(&archiveTicket).Error; err != nil {
		return err
	}
	if err := ArchiveTicketMessages(ticketId); err != nil {
		return err
	}
	db.Delete(&ticket)
	return nil
}