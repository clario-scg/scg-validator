package engine

import (
	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/message"
	"github.com/next-trace/scg-validator/parser"
	"github.com/next-trace/scg-validator/rules"
)

// Define constants to avoid magic strings and magic numbers
const (
	BailRuleName         = "bail"
	UnknownRuleErrorMsg  = "Unknown rule: "
	RuleCreationErrorMsg = "Rule creation error: "
)

// Engine implements the ValidationEngine interface
type Engine struct {
	Registry        contract.Registry
	MessageResolver contract.MessageResolver
}

// Ensure Engine implements contract.ValidationEngine
var _ contract.ValidationEngine = (*Engine)(nil)

// NewEngine creates a new validator engine
func NewEngine() *Engine {
	// Create a new registry with all default rules using options pattern
	reg := rules.NewRuleRegistry()

	return &Engine{
		Registry:        reg,
		MessageResolver: message.NewRequestScopedResolver(),
	}
}

// Execute validates data against the provided rules
func (e *Engine) Execute(data contract.DataProvider, rulesMap map[string]string) contract.Result {
	validationErrors := contract.NewValidationErrors()

	// Iterate over each field and corresponding rules
	for field, ruleString := range rulesMap {
		e.validateField(field, ruleString, data, validationErrors)
	}

	return validationErrors
}

// validateField validates a single field against its rules
func (e *Engine) validateField(
	field, ruleString string,
	data contract.DataProvider,
	validationErrors *contract.ValidationErrors,
) {
	parsedRules := parser.ParseRules(ruleString)
	value, _ := data.Get(field)
	allData := data.All()

	stopOnFailure := e.shouldStopOnFailure(parsedRules)

	for _, parsedRule := range parsedRules {
		if parsedRule.Name == BailRuleName {
			continue
		}

		if e.validateSingleRule(field, value, parsedRule, allData, validationErrors) && stopOnFailure {
			break
		}
	}
}

// shouldStopOnFailure checks if the bail rule is present in the parsed rules
func (e *Engine) shouldStopOnFailure(parsedRules []parser.ParsedRule) bool {
	for _, rule := range parsedRules {
		if rule.Name == BailRuleName {
			return true
		}
	}
	return false
}

// validateSingleRule validates a single rule and returns true if validation failed
func (e *Engine) validateSingleRule(
	field string,
	value interface{},
	parsedRule parser.ParsedRule,
	allData map[string]interface{},
	validationErrors *contract.ValidationErrors,
) bool {
	ruleName := parsedRule.Name

	// Fetch the rule creator from the registry
	ruleCreator, exists := e.Registry.Get(ruleName)
	if !exists {
		validationErrors.AddError(field, UnknownRuleErrorMsg+ruleName)
		return true
	}

	// Create the rule and handle any errors during creation
	rule, err := ruleCreator(parsedRule.Params)
	if err != nil {
		validationErrors.AddError(field, RuleCreationErrorMsg+err.Error())
		return true
	}

	// Create validation context and perform the validation
	ctx := contract.NewValidationContext(field, value, parsedRule.Params, allData)

	// Validate and handle error if validation fails
	if err := rule.Validate(ctx); err != nil {
		errorMessage := e.resolveErrorMessage(ruleName, field, parsedRule.Params, err)
		validationErrors.AddError(field, errorMessage)
		return true
	}

	return false
}

// resolveErrorMessage resolves the error message using the message resolver
func (e *Engine) resolveErrorMessage(ruleName, field string, params []string, originalError error) string {
	if e.MessageResolver != nil {
		return e.MessageResolver.Resolve(ruleName, field, params)
	}
	return originalError.Error()
}

// RegisterRule registers a new rule with the engine
func (e *Engine) RegisterRule(name string, creator contract.RuleCreator) error {
	return e.Registry.Register(name, creator)
}

// SetMessageResolver sets a custom message resolver
func (e *Engine) SetMessageResolver(resolver contract.MessageResolver) {
	e.MessageResolver = resolver
}

// SetCustomMessage sets a custom message for a rule
func (e *Engine) SetCustomMessage(rule string, message string) {
	if e.MessageResolver != nil {
		e.MessageResolver.SetCustomMessage(rule, message)
	}
}

// SetCustomAttribute sets a custom attribute name for a field
func (e *Engine) SetCustomAttribute(field string, attribute string) {
	if e.MessageResolver != nil {
		e.MessageResolver.SetCustomAttribute(field, attribute)
	}
}

// GetRegistry exposes the rule registry
func (e *Engine) GetRegistry() contract.Registry {
	return e.Registry
}

// GetMessageResolver returns the current message resolver
func (e *Engine) GetMessageResolver() contract.MessageResolver {
	return e.MessageResolver
}

// CloneWithResolver creates a new Engine that shares the same registry but uses the provided resolver
func (e *Engine) CloneWithResolver(resolver contract.MessageResolver) contract.ValidationEngine {
	return &Engine{
		Registry:        e.Registry,
		MessageResolver: resolver,
	}
}

// DataProvider implementation for map[strings]interface{}
type DataProvider struct {
	data map[string]interface{}
}

// NewDataProvider creates a new data provider from a map
func NewDataProvider(data map[string]interface{}) *DataProvider {
	return &DataProvider{data: data}
}

// Get retrieves a value by key
func (d *DataProvider) Get(key string) (interface{}, bool) {
	if d.data == nil {
		return nil, false
	}
	value, exists := d.data[key]
	return value, exists
}

// Has checks if a key exists
func (d *DataProvider) Has(key string) bool {
	if d.data == nil {
		return false
	}
	_, exists := d.data[key]
	return exists
}

// All returns all data
func (d *DataProvider) All() map[string]interface{} {
	return d.data
}
