package model

import (
	"errors"
	"fmt"
)

type Message struct {
	ID			uint	`gorm:"primary_key;AUTO_INCREMENT"`
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