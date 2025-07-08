package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/rules/comparison"

	"github.com/next-trace/scg-validator/contract"
)

func TestGteRule(t *testing.T) {
	rule, err := comparison.NewGteRule([]string{"10"})
	if err != nil {
		t.Fatalf("unexpected error creating gte rule: %v", err)
	}

	tests := []struct {
		name  string
		value any
		want  bool
	}{
		// Valid values
		{"int greater than threshold", 11, true},
		{"int equal to threshold", 10, true},
		{"float greater than threshold", 10.5, true},
		{"float equal to threshold", 10.0, true},
		{"string numeric > threshold", "11", true},
		{"string numeric = threshold", "10", true},

		// Invalid values
		{"int less than threshold", 9, false},
		{"float less than threshold", 9.99, false},
		{"string numeric < threshold", "9", false},
		{"non-numeric string (length >= threshold)", "not a number", true}, // 13 characters >= 10
		{"nil value", nil, false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.want && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.want && err == nil {
				t.Errorf("expected fail, but passed: value=%v", tt.value)
			}
		})
	}
}
