package conditional

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	prohibitedUnlessRuleName          = "prohibited_unless"
	prohibitedUnlessRuleDefaultMsg    = "the :attribute field is prohibited unless :other is in :values"
	prohibitedUnlessRuleParametersMsg = "prohibited_unless rule requires exactly two parameters: other_field and value"
)

// prohibitedUnlessRule checks if a field is prohibited unless another field has a specific value.
type prohibitedUnlessRule struct {
	common.BaseRule
	otherField string
	value      string
}

// NewProhibitedUnlessRule creates a new instance of prohibitedUnlessRule.
func NewProhibitedUnlessRule(params []string) (contract.Rule, error) {
	if len(params) < 2 {
		return nil, errors.New(prohibitedUnlessRuleParametersMsg)
	}

	return &prohibitedUnlessRule{
		BaseRule:   common.NewBaseRule(prohibitedUnlessRuleName, prohibitedUnlessRuleDefaultMsg, params),
		otherField: params[0],
		value:      params[1],
	}, nil
}

func (r *prohibitedUnlessRule) Name() string {
	return prohibitedUnlessRuleName
}

// Validate returns an error if the field is present and the other field does not match the allowed value.
func (r *prohibitedUnlessRule) Validate(ctx contract.RuleContext) error {
	data := ctx.Data()
	field := ctx.Field()

	otherValue, ok := data[r.otherField]

	if ok && fmt.Sprintf("%v", otherValue) == r.value {
		return nil // allowed: other field has allowed value
	}

	if _, present := data[field]; present {
		return errors.New(prohibitedUnlessRuleDefaultMsg)
	}

	return nil
}
