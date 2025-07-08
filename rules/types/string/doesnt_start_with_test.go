package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringrule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestDoesntStartWithRule(t *testing.T) {
	rule, err := stringrule.NewDoesntStartWithRule([]string{"foo", "bar"})
	if err != nil {
		t.Fatalf("failed to create rule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"no prefix match", "baz", true},
		{"starts with foo", "foobaz", false},
		{"starts with bar", "barbaz", false},
		{"starts with partial match", "foolish", false},
		{"non-string input", 123, false},
		{"empty string", "", true},
		{"nil input", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected fail, but passed: value=%v", tt.value)
			}
		})
	}
}
