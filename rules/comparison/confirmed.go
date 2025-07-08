package comparison

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	confirmedRuleName                       = "confirmed"
	confirmedRuleDefaultMsg                 = "the :attribute confirmation does not match"
	confirmedRuleConfirmationFieldMissedMsg = "the :attribute confirmation field is missing"
	confirmedRuleDataNotProvidedMsg         = "the :attribute has not provided data for confirmation validation"
)

// ConfirmedRule validates that a field matches its <field>_confirmation field.
type ConfirmedRule struct {
	common.BaseRule
}

// NewConfirmedRule creates a new instance of the ConfirmedRule.
func NewConfirmedRule() (contract.Rule, error) {
	return &ConfirmedRule{
		BaseRule: common.NewBaseRule(confirmedRuleName, confirmedRuleDefaultMsg, []string{}),
	}, nil
}

// Validate ensures the field matches the <field>_confirmation value in the input data.
func (r *ConfirmedRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if the value is nil
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Ensure we have data to validate
	data := ctx.Data()
	if data == nil {
		return errors.New(confirmedRuleDataNotProvidedMsg)
	}

	// Create the confirmation field name by appending "_confirmation" to the original field name
	confirmationField := ctx.Field() + "_confirmation"
	confirmationValue, exists := data[confirmationField]
	if !exists {
		return errors.New(confirmedRuleConfirmationFieldMissedMsg)
	}

	// Compare the field's value with the confirmation value
	if fmt.Sprintf("%v", ctx.Value()) != fmt.Sprintf("%v", confirmationValue) {
		return errors.New(confirmedRuleDefaultMsg)
	}

	return nil
}

func (r *ConfirmedRule) Name() string {
	return confirmedRuleName
}
