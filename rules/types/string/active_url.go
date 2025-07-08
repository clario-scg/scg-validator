package string

import (
	"errors"
	"net"
	"net/url"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
	"github.com/next-trace/scg-validator/rules/format"
)

const (
	activeURLRuleName          = "active_url"
	activeURLRuleDefaultMsg    = "the :attribute is not a valid, active url"
	activeURLRuleInvalidType   = "the :attribute must be a string url"
	activeURLRuleResolutionErr = "the :attribute could not be resolved"
)

// ActiveURLRule checks if a given string is a valid and resolvable URL.
type ActiveURLRule struct {
	common.BaseRule
}

// NewActiveURLRule creates a new ActiveURLRule.
func NewActiveURLRule() (contract.Rule, error) {
	return &ActiveURLRule{
		BaseRule: common.NewBaseRule(activeURLRuleName, activeURLRuleDefaultMsg, nil),
	}, nil
}

// Validate checks if the input is a well-formed URL and DNS-resolvable.
func (r *ActiveURLRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok {
		return errors.New(activeURLRuleInvalidType)
	}

	urlRule, _ := format.NewURLRule(ctx.Parameters())
	if err := urlRule.Validate(ctx); err != nil {
		return err
	}

	parsed, _ := url.Parse(val)
	if _, err := net.LookupHost(parsed.Host); err != nil {
		return errors.New(activeURLRuleResolutionErr)
	}

	return nil
}

func (r *ActiveURLRule) Name() string {
	return activeURLRuleName
}
