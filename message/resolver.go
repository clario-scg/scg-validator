package message

import (
	"strings"
	"sync"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/utils"
)

// Resolver implements the MessageResolver interface
// It provides request-scoped custom message and attribute resolution
type Resolver struct {
	customMessages   map[string]string
	customAttributes map[string]string
	defaultMessages  map[string]string
	mu               sync.RWMutex
}

// NewResolver creates a new message resolver instance
func NewResolver() *Resolver {
	return &Resolver{
		customMessages:   make(map[string]string),
		customAttributes: make(map[string]string),
		defaultMessages:  getDefaultMessages(),
	}
}

// NewRequestScopedResolver creates a new resolver for a specific request
// This ensures isolation between different validation requests
func NewRequestScopedResolver() *Resolver {
	return NewResolver()
}

// Resolve creates a validation error message for the given rule, field, and parameters
func (r *Resolver) Resolve(rule string, field string, parameters []string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Try to get custom message first
	if customMsg, exists := r.customMessages[rule]; exists {
		return r.formatMessage(customMsg, field, parameters)
	}

	// Try to get field-specific custom message (rule.field format)
	fieldSpecificKey := rule + "." + field
	if customMsg, exists := r.customMessages[fieldSpecificKey]; exists {
		return r.formatMessage(customMsg, field, parameters)
	}

	// Fall back to default message
	if defaultMsg, exists := r.defaultMessages[rule]; exists {
		return r.formatMessage(defaultMsg, field, parameters)
	}

	// Ultimate fallback
	return r.formatMessage("The :attribute field is invalid", field, parameters)
}

// SetCustomMessage sets a custom message for a rule
func (r *Resolver) SetCustomMessage(rule string, message string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.customMessages[rule] = message
}

// SetCustomAttribute sets a custom attribute name for a field
func (r *Resolver) SetCustomAttribute(field string, attribute string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.customAttributes[field] = attribute
}

// formatMessage formats the message by replacing placeholders
func (r *Resolver) formatMessage(message string, field string, parameters []string) string {
	// Replace :attribute with custom attribute name or field name
	attributeName := field
	if customAttr, exists := r.customAttributes[field]; exists {
		attributeName = customAttr
	}

	message = strings.ReplaceAll(message, ":attribute", attributeName)
	message = strings.ReplaceAll(message, ":field", field)

	// Replace parameter placeholders
	for i, param := range parameters {
		message = utils.ReplacePlaceholder(message, i, param)
	}

	return message
}

// Clone creates a copy of the resolver for request isolation
func (r *Resolver) Clone() contract.MessageResolver {
	r.mu.RLock()
	defer r.mu.RUnlock()

	newResolver := NewResolver()

	// Copy custom messages
	for k, v := range r.customMessages {
		newResolver.customMessages[k] = v
	}

	// Copy custom attributes
	for k, v := range r.customAttributes {
		newResolver.customAttributes[k] = v
	}

	return newResolver
}

// getDefaultMessages returns the default validation messages
func getDefaultMessages() map[string]string {
	return map[string]string{
		"accepted":             "The :attribute must be accepted",
		"accepted_if":          "The :attribute must be accepted when :param0 is :param1",
		"accepted_unless":      "The :attribute must be accepted unless :param0 is :param1",
		"accepted_with":        "The :attribute must be accepted when :param0 is present",
		"accepted_without":     "The :attribute must be accepted when :param0 is not present",
		"declined":             "The :attribute must be declined",
		"declined_if":          "The :attribute must be declined when :param0 is :param1",
		"declined_unless":      "The :attribute must be declined unless :param0 is :param1",
		"declined_with":        "The :attribute must be declined when :param0 is present",
		"declined_without":     "The :attribute must be declined when :param0 is not present",
		"boolean":              "The :attribute must be true or false",
		"between":              "The :attribute must be between :param0 and :param1",
		"different":            "The :attribute and :param0 must be different",
		"ends_with":            "The :attribute must end with one of the following: :param0",
		"bail":                 "Stop validation on first failure",
		"exists":               "The selected :attribute is invalid",
		"date":                 "The :attribute is not a valid date",
		"after":                "The :attribute must be a date after :param0",
		"after_or_equal":       "The :attribute must be a date after or equal to :param0",
		"before":               "The :attribute must be a date before :param0",
		"before_or_equal":      "The :attribute must be a date before or equal to :param0",
		"date_equals":          "The :attribute must be a date equal to :param0",
		"date_format":          "The :attribute does not match the format :param0",
		"decimal":              "The :attribute must have :param0 decimal places",
		"active_url":           "The :attribute must be a valid URL",
		"confirmed":            "The :attribute confirmation does not match",
		"alpha":                "The :attribute may only contain letters",
		"alphanum":             "The :attribute may only contain letters and numbers",
		"alpha_dash":           "The :attribute may only contain letters, numbers, dashes and underscores",
		"email":                "The :attribute must be a valid email address",
		"ascii":                "The :attribute must only contain ASCII characters",
		"current_password":     "The :attribute is incorrect",
		"doesnt_start_with":    "The :attribute must not start with one of the following: :param0",
		"doesnt_end_with":      "The :attribute must not end with one of the following: :param0",
		"required":             "The :attribute field is required",
		"required_if":          "The :attribute field is required when :param0 is :param1",
		"required_unless":      "The :attribute field is required unless :param0 is :param1",
		"required_with":        "The :attribute field is required when :param0 is present",
		"required_without":     "The :attribute field is required when :param0 is not present",
		"required_with_all":    "The :attribute field is required when :param0 are present",
		"required_without_all": "The :attribute field is required when none of :param0 are present",
		"prohibited":           "The :attribute field is prohibited",
		"prohibited_if":        "The :attribute field is prohibited when :param0 is :param1",
		"prohibited_unless":    "The :attribute field is prohibited unless :param0 is :param1",
		"prohibits":            "The :attribute field prohibits :param0 from being present",
		"filled":               "The :attribute field must have a value",
		"present":              "The :attribute field must be present",
		"sometimes":            "The :attribute field is sometimes required",
		"nullable":             "The :attribute field may be null",
		"numeric":              "The :attribute must be a number",
		"integer":              "The :attribute must be an integer",
		"multiple_of":          "The :attribute must be a multiple of :param0",
		"lowercase":            "The :attribute must be lowercase",
		"uppercase":            "The :attribute must be uppercase",
		"ulid":                 "The :attribute must be a valid ULID",
		"slug":                 "The :attribute must be a valid slug",
		"file":                 "The :attribute must be a file",
		"image":                "The :attribute must be an image",
		"mimes":                "The :attribute must be a file of type: :param0",
		"min":                  "The :attribute must be at least :param0",
		"max":                  "The :attribute may not be greater than :param0",
		"size":                 "The :attribute must be :param0",
		"gt":                   "The :attribute must be greater than :param0",
		"lt":                   "The :attribute must be less than :param0",
		"gte":                  "The :attribute must be greater than or equal to :param0",
		"lte":                  "The :attribute must be less than or equal to :param0",
		"same":                 "The :attribute and :param0 must match",
	}
}
