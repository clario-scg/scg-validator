package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/conditional"
)

func TestRequiredWithoutAllRule(t *testing.T) {
	rule, err := conditional.NewRequiredWithoutAllRule([]string{"foo", "bar"})
	if err != nil {
		t.Fatalf("Failed to create RequiredWithoutAllRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "passes when none of the other fields are present",
			value:      "test",
			data:       map[string]any{},
			shouldPass: true,
		},
		{
			name:       "passes when one of the other fields is present",
			value:      "test",
			data:       map[string]any{"foo": "value"},
			shouldPass: true,
		},
		{
			name:       "passes when all other fields are present and value is empty",
			value:      "",
			data:       map[string]any{"foo": "value", "bar": "value"},
			shouldPass: true,
		},
		{
			name:       "passes when all other fields are present and value is nil",
			value:      nil,
			data:       map[string]any{"foo": "value", "bar": "value"},
			shouldPass: true,
		},
		{
			name:       "passes when value is non-empty, and all other fields are present",
			value:      "test",
			data:       map[string]any{"foo": "value", "bar": "value"},
			shouldPass: true,
		},
		{
			name:       "fails when none of the other fields are present and value is empty",
			value:      "",
			data:       nil,
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, tt.data)
			err := rule.Validate(ctx)
			if tt.shouldPass && err != nil {
				t.Errorf("expected pass but got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected failure, got pass for value: %#v", tt.value)
			}
		})
	}
}

func TestRequiredWithoutAllRule_InvalidParameters(t *testing.T) {
	_, err := conditional.NewRequiredWithoutAllRule([]string{})
	if err == nil {
		t.Error("expected error when creating RequiredWithoutAllRule without parameters")
	}
}
