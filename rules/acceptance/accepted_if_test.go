package acceptance

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

func TestAcceptedIfRule(t *testing.T) {
	rule, err := NewAcceptedIfRule([]string{"status", "active"})
	if err != nil {
		t.Fatalf("Failed to create AcceptedIfRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{"accepted when condition met", "yes", map[string]any{"status": "active"}, true},
		{"not checked when condition not met", "no", map[string]any{"status": "inactive"}, true},
		{"fail when condition met but not accepted", "no", map[string]any{"status": "active"}, false},
		{"pass with nil data", "no", nil, true},
		{"pass when condition field missing", "no", map[string]any{}, true},
		{"accepted with int 1", 1, map[string]any{"status": "active"}, true},
		{"fail with int 0", 0, map[string]any{"status": "active"}, false},
		// Additional edge cases for string comparison
		{"accepted with string 'true' (case insensitive)", "TRUE", map[string]any{"status": "active"}, true},
		{"accepted with string 'true ' (extra space)", "true ", map[string]any{"status": "active"}, true},
		{"fail with invalid string 'maybe'", "maybe", map[string]any{"status": "active"}, false},
		{"pass with numeric string '1' when condition not met", "1", map[string]any{"status": "inactive"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("Expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("Expected fail, but passed: value=%v", tt.value)
			}
		})
	}
}

func TestAcceptedIfRule_InvalidParameters(t *testing.T) {
	tests := []struct {
		name   string
		params []string
	}{
		{"no parameters", []string{}},
		{"only one parameter", []string{"status"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAcceptedIfRule(tt.params)
			if err == nil {
				t.Errorf("Expected error for case: %s", tt.name)
			}
		})
	}
}
