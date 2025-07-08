package numeric

import (
	"errors"
	"math"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	integerRuleName        = "integer"
	integerRuleDefaultMsg  = "the :attribute must be an integer"
	integerRuleInvalidType = "the :attribute must be an integer"
)

// IntegerRule checks if a value is a valid integer type.
type IntegerRule struct {
	common.BaseRule
}

// NewIntegerRule creates an instance of IntegerRule.
func NewIntegerRule() (contract.Rule, error) {
	return &IntegerRule{
		BaseRule: common.NewBaseRule(integerRuleName, integerRuleDefaultMsg, nil),
	}, nil
}

// Validate confirms whether the input value is a valid integer.
func (r *IntegerRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	switch v := ctx.Value().(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return nil

	case string:
		if _, err := strconv.Atoi(v); err == nil {
			return nil
		}
		return errors.New(integerRuleInvalidType)

	case float32:
		if math.Mod(float64(v), 1) == 0 {
			return nil
		}
		return errors.New(integerRuleInvalidType)

	case float64:
		if math.Mod(v, 1) == 0 {
			return nil
		}
		return errors.New(integerRuleInvalidType)

	default:
		return errors.New(integerRuleInvalidType)
	}
}

func (r *IntegerRule) Name() string {
	return integerRuleName
}
