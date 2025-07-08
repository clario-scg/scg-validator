package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringRule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestAlphaDashRule(t *testing.T) {
	rule, err := stringRule.NewAlphaDashRule(nil)
	if err != nil {
		t.Fatalf("failed to create AlphaDashRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"letters only", "Hello", true},
		{"numbers only", "123456", true},
		{"letters and numbers", "Hello123", true},
		{"with dash", "Hello-123", true},
		{"with underscore", "Hello_123", true},
		{"mixed dash and underscore", "Test-Data_01", true},
		{"with space", "Hello 123", false},
		{"with dot", "Hello.123", false},
		{"with special char", "Hello@123", false},
		{"empty string", "", false},
		{"non-string input", 123, false},
		{"nil value", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("username", tt.value, nil, nil)
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

func TestNewAlphaDashRule_ShouldNotError(t *testing.T) {
	rule, err := stringRule.NewAlphaDashRule(nil)
	if err != nil {
		t.Fatalf("unexpected error creating rule: %v", err)
	}
	if rule == nil {
		t.Fatal("expected rule instance, got nil")
	}
}
