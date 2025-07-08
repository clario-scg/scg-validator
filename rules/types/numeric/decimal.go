package numeric

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	decimalRuleName                       = "decimal"
	decimalRuleDefaultMessage             = "the :attribute must have :decimal decimal places"
	decimalRuleErrMsgRequiresParam        = "decimal rule requires at least one parameter"
	decimalRuleErrMsgInvalidMinParam      = "invalid min parameter for decimal rule: %w"
	decimalRuleErrMsgInvalidMaxParam      = "invalid max parameter for decimal rule: %w"
	decimalRuleErrMsgInvalidType          = "the :attribute must be a float or string"
	decimalRuleErrMsgParseFailed          = "the :attribute must be a valid float"
	decimalRuleErrMsgExactDecimalMismatch = "the :attribute must have exactly :decimal decimal places"
	decimalRuleErrMsgRangeMismatch        = "the :attribute must have between :min and :max decimal places"
)

type DecimalRule struct {
	common.BaseRule
	minDecimals int
	maxDecimals int
}

func NewDecimalRule(params []string) (contract.Rule, error) {
	if len(params) == 0 {
		return nil, errors.New(decimalRuleErrMsgRequiresParam)
	}

	minVal, err := strconv.Atoi(params[0])
	if err != nil {
		return nil, fmt.Errorf(decimalRuleErrMsgInvalidMinParam, err)
	}

	maxVal := -1
	if len(params) > 1 {
		maxVal, err = strconv.Atoi(params[1])
		if err != nil {
			return nil, fmt.Errorf(decimalRuleErrMsgInvalidMaxParam, err)
		}
	}

	return &DecimalRule{
		BaseRule:    common.NewBaseRule(decimalRuleName, decimalRuleDefaultMessage, params),
		minDecimals: minVal,
		maxDecimals: maxVal,
	}, nil
}

func (r *DecimalRule) Validate(ctx contract.RuleContext) error {
	var decimals int

	switch v := ctx.Value().(type) {
	case float64:
		decimals = r.countDecimalPlaces(v)
	case float32:
		decimals = r.countDecimalPlaces(float64(v))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		// Integer types have 0 decimal places
		decimals = 0
	case string:
		// For strings, count decimal places from the original string to preserve trailing zeros
		decimals = r.countDecimalPlacesFromString(v)
		// Also validate that it's a valid number
		if _, err := strconv.ParseFloat(v, 64); err != nil {
			return errors.New(decimalRuleErrMsgParseFailed)
		}
	default:
		return errors.New(decimalRuleErrMsgInvalidType)
	}

	// No decimal part, valid only if minDecimals == 0
	if decimals == 0 && r.minDecimals == 0 {
		return nil
	}

	if r.maxDecimals == -1 {
		if decimals == r.minDecimals {
			return nil
		}
		return fmt.Errorf("%s", strings.ReplaceAll(decimalRuleErrMsgExactDecimalMismatch, ":decimal", r.label()))
	}

	if decimals >= r.minDecimals && decimals <= r.maxDecimals {
		return nil
	}

	msg := strings.ReplaceAll(decimalRuleErrMsgRangeMismatch, ":min", strconv.Itoa(r.minDecimals))
	msg = strings.ReplaceAll(msg, ":max", strconv.Itoa(r.maxDecimals))
	return fmt.Errorf("%s", msg)
}

func (r *DecimalRule) countDecimalPlaces(f float64) int {
	parts := strings.Split(strconv.FormatFloat(f, 'f', -1, 64), ".")
	if len(parts) == 2 {
		return len(parts[1])
	}
	return 0
}

func (r *DecimalRule) countDecimalPlacesFromString(s string) int {
	parts := strings.Split(s, ".")
	if len(parts) == 2 {
		return len(parts[1])
	}
	return 0
}

func (r *DecimalRule) label() string {
	if r.maxDecimals == -1 {
		return strconv.Itoa(r.minDecimals)
	}
	return fmt.Sprintf("between %d and %d", r.minDecimals, r.maxDecimals)
}

func (r *DecimalRule) Name() string {
	return decimalRuleName
}
