// Package facade provides a Laravel-style facade for the validator framework
package facade

import (
	"sync"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/engine"
	rules2 "github.com/next-trace/scg-validator/registry/rules"
	"github.com/next-trace/scg-validator/rules"
)

// globalValidator holds the global validator instance
var globalValidator *ValidatorFacade
var globalValidatorOnce sync.Once

// getGlobalValidator returns the global validator instance, initializing it if necessary
func getGlobalValidator() *ValidatorFacade {
	globalValidatorOnce.Do(func() {
		globalValidator = New()
	})
	return globalValidator
}

// ValidatorFacade provides a Laravel-style API for validator
type ValidatorFacade struct {
	engine *engine.Engine
}

// New creates a new ValidatorFacade instance
func New(options ...rules2.Option) *ValidatorFacade {
	registry := rules.NewRuleRegistry(options...)
	eng := &engine.Engine{
		Registry: registry,
	}

	return &ValidatorFacade{
		engine: eng,
	}
}

// ValidatorRequest represents a validator request similar to Laravel's Validator
type ValidatorRequest struct {
	data     contract.DataProvider
	rules    map[string][]string
	messages map[string]string
	engine   *engine.Engine
}

// Make creates a new validator instance (Laravel-style API)
// Usage: facade.Make(data, map[strings][]strings{"email": {"required", "email"}})
func Make(data contract.DataProvider, rules map[string][]string, messages ...map[string]string) *ValidatorRequest {
	return getGlobalValidator().Make(data, rules, messages...)
}

// Make creates a new validator instance
func (v *ValidatorFacade) Make(data contract.DataProvider, rules map[string][]string,
	messages ...map[string]string) *ValidatorRequest {
	var msgs map[string]string
	if len(messages) > 0 {
		msgs = messages[0]
	}

	return &ValidatorRequest{
		data:     data,
		rules:    rules,
		messages: msgs,
		engine:   v.engine,
	}
}

// Validate performs validator and returns errors if any (Laravel-style API)
// Usage: errors := facade.Validate(data, map[strings][]strings{"email": {"required", "email"}})
func Validate(data contract.DataProvider, rules map[string][]string,
	messages ...map[string]string) *contract.ValidationErrors {
	return getGlobalValidator().Validate(data, rules, messages...)
}

// Validate performs validator and returns errors if any
func (v *ValidatorFacade) Validate(data contract.DataProvider, rules map[string][]string,
	messages ...map[string]string) *contract.ValidationErrors {
	validator := v.Make(data, rules, messages...)
	return validator.Validate()
}

// ValidateMap provides a convenient way to validate map[strings]interface{} data (Laravel-style API)
// Usage: errors := facade.ValidateMap(map[strings]interface{}{"email": "test@example.com"},
//
//	map[strings][]strings{"email": {"required", "email"}})
func ValidateMap(data map[string]interface{}, rules map[string][]string,
	messages ...map[string]string) *contract.ValidationErrors {
	return getGlobalValidator().ValidateMap(data, rules, messages...)
}

// ValidateMap provides a convenient way to validate map[strings]interface{} data
func (v *ValidatorFacade) ValidateMap(data map[string]interface{}, rules map[string][]string,
	messages ...map[string]string) *contract.ValidationErrors {
	dataProvider := contract.NewSimpleDataProvider(data)
	return v.Validate(dataProvider, rules, messages...)
}

// Validate executes the validator and returns results
func (vr *ValidatorRequest) Validate() *contract.ValidationErrors {
	// Convert map[strings][]strings rules to map[strings]strings
	rulesMap := make(map[string]string)
	for field, fieldRules := range vr.rules {
		if len(fieldRules) > 0 {
			// Join multiple rules with pipe separator (Laravel-style)
			rulesString := ""
			for i, rule := range fieldRules {
				if i > 0 {
					rulesString += "|"
				}
				rulesString += rule
			}
			rulesMap[field] = rulesString
		}
	}

	// Execute validator using the engine
	result := vr.engine.Execute(vr.data, rulesMap)

	// Convert result to ValidationErrors
	validationErrors := contract.NewValidationErrors()
	if !result.IsValid() {
		// Add errors from the result
		errors := result.Errors()
		for field, fieldErrors := range errors {
			for _, err := range fieldErrors {
				validationErrors.AddError(field, err)
			}
		}
	}

	return validationErrors
}

// Fails returns true if validator failed (Laravel-style API)
func (vr *ValidatorRequest) Fails() bool {
	result := vr.Validate()
	return !result.IsValid()
}

// Passes returns true if validator passed (Laravel-style API)
func (vr *ValidatorRequest) Passes() bool {
	return !vr.Fails()
}

// Errors returns the validator errors (Laravel-style API)
func (vr *ValidatorRequest) Errors() *contract.ValidationErrors {
	return vr.Validate()
}

// GetMessageBag returns validator errors (Laravel-style alias)
func (vr *ValidatorRequest) GetMessageBag() *contract.ValidationErrors {
	return vr.Errors()
}

// WithRules allows adding additional rules (Laravel-style API)
func (vr *ValidatorRequest) WithRules(additionalRules map[string][]string) *ValidatorRequest {
	// Merge the rules
	newRules := make(map[string][]string)

	// Copy existing rules
	for field, fieldRules := range vr.rules {
		newRules[field] = make([]string, len(fieldRules))
		copy(newRules[field], fieldRules)
	}

	// Add new rules
	for field, fieldRules := range additionalRules {
		if existing, exists := newRules[field]; exists {
			newRules[field] = append(existing, fieldRules...)
		} else {
			newRules[field] = make([]string, len(fieldRules))
			copy(newRules[field], fieldRules)
		}
	}

	return &ValidatorRequest{
		data:     vr.data,
		rules:    newRules,
		messages: vr.messages,
		engine:   vr.engine,
	}
}

// WithMessages allows adding custom messages (Laravel-style API)
func (vr *ValidatorRequest) WithMessages(additionalMessages map[string]string) *ValidatorRequest {
	// Merge the messages
	newMessages := make(map[string]string)

	// Copy existing messages
	if vr.messages != nil {
		for key, msg := range vr.messages {
			newMessages[key] = msg
		}
	}

	// Add new messages
	for key, msg := range additionalMessages {
		newMessages[key] = msg
	}

	return &ValidatorRequest{
		data:     vr.data,
		rules:    vr.rules,
		messages: newMessages,
		engine:   vr.engine,
	}
}

// Extend allows registering custom rules (Laravel-style API)
// Usage: facade.Extend("custom_rule", func(parameters []strings) (contract.Rule, error) { ... })
func Extend(ruleName string, ruleCreator contract.RuleCreator) {
	getGlobalValidator().Extend(ruleName, ruleCreator)
}

// Extend allows registering custom rules
func (v *ValidatorFacade) Extend(ruleName string, ruleCreator contract.RuleCreator) {
	// Register the rule in the existing registry
	_ = v.engine.Registry.Register(ruleName, ruleCreator)
}

// ExtendImplicit allows registering custom implicit rules (Laravel-style API)
// Usage: facade.ExtendImplicit("sometimes", func(parameters []strings) (contract.Rule, error) { ... })
func ExtendImplicit(ruleName string, ruleCreator contract.RuleCreator) {
	// For now, treat implicit rules the same as regular rules
	// In the future, we could add special handling for implicit rules
	Extend(ruleName, ruleCreator)
}

// ExtendImplicit allows registering custom implicit rules
func (v *ValidatorFacade) ExtendImplicit(ruleName string, ruleCreator contract.RuleCreator) {
	v.Extend(ruleName, ruleCreator)
}

// Rules returns the list of available rules
func Rules() []string {
	return getGlobalValidator().Rules()
}

// Rules returns the list of available rules
func (v *ValidatorFacade) Rules() []string {
	if v.engine == nil || v.engine.Registry == nil {
		return []string{}
	}
	return v.engine.Registry.List()
}

// HasRule checks if a rule exists
func HasRule(ruleName string) bool {
	return getGlobalValidator().HasRule(ruleName)
}

// HasRule checks if a rule exists
func (v *ValidatorFacade) HasRule(ruleName string) bool {
	if v.engine == nil || v.engine.Registry == nil {
		return false
	}
	_, exists := v.engine.Registry.Get(ruleName)
	return exists
}

// Quick validator functions for common use cases

// Required validates that a field is required
func Required(data contract.DataProvider, field string) bool {
	rules := map[string][]string{field: {"required"}}
	result := Validate(data, rules)
	return result.IsValid()
}

// Email validates that a field is a valid email
func Email(data contract.DataProvider, field string) bool {
	rules := map[string][]string{field: {"email"}}
	result := Validate(data, rules)
	return result.IsValid()
}

// Numeric validates that a field is numeric
func Numeric(data contract.DataProvider, field string) bool {
	rules := map[string][]string{field: {"numeric"}}
	result := Validate(data, rules)
	return result.IsValid()
}

// Min validates that a field meets the minimum value requirement
func Min(data contract.DataProvider, field string, minVal string) bool {
	rules := map[string][]string{field: {"min:" + minVal}}
	result := Validate(data, rules)
	return result.IsValid()
}

// Max validates that a field meets the maximum value requirement
func Max(data contract.DataProvider, field string, maxVal string) bool {
	rules := map[string][]string{field: {"max:" + maxVal}}
	result := Validate(data, rules)
	return result.IsValid()
}

// Alpha validates that a field contains only alphabetic characters
func Alpha(data contract.DataProvider, field string) bool {
	rules := map[string][]string{field: {"alpha"}}
	result := Validate(data, rules)
	return result.IsValid()
}

// AlphaNum validates that a field contains only alphanumeric characters
func AlphaNum(data contract.DataProvider, field string) bool {
	rules := map[string][]string{field: {"alpha_num"}}
	result := Validate(data, rules)
	return result.IsValid()
}

// URL validates that a field is a valid URL
func URL(data contract.DataProvider, field string) bool {
	rules := map[string][]string{field: {"url"}}
	result := Validate(data, rules)
	return result.IsValid()
}

// UUID validates that a field is a valid UUID
func UUID(data contract.DataProvider, field string) bool {
	rules := map[string][]string{field: {"uuid"}}
	result := Validate(data, rules)
	return result.IsValid()
}
