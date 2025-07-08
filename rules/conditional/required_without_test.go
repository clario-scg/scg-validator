package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/conditional"
)

func TestRequiredWithoutRule(t *testing.T) {
	rule, err := conditional.NewRequiredWithoutRule([]string{"other_field"})
	if err != nil {
		t.Fatalf("Failed to create RequiredWithoutRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "passes when other field is present and value is empty",
			value:      "",
			data:       map[string]any{"other_field": "value"},
			shouldPass: true,
		},
		{
			name:       "passes when other field is present and value is nil",
			value:      nil,
			data:       map[string]any{"other_field": "value"},
			shouldPass: true,
		},
		{
			name:       "required when other field is not present and value is non-empty",
			value:      "test",
			data:       map[string]any{},
			shouldPass: true,
		},
		{
			name:       "fails when other field is not present and value is empty",
			value:      "",
			data:       map[string]any{},
			shouldPass: false,
		},
		{
			name:       "fails when other field is not present and value is nil",
			value:      nil,
			data:       map[string]any{},
			shouldPass: false,
		},
		{
			name:       "fails when no data and value is empty",
			value:      "",
			data:       nil,
			shouldPass: false,
		},
		{
			name:       "fails when no data and value is nil",
			value:      nil,
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

func TestRequiredWithoutRule_InvalidParameters(t *testing.T) {
	_, err := conditional.NewRequiredWithoutRule([]string{})
	if err == nil {
		t.Error("expected error when creating RequiredWithoutRule without parameters")
	}
}
