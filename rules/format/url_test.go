package format_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/format"
)

func TestURLRule(t *testing.T) {
	t.Parallel()

	rule, err := format.NewURLRule([]string{})
	if err != nil {
		t.Fatalf("failed to create URLRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// ✅ Valid cases
		{"valid - http URL", "http://example.com", true},
		{"valid - https URL", "https://example.com", true},
		{"valid - with path", "https://example.com/path", true},
		{"valid - with query", "https://example.com?query=value", true},

		// ❌ Invalid cases
		{"invalid - missing scheme", "example.com", false},
		{"invalid - missing host", "http://", false},
		{"invalid - malformed URL", "http:///path", false},
		{"invalid - empty string", "", false},
		{"invalid - non-string type", 123, false},
		{"invalid - nil value", nil, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := contract.NewValidationContext("url_field", tc.value, nil, nil)
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
