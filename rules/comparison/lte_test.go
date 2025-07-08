package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/rules/comparison"

	"github.com/next-trace/scg-validator/contract"
)

func TestLteRule(t *testing.T) {
	rule, err := comparison.NewLteRule([]string{"10"})
	if err != nil {
		t.Fatalf("unexpected error creating LteRule: %v", err)
	}

	cases := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// Integer cases
		{"int less than threshold", 9, true},
		{"int equal to threshold", 10, true},
		{"int greater than threshold", 11, false},

		// Float cases
		{"float less than threshold", 9.9, true},
		{"float equal to threshold", 10.0, true},
		{"float greater than threshold", 10.1, false},
		{"float padded zero", 10.0000, true},

		// String numeric
		{"string numeric less", "9", true},
		{"string numeric equal", "10", true},
		{"string numeric greater", "11", false},

		// Edge / invalid types
		{"invalid string input", "not a number", false},
		{"nil input", nil, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Helper()

			ctx := contract.NewValidationContext("field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected fail, but passed: value=%v", tc.value)
			}
		})
	}
}
