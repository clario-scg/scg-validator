package conditional

import (
	"errors"
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

func ctxWith(value any) contract.RuleContext {
	return contract.NewValidationContext("field", value, nil, map[string]any{"field": value})
}

func TestRequired_NilAndEmptyCases(t *testing.T) {
	rule, err := NewRequiredRule()
	if err != nil {
		t.Fatalf("unexpected error creating rule: %v", err)
	}

	cases := []struct {
		name    string
		val     any
		wantErr bool
	}{
		{"nil value", nil, true},
		{"empty string", "", true},
		{"non-empty string", "x", false},
		{"empty slice", []int{}, true},
		{"non-empty slice", []int{1}, false},
		{"zero int", 0, true},
		{"non-zero int", 5, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := rule.Validate(ctxWith(tc.val))
			if tc.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}

func TestRequired_PointerCases(t *testing.T) {
	rule, _ := NewRequiredRule()
	var pNil *int
	if err := rule.Validate(ctxWith(pNil)); err == nil {
		t.Fatalf("expected error for nil pointer")
	}
	pZero := new(int) // 0
	if err := rule.Validate(ctxWith(pZero)); err == nil {
		t.Fatalf("expected error for pointer to zero value")
	}
	v := 7
	pNonZero := &v
	if err := rule.Validate(ctxWith(pNonZero)); err != nil {
		t.Fatalf("expected no error for non-zero pointer, got %v", err)
	}
}

func TestRequired_ErrorMessage(t *testing.T) {
	rule, _ := NewRequiredRule()
	err := rule.Validate(ctxWith(nil))
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, errors.New("the :attribute field is required")) {
		// Compare strings instead of errors.Is on different instances
		if err.Error() != "the :attribute field is required" {
			t.Fatalf("unexpected error message: %q", err.Error())
		}
	}
}
