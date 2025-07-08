package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringrule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestLowercaseRule(t *testing.T) {
	rule, err := stringrule.NewLowercaseRule()
	if err != nil {
		t.Fatalf("Failed to create LowercaseRule: %v", err)
	}

	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{"all lowercase", "lowercase", true},
		{"starts with uppercase", "Lowercase", false},
		{"all uppercase", "LOWERCASE", false},
		{"non-string input", 123, false},
		{"empty string", "", true},
		{"nil input", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.want && err != nil {
				t.Errorf("Expected pass, got error: %v", err)
			}
			if !tt.want && err == nil {
				t.Errorf("Expected fail, but passed: value=%v", tt.value)
			}
		})
	}
}
