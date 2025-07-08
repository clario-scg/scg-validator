// Package validator provides a validator engine for schema-based validator.
// This package follows the Clean Architecture principles with clear separation of concerns:
// - contract/: Defines interfaces, types, and constants
// - engine/: Contains the validator engine
// - rules/: Contains built-in rule implementations
// - parser/: Implements rule parsing
package validator

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/engine"
	"github.com/next-trace/scg-validator/message"
)

// Validator is the main facade that provides a simple interface for validator
type Validator struct {
	engine contract.ValidationEngine
}

// New creates a new validator with all Laravel rules registered
func New() *Validator {
	eng := engine.NewEngine()

	return &Validator{
		engine: eng,
	}
}

// Validate validates data against the provided rules and returns an error
func (v *Validator) Validate(data any, rules map[string]string) error {
	result := v.ValidateWithResult(data, rules)
	if !result.IsValid() {
		if validationErrors, ok := result.(*contract.ValidationErrors); ok {
			return validationErrors
		}
		return errors.New("validator failed")
	}
	return nil
}

// ValidateWithResult validates data against the provided rules and returns the full result
func (v *Validator) ValidateWithResult(data any, rules map[string]string) contract.Result {
	// Convert data to map[strings]any if needed
	var dataMap map[string]any
	switch d := data.(type) {
	case map[string]any:
		dataMap = d
	default:
		// TODO: Handle other data types like structs
		dataMap = make(map[string]any)
	}

	// Create a request-scoped engine to ensure isolation between validation requests
	requestEngine := v.createRequestScopedEngine()

	dataProvider := engine.NewDataProvider(dataMap)
	return requestEngine.Execute(dataProvider, rules)
}

// AddRule adds a custom rule to the validator
func (v *Validator) AddRule(name string, creator contract.RuleCreator) error {
	return v.engine.RegisterRule(name, creator)
}

// HasRule checks if a rule exists
func (v *Validator) HasRule(name string) bool {
	// Use the registry from the engine to check if rule exists
	_, exists := v.engine.GetRegistry().Get(name)
	return exists
}

// GetAvailableRules returns a list of all available rule names
func (v *Validator) GetAvailableRules() []string {
	return v.engine.GetRegistry().List()
}

// ValidateMap is a convenience method for validating map data with array-style rules (Laravel-style)
func (v *Validator) ValidateMap(data map[string]interface{}, rules map[string][]string) contract.Result {
	// Convert array-style rules to pipe-separated strings
	stringRules := make(map[string]string)
	for field, fieldRules := range rules {
		if len(fieldRules) > 0 {
			rulesString := ""
			for i, rule := range fieldRules {
				if i > 0 {
					rulesString += "|"
				}
				rulesString += rule
			}
			stringRules[field] = rulesString
		}
	}

	return v.ValidateWithResult(data, stringRules)
}

// SetCustomMessage sets a custom message for a rule
func (v *Validator) SetCustomMessage(rule, message string) {
	v.engine.SetCustomMessage(rule, message)
}

// SetCustomAttribute sets a custom attribute name for a field
func (v *Validator) SetCustomAttribute(field, name string) {
	v.engine.SetCustomAttribute(field, name)
}

// createRequestScopedEngine creates a new engine instance with isolated message resolver
// This ensures that custom messages and attributes don't interfere between validation requests
func (v *Validator) createRequestScopedEngine() contract.ValidationEngine {
	// Clone the current message resolver to maintain custom messages while ensuring isolation
	var requestResolver contract.MessageResolver
	if mr := v.engine.GetMessageResolver(); mr != nil {
		requestResolver = mr.Clone()
	} else {
		requestResolver = message.NewRequestScopedResolver()
	}

	// Create a request-scoped engine that shares the same registry but uses the cloned resolver
	return v.engine.CloneWithResolver(requestResolver)
}
