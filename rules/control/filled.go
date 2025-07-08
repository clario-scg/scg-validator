package control

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	filledRuleName       = "filled"
	filledRuleDefaultMsg = "the :attribute field must have a value when it is present"
	filledRuleEmptyMsg   = "the :attribute field has not any provided data"
)

type filledRule struct {
	common.BaseRule
}

// NewFilledRule constructs a new filledRule.
func NewFilledRule() (contract.Rule, error) {
	return &filledRule{
		BaseRule: common.NewBaseRule(filledRuleName, filledRuleDefaultMsg, nil),
	}, nil
}

func (r *filledRule) Name() string {
	return filledRuleName
}

// Validate ensures the field is filled only if it is present in input data.
func (r *filledRule) Validate(ctx contract.RuleContext) error {
	_, fieldPresent := ctx.Data()[ctx.Field()]
	if !fieldPresent {
		return nil
	}

	value := ctx.Value()
	if value == nil {
		return errors.New(filledRuleEmptyMsg)
	}

	if str, ok := value.(string); ok && str == "" {
		return errors.New(filledRuleDefaultMsg)
	}

	return nil
}
