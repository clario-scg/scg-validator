package acceptance

import (
	"errors"
	"fmt"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	declinedRuleName       = "declined"
	declinedRuleDefaultMsg = "the :attribute must be declined"
)

var declinedAcceptedValues = []string{"no", "off", "0", "false"}

// DeclinedRule checks if the value represents a declined state.
type DeclinedRule struct {
	common.BaseRule
}

// NewDeclinedRule constructs a new DeclinedRule.
func NewDeclinedRule() (contract.Rule, error) {
	return &DeclinedRule{
		BaseRule: common.NewBaseRule(declinedRuleName, declinedRuleDefaultMsg, nil),
	}, nil
}

// Validate returns an error if the value is not considered declined.
func (r *DeclinedRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}
	val := ctx.Value()

	switch v := val.(type) {
	case bool:
		if !v {
			return nil
		}
	case string:
		for _, accepted := range declinedAcceptedValues {
			if strings.EqualFold(v, accepted) {
				return nil
			}
		}
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		if fmt.Sprintf("%v", v) == "0" {
			return nil
		}
	}

	return errors.New(declinedRuleDefaultMsg)
}

func (r *DeclinedRule) Name() string {
	return declinedRuleName
}
