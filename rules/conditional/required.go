package conditional

import (
	"errors"
	"reflect"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	requiredRuleName       = "required"
	requiredRuleDefaultMsg = "the :attribute field is required"
)

// requiredRule checks if a value is present and not empty.
type requiredRule struct {
	common.BaseRule
}

func NewRequiredRule() (contract.Rule, error) {
	return &requiredRule{
		BaseRule: common.NewBaseRule(requiredRuleName, requiredRuleDefaultMsg, nil),
	}, nil
}

func (r *requiredRule) Name() string {
	return requiredRuleName
}

// Validate checks for non-nil and non-empty value, idiomatic for all major Go types.
func (r *requiredRule) Validate(ctx contract.RuleContext) error {
	value := ctx.Value()
	if value == nil {
		return errors.New(requiredRuleDefaultMsg)
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if val.Len() == 0 {
			return errors.New(requiredRuleDefaultMsg)
		}
	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return errors.New(requiredRuleDefaultMsg)
		}
		// For non-nil pointers, check if the dereferenced value is zero
		if val.Kind() == reflect.Ptr {
			elem := val.Elem()
			if elem.IsZero() {
				return errors.New(requiredRuleDefaultMsg)
			}
		}
	default:
		if val.IsZero() {
			return errors.New(requiredRuleDefaultMsg)
		}
	}
	return nil
}
