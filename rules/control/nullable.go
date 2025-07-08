package control

import (
	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	nullableRuleName       = "nullable"
	nullableRuleDefaultMsg = "the :attribute field is nullable"
)

// nullableRule allows null or missing values for a field.
type nullableRule struct {
	common.BaseRule
}

// NewNullableRule creates a new instance of nullableRule.
func NewNullableRule() (contract.Rule, error) {
	return &nullableRule{
		BaseRule: common.NewBaseRule(nullableRuleName, nullableRuleDefaultMsg, nil),
	}, nil
}

func (r *nullableRule) Name() string {
	return nullableRuleName
}

// Validate always passes — indicates the field is allowed to be nil or missing.
func (r *nullableRule) Validate(_ contract.RuleContext) error {
	return nil
}
