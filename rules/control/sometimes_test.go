package control_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/control"
)

func TestSometimesRule(t *testing.T) {
	rule, err := control.NewSometimesRule()
	if err != nil {
		t.Fatalf("failed to create SometimesRule: %v", err)
	}

	tests := []struct {
		name  string
		value any
	}{
		{"with string", "test"},
		{"with nil", nil},
		{"with number", 123},
		{"with boolean", true},
		{"with slice", []int{1, 2, 3}},
		{"with empty string", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, map[string]any{})
			if err := rule.Validate(ctx); err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}
