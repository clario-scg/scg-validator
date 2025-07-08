package rules

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

type simpleRule struct{ contract.Rule }

func (s simpleRule) Name() string                          { return "custom_simple" }
func (s simpleRule) Validate(_ contract.RuleContext) error { return nil }

func TestNewRuleRegistry_Options(t *testing.T) {
	// include only a couple of rules
	reg := NewRuleRegistry(WithIncludeOnly(RuleRequired, RuleEmail))
	if _, ok := reg.Get(RuleRequired); !ok {
		t.Fatal("expected included rule")
	}
	if _, ok := reg.Get(RuleEmail); !ok {
		t.Fatal("expected included rule")
	}
	if _, ok := reg.Get(RuleMax); ok {
		t.Fatal("unexpected rule present when include-only set")
	}

	// exclude a rule
	reg2 := NewRuleRegistry(WithExcludeRules(RuleRequired))
	if _, ok := reg2.Get(RuleRequired); ok {
		t.Fatal("expected excluded rule to be absent")
	}

	// add custom rule
	creator := func(_ []string) (contract.Rule, error) { return simpleRule{}, nil }
	reg3 := NewRuleRegistry(WithCustomRule("custom_simple", creator))
	if _, ok := reg3.Get("custom_simple"); !ok {
		t.Fatal("expected custom rule present")
	}
}
