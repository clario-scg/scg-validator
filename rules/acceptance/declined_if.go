package acceptance

import (
	"errors"
	"fmt"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	declinedIfRuleName            = "declined_if"
	declinedIfRuleDefaultMsg      = "the :attribute must be declined when :other is :value."
	declinedIfRuleMissingParamMsg = "declined_if rule requires exactly 2 parameters"
	declinedIfRuleFailedMsgFormat = "the :attribute must be declined when %s is %s"
)

var declinedIfAcceptedValues = []string{"no", "off", "0", "false"}

// DeclinedIfRule validates that a field is declined when another field matches a specific value.
type DeclinedIfRule struct {
	common.BaseRule
	conditionField string
	conditionValue string
}

// NewDeclinedIfRule creates the DeclinedIfRule instance.
func NewDeclinedIfRule(parameters []string) (contract.Rule, error) {
	if len(parameters) != 2 {
		return nil, errors.New(declinedIfRuleMissingParamMsg)
	}

	return &DeclinedIfRule{
		BaseRule:       common.NewBaseRule(declinedIfRuleName, declinedIfRuleDefaultMsg, parameters),
		conditionField: parameters[0],
		conditionValue: parameters[1],
	}, nil
}

// Validate checks if the rule condition is met, and ensures the value is declined.
func (r *DeclinedIfRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}
	data := ctx.Data()
	val := ctx.Value()

	condVal, exists := data[r.conditionField]
	if !exists {
		return nil // No condition match, skip check
	}

	condStr := fmt.Sprintf("%v", condVal)
	if condStr != r.conditionValue {
		return nil // Condition not met
	}

	switch v := val.(type) {
	case bool:
		if !v {
			return nil
		}
	case string:
		for _, d := range declinedIfAcceptedValues {
			if strings.EqualFold(v, d) {
				return nil
			}
		}
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		if fmt.Sprintf("%v", v) == "0" {
			return nil
		}
	}

	return fmt.Errorf(declinedIfRuleFailedMsgFormat, r.conditionField, r.conditionValue)
}

func (r *DeclinedIfRule) Name() string {
	return declinedIfRuleName
}
