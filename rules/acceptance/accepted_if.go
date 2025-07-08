package acceptance

import (
	"errors"
	"fmt"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	acceptedIfRuleName            = "accepted_if"
	acceptedIfRuleDefaultMsg      = "the :attribute must be accepted when :other is :value"
	acceptedIfRuleMissingParamMsg = "accepted_if rule requires at least 2 parameters"
	acceptedIfRuleFailedMsgFormat = "the :attribute must be accepted when :other is one of the following: %s"
)

var acceptedIfAcceptedValues = map[string]bool{
	"yes":  true,
	"on":   true,
	"1":    true,
	"true": true,
}

// AcceptedIfRule checks if a field is accepted if another field has specific values.
type AcceptedIfRule struct {
	common.BaseRule
	conditionField  string
	conditionValues map[string]bool
}

// NewAcceptedIfRule creates a new AcceptedIfRule instance.
func NewAcceptedIfRule(parameters []string) (contract.Rule, error) {
	if len(parameters) < 2 {
		return nil, errors.New(acceptedIfRuleMissingParamMsg)
	}

	conditionValues := make(map[string]bool)
	for _, v := range parameters[1:] {
		conditionValues[v] = true
	}

	return &AcceptedIfRule{
		BaseRule:        common.NewBaseRule(acceptedIfRuleName, acceptedIfRuleDefaultMsg, parameters),
		conditionField:  parameters[0],
		conditionValues: conditionValues,
	}, nil
}

// Validate ensures the field is accepted if the condition is satisfied.
func (r *AcceptedIfRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}
	data := ctx.Data()
	fieldValue := ctx.Value()

	// Check if the condition field exists
	condVal, exists := data[r.conditionField]
	if !exists {
		return nil // Condition not met
	}

	condStr := fmt.Sprintf("%v", condVal)
	if _, ok := r.conditionValues[condStr]; !ok {
		return nil // Condition not met
	}

	// Validate the field based on the accepted values
	switch v := fieldValue.(type) {
	case bool:
		if v {
			return nil
		}
	case string:
		trimmed := strings.TrimSpace(v)
		if acceptedIfAcceptedValues[strings.ToLower(trimmed)] {
			return nil
		}
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		if fmt.Sprintf("%v", v) == "1" {
			return nil
		}
	}

	// Error message with dynamic field and condition values
	return fmt.Errorf(acceptedIfRuleFailedMsgFormat, strings.Join(r.conditionValuesToList(), ", "))
}

// helper function to convert condition values map to a list
func (r *AcceptedIfRule) conditionValuesToList() []string {
	var values []string
	for key := range r.conditionValues {
		values = append(values, key)
	}
	return values
}

func (r *AcceptedIfRule) Name() string {
	return acceptedIfRuleName
}
