package acceptance

import (
	"errors"
	"fmt"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	acceptedRuleName       = "accepted"
	acceptedRuleDefaultMsg = "the :attribute must be accepted"
)

// AcceptedRule checks whether a value is considered "accepted".
type AcceptedRule struct {
	common.BaseRule
}

// Accepted accepted values.
var acceptedValues = map[string]bool{
	"yes":  true,
	"on":   true,
	"1":    true,
	"true": true,
}

// NewAcceptedRule creates a new AcceptedRule instance.
func NewAcceptedRule() (contract.Rule, error) {
	return &AcceptedRule{
		BaseRule: common.NewBaseRule(acceptedRuleName, acceptedRuleDefaultMsg, nil),
	}, nil
}

// Validate checks if the given value is considered accepted.
func (r *AcceptedRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}
	val := ctx.Value()

	switch v := val.(type) {
	case bool:
		// A boolean `true` is accepted.
		if v {
			return nil
		}

	case string:
		// String matching with accepted values, case-insensitive and trimmed.
		trimmed := strings.TrimSpace(v)
		if acceptedValues[strings.ToLower(trimmed)] {
			return nil
		}

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		// For numeric types, check if the value equals "1" (truthy).
		if fmt.Sprintf("%v", v) == "1" {
			return nil
		}
	}

	return errors.New(acceptedRuleDefaultMsg)
}

func (r *AcceptedRule) Name() string {
	return acceptedRuleName
}
