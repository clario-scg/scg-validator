package string

import (
	"errors"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	lowercaseRuleName           = "lowercase"
	lowercaseRuleDefaultMsg     = "the :attribute must be lowercase"
	lowercaseRuleInvalidTypeMsg = "the :attribute must be a string"
	lowercaseRuleFailedMsg      = "the :attribute must be lowercase"
)

// LowercaseRule checks if a string is entirely lowercase.
type LowercaseRule struct {
	common.BaseRule
}

// NewLowercaseRule creates an instance of LowercaseRule.
func NewLowercaseRule() (contract.Rule, error) {
	return &LowercaseRule{
		BaseRule: common.NewBaseRule(lowercaseRuleName, lowercaseRuleDefaultMsg, nil),
	}, nil
}

// Validate checks if the string is fully lowercase.
func (r *LowercaseRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok {
		return errors.New(lowercaseRuleInvalidTypeMsg)
	}

	if val != strings.ToLower(val) {
		return errors.New(lowercaseRuleFailedMsg)
	}

	return nil
}

func (r *LowercaseRule) Name() string {
	return lowercaseRuleName
}
