package comparison

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	betweenRuleName         = "between"
	betweenRuleDefaultMsg   = "the :attribute must be between :min and :max"
	betweenRuleParamErr     = "between rule requires exactly two numeric parameters"
	betweenRuleMinParseFail = "between rule min parameter must be numeric: %w"
	betweenRuleMaxParseFail = "between rule max parameter must be numeric: %w"
	betweenRuleFailed       = "the :attribute must be between %v and %v"
	betweenRuleTypeErrorMsg = "value must be numeric"
)

// getAsFloat converts various types to a float64 for size comparison.
// For strings, it returns the rune count. For slices, arrays, and maps, it returns the length.
// For numeric types, it returns the float64 value.
func getAsFloat(value interface{}) (float64, error) {
	if value == nil {
		return 0, nil
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		return float64(utf8.RuneCountInString(val.String())), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return float64(val.Len()), nil
	}

	return 0, fmt.Errorf("unsupported type for comparison: %T", value)
}

// getAsComparable converts various types to a float64 for comparison rules.
// It tries to parse strings as numbers first, but falls back to length if not numeric.
// For collections (slices, maps, arrays), it returns the length.
// For numeric types, it returns the numeric value.
func getAsComparable(value interface{}) (float64, error) {
	if value == nil {
		return 0, errors.New("cannot convert nil to comparable value")
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		str := val.String()
		// Try to parse as number first
		if parsed, err := strconv.ParseFloat(str, 64); err == nil {
			return parsed, nil
		}
		// Fall back to string length (rune count)
		return float64(utf8.RuneCountInString(str)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return float64(val.Len()), nil
	}

	return 0, fmt.Errorf("unsupported type for comparison: %T", value)
}

// floatToString formats a float64 with 6 decimal places and trims trailing zeros
func floatToString(f float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.6f", f), "0"), ".")
}

// BetweenRule validates that a numeric value is between min and max (inclusive).
type BetweenRule struct {
	common.BaseRule
	min float64
	max float64
}

// NewBetweenRule creates a new BetweenRule with min and max parameters.
func NewBetweenRule(parameters []string) (contract.Rule, error) {
	if len(parameters) < 2 {
		return nil, errors.New(betweenRuleParamErr)
	}

	minVal, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, fmt.Errorf(betweenRuleMinParseFail, err)
	}
	maxVal, err := strconv.ParseFloat(parameters[1], 64)
	if err != nil {
		return nil, fmt.Errorf(betweenRuleMaxParseFail, err)
	}

	return &BetweenRule{
		BaseRule: common.NewBaseRule(betweenRuleName, betweenRuleDefaultMsg, parameters),
		min:      minVal,
		max:      maxVal,
	}, nil
}

// Validate checks whether the value is within the [min, max] range.
func (r *BetweenRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if value is nil or skipped by previous logic
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Convert value to float64
	value, err := getAsFloat(ctx.Value())
	if err != nil {
		return errors.New(betweenRuleTypeErrorMsg)
	}

	// Check if value is within the valid range
	if value >= r.min && value <= r.max {
		return nil
	}

	return fmt.Errorf(betweenRuleFailed, r.min, r.max)
}

func (r *BetweenRule) Name() string {
	return betweenRuleName
}
