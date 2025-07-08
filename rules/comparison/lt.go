package comparison

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
	"github.com/next-trace/scg-validator/utils"
)

const (
	ltRuleName         = "lt"
	ltRuleDefaultMsg   = "the :attribute must be less than :value"
	ltRuleTypeErrorMsg = "the :attribute must have a numeric value"
)

// LtRule validates that a numeric value is less than the specified comparison value.
type LtRule struct {
	common.BaseRule
	comparisonValue float64
}

// NewLtRule creates a new instance of LtRule with a comparison threshold.
func NewLtRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New("lt rule requires a value parameter")
	}

	// Parse the comparison value to float
	val, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value parameter for lt rule: %w", err)
	}

	return &LtRule{
		BaseRule:        common.NewBaseRule(ltRuleName, ltRuleDefaultMsg, parameters),
		comparisonValue: val,
	}, nil
}

// Validate checks if the given value is less than the comparison value.
func (r *LtRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if the value is nil or validation is skipped
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Convert the value to a float64
	value, err := utils.GetAsComparable(ctx.Value())
	if err != nil {
		return errors.New(ltRuleTypeErrorMsg)
	}

	// Check if the value is less than the comparison threshold
	if value < r.comparisonValue {
		return nil
	}

	// Return an error if the value is not less than the comparison value
	return errors.New(ltRuleDefaultMsg)
}

func (r *LtRule) Name() string {
	return ltRuleName
}
