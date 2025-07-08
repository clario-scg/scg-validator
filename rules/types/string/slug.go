package string

import (
	"errors"
	"regexp"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	slugRuleName           = "slug"
	slugRuleDefaultMsg     = "the :attribute must be a valid slug"
	slugRuleErrorMsg       = "the :attribute must be a valid slug"
	slugRuleInvalidTypeMsg = "the :attribute must be a string"
	slugRegexPattern       = `^[a-z0-9]+(?:-[a-z0-9]+)*$`
)

var slugCompiledRegex = regexp.MustCompile(slugRegexPattern)

// SlugRule checks if a string is a valid slug (lowercase letters, numbers, and dashes).
type SlugRule struct {
	common.BaseRule
}

// NewSlugRule creates a new instance of SlugRule.
func NewSlugRule() (contract.Rule, error) {
	return &SlugRule{
		BaseRule: common.NewBaseRule(slugRuleName, slugRuleDefaultMsg, nil),
	}, nil
}

// Validate returns an error if the value is not a valid slug.
func (r *SlugRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok {
		return errors.New(slugRuleInvalidTypeMsg)
	}

	if !slugCompiledRegex.MatchString(val) {
		return errors.New(slugRuleErrorMsg)
	}

	return nil
}

func (r *SlugRule) Name() string {
	return slugRuleName
}
