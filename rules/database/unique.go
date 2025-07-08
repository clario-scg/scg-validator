package database

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/registry/database"
)

const (
	uniqueRuleName              = "unique"
	uniqueRuleDefaultMsg        = "unique rule requires table and field parameters: unique:table,field"
	uniqueRuleNotImplementedMsg = "the presence verifier for table '%s' is not implemented; " +
		"please provide a '%s'PresenceVerifier"
	uniqueRuleFailedMsg = "the %s must be unique"
)

// uniqueRule checks if a value is unique in the specified table/field.
type uniqueRule struct{}

// NewUniqueRule constructs a new instance of uniqueRule.
// Usage: unique:users,email
func NewUniqueRule() (contract.Rule, error) {
	return &uniqueRule{}, nil
}

func (r *uniqueRule) Name() string {
	return uniqueRuleName
}

func (r *uniqueRule) Validate(ctx contract.RuleContext) error {
	params := ctx.Parameters()
	if len(params) < 2 {
		return errors.New(uniqueRuleDefaultMsg)
	}

	table := params[0]
	field := params[1]

	verifier, ok := database.FindPresenceVerifier(table)
	if !ok {
		return fmt.Errorf(uniqueRuleNotImplementedMsg, table, table)
	}

	isUnique, err := verifier.Unique(table, field, ctx.Value())
	if err != nil {
		return err
	}

	if !isUnique {
		return fmt.Errorf(uniqueRuleFailedMsg, field)
	}

	return nil
}
