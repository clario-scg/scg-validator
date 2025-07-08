package string

import (
	"errors"
	"fmt"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	doesntStartWithRuleName           = "doesnt_start_with"
	doesntStartWithRuleDefaultMsg     = "the :attribute must not start with one of the following: :values"
	doesntStartWithRuleMissingParam   = "doesnt_start_with rule requires at least one prefix"
	doesntStartWithRuleInvalidTypeMsg = "the :attribute must be a string"
	doesntStartWithRuleFailedMsg      = "the :attribute must not start with one of the following: %s"
)

// DoesntStartWithRule checks if a string does not start with any given prefixes.
type DoesntStartWithRule struct {
	common.BaseRule
	prefixes []string
}

// NewDoesntStartWithRule creates a new instance of the rule with given prefixes.
func NewDoesntStartWithRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(doesntStartWithRuleMissingParam)
	}

	return &DoesntStartWithRule{
		BaseRule: common.NewBaseRule(doesntStartWithRuleName, doesntStartWithRuleDefaultMsg, parameters),
		prefixes: parameters,
	}, nil
}

// Validate ensures the value does not start with any of the configured prefixes.
func (r *DoesntStartWithRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok {
		return errors.New(doesntStartWithRuleInvalidTypeMsg)
	}

	for _, prefix := range r.prefixes {
		if strings.HasPrefix(val, prefix) {
			return fmt.Errorf(doesntStartWithRuleFailedMsg, strings.Join(r.prefixes, ", "))
		}
	}

	return nil
}

func (r *DoesntStartWithRule) Name() string {
	return doesntStartWithRuleName
}
