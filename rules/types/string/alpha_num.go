package string

import (
	"errors"
	"unicode"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	alphaNumRuleName       = "alpha_num"
	alphaNumRuleDefaultMsg = "the :attribute must contain only alphanumeric characters"

	alphaNumErrInvalidType = "the :attribute must be a string"
	alphaNumErrEmpty       = "the :attribute must not be empty"
	alphaNumErrFormat      = "the :attribute must contain only alphanumeric characters"
)

// AlphaNumRule checks that a string contains only letters and digits (Unicode-aware).
type AlphaNumRule struct {
	common.BaseRule
}

// NewAlphaNumRule creates a new AlphaNumRule instance.
func NewAlphaNumRule(params []string, opts ...common.RuleOption) (contract.Rule, error) {
	return &AlphaNumRule{
		BaseRule: common.NewBaseRule(alphaNumRuleName, alphaNumRuleDefaultMsg, params, opts...),
	}, nil
}

// Validate ensures the value is a string of only letters and digits.
func (r *AlphaNumRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	str, ok := ctx.Value().(string)
	if !ok {
		return errors.New(alphaNumErrInvalidType)
	}
	if str == "" {
		return errors.New(alphaNumErrEmpty)
	}

	for _, ch := range str {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return errors.New(alphaNumErrFormat)
		}
	}

	return nil
}

func (r *AlphaNumRule) Name() string {
	return alphaNumRuleName
}
