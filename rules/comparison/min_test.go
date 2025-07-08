package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

func TestMinRule(t *testing.T) {
	rule, err := comparison.NewMinRule([]string{"5"})
	if err != nil {
		t.Fatalf("Failed to create MinRule: %v", err)
	}

	tests := []struct {
		name       string
		value      interface{}
		shouldPass bool
	}{
		// String length comparisons
		{"valid greater than min (string)", "hello world", true},
		{"valid equal to min (string)", "hello", true},
		{"invalid less than min (string)", "hi", false},

		// Numeric comparisons
		{"valid greater than min (numeric)", 10, true},
		{"valid equal to min (numeric)", 5, true},
		{"invalid less than min (numeric)", 3, false},

		// Slice length comparisons
		{"valid slice greater than min", []int{1, 2, 3, 4, 5, 6}, true},
		{"invalid slice less than min", []int{1, 2}, false},

		// Unsupported type comparisons
		{"invalid type (boolean)", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, []string{}, map[string]any{})
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("Expected validation to pass for value %v, but got error: %v", tt.value, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("Expected validation to fail for value %v, but it passed", tt.value)
			}
		})
	}
}

func TestMinRule_NoParameters(t *testing.T) {
	_, err := comparison.NewMinRule([]string{})
	if err == nil {
		t.Error("Expected error when creating MinRule without parameters")
	}
}

func TestMinRule_InvalidParameter(t *testing.T) {
	_, err := comparison.NewMinRule([]string{"abc"})
	if err == nil {
		t.Error("Expected error when creating MinRule with non-numeric parameter")
	}
}
