package acceptance

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

func TestAcceptedRule(t *testing.T) {
	rule, err := NewAcceptedRule()
	if err != nil {
		t.Fatalf("Failed to create AcceptedRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// Boolean test cases
		{"bool true", true, true},
		{"bool false", false, false},

		// Accepted string values (case-insensitive)
		{"string yes", "yes", true},
		{"string on", "on", true},
		{"string 1", "1", true},
		{"string true", "true", true},
		{"string uppercase YES", "YES", true},

		// Non-accepted string values
		{"string no", "no", false},
		{"string 0", "0", false},
		{"string false", "false", false},
		{"string maybe", "maybe", false},

		// Numeric tests
		{"int 1", 1, true},
		{"int 0", 0, false},
		{"int 2", 2, false},

		{"uint 1", uint(1), true},
		{"uint 0", uint(0), false},

		// Nil value test case
		{"nil value", nil, false},

		// Additional Edge Cases
		{"empty string", "", false},                   // Empty string should fail
		{"string with extra spaces", "  yes  ", true}, // String with spaces should pass
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("accepted", tt.value, nil, nil)
			err := rule.Validate(ctx)

			// Check if the validation passed or failed
			if tt.shouldPass && err != nil {
				t.Errorf("Expected pass for %v, but got error: %v", tt.value, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("Expected fail for %v, but validation passed", tt.value)
			}
		})
	}
}
