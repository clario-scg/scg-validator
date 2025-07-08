package control_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/control"
)

func TestNullableRule(t *testing.T) {
	rule, err := control.NewNullableRule()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name  string
		value any
	}{
		{"nil value", nil},
		{"empty string", ""},
		{"non-empty string", "hello"},
		{"number", 42},
		{"boolean", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)

			if err := rule.Validate(ctx); err != nil {
				t.Errorf("expected no error for value: %v, got: %v", tt.value, err)
			}
		})
	}
}
