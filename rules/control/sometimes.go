package control

import (
	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	sometimesRuleName       = "sometimes"
	sometimesRuleDefaultMsg = "the :attribute field is conditionally validated"
)

// sometimesRule is a control directive that tells the engine to validate only if the field is present.
type sometimesRule struct {
	common.BaseRule
}

// NewSometimesRule creates a new instance of sometimesRule.
func NewSometimesRule() (contract.Rule, error) {
	return &sometimesRule{
		BaseRule: common.NewBaseRule(sometimesRuleName, sometimesRuleDefaultMsg, nil),
	}, nil
}

func (r *sometimesRule) Name() string {
	return sometimesRuleName
}

// Validate is a no-op; it's handled by the validation engine.
func (r *sometimesRule) Validate(_ contract.RuleContext) error {
	return nil
}
