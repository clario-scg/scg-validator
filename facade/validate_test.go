package facade

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

// All facade validator tests are consolidated in this file to follow a single-pattern convention.

func TestMakeAndWithMethods(t *testing.T) {
	v := New()
	data := contract.NewSimpleDataProvider(map[string]any{"name": "Bob"})
	vr := v.Make(data, map[string][]string{"name": {"required"}}, map[string]string{"required.name": "Name is required"})
	if vr == nil {
		t.Fatal("expected validator request")
	}

	vr2 := vr.WithRules(map[string][]string{"age": {"numeric"}})
	if _, ok := vr.rules["age"]; ok {
		t.Fatal("original rules should not be modified")
	}
	if _, ok := vr2.rules["age"]; !ok {
		t.Fatal("expected rules to be merged in new request")
	}

	vr3 := vr.WithMessages(map[string]string{"required.name": "Custom"})
	if vr.messages["required.name"] != "Name is required" {
		t.Fatal("original messages should stay intact")
	}
	if vr3.messages["required.name"] != "Custom" {
		t.Fatal("expected message override in new request")
	}
}

func TestValidateMap_EndToEnd(t *testing.T) {
	data := map[string]any{"email": "test@example.com"}
	rules := map[string][]string{"email": {"required", "email"}}
	errs := ValidateMap(data, rules)
	if !errs.IsValid() {
		t.Fatalf("expected valid, got: %v", errs.Errors())
	}

	data2 := map[string]any{"email": "bad"}
	errs2 := ValidateMap(data2, rules)
	if errs2.IsValid() || !errs2.HasFieldError("email") {
		t.Fatalf("expected error for invalid email, got: %v", errs2.Errors())
	}
}

func TestFacadeQuickHelpers(t *testing.T) {
	dp := contract.NewSimpleDataProvider(map[string]any{"name": "A", "age": 20})
	if !Required(dp, "name") {
		t.Fatal("required helper failed")
	}
	if !Numeric(dp, "age") {
		t.Fatal("numeric helper failed")
	}
	if !Min(dp, "age", "10") || !Max(dp, "age", "25") {
		t.Fatal("min/max helpers failed")
	}
}

func TestFacadeRulesAndHasRule(t *testing.T) {
	v := New()
	list := v.Rules()
	if len(list) == 0 {
		t.Fatal("expected some registered rules")
	}
	if !v.HasRule("required") {
		t.Fatal("expected required rule to exist")
	}
	if v.HasRule("__not_a_rule__") {
		t.Fatal("unexpected non-existent rule reported as existing")
	}
}
