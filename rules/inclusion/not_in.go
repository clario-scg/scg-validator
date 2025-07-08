package inclusion

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	notInRuleName                    = "not_in"
	notInRuleDefaultMessageTemplate  = "the :attribute field must not be one of: :param0"
	notInRuleMissingParamsError      = "not_in rule requires at least one parameter"
	notInRuleValidationFailedMessage = "the :attribute is in the forbidden list"
)

// NotInRule checks if a value is NOT in a set of forbidden values.
type NotInRule struct {
	common.BaseRule
	forbidden []string
}

// NewNotInRule creates a new instance of NotInRule.
func NewNotInRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(notInRuleMissingParamsError)
	}

	return &NotInRule{
		BaseRule:  common.NewBaseRule(notInRuleName, notInRuleDefaultMessageTemplate, parameters, options...),
		forbidden: parameters,
	}, nil
}

// Validate checks whether the value is NOT in the forbidden list.
func (r *NotInRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	value := fmt.Sprintf("%v", ctx.Value())
	for _, f := range r.forbidden {
		if value == f {
			return errors.New(notInRuleValidationFailedMessage)
		}
	}
	return nil
}

func (r *NotInRule) Name() string {
	return notInRuleName
}
