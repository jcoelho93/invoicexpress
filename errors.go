package invoicexpress

import "fmt"

type Error struct {
	Error string `json:"error"`
}

type APIError struct {
	StatusCode int
	Errors     []Error `json:"errors,omitempty"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Errors[0].Error)
}
