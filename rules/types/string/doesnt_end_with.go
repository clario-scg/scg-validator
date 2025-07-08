package string

import (
	"errors"
	"fmt"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	doesntEndWithRuleName           = "doesnt_end_with"
	doesntEndWithRuleDefaultMsg     = "the :attribute must not end with one of the following: :values"
	doesntEndWithRuleMissingParam   = "doesnt_end_with rule requires at least one suffix"
	doesntEndWithRuleInvalidTypeMsg = "the :attribute must be a string"
	doesntEndWithRuleFailedMsg      = "the :attribute must not end with one of the following: %s"
)

// DoesntEndWithRule checks if a string does not end with any of the specified suffixes.
type DoesntEndWithRule struct {
	common.BaseRule
	suffixes []string
}

// NewDoesntEndWithRule creates a new instance of the rule.
func NewDoesntEndWithRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(doesntEndWithRuleMissingParam)
	}

	return &DoesntEndWithRule{
		BaseRule: common.NewBaseRule(doesntEndWithRuleName, doesntEndWithRuleDefaultMsg, parameters, options...),
		suffixes: parameters,
	}, nil
}

// Validate performs the actual check for suffix presence.
func (r *DoesntEndWithRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	value, ok := ctx.Value().(string)
	if !ok {
		return errors.New(doesntEndWithRuleInvalidTypeMsg)
	}

	for _, suffix := range r.suffixes {
		if strings.HasSuffix(value, suffix) {
			return fmt.Errorf(doesntEndWithRuleFailedMsg, strings.Join(r.suffixes, ", "))
		}
	}

	return nil
}

func (r *DoesntEndWithRule) Name() string {
	return doesntEndWithRuleName
}
