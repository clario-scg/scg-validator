// Package boolean provides the boolean validator rule.
package boolean

import (
	"errors"
	"reflect"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	booleanRuleName         = "boolean"
	booleanRuleDefaultMsg   = "The :attribute must be a boolean value."
	booleanRuleInvalidValue = "the value is not a recognizable boolean"
)

// acceptedBooleanStrings defines accepted string representations of booleans.
var acceptedBooleanStrings = map[string]struct{}{
	"true":  {},
	"false": {},
	"1":     {},
	"0":     {},
	"yes":   {},
	"no":    {},
	"on":    {},
	"off":   {},
}

// Rule implements a rule to validate boolean values or representations.
type Rule struct {
	common.BaseRule
}

// NewBooleanRule creates a new Rule instance.
func NewBooleanRule() (contract.Rule, error) {
	return &Rule{
		BaseRule: common.NewBaseRule(booleanRuleName, booleanRuleDefaultMsg, nil),
	}, nil
}

// Validate checks if the value is a boolean or valid boolean representation.
func (r *Rule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val := ctx.Value()

	switch v := val.(type) {
	case bool:
		return nil
	case string:
		if _, ok := acceptedBooleanStrings[strings.ToLower(v)]; ok {
			return nil
		}
	case int, int8, int16, int32, int64:
		i := reflect.ValueOf(v).Int()
		if i == 0 || i == 1 {
			return nil
		}
		return errors.New(booleanRuleInvalidValue)
	case uint, uint8, uint16, uint32, uint64:
		u := reflect.ValueOf(v).Uint()
		if u == 0 || u == 1 {
			return nil
		}
		return errors.New(booleanRuleInvalidValue)
	case float32, float64:
		f := reflect.ValueOf(v).Float()
		if f == 0.0 || f == 1.0 {
			return nil
		}
		return errors.New(booleanRuleInvalidValue)
	}

	return errors.New(booleanRuleInvalidValue)
}

func (r *Rule) Name() string {
	return booleanRuleName
}
