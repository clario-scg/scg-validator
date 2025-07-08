package string

import (
	"errors"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	uppercaseRuleName = "uppercase"

	uppercaseRuleDefaultMessage = "value must be uppercase"
	uppercaseRuleTypeError      = "value must be a string"
	uppercaseRuleFailMessage    = "value must be uppercase"
)

// UppercaseRule validates that a string is entirely uppercase.
type UppercaseRule struct {
	common.BaseRule
}

// NewUppercaseRule creates a new instance of UppercaseRule.
func NewUppercaseRule() (contract.Rule, error) {
	return &UppercaseRule{
		BaseRule: common.NewBaseRule(uppercaseRuleName, uppercaseRuleDefaultMessage, nil),
	}, nil
}

// Validate checks if the string is fully uppercase.
func (r *UppercaseRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	str, ok := ctx.Value().(string)
	if !ok {
		return errors.New(uppercaseRuleTypeError)
	}

	if str != strings.ToUpper(str) {
		return errors.New(uppercaseRuleFailMessage)
	}

	return nil
}

func (r *UppercaseRule) Name() string {
	return uppercaseRuleName
}
