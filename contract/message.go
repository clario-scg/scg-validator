package contract

// MessageResolver resolves validator error messages
type MessageResolver interface {
	// Resolve creates a validator error message
	Resolve(rule string, field string, parameters []string) string

	// SetCustomMessage sets a custom message for a rule
	SetCustomMessage(rule string, message string)

	// SetCustomAttribute sets a custom attribute name for a field
	SetCustomAttribute(field string, attribute string)

	// Clone creates a copy of the message resolver for request isolation
	Clone() MessageResolver
}
