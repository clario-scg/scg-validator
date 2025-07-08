package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringrule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestEndsWithRule(t *testing.T) {
	rule, err := stringrule.NewEndsWithRule([]string{"foo", "bar"})
	if err != nil {
		t.Fatalf("failed to create ends_with rule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"ends with foo", "bazfoo", true},
		{"ends with bar", "bazbar", true},
		{"no matching suffix", "baz", false},
		{"non-string input", 123, false},
		{"empty string", "", false},
		{"nil value", nil, false},
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
