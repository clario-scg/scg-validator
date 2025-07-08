package control

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	presentRuleName       = "present"
	presentRuleDefaultMsg = "the :attribute field must be present"
)

type presentRule struct {
	common.BaseRule
}

// NewPresentRule creates a new instance of presentRule.
func NewPresentRule() (contract.Rule, error) {
	return &presentRule{
		BaseRule: common.NewBaseRule(presentRuleName, presentRuleDefaultMsg, nil),
	}, nil
}

func (r *presentRule) Name() string {
	return presentRuleName
}

// Validate fails if the field is not present in the data.
func (r *presentRule) Validate(ctx contract.RuleContext) error {
	if _, ok := ctx.Data()[ctx.Field()]; !ok {
		return errors.New(presentRuleDefaultMsg)
	}
	return nil
}
