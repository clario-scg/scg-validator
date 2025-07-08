package conditional

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	requiredWithAllRuleName           = "required_with_all"
	requiredWithAllRuleDefaultMsg     = "the :attribute field is required when :values are present"
	requiredWithAllRuleInvalidDataMsg = "the :attribute field has not any provided data"
	requiredWithAllRuleParamErrorMsg  = "required_with_all rule requires at least one field"
)

// requiredWithAllRule checks if a field is required when all other specified fields are present.
type requiredWithAllRule struct {
	common.BaseRule
	otherFields []string
}

// NewRequiredWithAllRule creates a new instance of requiredWithAllRule.
func NewRequiredWithAllRule(params []string) (contract.Rule, error) {
	if len(params) == 0 {
		return nil, errors.New(requiredWithAllRuleParamErrorMsg)
	}
	return &requiredWithAllRule{
		BaseRule:    common.NewBaseRule(requiredWithAllRuleName, requiredWithAllRuleDefaultMsg, params),
		otherFields: params,
	}, nil
}

func (r *requiredWithAllRule) Name() string {
	return requiredWithAllRuleName
}

// Validate checks if this field is required when all the other fields are present.
func (r *requiredWithAllRule) Validate(ctx contract.RuleContext) error {
	data := ctx.Data()
	for _, field := range r.otherFields {
		if _, ok := data[field]; !ok {
			return nil // If any required field is missing, this rule passes
		}
	}

	// All required fields are present; now this field must be non-empty
	value := ctx.Value()
	if value == nil {
		return errors.New(requiredWithAllRuleInvalidDataMsg)
	}
	if s, ok := value.(string); ok && s == "" {
		return errors.New(requiredWithAllRuleDefaultMsg)
	}
	return nil
}
