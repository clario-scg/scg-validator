package inclusion_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/inclusion"
)

func TestNotInRule(t *testing.T) {
	t.Parallel()

	rule, err := inclusion.NewNotInRule([]string{"foo", "bar", "baz"})
	if err != nil {
		t.Fatalf("failed to create NotInRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"forbidden - foo", "foo", false},
		{"forbidden - bar", "bar", false},
		{"forbidden - baz", "baz", false},
		{"allowed - qux", "qux", true},
		{"allowed - int", 123, true},
		{"allowed - bool", true, true},
		{"allowed - nil", nil, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := contract.NewValidationContext("field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass for value %v, got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected fail for value %v, but got no error", tc.value)
			}
		})
	}
}

func TestNotInRule_InvalidConstruction(t *testing.T) {
	t.Parallel()

	_, err := inclusion.NewNotInRule([]string{})
	if err == nil {
		t.Error("expected error when creating NotInRule with no forbidden values")
	}
}
