package conditional

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	requiredWithoutRuleName           = "required_without"
	requiredWithoutRuleDefaultMsg     = "the :attribute field is required when :values is not present"
	requiredWithoutRuleInvalidDataMsg = "the :attribute field has not any provided data"
	requiredWithoutRuleParamErrorMsg  = "required_without rule requires at least one parameter"
)

// requiredWithoutRule checks if a field is required when any of the other fields are not present.
type requiredWithoutRule struct {
	common.BaseRule
	otherFields []string
}

// NewRequiredWithoutRule creates a new instance of requiredWithoutRule.
func NewRequiredWithoutRule(params []string) (contract.Rule, error) {
	if len(params) == 0 {
		return nil, errors.New(requiredWithoutRuleParamErrorMsg)
	}

	return &requiredWithoutRule{
		BaseRule:    common.NewBaseRule(requiredWithoutRuleName, requiredWithoutRuleDefaultMsg, params),
		otherFields: params,
	}, nil
}

func (r *requiredWithoutRule) Name() string {
	return requiredWithoutRuleName
}

// Validate checks if the field is required when none of the other fields are present.
func (r *requiredWithoutRule) Validate(ctx contract.RuleContext) error {
	data := ctx.Data()
	anyFieldPresent := false
	for _, field := range r.otherFields {
		if _, ok := data[field]; ok {
			anyFieldPresent = true
			break
		}
	}
	if anyFieldPresent {
		return nil // Not required if any of the other fields are present
	}

	// If none of the other fields are present, this field is required.
	value := ctx.Value()
	if value == nil {
		return errors.New(requiredWithoutRuleInvalidDataMsg)
	}
	if s, ok := value.(string); ok && s == "" {
		return errors.New(requiredWithoutRuleInvalidDataMsg)
	}
	return nil
}
