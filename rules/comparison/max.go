package comparison

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	maxRuleName         = "max"
	maxRuleDefaultMsg   = "the :attribute must not be greater than :value"
	maxRuleTypeErrorMsg = "the :attribute must have a numeric value"
)

// MaxRule is a validation rule that checks if a numeric value is not greater than a given value.
type MaxRule struct {
	common.BaseRule
	comparisonValue float64
}

// NewMaxRule creates a new MaxRule with a comparison threshold.
func NewMaxRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New("max rule requires a value parameter")
	}

	// Parse the threshold value
	val, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value parameter for max rule: %w", err)
	}

	return &MaxRule{
		BaseRule:        common.NewBaseRule(maxRuleName, maxRuleDefaultMsg, parameters),
		comparisonValue: val,
	}, nil
}

// Validate checks if the given value is not greater than the comparison value.
func (r *MaxRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if the value is nil or the validation is skipped
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Convert the value to a float
	value, err := getAsFloat(ctx.Value())
	if err != nil {
		return errors.New(maxRuleTypeErrorMsg)
	}

	// Check if the value is less than or equal to the maximum allowed
	if value <= r.comparisonValue {
		return nil
	}

	// Return error if the value exceeds the allowed maximum
	return errors.New(maxRuleDefaultMsg)
}

func (r *MaxRule) Name() string {
	return maxRuleName
}
