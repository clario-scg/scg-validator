package date

import (
	"errors"
	"fmt"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	beforeRuleName                    = "before"
	beforeRuleDefaultTemplate         = "The :attribute must be a date before :date."
	beforeRuleMissingParamError       = "before rule requires a date parameter"
	beforeRuleInvalidFormatError      = "invalid date format for before rule: %w"
	beforeRuleInvalidTypeError        = "the value must be a string to validate as a date"
	beforeRuleValidationFailedMessage = "the date must be before the comparison date"
)

// BeforeRule checks if the given value is before a specific comparison date.
type BeforeRule struct {
	common.BaseRule
	comparisonDate time.Time
	format         string
}

// NewBeforeRule constructs a new BeforeRule.
// parameters[0] = comparison date string
// parameters[1] = optional format (defaults to RFC3339)
func NewBeforeRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(beforeRuleMissingParamError)
	}

	format := time.RFC3339
	if len(parameters) > 1 {
		format = parameters[1]
	}

	refDate, err := time.Parse(format, parameters[0])
	if err != nil {
		return nil, fmt.Errorf(beforeRuleInvalidFormatError, err)
	}

	return &BeforeRule{
		BaseRule:       common.NewBaseRule(beforeRuleName, beforeRuleDefaultTemplate, parameters),
		comparisonDate: refDate,
		format:         format,
	}, nil
}

// Validate ensures the value is a string and a date before the comparison date.
func (r *BeforeRule) Validate(ctx contract.RuleContext) error {
	raw := ctx.Value()
	strVal, ok := raw.(string)
	if !ok {
		return errors.New(beforeRuleInvalidTypeError)
	}

	parsed, err := time.Parse(r.format, strVal)
	if err != nil {
		return errors.New(beforeRuleInvalidTypeError)
	}

	if parsed.Before(r.comparisonDate) {
		return nil
	}

	return errors.New(beforeRuleValidationFailedMessage)
}

func (r *BeforeRule) Name() string {
	return beforeRuleName
}
