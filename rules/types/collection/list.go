package collection

import (
	"errors"
	"reflect"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	listRuleName               = "list"
	listRuleDefaultMessage     = "the :attribute field must be a list (slice or array)"
	listRuleInvalidTypeMessage = "the :attribute must be a slice or array type"
)

// ListRule validates that a value is a slice or an array.
type ListRule struct {
	common.BaseRule
}

// NewListRule creates a new ListRule instance.
func NewListRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	return &ListRule{
		BaseRule: common.NewBaseRule(listRuleName, listRuleDefaultMessage, parameters, options...),
	}, nil
}

// Validate checks whether the value is a slice or array.
func (r *ListRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	value := ctx.Value()
	if value == nil {
		return errors.New(listRuleInvalidTypeMessage)
	}

	kind := reflect.TypeOf(value).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return errors.New(listRuleInvalidTypeMessage)
	}

	return nil
}

func (r *ListRule) Name() string {
	return listRuleName
}
