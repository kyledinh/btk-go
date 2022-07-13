package moxerr

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrResourceNotFound = errors.New("RESOURCE_NOT_FOUND")
	ErrResourceRead     = errors.New("RESOURCE_NOT_READ")
	ErrConversionFormat = errors.New("CONVERSION_FORMAT_FAILED")
	ErrCLIAction        = errors.New("CLI_ACTION_FAILED")
	ErrWriteFile        = errors.New("WRITE_FILE_FAILED")
)

type WrappedError struct {
	Message string
	MoxErr  *error
}

func (we *WrappedError) Error() string {
	return fmt.Sprintf("wrapped message: %s", we.Message)
}

func NewWrappedError(message string, err *error) WrappedError {
	wrappedError := WrappedError{
		Message: message,
		MoxErr:  err,
	}
	return wrappedError
}
