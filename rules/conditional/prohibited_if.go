package conditional

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	prohibitedIfRuleName       = "prohibited_if"
	prohibitedIfRuleDefaultMsg = "the :attribute field is prohibited when :other is :value"
)

// prohibitedIfRule checks if the field is prohibited when another field has a specific value.
type prohibitedIfRule struct {
	common.BaseRule
	otherField string
	value      string
}

// NewProhibitedIfRule creates a new instance of prohibitedIfRule.
func NewProhibitedIfRule(params []string) (contract.Rule, error) {
	if len(params) < 2 {
		return nil, errors.New("prohibited_if rule requires exactly two parameters: other_field and value")
	}

	return &prohibitedIfRule{
		BaseRule:   common.NewBaseRule(prohibitedIfRuleName, prohibitedIfRuleDefaultMsg, params),
		otherField: params[0],
		value:      params[1],
	}, nil
}

func (r *prohibitedIfRule) Name() string {
	return prohibitedIfRuleName
}

// Validate returns an error if the other field has the expected value and the current field is present.
func (r *prohibitedIfRule) Validate(ctx contract.RuleContext) error {
	data := ctx.Data()
	field := ctx.Field()

	otherValue, ok := data[r.otherField]
	if !ok || fmt.Sprintf("%v", otherValue) != r.value {
		return nil // Other field is not present or doesn't match → pass
	}

	if _, present := data[field]; present {
		return errors.New(prohibitedIfRuleDefaultMsg)
	}

	return nil
}
