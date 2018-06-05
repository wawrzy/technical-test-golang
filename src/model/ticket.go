package model

type Ticket struct {
	ID		uint	`gorm:"primary_key;AUTO_INCREMENT"`
	Author	string
	Status	string
	Title	string
}

func CreateTicket(author string, status string, title string) error {
	ticket := Ticket{Author: author, Status: status, Title: title }

	if err := db.Create(&ticket).Error; err != nil {
		return err
	}

	return nil
}
