package control

import (
	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	BailRuleName       = "bail"
	bailRuleDefaultMsg = "stop on first validation failure"
)

// bailRule is a control directive to stop validation on the first failure.
type bailRule struct {
	common.BaseRule
}

// NewBailRule creates a new instance of bailRule.
func NewBailRule() (contract.Rule, error) {
	return &bailRule{
		BaseRule: common.NewBaseRule(BailRuleName, bailRuleDefaultMsg, nil),
	}, nil
}

func (r *bailRule) Name() string {
	return BailRuleName
}

// Validate is a no-op; the bail rule behavior is enforced in the validation engine.
func (r *bailRule) Validate(_ contract.RuleContext) error {
	return nil
}
