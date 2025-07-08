package date

import (
	"github.com/next-trace/scg-validator/contract"
)

const (
	beforeOrEqualRuleName             = "before_or_equal"
	beforeOrEqualRuleDefaultMsg       = "The :attribute must be a date before or equal to :date."
	beforeOrEqualRuleMissingParamErr  = "before_or_equal rule requires a date parameter"
	beforeOrEqualRuleInvalidFormatErr = "invalid date format for before_or_equal rule: %w"
	beforeOrEqualRuleInvalidTypeErr   = "the value must be a string to validate as a date"
	beforeOrEqualRuleFailedErr        = "the date must be before or equal to the comparison date"
)

// BeforeOrEqualRule checks if a date is before or equal to a given date.
type BeforeOrEqualRule struct {
	*BaseDateComparisonRule
}

// NewBeforeOrEqualRule creates the rule instance.
func NewBeforeOrEqualRule(parameters []string) (contract.Rule, error) {
	base, err := NewBaseDateComparisonRule(
		beforeOrEqualRuleName,
		beforeOrEqualRuleDefaultMsg,
		beforeOrEqualRuleMissingParamErr,
		beforeOrEqualRuleInvalidFormatErr,
		beforeOrEqualRuleInvalidTypeErr,
		beforeOrEqualRuleFailedErr,
		ComparisonBeforeOrEqual,
		parameters,
	)
	if err != nil {
		return nil, err
	}

	return &BeforeOrEqualRule{
		BaseDateComparisonRule: base,
	}, nil
}
