package model

import "errors"

type User struct {
	Email		string `gorm:"primary_key"`
	Firstname	string
	Lastname	string
	Type		string
}

func CreateUser(email string, firstname string, lastname string, role string) error {
	user := User{Email: email, Firstname: firstname, Lastname: lastname, Type: role}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(old_email string, email string, firstname string, lastname string, role string) error {
	user := User{Email: old_email}

	if db.First(&user).RecordNotFound() {
		return errors.New("user with email " + old_email + " not found")
	}

	user.Email = email
	user.Firstname = firstname
	user.Lastname = lastname
	user.Type = role

	if email != old_email {
		db.Delete(User{Email: old_email})
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	} else if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
