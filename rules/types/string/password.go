package string

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	passwordRuleName       = "password"
	passwordRuleDefaultMsg = "the :attribute is not strong enough"

	// parameter keys
	passwordRuleParamMin       = "min"
	passwordRuleParamSymbols   = "symbols"
	passwordRuleParamNumbers   = "numbers"
	passwordRuleParamLetters   = "letters"
	passwordRuleParamMixed     = "mixedcase"
	passwordRuleParamUppercase = "uppercase"
	passwordRuleParamLowercase = "lowercase"

	// error messages
	passwordRuleErrMin    = "password must be at least %d characters long"
	passwordRuleErrLetter = "password must contain at least one letter"
	passwordRuleErrMixed  = "password must contain both uppercase and lowercase letters"
	passwordRuleErrUpper  = "password must contain at least one uppercase letter"
	passwordRuleErrLower  = "password must contain at least one lowercase letter"
	passwordRuleErrNumber = "password must contain at least one number"
	passwordRuleErrSymbol = "password must contain at least one symbol"

	passwordRuleInvalidMin = "invalid min value for password rule: %w" // #nosec G101

	passwordRuleDefaultMinLength = 8
)

// PasswordRule validates a password against complexity constraints.
type PasswordRule struct {
	common.BaseRule
	minLength      int
	requireLetters bool
	requireMixed   bool
	requireUpper   bool
	requireLower   bool
	requireNumbers bool
	requireSymbols bool
}

// NewPasswordRule constructs the PasswordRule from parameters.
func NewPasswordRule(parameters []string, options ...common.RuleOption) (contract.Rule, error) {
	r := &PasswordRule{
		minLength: passwordRuleDefaultMinLength,
	}

	for _, param := range parameters {
		parts := strings.SplitN(param, ":", 2)
		key := strings.ToLower(parts[0])

		switch key {
		case passwordRuleParamMin:
			if len(parts) == 2 {
				if val, err := strconv.Atoi(parts[1]); err == nil {
					r.minLength = val
				} else {
					return nil, fmt.Errorf(passwordRuleInvalidMin, err)
				}
			}
		case passwordRuleParamLetters:
			r.requireLetters = true
		case passwordRuleParamMixed:
			r.requireMixed = true
		case passwordRuleParamUppercase:
			r.requireUpper = true
		case passwordRuleParamLowercase:
			r.requireLower = true
		case passwordRuleParamNumbers:
			r.requireNumbers = true
		case passwordRuleParamSymbols:
			r.requireSymbols = true
		}
	}

	r.BaseRule = common.NewBaseRule(passwordRuleName, passwordRuleDefaultMsg, parameters, options...)
	return r, nil
}

// Validate applies complexity checks to the password.
func (r *PasswordRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	val, ok := ctx.Value().(string)
	if !ok || val == "" {
		return errors.New(passwordRuleDefaultMsg)
	}

	if len(val) < r.minLength {
		return fmt.Errorf(passwordRuleErrMin, r.minLength)
	}

	var hasLetter, hasUpper, hasLower, hasNumber, hasSymbol bool

	for _, c := range val {
		switch {
		case unicode.IsLetter(c):
			hasLetter = true
			if unicode.IsUpper(c) {
				hasUpper = true
			}
			if unicode.IsLower(c) {
				hasLower = true
			}
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c), unicode.IsSymbol(c):
			hasSymbol = true
		}
	}

	if r.requireLetters && !hasLetter {
		return errors.New(passwordRuleErrLetter)
	}
	if r.requireMixed && (!hasUpper || !hasLower) {
		return errors.New(passwordRuleErrMixed)
	}
	if r.requireUpper && !hasUpper {
		return errors.New(passwordRuleErrUpper)
	}
	if r.requireLower && !hasLower {
		return errors.New(passwordRuleErrLower)
	}
	if r.requireNumbers && !hasNumber {
		return errors.New(passwordRuleErrNumber)
	}
	if r.requireSymbols && !hasSymbol {
		return errors.New(passwordRuleErrSymbol)
	}

	return nil
}

func (r *PasswordRule) Name() string {
	return passwordRuleName
}
