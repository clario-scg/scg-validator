package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/rules/conditional"

	"github.com/next-trace/scg-validator/contract"
)

func TestRequiredWithRule(t *testing.T) {
	rule, err := conditional.NewRequiredWithRule([]string{"other_field"})
	if err != nil {
		t.Fatalf("Failed to create RequiredWithRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "passes when other field is present and value is non-empty",
			value:      "test",
			data:       map[string]any{"other_field": "value"},
			shouldPass: true,
		},
		{
			name:       "fails when other field is present and value is empty string",
			value:      "",
			data:       map[string]any{"other_field": "value"},
			shouldPass: false,
		},
		{
			name:       "fails when other field is present and value is nil",
			value:      nil,
			data:       map[string]any{"other_field": "value"},
			shouldPass: false,
		},
		{
			name:       "passes when other field is not present and value is empty string",
			value:      "",
			data:       map[string]any{},
			shouldPass: true,
		},
		{
			name:       "passes when other field is not present and value is nil",
			value:      nil,
			data:       map[string]any{},
			shouldPass: true,
		},
		{
			name:       "passes when data is nil",
			value:      "",
			data:       nil,
			shouldPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected failure, got pass for value: %#v", tt.value)
			}
		})
	}
}

func TestRequiredWithRule_InvalidParameters(t *testing.T) {
	_, err := conditional.NewRequiredWithRule([]string{})
	if err == nil {
		t.Error("expected error when creating RequiredWithRule without parameters")
	}
}
