package conditional

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	requiredWithoutAllRuleName           = "required_without_all"
	requiredWithoutAllRuleDefaultMsg     = "the :attribute field is required when none of :values are present"
	requiredWithoutAllRuleInvalidDataMsg = "the :attribute field has not any provided data"
	requiredWithoutAllRuleParamErrorMsg  = "required_without_all rule requires at least one field"
)

// requiredWithoutAllRule checks if the field is required when all other specified fields are not present.
type requiredWithoutAllRule struct {
	common.BaseRule
	otherFields []string
}

// NewRequiredWithoutAllRule creates a new instance of requiredWithoutAllRule.
func NewRequiredWithoutAllRule(params []string) (contract.Rule, error) {
	if len(params) == 0 {
		return nil, errors.New(requiredWithoutAllRuleParamErrorMsg)
	}
	return &requiredWithoutAllRule{
		BaseRule:    common.NewBaseRule(requiredWithoutAllRuleName, requiredWithoutAllRuleDefaultMsg, params),
		otherFields: params,
	}, nil
}

// Validate checks if the field is required when all other fields are not present.
func (r *requiredWithoutAllRule) Validate(ctx contract.RuleContext) error {
	data := ctx.Data()
	value := ctx.Value()

	// Count how many of the other fields are present
	presentCount := 0
	if data != nil {
		for _, field := range r.otherFields {
			if _, ok := data[field]; ok {
				presentCount++
			}
		}
	}

	// If ALL other fields are missing (presentCount == 0), then this field is required
	if presentCount == 0 {
		// Field is required - check if it has a value
		if value == nil {
			return errors.New(requiredWithoutAllRuleInvalidDataMsg)
		}
		if s, ok := value.(string); ok && s == "" {
			return errors.New(requiredWithoutAllRuleInvalidDataMsg)
		}
	}

	// If any other field is present, this field is not required, so it always passes
	return nil
}
