package acceptance

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

func TestDeclinedRule(t *testing.T) {
	rule, err := NewDeclinedRule()
	if err != nil {
		t.Fatalf("Failed to create DeclinedRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// Boolean tests
		{"declined boolean false", false, true},
		{"not declined boolean true", true, false},

		// String tests for declined values
		{"declined string 'no'", "no", true},
		{"declined string 'off'", "off", true},
		{"declined string '0'", "0", true},
		{"declined string 'false'", "false", true},
		{"declined string 'NO' (uppercase)", "NO", true},

		// Not declined strings
		{"not declined string 'yes'", "yes", false},
		{"not declined string '1'", "1", false},
		{"not declined string 'true'", "true", false},
		{"not declined string 'maybe'", "maybe", false},

		// Integer tests for declined values
		{"declined int 0", 0, true},
		{"not declined int 1", 1, false},
		{"not declined int 2", 2, false},

		// Unsigned int tests for declined values
		{"declined uint 0", uint(0), true},
		{"not declined uint 1", uint(1), false},

		// Edge case: nil value should not be accepted
		{"nil value", nil, false},
		{"empty string", "", false}, // Test case for empty string
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("declined", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("Test %s failed: expected pass, got error: %v", tt.name, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("Test %s failed: expected fail, but passed for value: %v", tt.name, tt.value)
			}
		})
	}
}
