package numeric

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
	"github.com/next-trace/scg-validator/utils"
)

const (
	multipleOfRuleName                    = "multiple_of"
	multipleOfRuleDefaultMsg              = "the :attribute must be a multiple of :value"
	multipleOfRuleMissingParamMsg         = "multiple_of rule requires a value parameter"
	multipleOfRuleInvalidParamMsg         = "invalid value parameter for multiple_of rule: %w"
	multipleOfRuleZeroParamMsg            = "the multiple_of parameter cannot be zero"
	multipleOfRuleFailedMsg               = "the :attribute must be a multiple of %s"
	multipleOfRuleInvalidInputMsg         = "the :attribute must be a numeric value"
	floatEqualityTolerance        float64 = 1e-9
)

// MultipleOfRule checks if a value is a multiple of a given float.
type MultipleOfRule struct {
	common.BaseRule
	multiple float64
}

// NewMultipleOfRule constructs a new rule for checking multiples.
func NewMultipleOfRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(multipleOfRuleMissingParamMsg)
	}

	val, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf(multipleOfRuleInvalidParamMsg, err)
	}

	if val == 0 {
		return nil, errors.New(multipleOfRuleZeroParamMsg)
	}

	return &MultipleOfRule{
		BaseRule: common.NewBaseRule(multipleOfRuleName, multipleOfRuleDefaultMsg, parameters),
		multiple: val,
	}, nil
}

// Validate verifies whether the context value is a multiple of configured `r.multiple`.
func (r *MultipleOfRule) Validate(ctx contract.RuleContext) error {
	// Handle nil values explicitly
	if ctx.Value() == nil {
		return errors.New(multipleOfRuleInvalidInputMsg)
	}

	var value float64
	var err error

	switch v := ctx.Value().(type) {
	case float64:
		value = v
	case float32:
		value = float64(v)
	case int:
		value = float64(v)
	case int8:
		value = float64(v)
	case int16:
		value = float64(v)
	case int32:
		value = float64(v)
	case int64:
		value = float64(v)
	case uint:
		value = float64(v)
	case uint8:
		value = float64(v)
	case uint16:
		value = float64(v)
	case uint32:
		value = float64(v)
	case uint64:
		value = float64(v)
	case string:
		// Parse string as numeric value
		value, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return errors.New(multipleOfRuleInvalidInputMsg)
		}
	default:
		return errors.New(multipleOfRuleInvalidInputMsg)
	}

	if isMultiple(value, r.multiple) {
		return nil
	}

	return fmt.Errorf(multipleOfRuleFailedMsg, utils.FloatToString(r.multiple))
}

// isMultiple checks if `value` is a multiple of `factor` using a tolerance.
func isMultiple(value, factor float64) bool {
	if factor == 0 {
		return false
	}
	result := value / factor
	return math.Abs(result-math.Round(result)) < floatEqualityTolerance
}

func (r *MultipleOfRule) Name() string {
	return multipleOfRuleName
}
