package format

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	uuidRuleName                 = "uuid"
	uuidRuleDefaultMessage       = "the :attribute field must be a valid UUID"
	uuidRuleInvalidTypeMessage   = "the :attribute must be a string to validate as UUID"
	uuidRuleInvalidFormatMessage = "the :attribute must be a valid UUID format"
)

// Strict UUID pattern that requires dashes in the correct positions
var uuidPattern = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// UUIDRule validates that a value is a valid UUID string.
type UUIDRule struct {
	common.BaseRule
}

// NewUUIDRule creates a new UUID validation rule.
func NewUUIDRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	return &UUIDRule{
		BaseRule: common.NewBaseRule(uuidRuleName, uuidRuleDefaultMessage, parameters, options...),
	}, nil
}

// Validate checks whether the value is a valid UUID string.
func (r *UUIDRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val := ctx.Value()
	str, ok := val.(string)
	if !ok || str == "" {
		return errors.New(uuidRuleInvalidTypeMessage)
	}

	// First check strict format with dashes
	if !uuidPattern.MatchString(str) {
		return errors.New(uuidRuleInvalidFormatMessage)
	}

	// Then validate with uuid.Parse for additional validation
	if _, err := uuid.Parse(str); err != nil {
		return errors.New(uuidRuleInvalidFormatMessage)
	}

	return nil
}

func (r *UUIDRule) Name() string {
	return uuidRuleName
}
