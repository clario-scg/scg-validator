package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

func TestLtRule(t *testing.T) {
	rule, err := comparison.NewLtRule([]string{"5"})
	if err != nil {
		t.Fatalf("failed to create LtRule: %v", err)
	}

	testCases := []struct {
		name       string
		value      interface{}
		shouldPass bool
	}{
		// Numeric comparisons
		{"numeric < 5", 4, true},
		{"numeric = 5", 5, false},
		{"numeric > 5", 6, false},
		{"float < 5", 4.5, true},
		{"float = 5", 5.0, false},

		// String length comparisons (threshold is 5)
		{"string length < 5", "hi", true},           // 2 characters < 5
		{"string length = 5", "hello", false},       // 5 characters = 5
		{"string length > 5", "hello world", false}, // 11 characters > 5

		// Slice length comparisons (threshold is 5)
		{"slice length < 5", []int{1, 2, 3}, true},           // 3 elements < 5
		{"slice length = 5", []int{1, 2, 3, 4, 5}, false},    // 5 elements = 5
		{"slice length > 5", []int{1, 2, 3, 4, 5, 6}, false}, // 6 elements > 5

		// Map length comparisons (threshold is 5)
		{"map length < 5", map[string]int{"a": 1, "b": 2, "c": 3}, true},                  // 3 elements < 5
		{"map length = 5", map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, false}, // 5 elements = 5

		// Unsupported type
		{"unsupported bool", true, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tc.value, nil, map[string]interface{}{})
			err := rule.Validate(ctx)

			// Test assertion for passing and failing cases
			if tc.shouldPass && err != nil {
				t.Errorf("expected pass but got error: %v", err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure but got no error for value: %v", tc.value)
			}
		})
	}
}
