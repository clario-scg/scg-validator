package aggregate

import (
	"fmt"
	"strings"
)

// Errors is a map of field names to error messages.
// It's used to collect and manage multiple validator errors across different fields.
type Errors map[string][]string

// Add adds an error message for a field.
func (e Errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns all error messages for a field, or an empty slice if there are no errors.
func (e Errors) Get(field string) []string {
	if messages, ok := e[field]; ok {
		return messages
	}
	return []string{}
}

// First returns the first error message for a field, or an empty strings if there are no errors.
func (e Errors) First(field string) string {
	if messages := e.Get(field); len(messages) > 0 {
		return messages[0]
	}
	return ""
}

// Has returns true if there are any error messages for a field.
func (e Errors) Has(field string) bool {
	_, ok := e[field]
	return ok
}

// Error implements the error interface, allowing Errors to be returned as an error.
func (e Errors) Error() string {
	var builder strings.Builder
	builder.WriteString("validator failed with the following errors:\n")
	for field, messages := range e {
		builder.WriteString(fmt.Sprintf("  - field '%s': %s\n", field, strings.Join(messages, ", ")))
	}
	return builder.String()
}
