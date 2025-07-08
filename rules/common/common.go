package common

import (
	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/utils"
)

// RuleOption is a functional option for configuring a rule.
type RuleOption func(*Config)

// Config holds the configuration values for a rule.
type Config struct {
	Name       string   // Rule name, e.g., "required", "email"
	Message    string   // Default or custom error message
	Parameters []string // Rule parameters, e.g., min, max
	Nullable   bool     // If true, nil values are accepted
	StopOnFail bool     // If true, stops validation on the first failure
}

// WithMessage sets a custom message for the rule.
func WithMessage(message string) RuleOption {
	return func(cfg *Config) {
		cfg.Message = message
	}
}

// WithNullable allows the rule to accept nil values (nullable).
func WithNullable(nullable bool) RuleOption {
	return func(cfg *Config) {
		cfg.Nullable = nullable
	}
}

// WithStopOnFail determines if validation should stop after the rule fails.
func WithStopOnFail(stop bool) RuleOption {
	return func(cfg *Config) {
		cfg.StopOnFail = stop
	}
}

// BaseRule provides a reusable, configurable foundation for validation rules.
type BaseRule struct {
	config Config
}

// NewBaseRule creates a new BaseRule with the specified options.
func NewBaseRule(name, defaultMessage string, parameters []string, options ...RuleOption) BaseRule {
	cfg := Config{
		Name:       name,
		Message:    defaultMessage,
		Parameters: parameters,
		Nullable:   false,
		StopOnFail: false,
	}

	// Apply functional options
	for _, opt := range options {
		if opt != nil {
			opt(&cfg)
		}
	}

	return BaseRule{config: cfg}
}

// Name returns the rule name.
func (r BaseRule) Name() string {
	return r.config.Name
}

// GetMessage returns the rule's validation error message.
func (r BaseRule) GetMessage() string {
	return r.config.Message
}

// Parameters returns the list of parameters associated with the rule.
func (r BaseRule) Parameters() []string {
	return r.config.Parameters
}

// IsNullable checks if the rule accepts nil values.
func (r BaseRule) IsNullable() bool {
	return r.config.Nullable
}

// ShouldSkipValidation checks if the rule should be skipped.
func (r BaseRule) ShouldSkipValidation(value interface{}) bool {
	// Skip if value is nil and rule is nullable
	return value == nil && r.config.Nullable
}

// SimpleRule is a wrapper for rules with a single validator function.
type SimpleRule struct {
	BaseRule
	validator func(ctx contract.RuleContext) error
}

// Validate invokes the validator function, skipping if the rule is nullable and value is nil.
func (r *SimpleRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}
	return r.validator(ctx)
}

// Message applies placeholders in the message (e.g., :attribute, :value).
func (r *SimpleRule) Message() string {
	msg := r.GetMessage()
	for i, param := range r.Parameters() {
		// Replace placeholders with actual parameter values
		msg = utils.ReplacePlaceholder(msg, i, param)
	}
	return msg
}
