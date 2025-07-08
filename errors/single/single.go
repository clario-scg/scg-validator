package single

import (
	"errors"
	"fmt"
)

// ErrValidationFailed is returned when a validator rule fails.
// It's used by validator functions to indicate failure.
var ErrValidationFailed = errors.New("validator failed")

// ValidationError represents a validator error with a custom message.
type ValidationError struct {
	Message string
}

// Error implements the error interface.
func (e ValidationError) Error() string {
	return e.Message
}

// NewValidationError creates a new ValidationError with the given message.
func NewValidationError(format string, args ...interface{}) error {
	return ValidationError{
		Message: fmt.Sprintf(format, args...),
	}
}

// IsValidationError checks if an error is a ValidationError.
func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}

// IsValidationFailed checks if an error is ErrValidationFailed.
func IsValidationFailed(err error) bool {
	return errors.Is(err, ErrValidationFailed)
}
