package date

import (
	"errors"
	"fmt"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

// ComparisonType defines the type of date comparison to perform
type ComparisonType int

const (
	ComparisonAfter ComparisonType = iota
	ComparisonAfterOrEqual
	ComparisonBefore
	ComparisonBeforeOrEqual
	ComparisonEqual
)

// BaseDateComparisonRule provides common functionality for date comparison rules
type BaseDateComparisonRule struct {
	common.BaseRule
	comparisonDate     time.Time
	format             string
	comparisonType     ComparisonType
	ruleName           string
	typeErrorMsg       string
	validationErrorMsg string
}

// NewBaseDateComparisonRule creates a new base date comparison rule
func NewBaseDateComparisonRule(
	ruleName, defaultTemplate, missingParamError, parseError, typeError, validationError string,
	comparisonType ComparisonType,
	parameters []string,
) (*BaseDateComparisonRule, error) {
	if len(parameters) == 0 {
		return nil, errors.New(missingParamError)
	}

	format := time.RFC3339
	if len(parameters) > 1 {
		format = parameters[1]
	}

	parsedDate, err := time.Parse(format, parameters[0])
	if err != nil {
		return nil, fmt.Errorf(parseError, err)
	}

	return &BaseDateComparisonRule{
		BaseRule:           common.NewBaseRule(ruleName, defaultTemplate, parameters),
		comparisonDate:     parsedDate,
		format:             format,
		comparisonType:     comparisonType,
		ruleName:           ruleName,
		typeErrorMsg:       typeError,
		validationErrorMsg: validationError,
	}, nil
}

// Validate performs the date comparison validation
func (r *BaseDateComparisonRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok {
		return errors.New(r.typeErrorMsg)
	}

	parsedVal, err := time.Parse(r.format, val)
	if err != nil {
		return errors.New(r.typeErrorMsg)
	}

	if r.compareDate(parsedVal) {
		return nil
	}

	return errors.New(r.validationErrorMsg)
}

// compareDate performs the actual date comparison based on the comparison type
func (r *BaseDateComparisonRule) compareDate(value time.Time) bool {
	switch r.comparisonType {
	case ComparisonAfter:
		return value.After(r.comparisonDate)
	case ComparisonAfterOrEqual:
		return value.After(r.comparisonDate) || value.Equal(r.comparisonDate)
	case ComparisonBefore:
		return value.Before(r.comparisonDate)
	case ComparisonBeforeOrEqual:
		return value.Before(r.comparisonDate) || value.Equal(r.comparisonDate)
	case ComparisonEqual:
		return value.Equal(r.comparisonDate)
	default:
		return false
	}
}

func (r *BaseDateComparisonRule) Name() string {
	return r.ruleName
}
