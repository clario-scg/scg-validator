package contract

// This file is correctly named and domain-agnostic. No changes needed.

type Validator interface {
	Validate(data any, rules map[string]string) error
	ValidateWithResult(data any, rules map[string]string) Result
}
