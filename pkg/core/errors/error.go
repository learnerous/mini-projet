package errors

import (
	"fmt"
)

const (
	DB        = "DB"
	INTERNALS = "INTERNALS"
)

type CustomError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    string `json:"code"`
	Err     error  `json:"causedBy"`
}

func (e CustomError) Error() string {
	logFormat := "{\"type\":\"%s\",\"message\":\"%s\",\"code\":\"%s\",\"causedBy\":\"%s\"}"
	if e.Err != nil {
		return fmt.Sprintf(logFormat, e.Type, e.Message, e.Code, e.Err.Error())
	}
	return fmt.Sprintf(logFormat, e.Type, e.Message, e.Code, "")
}
