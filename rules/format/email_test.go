package format_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/format"
	"github.com/next-trace/scg-validator/utils"
)

func TestEmailRule(t *testing.T) {
	rule, err := format.NewEmailRule([]string{})
	if err != nil {
		t.Fatalf("Failed to create EmailRule: %v", err)
	}

	tests := []struct {
		name       string
		value      interface{}
		shouldPass bool
	}{
		{"valid email", "test@example.com", true},
		{"valid email with subdomain", "user@mail.example.com", true},
		{"valid email with plus", "test+tag@example.com", true},
		{"invalid no domain", "test@", false},
		{"invalid no @", "testexample.com", false},
		{"invalid no TLD", "test@example", false},
		{"invalid leading dot", "test@.example.com", false},
		{"invalid trailing dot", "test@example.com.", false},
		{"invalid multiple @", "test@@example.com", false},
		{"empty string", "", false},
		{"non-string type", 123, false},
		{"nil value", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, []string{}, map[string]any{})
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("Expected validation to pass for %v, but got error: %v", tt.value, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("Expected validation to fail for %v, but it passed", tt.value)
			}
		})
	}
}

func TestContainsDot(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"example.com", true},
		{"example", false},
		{"sub.example.com", true},
		{"", false},
	}

	for _, tt := range tests {
		result := utils.ContainsDot(tt.input)
		if result != tt.expected {
			t.Errorf("containsDot(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}
