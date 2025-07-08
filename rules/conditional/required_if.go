package conditional

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	requiredIfRuleName           = "required_if"
	requiredIfRuleDefaultMsg     = "the :attribute field is required when :other is :value"
	requiredIfRuleInvalidDataMsg = "the :attribute field has not any provided data"
	requiredIfRuleParamErrorMsg  = "required_if rule requires at least 2 parameters"
)

// requiredIfRule checks if a field is required when another field has a specific value.
type requiredIfRule struct {
	common.BaseRule
	conditionField string
	conditionValue string
}

// NewRequiredIfRule creates a new instance of requiredIfRule.
func NewRequiredIfRule(params []string) (contract.Rule, error) {
	if len(params) < 2 {
		return nil, errors.New(requiredIfRuleParamErrorMsg)
	}
	return &requiredIfRule{
		BaseRule:       common.NewBaseRule(requiredIfRuleName, requiredIfRuleDefaultMsg, params),
		conditionField: params[0],
		conditionValue: params[1],
	}, nil
}

func (r *requiredIfRule) Name() string {
	return requiredIfRuleName
}

// Validate checks if the field is required when the condition is met.
func (r *requiredIfRule) Validate(ctx contract.RuleContext) error {
	otherValue, exists := ctx.Data()[r.conditionField]
	if !exists || fmt.Sprintf("%v", otherValue) != r.conditionValue {
		return nil // Condition not met → field not required
	}

	value := ctx.Value()
	if value == nil {
		return errors.New(requiredIfRuleInvalidDataMsg)
	}
	if s, ok := value.(string); ok && s == "" {
		return errors.New(requiredIfRuleDefaultMsg)
	}
	return nil
}
