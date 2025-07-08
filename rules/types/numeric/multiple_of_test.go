package numeric_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/numeric"
)

func TestMultipleOfRule_IntegerValues(t *testing.T) {
	rule, err := numeric.NewMultipleOfRule([]string{"4"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{"valid int", 8, true},
		{"invalid int", 10, false},
		{"valid string", "12", true},
		{"invalid string", "13", false},
		{"non-numeric string", "abc", false},
		{"zero", 0, true},
		{"nil value", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.want && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.want && err == nil {
				t.Errorf("expected fail, got success for value: %v", tt.value)
			}
		})
	}
}

func TestMultipleOfRule_FloatValues(t *testing.T) {
	rule, err := numeric.NewMultipleOfRule([]string{"2.5"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{"valid float", 5.0, true},
		{"invalid float", 5.1, false},
		{"valid string float", "7.5", true},
		{"invalid string float", "7.6", false},
		{"slightly under", 9.999999999, true},
		{"slightly over", 10.0000001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("amount", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.want && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.want && err == nil {
				t.Errorf("expected fail, got success for value: %v", tt.value)
			}
		})
	}
}

func TestMultipleOfRule_InvalidConstruction(t *testing.T) {
	t.Run("empty parameter", func(t *testing.T) {
		_, err := numeric.NewMultipleOfRule([]string{})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("non-numeric parameter", func(t *testing.T) {
		_, err := numeric.NewMultipleOfRule([]string{"invalid"})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("zero as parameter", func(t *testing.T) {
		_, err := numeric.NewMultipleOfRule([]string{"0"})
		if err == nil {
			t.Error("expected error for zero value, got nil")
		}
	})
}
