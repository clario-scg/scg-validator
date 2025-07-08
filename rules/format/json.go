package format

import (
	"encoding/json"
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	jsonRuleName                 = "json"
	jsonRuleDefaultMessage       = "the :attribute field must be a valid JSON string"
	jsonRuleInvalidTypeMessage   = "the :attribute must be a string to validate as JSON"
	jsonRuleInvalidFormatMessage = "the :attribute must be a valid JSON string"
)

// JSONRule validates that a field contains a valid JSON string.
type JSONRule struct {
	common.BaseRule
}

// NewJSONRule constructs a new JSONRule.
func NewJSONRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	return &JSONRule{
		BaseRule: common.NewBaseRule(jsonRuleName, jsonRuleDefaultMessage, parameters, options...),
	}, nil
}

// Validate checks whether the input is a valid JSON string.
func (r *JSONRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	raw := ctx.Value()
	str, ok := raw.(string)
	if !ok {
		return errors.New(jsonRuleInvalidTypeMessage)
	}

	var js any
	if err := json.Unmarshal([]byte(str), &js); err != nil {
		return errors.New(jsonRuleInvalidFormatMessage)
	}

	return nil
}

func (r *JSONRule) Name() string {
	return jsonRuleName
}
