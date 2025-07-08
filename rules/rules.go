// Package rules provides the main rule registry and initialization functions following functional approach
package rules

import (
	"fmt"

	"github.com/next-trace/scg-validator/registry/rules"
	"github.com/next-trace/scg-validator/rules/authentication"
	"github.com/next-trace/scg-validator/rules/file"
	"github.com/next-trace/scg-validator/rules/format"
	dateRules "github.com/next-trace/scg-validator/rules/types/date"
	stringRules "github.com/next-trace/scg-validator/rules/types/string"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/acceptance"
	"github.com/next-trace/scg-validator/rules/comparison"
	"github.com/next-trace/scg-validator/rules/conditional"
	"github.com/next-trace/scg-validator/rules/control"
	"github.com/next-trace/scg-validator/rules/types/boolean"
	"github.com/next-trace/scg-validator/rules/types/numeric"
)

// Rule names constants to avoid magic strings
// These are the rule names used throughout the validation engine,
// grouped by category for easier understanding and management.
const (
	// Acceptance Rules
	RuleAccepted   = "accepted"
	RuleDeclined   = "declined"
	RuleAcceptedIf = "accepted_if"
	RuleDeclinedIf = "declined_if"

	// Boolean Rules
	RuleBoolean = "boolean"

	// Comparison Rules
	RuleMin       = "min"
	RuleMax       = "max"
	RuleSize      = "size"
	RuleBetween   = "between"
	RuleGt        = "gt"
	RuleLt        = "lt"
	RuleGte       = "gte"
	RuleLte       = "lte"
	RuleSame      = "same"
	RuleDifferent = "different"
	RuleConfirmed = "confirmed"

	// Conditional Rules
	RuleRequired           = "required"
	RuleRequiredIf         = "required_if"
	RuleRequiredUnless     = "required_unless"
	RuleRequiredWith       = "required_with"
	RuleRequiredWithout    = "required_without"
	RuleRequiredWithAll    = "required_with_all"
	RuleRequiredWithoutAll = "required_without_all"

	// Prohibited Rules
	RuleProhibited       = "prohibited"
	RuleProhibitedIf     = "prohibited_if"
	RuleProhibitedUnless = "prohibited_unless"
	RuleProhibits        = "prohibits"

	// Control Rules
	RuleBail      = "bail"
	RuleFilled    = "filled"
	RulePresent   = "present"
	RuleSometimes = "sometimes"
	RuleNullable  = "nullable"

	// Date Rules
	RuleAfter         = "after"
	RuleBefore        = "before"
	RuleBeforeOrEqual = "before_or_equal"
	RuleAfterOrEqual  = "after_or_equal"
	RuleDate          = "date"
	RuleDateEquals    = "date_equals"

	// Numeric Rules
	RuleNumeric    = "numeric"
	RuleInteger    = "integer"
	RuleDecimal    = "decimal"
	RuleMultipleOf = "multiple_of"

	// String Rules
	RuleAlpha     = "alpha"
	RuleAlphaNum  = "alpha_num"
	RuleAlphaDash = "alpha_dash"
	RuleEmail     = "email"
	RuleUUID      = "uuid"
	RuleURL       = "url"
	RuleActiveURL = "active_url"
	RuleJSON      = "json"
	RuleRegex     = "regex"
	RuleIP        = "ip"
	RuleIPv4      = "ipv4"
	RuleIPv6      = "ipv6"
	RuleMAC       = "mac"

	// File Validation Rules
	RuleFile  = "file"
	RuleImage = "image"
	RuleMimes = "mimes"

	// Special String Rules
	RuleLowercase       = "lowercase"
	RuleUppercase       = "uppercase"
	RuleASCII           = "ascii"
	RuleUlid            = "ulid"
	RuleSlug            = "slug"
	RuleDoesntStartWith = "doesnt_start_with"
	RuleDoesntEndWith   = "doesnt_end_with"

	// Auth Rules
	RuleCurrentPassword = "current_password"
)

// WithCustomRule adds a custom rule to the registry
func WithCustomRule(name string, creator contract.RuleCreator) rules.Option {
	return func(config *contract.Config) {
		if config.CustomRules == nil {
			config.CustomRules = make(map[string]contract.RuleCreator)
		}
		config.CustomRules[name] = creator
	}
}

// WithCustomMessage sets a custom message for a rule
func WithCustomMessage(ruleName string, message string) rules.Option {
	return func(config *contract.Config) {
		if config.CustomMessages == nil {
			config.CustomMessages = make(map[string]string)
		}
		config.CustomMessages[ruleName] = message
	}
}

// WithExcludeRules excludes specific default rules from being registered
func WithExcludeRules(ruleNames ...string) rules.Option {
	return func(config *contract.Config) {
		if config.ExcludeRules == nil {
			config.ExcludeRules = make(map[string]bool)
		}
		for _, name := range ruleNames {
			config.ExcludeRules[name] = true
		}
	}
}

// WithIncludeOnly registers only the specified default rules
func WithIncludeOnly(ruleNames ...string) rules.Option {
	return func(config *contract.Config) {
		if config.IncludeOnly == nil {
			config.IncludeOnly = make(map[string]bool)
		}
		for _, name := range ruleNames {
			config.IncludeOnly[name] = true
		}
	}
}

// NewRuleRegistry creates a new rule registry with all default rules and applies options
func NewRuleRegistry(options ...rules.Option) contract.Registry {
	reg := rules.NewRegistry()

	// Create config and apply options
	config := &contract.Config{}
	for _, option := range options {
		option(config)
	}

	// Register all default rules using constants with filtering
	if err := registerDefaultRules(reg, config); err != nil {
		panic(fmt.Sprintf("failed to register default rules: %v", err))
	}

	// Register custom rules from config
	for name, creator := range config.CustomRules {
		_ = reg.Register(name, creator)
	}

	return reg
}

// registerDefaultRules registers all available rules from different packages
func registerDefaultRules(reg contract.Registry, config *contract.Config) error {
	rules := map[string]contract.RuleCreator{
		// Acceptance rules
		RuleAccepted:   func(_ []string) (contract.Rule, error) { return acceptance.NewAcceptedRule() },
		RuleDeclined:   func(_ []string) (contract.Rule, error) { return acceptance.NewDeclinedRule() },
		RuleAcceptedIf: acceptance.NewAcceptedIfRule,
		RuleDeclinedIf: acceptance.NewDeclinedIfRule,

		// Boolean rules
		RuleBoolean: func(_ []string) (contract.Rule, error) { return boolean.NewBooleanRule() },

		// Comparison rules
		RuleMin:       comparison.NewMinRule,
		RuleMax:       comparison.NewMaxRule,
		RuleSize:      comparison.NewSizeRule,
		RuleBetween:   comparison.NewBetweenRule,
		RuleGt:        comparison.NewGtRule,
		RuleLt:        comparison.NewLtRule,
		RuleGte:       comparison.NewGteRule,
		RuleLte:       comparison.NewLteRule,
		RuleSame:      comparison.NewSameRule,
		RuleDifferent: comparison.NewDifferentRule,
		RuleConfirmed: func(_ []string) (contract.Rule, error) { return comparison.NewConfirmedRule() },

		// Conditional rules
		RuleRequired:           func(_ []string) (contract.Rule, error) { return conditional.NewRequiredRule() },
		RuleRequiredIf:         conditional.NewRequiredIfRule,
		RuleRequiredUnless:     conditional.NewRequiredUnlessRule,
		RuleRequiredWith:       conditional.NewRequiredWithRule,
		RuleRequiredWithout:    conditional.NewRequiredWithoutRule,
		RuleRequiredWithAll:    conditional.NewRequiredWithAllRule,
		RuleRequiredWithoutAll: conditional.NewRequiredWithoutAllRule,

		// Prohibited rules
		RuleProhibited:       func(_ []string) (contract.Rule, error) { return conditional.NewProhibitedRule() },
		RuleProhibitedIf:     conditional.NewProhibitedIfRule,
		RuleProhibitedUnless: conditional.NewProhibitedUnlessRule,
		RuleProhibits:        conditional.NewProhibitsRule,

		// Control rules
		RuleBail:      func(_ []string) (contract.Rule, error) { return control.NewBailRule() },
		RuleFilled:    func(_ []string) (contract.Rule, error) { return control.NewFilledRule() },
		RulePresent:   func(_ []string) (contract.Rule, error) { return control.NewPresentRule() },
		RuleSometimes: func(_ []string) (contract.Rule, error) { return control.NewSometimesRule() },
		RuleNullable:  func(_ []string) (contract.Rule, error) { return control.NewNullableRule() },

		// Date rules
		RuleAfter:         dateRules.NewAfterRule,
		RuleBefore:        dateRules.NewBeforeRule,
		RuleBeforeOrEqual: dateRules.NewBeforeOrEqualRule,
		RuleAfterOrEqual:  dateRules.NewAfterOrEqualRule,

		// Numeric rules
		RuleNumeric:    func(_ []string) (contract.Rule, error) { return numeric.NewNumericRule() },
		RuleInteger:    func(_ []string) (contract.Rule, error) { return numeric.NewIntegerRule() },
		RuleDecimal:    numeric.NewDecimalRule,
		RuleMultipleOf: numeric.NewMultipleOfRule,

		// String rules
		RuleAlpha:           func(p []string) (contract.Rule, error) { return stringRules.NewAlphaRule(p) },
		RuleAlphaNum:        func(p []string) (contract.Rule, error) { return stringRules.NewAlphaNumRule(p) },
		RuleAlphaDash:       func(p []string) (contract.Rule, error) { return stringRules.NewAlphaDashRule(p) },
		RuleLowercase:       func(_ []string) (contract.Rule, error) { return stringRules.NewLowercaseRule() },
		RuleUppercase:       func(_ []string) (contract.Rule, error) { return stringRules.NewUppercaseRule() },
		RuleASCII:           func(_ []string) (contract.Rule, error) { return stringRules.NewASCIIRule() },
		RuleUlid:            func(_ []string) (contract.Rule, error) { return stringRules.NewUlidRule() },
		RuleSlug:            func(_ []string) (contract.Rule, error) { return stringRules.NewSlugRule() },
		RuleDoesntStartWith: stringRules.NewDoesntStartWithRule,
		RuleDoesntEndWith:   func(p []string) (contract.Rule, error) { return stringRules.NewDoesntEndWithRule(p) },

		// Format rules
		RuleEmail: func(p []string) (contract.Rule, error) { return format.NewEmailRule(p) },
		RuleURL:   func(p []string) (contract.Rule, error) { return format.NewURLRule(p) },

		// File rules
		RuleFile:  func(_ []string) (contract.Rule, error) { return file.NewFileRule() },
		RuleImage: func(_ []string) (contract.Rule, error) { return file.NewImageRule() },
		RuleMimes: file.NewMimesRule,

		// Auth rules
		RuleCurrentPassword: func(_ []string) (contract.Rule, error) { return authentication.NewCurrentPasswordRule() },
	}

	// Apply filtering based on config
	filteredRules := make(map[string]contract.RuleCreator)

	// If IncludeOnly is specified, only include those rules
	if len(config.IncludeOnly) > 0 {
		for name, creator := range rules {
			if config.IncludeOnly[name] {
				filteredRules[name] = creator
			}
		}
	} else {
		// Otherwise, include all rules except excluded ones
		for name, creator := range rules {
			if !config.ExcludeRules[name] {
				filteredRules[name] = creator
			}
		}
	}

	// Register filtered rules to the registry
	for name, creator := range filteredRules {
		if err := reg.Register(name, creator); err != nil {
			return fmt.Errorf("failed to register rule %s: %w", name, err)
		}
	}

	return nil
}
