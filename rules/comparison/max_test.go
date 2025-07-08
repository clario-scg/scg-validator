package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

func TestMaxRule(t *testing.T) {
	rule, err := comparison.NewMaxRule([]string{"10"})
	if err != nil {
		t.Fatalf("Failed to create MaxRule: %v", err)
	}

	tests := []struct {
		name       string
		value      interface{}
		shouldPass bool
	}{
		{"valid less than max", "hello", true},
		{"valid equal to max", "hello12345", true},
		{"invalid greater than max", "hello world this is too long", false},
		{"valid numeric less", 5, true},
		{"valid numeric equal", 10, true},
		{"invalid numeric greater", 15, false},
		{"valid slice less", []int{1, 2, 3}, true},
		{"invalid slice greater", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, false},
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
