package date

import (
	"github.com/next-trace/scg-validator/contract"
)

const (
	afterOrEqualRuleName              = "after_or_equal"
	afterOrEqualRuleDefaultTemplate   = "The :attribute must be a date after or equal to :date."
	afterOrEqualRuleMissingParamError = "after_or_equal rule requires a date parameter"
	afterOrEqualRuleParseError        = "invalid date format for after_or_equal rule: %w"
	afterOrEqualRuleTypeError         = "the value must be a string to validate as a date"
	afterOrEqualRuleValidationError   = "the date must be after or equal to the comparison date"
)

// AfterOrEqualRule checks if the input date is after or equal to a specific date.
type AfterOrEqualRule struct {
	*BaseDateComparisonRule
}

// NewAfterOrEqualRule creates a new rule instance.
// parameters[0] = reference date, parameters[1] = optional format (defaults to RFC3339)
func NewAfterOrEqualRule(parameters []string) (contract.Rule, error) {
	base, err := NewBaseDateComparisonRule(
		afterOrEqualRuleName,
		afterOrEqualRuleDefaultTemplate,
		afterOrEqualRuleMissingParamError,
		afterOrEqualRuleParseError,
		afterOrEqualRuleTypeError,
		afterOrEqualRuleValidationError,
		ComparisonAfterOrEqual,
		parameters,
	)
	if err != nil {
		return nil, err
	}

	return &AfterOrEqualRule{
		BaseDateComparisonRule: base,
	}, nil
}
