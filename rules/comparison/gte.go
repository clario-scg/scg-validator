package comparison

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	gteRuleName                 = "gte"
	gteRuleDefaultMessage       = "the :attribute must be greater than or equal to :value"
	gteRuleMissingParamError    = "gte rule requires a value parameter"
	gteRuleInvalidParamError    = "invalid value parameter for gte rule: %w"
	gteRuleInvalidInputType     = "the :attribute must be a numeric value"
	gteRuleValidationFailedFmt  = "the :attribute must be greater than or equal to %s"
	gteRuleFloatFormatPrecision = "%.6f"
)

// GteRule validates that the input is >= to the given comparison value.
type GteRule struct {
	common.BaseRule
	comparisonValue float64
}

// NewGteRule initializes a new GteRule from parameters.
func NewGteRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(gteRuleMissingParamError)
	}

	val, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf(gteRuleInvalidParamError, err)
	}

	return &GteRule{
		BaseRule:        common.NewBaseRule(gteRuleName, gteRuleDefaultMessage, parameters),
		comparisonValue: val,
	}, nil
}

// Validate checks if the input value is greater than or equal to the comparison value.
func (r *GteRule) Validate(ctx contract.RuleContext) error {
	value, err := getAsComparable(ctx.Value())
	if err != nil {
		return errors.New(gteRuleInvalidInputType)
	}

	if value >= r.comparisonValue {
		return nil
	}

	return fmt.Errorf(gteRuleValidationFailedFmt, floatToString(r.comparisonValue))
}

func (r *GteRule) Name() string {
	return gteRuleName
}
