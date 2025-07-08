package contract

// ValidationEngine is the abstraction the validator facade depends on.
// It enables swapping the underlying engine implementation without changing
// the facade or its consumers (DIP / code-to-interfaces).
type ValidationEngine interface {
	// Execute validates data against the provided rules and returns a Result.
	Execute(data DataProvider, rules map[string]string) Result

	// RegisterRule registers a new rule.
	RegisterRule(name string, creator RuleCreator) error

	// GetRegistry exposes the rule registry (read-only usage by facade).
	GetRegistry() Registry

	// GetMessageResolver returns the message resolver in use.
	GetMessageResolver() MessageResolver

	// SetCustomMessage sets a custom message for a rule.
	SetCustomMessage(rule string, message string)

	// SetCustomAttribute sets a custom attribute for a field.
	SetCustomAttribute(field string, attribute string)

	// CloneWithResolver returns a new, request-scoped engine instance
	// that shares the same registry but uses the provided message resolver.
	CloneWithResolver(resolver MessageResolver) ValidationEngine
}
