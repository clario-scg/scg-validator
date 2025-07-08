package comparison

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	differentRuleName       = "different"
	differentRuleDefaultMsg = "the :attribute and :other must be different"
	differentRuleParamError = "different rule requires a field to compare against"
)

// DifferentRule checks if a field is different from another field.
type DifferentRule struct {
	common.BaseRule
	otherField string
}

// NewDifferentRule creates a new DifferentRule.
func NewDifferentRule(parameters []string) (contract.Rule, error) {
	if len(parameters) < 1 {
		return nil, errors.New(differentRuleParamError)
	}
	return &DifferentRule{
		BaseRule:   common.NewBaseRule(differentRuleName, differentRuleDefaultMsg, parameters),
		otherField: parameters[0],
	}, nil
}

// Validate ensures the field is different from the specified other field.
func (r *DifferentRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if the value is nil or validation is skipped
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Retrieve the value of the other field to compare against
	otherVal, ok := ctx.Data()[r.otherField]
	if !ok {
		return errors.New(differentRuleParamError)
	}

	// Compare the values, return error if they are equal
	if fmt.Sprintf("%v", ctx.Value()) == fmt.Sprintf("%v", otherVal) {
		return errors.New(differentRuleDefaultMsg)
	}

	return nil
}

func (r *DifferentRule) Name() string {
	return differentRuleName
}
