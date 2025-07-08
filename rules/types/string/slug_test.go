package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringrule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestSlugRule(t *testing.T) {
	rule, err := stringrule.NewSlugRule()
	if err != nil {
		t.Fatalf("failed to create slug rule: %v", err)
	}

	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{"valid: simple slug", "hello-world", true},
		{"valid: with numbers", "hello-world-123", true},
		{"valid: no dashes", "helloworld", true},
		{"invalid: contains uppercase", "Hello-world", false},
		{"invalid: contains space", "hello world", false},
		{"invalid: contains symbol", "hello-world!", false},
		{"invalid: leading dash", "-hello-world", false},
		{"invalid: trailing dash", "hello-world-", false},
		{"invalid: double dash", "hello--world", false},
		{"invalid: non-string", 123, false},
		{"invalid: nil value", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("slug", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.want && err != nil {
				t.Errorf("expected pass but got error: %v", err)
			}
			if !tt.want && err == nil {
				t.Errorf("expected fail but got pass for value: %v", tt.value)
			}
		})
	}
}
