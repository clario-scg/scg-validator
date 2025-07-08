// Package contract defines the core interfaces and types for validator
package contract

import (
	"errors"
	"fmt"

	validatorErrors "github.com/next-trace/scg-validator/errors"
)

// Common validator errors
var (
	ErrRuleNotFound = errors.New("rule not found")
	ErrInvalidRule  = errors.New("invalid rule")
	ErrInvalidData  = errors.New("invalid data")
)

// IsValidationFailed checks if an error is a validator failure
func IsValidationFailed(err error) bool {
	return err != nil && err.Error() == validatorErrors.ErrValidationFailed.Error()
}

// NewValidationError creates a new validator error with formatting
func NewValidationError(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
