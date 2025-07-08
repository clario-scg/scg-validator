package comparison

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	minRuleName         = "min"
	minRuleDefaultMsg   = "the :attribute must be at least :value"
	minRuleTypeErrorMsg = "the :attribute must have a numeric value"
)

// MinRule validates that a numeric value is at least the specified minimum value.
type MinRule struct {
	common.BaseRule
	comparisonValue float64
}

// NewMinRule creates a new MinRule with a comparison threshold.
func NewMinRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New("min rule requires a value parameter")
	}

	// Parse the threshold value
	val, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value parameter for min rule: %w", err)
	}

	return &MinRule{
		BaseRule:        common.NewBaseRule(minRuleName, minRuleDefaultMsg, parameters),
		comparisonValue: val,
	}, nil
}

// Validate checks if the given value is at least the comparison value.
func (r *MinRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if the value is nil or validation is skipped
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Convert the value to a float
	value, err := getAsFloat(ctx.Value())
	if err != nil {
		return errors.New(minRuleTypeErrorMsg)
	}

	// Check if the value is greater than or equal to the minimum value
	if value >= r.comparisonValue {
		return nil
	}

	// Return error if the value is less than the allowed minimum
	return errors.New(minRuleDefaultMsg)
}

func (r *MinRule) Name() string {
	return minRuleName
}
