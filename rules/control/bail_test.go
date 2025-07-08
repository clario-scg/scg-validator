package control_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/control"
)

func TestBailRule(t *testing.T) {
	rule, err := control.NewBailRule()
	if err != nil {
		t.Fatalf("failed to create BailRule: %v", err)
	}

	tests := []struct {
		name  string
		value any
	}{
		{"with string", "test"},
		{"with nil", nil},
		{"with number", 123},
		{"with bool", true},
		{"with slice", []int{1, 2, 3}},
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
