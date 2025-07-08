package string

import (
	"errors"
	"fmt"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	startsWithRuleName              = "starts_with"
	startsWithRuleDefaultMessage    = "value must start with one of the allowed prefixes"
	startsWithRuleMissingParamError = "starts_with rule requires at least one prefix"
	startsWithRuleInvalidTypeError  = "value must be a string"
	startsWithRuleFailedFormat      = "value must start with one of the following: %s"
)

// StartsWithRule validates that a string starts with any of the given prefixes.
type StartsWithRule struct {
	common.BaseRule
	prefixes []string
}

// NewStartsWithRule constructs a StartsWithRule with the specified prefixes.
func NewStartsWithRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(startsWithRuleMissingParamError)
	}

	return &StartsWithRule{
		BaseRule: common.NewBaseRule(startsWithRuleName, startsWithRuleDefaultMessage, parameters, options...),
		prefixes: parameters,
	}, nil
}

// Validate checks whether the input value starts with any of the configured prefixes.
func (r *StartsWithRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	value, ok := ctx.Value().(string)
	if !ok {
		return errors.New(startsWithRuleInvalidTypeError)
	}

	for _, prefix := range r.prefixes {
		if strings.HasPrefix(value, prefix) {
			return nil
		}
	}

	return fmt.Errorf(startsWithRuleFailedFormat, strings.Join(r.prefixes, ", "))
}

func (r *StartsWithRule) Name() string {
	return startsWithRuleName
}
