package format

import (
	"errors"
	"net/url"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	urlRuleName                 = "url"
	urlRuleDefaultMessage       = "the :attribute field must be a valid URL"
	urlRuleInvalidTypeMessage   = "the :attribute must be a string to validate as URL"
	urlRuleInvalidFormatMessage = "the :attribute must be a properly formatted URL (with scheme and host)"
)

// URLRule validates that a field is a valid absolute URL.
type URLRule struct {
	common.BaseRule
}

// NewURLRule creates a new instance of URLRule.
func NewURLRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	return &URLRule{
		BaseRule: common.NewBaseRule(urlRuleName, urlRuleDefaultMessage, parameters, options...),
	}, nil
}

// Validate performs the URL format validation.
func (r *URLRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	raw := ctx.Value()
	str, ok := raw.(string)
	if !ok || str == "" {
		return errors.New(urlRuleInvalidTypeMessage)
	}

	parsed, err := url.ParseRequestURI(str)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return errors.New(urlRuleInvalidFormatMessage)
	}

	return nil
}

func (r *URLRule) Name() string {
	return urlRuleName
}
