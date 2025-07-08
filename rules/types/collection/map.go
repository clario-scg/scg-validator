package collection

import (
	"errors"
	"reflect"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	mapRuleName               = "map"
	mapRuleDefaultMessage     = "the :attribute field must be a map"
	mapRuleInvalidTypeMessage = "the :attribute must be a map type"
)

// MapRule validates that a value is a map.
type MapRule struct {
	common.BaseRule
}

// NewMapRule creates a new MapRule instance.
func NewMapRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	return &MapRule{
		BaseRule: common.NewBaseRule(mapRuleName, mapRuleDefaultMessage, parameters, options...),
	}, nil
}

// Validate checks whether the value is a map.
func (r *MapRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	value := ctx.Value()
	if value == nil {
		return errors.New(mapRuleInvalidTypeMessage)
	}

	if reflect.TypeOf(value).Kind() != reflect.Map {
		return errors.New(mapRuleInvalidTypeMessage)
	}

	return nil
}

func (r *MapRule) Name() string {
	return mapRuleName
}
