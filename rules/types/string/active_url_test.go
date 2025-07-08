package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringRule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestActiveURLRule(t *testing.T) {
	rule, err := stringRule.NewActiveURLRule()
	if err != nil {
		t.Fatalf("failed to create ActiveURLRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// ✅ Valid URLs (resolvable domains only — assumes test env has internet)
		{"valid http", "http://example.com", true},
		{"valid https", "https://example.com", true},
		{"url with path", "https://example.com/path/to/page", true},
		{"url with query", "https://example.com?foo=bar", true},

		// ❌ Invalid formats
		{"missing scheme", "example.com", false},
		{"missing host", "http://", false},
		{"malformed scheme", "http:///abc", false},
		{"empty string", "", false},

		// ❌ Invalid types
		{"non-string int", 123, false},
		{"non-string bool", true, false},
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
