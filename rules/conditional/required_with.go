package conditional

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	requiredWithRuleName           = "required_with"
	requiredWithRuleDefaultMsg     = "the :attribute field is required when :values is present"
	requiredWithRuleInvalidDataMsg = "the :attribute field has not any provided data"
	requiredWithRuleParamErrorMsg  = "required_with rule requires at least one parameter"
)

// requiredWithRule checks if a field is required when another field is present.
type requiredWithRule struct {
	common.BaseRule
	otherFields []string
}

// NewRequiredWithRule creates a new instance of requiredWithRule.
func NewRequiredWithRule(params []string) (contract.Rule, error) {
	if len(params) == 0 {
		return nil, errors.New(requiredWithRuleParamErrorMsg)
	}
	return &requiredWithRule{
		BaseRule:    common.NewBaseRule(requiredWithRuleName, requiredWithRuleDefaultMsg, params),
		otherFields: params,
	}, nil
}

func (r *requiredWithRule) Name() string {
	return requiredWithRuleName
}

// Validate checks if this field is required when any of the other fields are present.
func (r *requiredWithRule) Validate(ctx contract.RuleContext) error {
	data := ctx.Data()
	anyFieldPresent := false
	for _, field := range r.otherFields {
		if _, ok := data[field]; ok {
			anyFieldPresent = true
			break
		}
	}
	if !anyFieldPresent {
		return nil // None of the other fields are present; not required
	}

	// If any of the other fields are present, this field is required.
	value := ctx.Value()
	if value == nil {
		return errors.New(requiredWithRuleInvalidDataMsg)
	}
	if s, ok := value.(string); ok && s == "" {
		return errors.New(requiredWithRuleDefaultMsg)
	}
	return nil
}
