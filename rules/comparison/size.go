package comparison

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	sizeRuleName       = "size"
	sizeRuleDefaultMsg = "the :attribute must be :value"
	sizeRuleParamError = "size rule parameter must be numeric"
)

// SizeRule is a validation rule that checks if a value's size matches a given value.
type SizeRule struct {
	common.BaseRule
	size float64
}

// NewSizeRule creates a new SizeRule.
func NewSizeRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New("size rule requires a value parameter")
	}

	// Parse the provided size parameter
	val, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf("size rule parameter must be numeric: %v", err)
	}

	return &SizeRule{
		BaseRule: common.NewBaseRule(sizeRuleName, sizeRuleDefaultMsg, parameters),
		size:     val,
	}, nil
}

// Validate checks if the given value's size matches the specified size.
func (r *SizeRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if needed
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Get the actual value as a float
	actualValue, err := getAsFloat(ctx.Value())
	if err != nil {
		return errors.New(sizeRuleParamError)
	}

	// Check if the size matches the specified size
	if actualValue != r.size {
		return errors.New(sizeRuleDefaultMsg)
	}

	return nil
}

func (r *SizeRule) Name() string {
	return sizeRuleName
}
