package errors

import (
	"errors"

	"github.com/next-trace/scg-validator/errors/aggregate"
	"github.com/next-trace/scg-validator/errors/single"
)

// Re-export single error types
var (
	ErrValidationFailed = single.ErrValidationFailed
	NewValidationError  = single.NewValidationError
	IsValidationFailed  = single.IsValidationFailed
	ErrRuleNotFound     = errors.New("rule not found")
	ErrInvalidRule      = errors.New("invalid rule")
)

// Re-export aggregate error types
type (
	ValidationErrors = aggregate.Errors
)

// ValidationError represents a single validator error
type ValidationError = single.ValidationError
