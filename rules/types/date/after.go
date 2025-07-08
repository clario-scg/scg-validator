package date

import (
	"errors"
	"fmt"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	afterRuleName                  = "after"
	afterRuleDefaultTemplate       = "the :attribute must be a date after :date"
	afterRuleMissingParamError     = "after rule requires a date parameter"
	afterRuleInvalidFormatError    = "invalid date format for after rule: %w"
	afterRuleValueMustBeDateError  = "the value must be a date string in the expected format"
	afterRuleComparisonFailedError = "the date must be after the comparison date"
)

// AfterRule validates that a value is a date after a given reference date.
type AfterRule struct {
	common.BaseRule
	comparisonDate time.Time
	format         string
}

// NewAfterRule constructs a new AfterRule.
// parameters[0] = comparison date string (required)
// parameters[1] = format (optional, defaults to RFC3339)
func NewAfterRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(afterRuleMissingParamError)
	}

	format := time.RFC3339
	if len(parameters) > 1 {
		format = parameters[1]
	}

	parsedDate, err := time.Parse(format, parameters[0])
	if err != nil {
		return nil, fmt.Errorf(afterRuleInvalidFormatError, err)
	}

	return &AfterRule{
		BaseRule:       common.NewBaseRule(afterRuleName, afterRuleDefaultTemplate, parameters),
		comparisonDate: parsedDate,
		format:         format,
	}, nil
}

// Validate ensures the given value is a valid date string and is after the configured comparison date.
func (r *AfterRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	str, ok := ctx.Value().(string)
	if !ok {
		return errors.New(afterRuleValueMustBeDateError)
	}

	parsedValue, err := time.Parse(r.format, str)
	if err != nil {
		return errors.New(afterRuleValueMustBeDateError)
	}

	if !parsedValue.After(r.comparisonDate) {
		return errors.New(afterRuleComparisonFailedError)
	}

	return nil
}

func (r *AfterRule) Name() string {
	return afterRuleName
}
