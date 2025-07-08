package acceptance

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

func TestDeclinedIfRule(t *testing.T) {
	rule, err := NewDeclinedIfRule([]string{"status", "inactive"})
	if err != nil {
		t.Fatalf("Failed to create DeclinedIfRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		// Condition met, value is "no" (declined)
		{"declined when condition met", "no", map[string]any{"status": "inactive"}, true},

		// Condition not met, value is "yes", not checked
		{"not checked when condition not met", "yes", map[string]any{"status": "active"}, true},

		// Condition met, but value is not declined ("yes" instead of "no")
		{"fail - condition met but not declined", "yes", map[string]any{"status": "inactive"}, false},

		// Pass with nil data (no condition to check)
		{"pass - nil data", "yes", nil, true},

		// Pass when condition field is missing (no condition to check)
		{"pass - condition field missing", "yes", map[string]any{}, true},

		// Acceptable value (0, int) passes, as it's considered declined
		{"declined int value passes", 0, map[string]any{"status": "inactive"}, true},

		// Unacceptable value (1, int) fails when condition is met
		{"accepted int value fails when condition met", 1, map[string]any{"status": "inactive"}, false},

		// Declined value (false, bool) passes when condition is met
		{"declined bool false passes", false, map[string]any{"status": "inactive"}, true},

		// Edge case: value is nil and condition is met, so it fails
		{"fail - condition met but value is nil", nil, map[string]any{"status": "inactive"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("declined", tt.value, nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("Test failed: %s. Expected pass, but got error: %v", tt.name, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("Test failed: %s. Expected fail, but passed: value=%v", tt.name, tt.value)
			}
		})
	}
}

func TestDeclinedIfRule_InvalidParameters(t *testing.T) {
	tests := []struct {
		name   string
		params []string
	}{
		{"empty param list", []string{}},                                        // Missing condition
		{"only condition field", []string{"status"}},                            // Only one parameter (missing value)
		{"multiple condition fields", []string{"status", "inactive", "active"}}, // Too many parameters, should fail
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDeclinedIfRule(tt.params)
			if err == nil {
				t.Errorf("Expected error for invalid params in case: %s", tt.name)
			}
		})
	}
}
