package engine

import (
	"errors"
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

type alwaysFailRule struct{}

func (r *alwaysFailRule) Name() string { return "custom_fail" }
func (r *alwaysFailRule) Validate(ctx contract.RuleContext) error {
	return errors.New("custom error message")
}

func TestEngine_BailVsNoBail(t *testing.T) {
	e := NewEngine()
	data := NewDataProvider(map[string]any{"name": ""})

	// With bail -> only first failure should be recorded
	res1 := e.Execute(data, map[string]string{"name": "bail|required|min:3"})
	if res1.IsValid() {
		t.Fatalf("expected invalid result")
	}
	if got := len(res1.Errors()["name"]); got != 1 {
		t.Fatalf("expected 1 error with bail, got %d", got)
	}

	// Without bail -> multiple failures should be recorded
	res2 := e.Execute(data, map[string]string{"name": "required|min:3"})
	if res2.IsValid() {
		t.Fatalf("expected invalid result")
	}
	if got := len(res2.Errors()["name"]); got < 2 {
		t.Fatalf("expected at least 2 errors without bail, got %d", got)
	}
}

func TestEngine_ResolveErrorMessage_FallbackToOriginal(t *testing.T) {
	e := NewEngine()
	// Disable message resolver to force fallback to original error
	e.MessageResolver = nil
	_ = e.Registry.Register("custom_fail", func(_ []string) (contract.Rule, error) { return &alwaysFailRule{}, nil })

	data := NewDataProvider(map[string]any{"f": 1})
	res := e.Execute(data, map[string]string{"f": "custom_fail"})
	if res.IsValid() {
		t.Fatalf("expected invalid result")
	}
	if got := res.FieldError("f"); got != "custom error message" {
		t.Fatalf("unexpected message: %q", got)
	}
}
