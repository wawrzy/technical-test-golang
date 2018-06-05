package model

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Message struct {
	ID			uint	`gorm:"primary_key;AUTO_INCREMENT"`
	Ticket		uint
	Message		string
	Author		string
}

type MessageArchive struct {
	ID			uint
	Ticket		uint
	Message		string
	Author		string
}

type MessageJson struct {
	ID			uint	`json:"id"`
	Ticket		uint	`json:"ticket"`
	Message		string	`json:"message"`
	Author		string	`json:"author"`
}

func CreateMessage(ticket uint, author string, message string) error {
	newMessage := Message{Author: author, Message: message, Ticket: ticket }

	if err := db.Create(&newMessage).Error; err != nil {
		return err
	}

	updateTicket := Ticket{ID: ticket}
	db.First(&updateTicket)
	updateTicket.Status = "pending reply"

	if err := db.Save(&updateTicket).Error; err != nil {
		return err
	}

	return nil
}

func UpdateMessage(message_id uint, ticket uint, author string, message string) error {
	newMessage := Message{ID: message_id}

	if db.First(&newMessage).RecordNotFound() {
		return errors.New(fmt.Sprintf("message with id %d not found\n", message_id))
	}

	newMessage.Author = author
	newMessage.Ticket = ticket
	newMessage.Message = message

	if err := db.Save(&newMessage).Error; err != nil {
		return err
	}

	return nil
}

func getSingleMessage(messageId uint) (interface{}, error) {
	message := Message{ID: messageId}
	if db.First(&message).RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("message with id %d not found\n", messageId))
	}
	response := MessageJson{ID: message.ID, Author: message.Author, Ticket: message.Ticket, Message: message.Message}

	return response, nil
}

func getTicketMessages(ticketId uint) (interface{}, error) {
	var messages []Message
	var response []MessageJson

	db.Find(&messages, "ticket = ?", ticketId)
	for _, message := range messages {
		singleMessage := MessageJson{ID: message.ID, Author: message.Author, Ticket: message.Ticket, Message: message.Message}
		response = append(response, singleMessage)
	}
	return response, nil
}

func GetMessage(r *http.Request) (interface{}, error) {
	messageId_str := r.URL.Query().Get("message_id")
	ticketId_str := r.URL.Query().Get("ticket_id")

	if len(ticketId_str) != 0 {
		var ticketId uint64
		var err error
		if ticketId, err = strconv.ParseUint(ticketId_str, 10, 64); err != nil {
			return nil, errors.New("ticket_id query param should be an integer")
		}
		return getTicketMessages(uint(ticketId))
	} else if len(messageId_str) != 0 {
		var messageId uint64
		var err error
		if messageId, err = strconv.ParseUint(messageId_str, 10, 64); err != nil {
			return nil, errors.New("ticket_id query param should be an integer")
		}
		return getSingleMessage(uint(messageId))
	}
	return nil, errors.New("query param missing")
}


func ArchiveTicketMessages(ticketId uint) error {
	var messages []Message

	db.Find(&messages, "ticket = ?", ticketId)

	for _, message := range messages {
		messageArchive := MessageArchive{Ticket: ticketId, Author: message.Author, Message: message.Author, ID: message.ID}
		if err := db.Create(&messageArchive).Error; err != nil {
			return err
		}
	}
	if len(messages) > 0 {
		db.Delete(&messages)
	}
	return nil
}