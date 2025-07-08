package authentication

import (
	"errors"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/registry/password"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	currentPasswordRuleName       = "current_password"
	currentPasswordRuleDefaultMsg = "the :attribute is incorrect"
	currentPasswordMissingMsg     = "no PasswordVerifier registered or provided. " +
		"Please register one via registry.RegisterPasswordVerifier or inject it via context"
)

// CurrentPasswordRule checks if the provided password matches the current user’s password.
type CurrentPasswordRule struct {
	common.BaseRule
	verifier contract.PasswordVerifier
}

// NewCurrentPasswordRule returns a new instance of CurrentPasswordRule.
func NewCurrentPasswordRule() (contract.Rule, error) {
	return &CurrentPasswordRule{
		BaseRule: common.NewBaseRule(currentPasswordRuleName, currentPasswordRuleDefaultMsg, nil),
	}, nil
}

// SetVerifier allows manual injection of a PasswordVerifier.
func (r *CurrentPasswordRule) SetVerifier(verifier contract.PasswordVerifier) {
	r.verifier = verifier
}

// Validate checks whether the input matches the authenticated user’s current password.
func (r *CurrentPasswordRule) Validate(ctx contract.RuleContext) error {
	// Skip validation if the rule is marked as nullable and the value is nil.
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	// Ensure the provided value is a non-empty string.
	val, ok := ctx.Value().(string)
	if !ok || val == "" {
		return errors.New(currentPasswordRuleDefaultMsg)
	}

	// Resolve the PasswordVerifier to use for validation
	verifier := r.resolveVerifier(ctx)
	if verifier == nil {
		return errors.New(currentPasswordMissingMsg)
	}

	// Verify if the provided password matches the stored password
	match, err := verifier.Verify(val)
	if err != nil || !match {
		return errors.New(currentPasswordRuleDefaultMsg)
	}

	return nil
}

// resolveVerifier attempts to retrieve a PasswordVerifier from context or registry.
func (r *CurrentPasswordRule) resolveVerifier(ctx contract.RuleContext) contract.PasswordVerifier {
	// Use the injected PasswordVerifier if available
	if r.verifier != nil {
		return r.verifier
	}

	// Try to retrieve PasswordVerifier from the context (if available)
	if vCtx, ok := ctx.(interface {
		PasswordVerifier() contract.PasswordVerifier
	}); ok {
		return vCtx.PasswordVerifier()
	}

	// Fallback to global registry if no PasswordVerifier is provided or found in context
	if v, ok := password.FindPasswordVerifier("default"); ok {
		return v
	}

	// Return nil if no PasswordVerifier is found
	return nil
}

func (r *CurrentPasswordRule) Name() string {
	return currentPasswordRuleName
}
