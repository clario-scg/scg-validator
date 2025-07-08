package date

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	dateEqualsRuleName           = "date_equals"
	dateEqualsDefaultMsg         = "the :attribute must be a date equal to :date"
	dateEqualsMissingParamMsg    = "date_equals rule requires a date parameter"
	dateEqualsInvalidFormatError = "invalid date format for date_equals rule: %w"
)

// EqualsRule validates that a value is a date equal to a target date.
type EqualsRule struct {
	common.BaseRule
	comparisonDate time.Time
	format         string
}

// NewDateEqualsRule creates a new EqualsRule with the given parameters.
func NewDateEqualsRule(parameters []string) (contract.Rule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(dateEqualsMissingParamMsg)
	}

	dateStr := parameters[0]
	format := time.RFC3339
	if len(parameters) > 1 && parameters[1] != "" {
		format = parameters[1]
	}

	parsedDate, err := time.Parse(format, dateStr)
	if err != nil {
		return nil, fmt.Errorf(dateEqualsInvalidFormatError, err)
	}

	return &EqualsRule{
		BaseRule:       common.NewBaseRule(dateEqualsRuleName, dateEqualsDefaultMsg, parameters),
		comparisonDate: parsedDate,
		format:         format,
	}, nil
}

// Validate checks if the context value equals the comparison date.
func (r *EqualsRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	valueStr, ok := ctx.Value().(string)
	if !ok {
		return errors.New(dateEqualsDefaultMsg)
	}

	parsedValue, err := time.Parse(r.format, valueStr)
	if err != nil {
		return fmt.Errorf(dateEqualsInvalidFormatError, err)
	}

	if parsedValue.Equal(r.comparisonDate) {
		return nil
	}

	return fmt.Errorf("%s", strings.Replace(dateEqualsDefaultMsg, ":date", r.Parameters()[0], 1))
}

func (r *EqualsRule) Name() string {
	return dateEqualsRuleName
}
