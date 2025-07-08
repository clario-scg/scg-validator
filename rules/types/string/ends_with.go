package string

import (
	"errors"
	"fmt"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	endsWithRuleName           = "ends_with"
	endsWithRuleDefaultMsg     = "the :attribute must end with one of the following: :values"
	endsWithRuleMissingParam   = "ends_with rule requires at least one suffix"
	endsWithRuleInvalidTypeMsg = "the :attribute must be a string"
	endsWithRuleFailedMsg      = "the :attribute must end with one of the following: %s"
)

// EndsWithRule checks if a string ends with one of the given suffixes.
type EndsWithRule struct {
	common.BaseRule
	suffixes []string
}

// NewEndsWithRule creates a new instance of EndsWithRule.
func NewEndsWithRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(endsWithRuleMissingParam)
	}

	return &EndsWithRule{
		BaseRule: common.NewBaseRule(endsWithRuleName, endsWithRuleDefaultMsg, parameters),
		suffixes: parameters,
	}, nil
}

// Validate checks if the input ends with one of the allowed suffixes.
func (r *EndsWithRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	value, ok := ctx.Value().(string)
	if !ok {
		return errors.New(endsWithRuleInvalidTypeMsg)
	}

	for _, suffix := range r.suffixes {
		if strings.HasSuffix(value, suffix) {
			return nil
		}
	}

	return fmt.Errorf(endsWithRuleFailedMsg, strings.Join(r.suffixes, ", "))
}

func (r *EndsWithRule) Name() string {
	return endsWithRuleName
}
