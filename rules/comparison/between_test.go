package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

func TestBetweenRule(t *testing.T) {
	rule, err := comparison.NewBetweenRule([]string{"3", "10"})
	if err != nil {
		t.Fatalf("Failed to create BetweenRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"valid string between", "hello", true},     // len = 5
		{"valid string at min", "hi!", true},        // len = 3
		{"valid string at max", "hello12345", true}, // len = 10
		{"invalid string too short", "hi", false},   // len = 2
		{"invalid string too long", "hello world this is long", false},

		{"valid numeric between", 5, true},
		{"valid numeric at min", 3, true},
		{"valid numeric at max", 10, true},
		{"invalid numeric less", 2, false},
		{"invalid numeric greater", 15, false},

		{"valid slice between", []int{1, 2, 3, 4, 5}, true}, // len = 5
		{"invalid slice too small", []int{1, 2}, false},     // len = 2
		{"invalid slice too large", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, false},

		{"invalid type", true, false}, // non-numeric, non-len-compatible
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, []string{}, nil)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("[%s] Expected validation to pass but got error: %v", tt.name, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("[%s] Expected validation to fail but passed", tt.name)
			}
		})
	}
}

func TestBetweenRule_InvalidParameters(t *testing.T) {
	tests := []struct {
		name   string
		params []string
	}{
		{"no parameters", []string{}},
		{"one parameter", []string{"5"}},
		{"invalid min", []string{"abc", "10"}},
		{"invalid max", []string{"5", "abc"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := comparison.NewBetweenRule(tt.params)
			if err == nil {
				t.Errorf("[%s] Expected error, got nil", tt.name)
			}
		})
	}
}
