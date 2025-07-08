package conditional

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	requiredUnlessRuleName           = "required_unless"
	requiredUnlessRuleDefaultMsg     = "the :attribute field is required unless :other is :value"
	requiredUnlessRuleInvalidDataMsg = "the :attribute field has not any provided data"
	requiredUnlessRuleParamErrorMsg  = "required_unless rule requires at least 2 parameters"
)

// requiredUnlessRule checks if a field is required unless another field has a specific value.
type requiredUnlessRule struct {
	common.BaseRule
	conditionField string
	conditionValue string
}

// NewRequiredUnlessRule creates a new instance of requiredUnlessRule.
func NewRequiredUnlessRule(params []string) (contract.Rule, error) {
	if len(params) < 2 {
		return nil, errors.New(requiredUnlessRuleParamErrorMsg)
	}
	return &requiredUnlessRule{
		BaseRule:       common.NewBaseRule(requiredUnlessRuleName, requiredUnlessRuleDefaultMsg, params),
		conditionField: params[0],
		conditionValue: params[1],
	}, nil
}

func (r *requiredUnlessRule) Name() string {
	return requiredUnlessRuleName
}

// Validate checks if the field is required unless the other field has the specified value.
func (r *requiredUnlessRule) Validate(ctx contract.RuleContext) error {
	otherValue, exists := ctx.Data()[r.conditionField]
	if exists && fmt.Sprintf("%v", otherValue) == r.conditionValue {
		return nil // condition met, field not required
	}

	value := ctx.Value()
	if value == nil {
		return errors.New(requiredUnlessRuleInvalidDataMsg)
	}
	if s, ok := value.(string); ok && s == "" {
		return errors.New(requiredUnlessRuleDefaultMsg)
	}
	return nil
}
