package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

// TestSameRule tests the SameRule for various scenarios.
func TestSameRule(t *testing.T) {
	rule, err := comparison.NewSameRule([]string{"test"})
	if err != nil {
		t.Fatalf("Failed to create SameRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		params     []string
		data       map[string]any
		shouldPass bool
	}{
		{"same string", "test", []string{"test"}, map[string]any{"test": "test"}, true},
		{"different string", "test", []string{"test"}, map[string]any{"test": "different"}, false},
		{"same int", 42, []string{"test"}, map[string]any{"test": 42}, true},
		{"different int", 42, []string{"test"}, map[string]any{"test": 43}, false},
		{"same float", 42.0, []string{"test"}, map[string]any{"test": 42.0}, true},
		{"different float", 42.0, []string{"test"}, map[string]any{"test": 43.0}, false},
		{"same nil", nil, []string{"test"}, map[string]any{"test": nil}, true},
		{"different nil", nil, []string{"test"}, map[string]any{"test": 42}, false},
		{"empty string vs nil", "", []string{"test"}, map[string]any{"test": nil}, false}, // Added edge case
		{"nil vs empty string", nil, []string{"test"}, map[string]any{"test": ""}, false}, // Added edge case
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, tt.params, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass but got error: %v for value: %v", err, tt.value)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected failure but got no error for value: %v", tt.value)
			}
		})
	}
}
