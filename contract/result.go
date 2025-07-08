package contract

// Validator defines the main validator interface

// Result represents the outcome of a validator operation
type Result interface {
	// IsValid reports whether validator passed without errors
	IsValid() bool

	// Errors returns all validator errors grouped by field
	Errors() map[string][]string

	// FirstError returns the first validator error, if any
	FirstError() string

	// FieldError returns the first error for a specific field
	FieldError(field string) string

	// HasFieldError reports whether a field has validator errors
	HasFieldError(field string) bool
}

// ValidationErrors is a concrete implementation of Result
type ValidationErrors struct {
	errors map[string][]string
}

// NewValidationErrors creates a new ValidationErrors instance
func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		errors: make(map[string][]string),
	}
}

// AddError adds an error for a specific field
func (ve *ValidationErrors) AddError(field, message string) {
	ve.errors[field] = append(ve.errors[field], message)
}

// IsValid reports whether validator passed without errors
func (ve *ValidationErrors) IsValid() bool {
	return len(ve.errors) == 0
}

// Errors returns all validator errors grouped by field
func (ve *ValidationErrors) Errors() map[string][]string {
	return ve.errors
}

// FirstError returns the first validator error, if any
func (ve *ValidationErrors) FirstError() string {
	for _, fieldErrors := range ve.errors {
		if len(fieldErrors) > 0 {
			return fieldErrors[0]
		}
	}
	return ""
}

// FieldError returns the first error for a specific field
func (ve *ValidationErrors) FieldError(field string) string {
	if errors, exists := ve.errors[field]; exists && len(errors) > 0 {
		return errors[0]
	}
	return ""
}

// HasFieldError reports whether a field has validator errors
func (ve *ValidationErrors) HasFieldError(field string) bool {
	errors, exists := ve.errors[field]
	return exists && len(errors) > 0
}

// Error implements the error interface
func (ve *ValidationErrors) Error() string {
	return ve.FirstError()
}
