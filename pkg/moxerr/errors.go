package moxerr

import (
	"errors"
	"fmt"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrResourceRead     = errors.New("resource not read")
	ErrConversionFormat = errors.New("format conversion failed")
	ErrCLIAction        = errors.New("cli action failed to execute")
)

type WrappedError struct {
	Message string
	MoxErr  error
}

func (we *WrappedError) Error() string {
	return fmt.Sprintf("message: %s", we.Message)
}

func NewWrappedError(message string, err error) *WrappedError {
	return &WrappedError{
		Message: message,
		MoxErr:  err,
	}
}
