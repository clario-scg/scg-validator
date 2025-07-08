package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

func TestGtRule(t *testing.T) {
	rule, err := comparison.NewGtRule([]string{"5"})
	if err != nil {
		t.Fatalf("failed to create GtRule: %v", err)
	}

	testCases := []struct {
		name       string
		value      interface{}
		shouldPass bool
	}{
		// Numeric comparisons
		{"numeric > 5", 6, true},
		{"numeric = 5", 5, false},
		{"numeric < 5", 4, false},
		{"float > 5", 5.5, true},
		{"float = 5", 5.0, false},

		// String length comparisons
		{"string length > 5", "hello world", true},
		{"string length = 5", "hello", false},
		{"string length < 5", "hi", false},

		// Slice length comparisons
		{"slice length > 5", []int{1, 2, 3, 4, 5, 6}, true},
		{"slice length = 5", []int{1, 2, 3, 4, 5}, false},
		{"slice length < 5", []int{1, 2, 3}, false},

		// Map length comparisons
		{"map length > 5", map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}, true},
		{"map length = 5", map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, false},

		// Unsupported type (non-numeric)
		{"unsupported bool", true, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tc.value, nil, map[string]interface{}{})
			err := rule.Validate(ctx)

			// Validate if the result matches the expected outcome
			if tc.shouldPass && err != nil {
				t.Errorf("expected pass but got error: %v", err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure but got no error for value: %v", tc.value)
			}
		})
	}
}
