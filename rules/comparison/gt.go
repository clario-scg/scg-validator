package comparison

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	gtRuleName         = "gt"
	gtRuleDefaultMsg   = "the :attribute must be greater than :value"
	gtRuleTypeErrorMsg = "the :attribute must have a numeric value"
)

// GtRule validates that a numeric value is greater than a specified threshold.
type GtRule struct {
	common.BaseRule
	threshold float64
}

// NewGtRule creates a new GtRule with a comparison threshold.
func NewGtRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New("gt rule requires a value parameter")
	}

	val, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value parameter for gt rule: %w", err)
	}

	return &GtRule{
		BaseRule:  common.NewBaseRule(gtRuleName, gtRuleDefaultMsg, parameters),
		threshold: val,
	}, nil
}

// Validate checks if the input value is greater than the configured threshold.
func (r *GtRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if the value is nil or validation is skipped
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Convert value to float64
	value, err := getAsFloat(ctx.Value())
	if err != nil {
		return errors.New(gtRuleTypeErrorMsg)
	}

	// Compare the value with the threshold
	if value <= r.threshold {
		return errors.New(gtRuleDefaultMsg)
	}

	return nil
}

func (r *GtRule) Name() string {
	return gtRuleName
}
