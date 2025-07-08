package comparison

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	sameRuleName                     = "same"
	sameRuleDefaultMsg               = "the :attribute and :other must match"
	sameRuleFieldForCompareMissedMsg = "the :attribute field for same rule requires another field to compare against"
)

// SameRule checks if a field is the same as another field.
type SameRule struct {
	common.BaseRule
	otherField string
}

// NewSameRule creates a new SameRule.
func NewSameRule(parameters []string) (contract.Rule, error) {
	if len(parameters) < 1 {
		return nil, errors.New("same rule requires a field to compare against")
	}
	otherField := parameters[0]

	return &SameRule{
		BaseRule:   common.NewBaseRule(sameRuleName, sameRuleDefaultMsg, parameters),
		otherField: otherField,
	}, nil
}

// Validate checks if the field's value is the same as the other field's value.
func (r *SameRule) Validate(ctx contract.RuleContext) error {
	otherValue, ok := ctx.Data()[r.otherField]
	if !ok {
		return errors.New(sameRuleFieldForCompareMissedMsg)
	}

	// Using fmt.Sprintf to handle different types consistently during comparison.
	if fmt.Sprintf("%v", ctx.Value()) != fmt.Sprintf("%v", otherValue) {
		return errors.New(sameRuleDefaultMsg)
	}

	return nil
}

func (r *SameRule) Name() string {
	return sameRuleName
}
