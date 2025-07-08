package date

import (
	"errors"
	"fmt"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	dateRuleName        = "date"
	dateRuleDefaultMsg  = "the :attribute is not a valid date"
	dateRuleParseErrMsg = "invalid date format for date rule: %w"
)

// Rule checks if a value is a valid date string based on a given format.
type Rule struct {
	common.BaseRule
	format string
}

// NewDateRule creates a new Rule with an optional custom time format.
func NewDateRule(parameters []string) (contract.Rule, error) {
	format := time.RFC3339
	if len(parameters) > 0 && parameters[0] != "" {
		format = parameters[0]
	}

	return &Rule{
		BaseRule: common.NewBaseRule(dateRuleName, dateRuleDefaultMsg, parameters),
		format:   format,
	}, nil
}

// Validate checks if the context value is a string and a valid date in the specified format.
func (r *Rule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	strVal, ok := ctx.Value().(string)
	if !ok {
		return errors.New(dateRuleDefaultMsg)
	}

	if _, err := time.Parse(r.format, strVal); err != nil {
		return fmt.Errorf(dateRuleParseErrMsg, err)
	}

	return nil
}

func (r *Rule) Name() string {
	return dateRuleName
}
