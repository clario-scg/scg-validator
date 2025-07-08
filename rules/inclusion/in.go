package inclusion

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	inRuleName                    = "in"
	inRuleDefaultMessageTemplate  = "the :attribute field must be one of: :param0"
	inRuleMissingParamsError      = "in rule requires at least one parameter"
	inRuleValidationFailedMessage = "the :attribute is not in the allowed list"
)

// InRule checks if a value is in a predefined list of allowed values.
type InRule struct {
	common.BaseRule
	allowed []string
}

// NewInRule creates a new instance of InRule with allowed parameters.
func NewInRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(inRuleMissingParamsError)
	}
	return &InRule{
		BaseRule: common.NewBaseRule(inRuleName, inRuleDefaultMessageTemplate, parameters, options...),
		allowed:  parameters,
	}, nil
}

// Validate checks whether the value exists in the allowed list.
func (r *InRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	value := fmt.Sprintf("%v", ctx.Value())
	for _, allowed := range r.allowed {
		if value == allowed {
			return nil
		}
	}

	return errors.New(inRuleValidationFailedMessage)
}

func (r *InRule) Name() string {
	return inRuleName
}
