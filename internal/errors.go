package internal

import (
	"fmt"
)

type PVEError struct {
	Code    int
	Message string
}

func NewPVEError(code int, message string) error {
	return &PVEError{code, message}
}

func (e *PVEError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}
