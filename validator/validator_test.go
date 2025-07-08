package validator

import (
	"errors"
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

func TestValidator_Validate_Success(t *testing.T) {
	v := New()
	data := map[string]any{"name": "John"}
	rules := map[string]string{"name": "required"}
	if err := v.Validate(data, rules); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidator_Validate_FailureAndType(t *testing.T) {
	v := New()
	data := map[string]any{"name": ""}
	rules := map[string]string{"name": "required"}
	err := v.Validate(data, rules)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	var ve *contract.ValidationErrors
	if !errors.As(err, &ve) {
		t.Fatalf("expected *contract.ValidationErrors, got %T", err)
	}
	if got := ve.FieldError("name"); got == "" {
		t.Fatalf("expected field error for name, got empty")
	}
}

func TestValidator_ValidateMap_BailStopsOnFirstFailure(t *testing.T) {
	v := New()
	data := map[string]any{"name": ""}
	rules := map[string][]string{
		"name": {"bail", "required", "min:3"},
	}
	res := v.ValidateMap(data, rules)
	if res.IsValid() {
		t.Fatalf("expected invalid result")
	}
	errs := res.Errors()["name"]
	if len(errs) != 1 {
		t.Fatalf("expected exactly one error due to bail, got %v", errs)
	}
}

func TestValidator_CustomMessageAndAttribute(t *testing.T) {
	v := New()
	v.SetCustomMessage("required", "Please fill in the :attribute")
	v.SetCustomAttribute("email", "Email Address")

	data := map[string]any{"email": ""}
	rules := map[string]string{"email": "required"}
	res := v.ValidateWithResult(data, rules)
	if res.IsValid() {
		t.Fatalf("expected invalid result")
	}
	got := res.FieldError("email")
	if got != "Please fill in the Email Address" {
		t.Fatalf("unexpected message: %q", got)
	}
}

func TestValidator_HasRuleAndAvailable(t *testing.T) {
	v := New()
	if !v.HasRule("required") {
		t.Fatalf("expected required rule to exist")
	}
	if v.HasRule("definitely_not_a_rule") {
		t.Fatalf("did not expect nonexistent rule to exist")
	}
	list := v.GetAvailableRules()
	found := false
	for _, r := range list {
		if r == "required" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected required in available rules, got %v", list)
	}
}

func TestValidator_RuleCreationError_AcceptedIf_MissingParams(t *testing.T) {
	v := New()
	data := map[string]any{"status": "active", "tos": "no"}
	// accepted_if requires at least 2 params; give only 1 to trigger creation error
	rules := map[string]string{"tos": "accepted_if:status"}
	res := v.ValidateWithResult(data, rules)
	if res.IsValid() {
		t.Fatalf("expected invalid result")
	}
	msg := res.FieldError("tos")
	if msg == "" {
		t.Fatalf("expected error message for tos")
	}
	if want := "Rule creation error:"; len(msg) < len(want) || msg[:len(want)] != want {
		t.Fatalf("expected rule creation error, got %q", msg)
	}
}

func TestValidator_UnknownRule(t *testing.T) {
	v := New()
	data := map[string]any{"field": 1}
	rules := map[string]string{"field": "nonexistent"}
	res := v.ValidateWithResult(data, rules)
	if res.IsValid() {
		t.Fatalf("expected invalid result")
	}
	if got := res.FieldError("field"); got != "Unknown rule: nonexistent" {
		t.Fatalf("unexpected message: %q", got)
	}
}

// Integrated from integration_rules_test.go to keep all validator facade tests in one file.
func TestValidator_MultiRuleIntegration(t *testing.T) {
	v := New()
	data := map[string]any{
		"alphaOnly":             "abc123",       // alpha -> fail
		"alpha_num":             "abc123",       // alpha_num -> pass
		"email":                 "not-an-email", // email -> fail
		"age":                   "abc",          // numeric -> fail
		"count":                 2,              // min:3 -> fail
		"title":                 "hello",        // max:3 -> fail
		"pin":                   "1234",         // size:4 -> pass
		"gt":                    5,              // gt:10 -> fail
		"lt":                    5,              // lt:3 -> fail
		"lower":                 "ABC",          // lowercase -> fail
		"upper":                 "abc",          // uppercase -> fail
		"password":              "secret",       // confirmed -> mismatch
		"password_confirmation": "secret123",
		"sameA":                 "x",
		"sameB":                 "y", // same:sameA -> fail
		"diffA":                 "x",
		"diffB":                 "x", // different:diffA -> fail
		"numBetweenPass":        5,
		"numBetweenFail":        100,
	}
	rules := map[string]string{
		"alphaOnly":      "alpha",
		"alpha_num":      "alpha_num",
		"email":          "email",
		"age":            "numeric",
		"count":          "min:3",
		"title":          "max:3",
		"pin":            "size:4",
		"gt":             "gt:10",
		"lt":             "lt:3",
		"lower":          "lowercase",
		"upper":          "uppercase",
		"password":       "confirmed",
		"sameB":          "same:sameA",
		"diffB":          "different:diffA",
		"numBetweenPass": "between:1,10",
		"numBetweenFail": "between:1,10",
	}
	res := v.ValidateWithResult(data, rules)
	if res.IsValid() {
		t.Fatalf("expected invalid result")
	}
	// Spot check a few messages exist for expected failures
	for _, field := range []string{"alphaOnly", "email", "age", "count", "title", "gt", "lt", "lower", "upper", "password", "sameB", "diffB"} {
		if msg := res.FieldError(field); msg == "" {
			t.Fatalf("expected error for %s", field)
		}
	}
	// And passes should have no errors
	for _, field := range []string{"alpha_num", "pin", "numBetweenPass"} {
		if res.HasFieldError(field) {
			t.Fatalf("did not expect error for %s", field)
		}
	}
}
