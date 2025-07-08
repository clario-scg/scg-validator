package string

import (
	"errors"
	"regexp"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	asciiRuleName           = "ascii"
	asciiRuleDefaultMsg     = "the :attribute must only contain ASCII characters"
	asciiRuleErrorMsg       = "the :attribute must only contain ASCII characters"
	asciiRuleInvalidTypeMsg = "the :attribute must be a string"
	asciiRegexPattern       = `^[[:ascii:]]*$`
)

// ASCIIRule checks if a string only contains ASCII characters.
type ASCIIRule struct {
	common.BaseRule
}

// NewASCIIRule creates a new instance of ASCIIRule.
func NewASCIIRule() (contract.Rule, error) {
	return &ASCIIRule{
		BaseRule: common.NewBaseRule(asciiRuleName, asciiRuleDefaultMsg, nil),
	}, nil
}

// Validate ensures the input value contains only ASCII characters.
func (r *ASCIIRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok {
		return errors.New(asciiRuleInvalidTypeMsg)
	}

	matched, err := regexp.MatchString(asciiRegexPattern, val)
	if err != nil || !matched {
		return errors.New(asciiRuleErrorMsg)
	}

	return nil
}

func (r *ASCIIRule) Name() string {
	return asciiRuleName
}
