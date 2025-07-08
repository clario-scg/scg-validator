package inclusion_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/inclusion"
)

func TestInRule(t *testing.T) {
	t.Parallel()

	rule, err := inclusion.NewInRule([]string{"foo", "bar", "baz"})
	if err != nil {
		t.Fatalf("failed to create InRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"match - foo", "foo", true},
		{"match - bar", "bar", true},
		{"match - baz", "baz", true},
		{"no match - qux", "qux", false},
		{"no match - integer", 123, false},
		{"no match - bool", true, false},
		{"no match - nil", nil, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := contract.NewValidationContext("field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass for value %v, but got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure for value %v, but got none", tc.value)
			}
		})
	}
}

func TestInRule_InvalidConstruction(t *testing.T) {
	t.Parallel()

	_, err := inclusion.NewInRule([]string{})
	if err == nil {
		t.Error("expected error when creating InRule with no allowed values")
	}
}
