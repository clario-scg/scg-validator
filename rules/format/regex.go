package format

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	regexRuleName             = "regex"
	regexRuleDefaultMessage   = "the :attribute field format is invalid"
	regexRuleInvalidParamMsg  = "regex rule expects a valid pattern as its first parameter"
	regexRuleInvalidTypeMsg   = "the :attribute must be a string to validate with regex"
	regexRuleMismatchErrorMsg = "the :attribute does not match the required pattern"
)

// RegexRule validates a string against a regular expression pattern.
type RegexRule struct {
	common.BaseRule
	pattern *regexp.Regexp
}

// NewRegexRule creates a new RegexRule with the given pattern parameter.
func NewRegexRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(regexRuleInvalidParamMsg)
	}

	pat, err := regexp.Compile(parameters[0])
	if err != nil {
		return nil, fmt.Errorf(regexRuleInvalidParamMsg+": %w", err)
	}

	return &RegexRule{
		BaseRule: common.NewBaseRule(regexRuleName, regexRuleDefaultMessage, parameters, options...),
		pattern:  pat,
	}, nil
}

// Validate checks whether the value matches the regex pattern.
func (r *RegexRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	str, ok := ctx.Value().(string)
	if !ok {
		return errors.New(regexRuleInvalidTypeMsg)
	}

	if !r.pattern.MatchString(str) {
		return errors.New(regexRuleMismatchErrorMsg)
	}

	return nil
}

func (r *RegexRule) Name() string {
	return regexRuleName
}
