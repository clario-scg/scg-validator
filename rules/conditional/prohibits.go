package conditional

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	prohibitsRuleName       = "prohibits"
	prohibitsRuleDefaultMsg = "the :attribute field prohibits :other from being present"
	prohibitsRuleParamErr   = "prohibits rule requires at least one parameter"
)

// prohibitsRule fails if any of the specified fields are present when the validated field is present.
type prohibitsRule struct {
	common.BaseRule
	otherFields []string
}

// NewProhibitsRule constructs a prohibitsRule with given other fields.
func NewProhibitsRule(params []string) (contract.Rule, error) {
	if len(params) == 0 {
		return nil, errors.New(prohibitsRuleParamErr)
	}

	return &prohibitsRule{
		BaseRule:    common.NewBaseRule(prohibitsRuleName, prohibitsRuleDefaultMsg, params),
		otherFields: params,
	}, nil
}

func (r *prohibitsRule) Name() string {
	return prohibitsRuleName
}

// Validate checks that none of the other fields are present if the current field is set.
func (r *prohibitsRule) Validate(ctx contract.RuleContext) error {
	data := ctx.Data()
	field := ctx.Field()

	// If the current field is absent, rule passes
	if _, exists := data[field]; !exists {
		return nil
	}

	// Current field is present → check that none of the prohibited fields are present
	for _, other := range r.otherFields {
		if _, conflict := data[other]; conflict {
			return errors.New(prohibitsRuleDefaultMsg)
		}
	}

	return nil
}
