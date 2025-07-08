package format

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	EmailRuleName                                     = "email"
	EmailRuleDefaultMessage                           = "the :attribute must be a valid email address"
	EmailRuleDataNotProvideOrIsEmptyMSgDefaultMessage = "the :attribute must provide a valid email address, " +
		"but it is empty or not provided"
	EmailRuleProvidedDataCanNotBeAnEmailDefaultMessage = "the :attribute must provide a valid email address, " +
		"but it is not a valid email format"
	EmailRuleProvidedDataHasInvalidDomainMessage = "the :attribute must provide a valid email address, " +
		"but the domain part is invalid"
)

// extractDomain extracts the domain part from an email address
func extractDomain(email string) string {
	at := strings.LastIndex(email, "@")
	if at == -1 || at+1 >= len(email) {
		return ""
	}
	return email[at+1:]
}

// isValidDomain checks if a domain is valid
func isValidDomain(domain string) bool {
	return domain != "" &&
		!strings.HasPrefix(domain, ".") &&
		!strings.HasSuffix(domain, ".") &&
		strings.Contains(domain, ".")
}

type EmailRule struct {
	common.BaseRule
}

func NewEmailRule(params []string, opts ...common.RuleOption) (contract.Rule, error) {
	return &EmailRule{
		BaseRule: common.NewBaseRule(EmailRuleName, EmailRuleDefaultMessage, params, opts...),
	}, nil
}

func (r *EmailRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}
	val, ok := ctx.Value().(string)
	if !ok || strings.TrimSpace(val) == "" {
		return errors.New(EmailRuleDataNotProvideOrIsEmptyMSgDefaultMessage)
	}

	addr, err := mail.ParseAddress(val)
	if err != nil || addr.Address != val {
		return errors.New(EmailRuleProvidedDataCanNotBeAnEmailDefaultMessage)
	}

	domain := extractDomain(val)
	if !isValidDomain(domain) {
		return errors.New(EmailRuleProvidedDataHasInvalidDomainMessage)
	}

	return nil
}

func (r *EmailRule) Name() string {
	return EmailRuleName
}
