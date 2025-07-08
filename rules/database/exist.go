package database

import (
	"errors"
	"fmt"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/registry/database"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	existRuleName              = "exist"
	existRuleDefaultMsg        = "exist rule requires table and field parameters: exist:table,field"
	existRuleNotImplementedMsg = "the presence verifier for table '%s' is not implemented; " +
		"please provide a '%s'PresenceVerifier"
	existRuleMissingTableMsg = "exist rule requires a table name parameter"
	existRuleFailedMsg       = "%v does not exist in %s.%s"
)

type existRule struct {
	common.BaseRule
	table string
}

// NewExistRule initializes an existRule instance.
// Usage: exist:table,field
func NewExistRule(params []string) (contract.Rule, error) {
	if len(params) < 1 {
		return nil, errors.New(existRuleMissingTableMsg)
	}

	r := &existRule{
		table: params[0],
	}
	r.BaseRule = common.NewBaseRule(existRuleName, existRuleDefaultMsg, params)
	return r, nil
}

func (r *existRule) Name() string {
	return existRuleName
}

func (r *existRule) Validate(ctx contract.RuleContext) error {
	params := ctx.Parameters()
	if len(params) < 2 {
		return errors.New(existRuleDefaultMsg)
	}

	table := params[0]
	field := params[1]

	verifier, ok := database.FindPresenceVerifier(table)
	if !ok {
		return fmt.Errorf(existRuleNotImplementedMsg, table, table)
	}

	found, err := verifier.Exists(table, field, ctx.Value())
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf(existRuleFailedMsg, ctx.Value(), table, field)
	}

	return nil
}
