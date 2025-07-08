package comparison

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	lteRuleName                = "lte"
	lteRuleDefaultMessage      = "the :attribute must be less than or equal to :value"
	lteRuleErrMissingParam     = "lte rule requires a value parameter"
	lteRuleErrInvalidParam     = "invalid value parameter for lte rule: %w"
	lteRuleErrInvalidInputType = "the :attribute must be a numeric value"
	lteRuleErrFailed           = "the :attribute must be less than or equal to %s"
)

// LteRule checks if a numeric value is <= comparisonValue.
type LteRule struct {
	common.BaseRule
	comparisonValue float64
}

// NewLteRule constructs a new lteRule instance.
func NewLteRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(lteRuleErrMissingParam)
	}

	val, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf(lteRuleErrInvalidParam, err)
	}

	return &LteRule{
		BaseRule:        common.NewBaseRule(lteRuleName, lteRuleDefaultMessage, parameters),
		comparisonValue: val,
	}, nil
}

// Validate checks if value <= comparisonValue.
func (r *LteRule) Validate(ctx contract.RuleContext) error {
	value, err := getAsComparable(ctx.Value())
	if err != nil {
		return errors.New(lteRuleErrInvalidInputType)
	}

	if value <= r.comparisonValue {
		return nil
	}

	return fmt.Errorf(lteRuleErrFailed, floatToString(r.comparisonValue))
}

func (r *LteRule) Name() string {
	return lteRuleName
}
