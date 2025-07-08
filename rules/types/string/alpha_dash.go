package string

import (
	"errors"
	"unicode"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	alphaDashRuleName           = "alpha_dash"
	alphaDashRuleDefaultMsg     = "the :attribute may only contain letters, numbers, dashes, and underscores"
	alphaDashRuleInvalidTypeMsg = "the :attribute must be a string"
	alphaDashRuleEmptyStringMsg = "the :attribute must not be empty"
	alphaDashRuleInvalidCharMsg = "the :attribute may only contain letters, numbers, dashes, and underscores"
)

// AlphaDashRule checks if a string only contains letters, digits, dashes, or underscores.
type AlphaDashRule struct {
	common.BaseRule
}

// NewAlphaDashRule creates a new instance of AlphaDashRule.
func NewAlphaDashRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	return &AlphaDashRule{
		BaseRule: common.NewBaseRule(alphaDashRuleName, alphaDashRuleDefaultMsg, parameters, options...),
	}, nil
}

// Validate ensures the string only contains allowed characters.
func (r *AlphaDashRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	str, ok := ctx.Value().(string)
	if !ok {
		return errors.New(alphaDashRuleInvalidTypeMsg)
	}
	if str == "" {
		return errors.New(alphaDashRuleEmptyStringMsg)
	}

	for _, ch := range str {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '-' && ch != '_' {
			return errors.New(alphaDashRuleInvalidCharMsg)
		}
	}

	return nil
}

func (r *AlphaDashRule) Name() string {
	return alphaDashRuleName
}
