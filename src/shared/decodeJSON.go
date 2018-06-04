package shared

import (
	"net/http"
	"encoding/json"
)

type ErrMissingField string

func (e ErrMissingField) Error() string {
	return string(e) + " is required"
}

// ok represents types capable of validating
// themselves.
type ok interface {
	OK() error
}

// decode can be this simple to start with, but can be extended later
// to support different formats and behaviours without changing
// the interface.
func DecodeJSON(r *http.Request, v ok) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.OK()
}
