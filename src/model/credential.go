package model

import (
	"fmt"
	"crypto/rand"
	"errors"
)

type Credential struct {
	Username	string `gorm:"primary_key"`
	Password	string
	Token		string
}

func CreateCredential(username string, password string) error {
	credential := Credential{Username: username, Password: password}

	if err := db.Create(&credential).Error; err != nil {
		return err
	}

	return nil
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func UpdateTokenCredential(username string, password string) (string, error) {
	credential := Credential{Username: username, Password: password}

	if db.First(&credential).RecordNotFound() {
		return "", errors.New("credentials not found")
	}

	credential.Token = randToken()

	if err := db.Save(&credential).Error; err != nil {
		return "", err
	}

	return credential.Token, nil
}

func FindToken(token string) bool {
	credential := Credential{Token: token}

	if db.First(&credential).RecordNotFound() {
		return false
	}

	return true
}
