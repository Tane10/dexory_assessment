package api

import "fmt"

type CustomError struct {
	Message string
	Err     string
}

func NewCustomError(message string, err string) string {
	e := &CustomError{
		Message: message,
		Err:     err,
	}

	return fmt.Sprintf("%s: %v", e.Message, e.Err)

}
