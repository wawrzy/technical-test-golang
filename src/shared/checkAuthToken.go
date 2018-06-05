package shared

import "net/http"
import (
	"../model"
	"errors"
)

func CheckAuthToken(r *http.Request) error {
	token := r.Header.Get("Authorization")

	if len(token) == 0 || !model.FindToken(token) {
		return errors.New("invalid token")
	}

	return nil
}
