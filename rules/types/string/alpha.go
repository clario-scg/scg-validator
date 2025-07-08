package string

import (
	"errors"
	"unicode"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
	"golang.org/x/text/unicode/norm"
)

const (
	alphaRuleName           = "alpha"
	alphaRuleDefaultMsg     = "The :attribute must contain only alphabetic characters."
	alphaRuleInvalidTypeMsg = "the :attribute must be a string"
	alphaRuleEmptyStringMsg = "the :attribute must not be empty"
	alphaRuleInvalidCharMsg = "the :attribute must contain only letters"
)

// AlphaRule checks if a string contains only alphabetic characters (Unicode-aware).
type AlphaRule struct {
	common.BaseRule
}

// NewAlphaRule creates an AlphaRule instance.
func NewAlphaRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	return &AlphaRule{
		BaseRule: common.NewBaseRule(alphaRuleName, alphaRuleDefaultMsg, parameters, options...),
	}, nil
}

// Validate checks if the input is a string containing only alphabetic characters.
func (r *AlphaRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	str, ok := ctx.Value().(string)
	if !ok {
		return errors.New(alphaRuleInvalidTypeMsg)
	}
	if str == "" {
		return errors.New(alphaRuleEmptyStringMsg)
	}

	normStr := norm.NFC.String(str)
	for _, ch := range normStr {
		if !unicode.IsLetter(ch) && !unicode.IsMark(ch) {
			return errors.New(alphaRuleInvalidCharMsg)
		}
	}

	return nil
}

func (r *AlphaRule) Name() string {
	return alphaRuleName
}
