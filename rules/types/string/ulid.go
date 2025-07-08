package string

import (
	"errors"
	"regexp"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	ulidRuleName             = "ulid"
	ulidRuleDefaultMessage   = "the :attribute must be a valid ULID"
	ulidRuleInvalidTypeError = "ulid validation failed: value must be a string"
	ulidRuleValidationFailed = "ulid validation failed: value does not match ULID format"
	ulidRulePattern          = `^[0-7][0-9A-HJKMNP-TV-Z]{25}$`
)

// UlidRule checks if a string is a valid ULID.
type UlidRule struct {
	common.BaseRule
}

// NewUlidRule creates a new UlidRule instance.
func NewUlidRule() (contract.Rule, error) {
	return &UlidRule{
		BaseRule: common.NewBaseRule(ulidRuleName, ulidRuleDefaultMessage, nil),
	}, nil
}

// Validate checks if the value conforms to the ULID format.
func (r *UlidRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok {
		return errors.New(ulidRuleInvalidTypeError)
	}

	matched, err := regexp.MatchString(ulidRulePattern, val)
	if err != nil || !matched {
		return errors.New(ulidRuleValidationFailed)
	}

	return nil
}
func (r *UlidRule) Name() string {
	return ulidRuleName
}
