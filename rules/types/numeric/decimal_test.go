package numeric_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/numeric"
)

func runDecimalTests(t *testing.T, rule contract.Rule, cases []struct {
	name  string
	value any
	want  bool
}) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err := rule.Validate(ctx)
			passed := err == nil

			if passed != tt.want {
				t.Errorf("test %q: value=%v, expected=%v, got error=%v", tt.name, tt.value, tt.want, err)
			}
		})
	}
}

func TestDecimalRule_ExactMatch(t *testing.T) {
	rule, err := numeric.NewDecimalRule([]string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	runDecimalTests(t, rule, []struct {
		name  string
		value any
		want  bool
	}{
		{"valid float with 2 decimals", 12.34, true},
		{"valid string with 2 decimals", "12.34", true},
		{"invalid float with 1 decimal", 12.3, false},
		{"invalid float with 3 decimals", 12.345, false},
		{"valid string with .00", "12.00", true},
		{"integer value", 12, false},
		{"string integer", "12", false},
		{"invalid string", "not a number", false},
	})
}

func TestDecimalRule_Range(t *testing.T) {
	rule, err := numeric.NewDecimalRule([]string{"2", "4"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	runDecimalTests(t, rule, []struct {
		name  string
		value any
		want  bool
	}{
		{"2 decimals", 12.34, true},
		{"3 decimals", 12.345, true},
		{"4 decimals", 12.3456, true},
		{"string 3 decimals", "12.345", true},
		{"1 decimal", 12.3, false},
		{"5 decimals", 12.34567, false},
		{"integer", 12, false},
	})
}

func TestDecimalRule_ZeroDecimals(t *testing.T) {
	rule, err := numeric.NewDecimalRule([]string{"0"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	runDecimalTests(t, rule, []struct {
		name  string
		value any
		want  bool
	}{
		{"integer", 12, true},
		{"float zero decimals", 12.0, true},
		{"string zero decimals", "12.0", false}, // "12.0" has 1 decimal place, not 0
		{"string integer", "12", true},
		{"float with decimals", 12.34, false},
		{"string with decimals", "12.34", false},
	})
}
