package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

func TestSizeRule(t *testing.T) {
	rule, err := comparison.NewSizeRule([]string{"5"})
	if err != nil {
		t.Fatalf("Failed to create SizeRule: %v", err)
	}

	tests := []struct {
		name       string
		value      interface{}
		shouldPass bool
	}{
		// String size comparisons
		{"valid string size", "hello", true},              // len("hello") == 5
		{"invalid string too short", "hi", false},         // len("hi") < 5
		{"invalid string too long", "hello world", false}, // len("hello world") > 5

		// Numeric comparisons
		{"valid numeric size", 5, true},        // 5 == 5
		{"invalid numeric less", 3, false},     // 3 < 5
		{"invalid numeric greater", 10, false}, // 10 > 5

		// Slice length comparisons
		{"valid slice size", []int{1, 2, 3, 4, 5}, true},    // len([1,2,3,4,5]) == 5
		{"invalid slice wrong size", []int{1, 2, 3}, false}, // len([1,2,3]) < 5

		// Unsupported types
		{"invalid type (boolean)", true, false}, // boolean, not a valid size type
		{"invalid type (nil)", nil, false},      // nil should fail
		{"empty string", "", false},             // empty string, len("") < 5
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
