package conditional

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	prohibitedRuleName       = "prohibited"
	prohibitedRuleDefaultMsg = "the :attribute field is prohibited"
)

// prohibitedRule fails if the field is present in the input data.
type prohibitedRule struct {
	common.BaseRule
}

// NewProhibitedRule creates a new instance of prohibitedRule.
func NewProhibitedRule() (contract.Rule, error) {
	return &prohibitedRule{
		BaseRule: common.NewBaseRule(prohibitedRuleName, prohibitedRuleDefaultMsg, nil),
	}, nil
}

func (r *prohibitedRule) Name() string {
	return prohibitedRuleName
}

// Validate returns an error if the field exists in the data, regardless of value.
func (r *prohibitedRule) Validate(ctx contract.RuleContext) error {
	if _, exists := ctx.Data()[ctx.Field()]; exists {
		return errors.New(prohibitedRuleDefaultMsg)
	}
	return nil
}
