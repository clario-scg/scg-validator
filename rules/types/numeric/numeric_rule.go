package numeric

import (
	"errors"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	numericRuleName       = "numeric"
	numericRuleDefaultMsg = "the :attribute must be numeric"
	numericRuleErrorMsg   = "the :attribute must be numeric"
)

// Rule checks whether a value is numeric (int, float, or numeric string).
type Rule struct {
	common.BaseRule
}

// NewNumericRule creates a new instance of Rule.
func NewNumericRule() (contract.Rule, error) {
	return &Rule{
		BaseRule: common.NewBaseRule(numericRuleName, numericRuleDefaultMsg, nil),
	}, nil
}

// Validate checks if the value is numeric.
func (r *Rule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	switch v := ctx.Value().(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return nil
	case string:
		if _, err := strconv.ParseFloat(v, 64); err == nil {
			return nil
		}
	}

	return errors.New(numericRuleErrorMsg)
}

func (r *Rule) Name() string {
	return numericRuleName
}
