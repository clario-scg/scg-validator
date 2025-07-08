package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/rules/conditional"

	"github.com/next-trace/scg-validator/contract"
)

func TestRequiredIfRule(t *testing.T) {
	rule, err := conditional.NewRequiredIfRule([]string{"status", "active"})
	if err != nil {
		t.Fatalf("Failed to create RequiredIfRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "passes when condition met and field is present",
			value:      "test",
			data:       map[string]any{"status": "active"},
			shouldPass: true,
		},
		{
			name:       "fails when condition met but field is empty string",
			value:      "",
			data:       map[string]any{"status": "active"},
			shouldPass: false,
		},
		{
			name:       "fails when condition met but field is nil",
			value:      nil,
			data:       map[string]any{"status": "active"},
			shouldPass: false,
		},
		{
			name:       "passes when condition not met and field is empty string",
			value:      "",
			data:       map[string]any{"status": "inactive"},
			shouldPass: true,
		},
		{
			name:       "passes when condition not met and field is nil",
			value:      nil,
			data:       map[string]any{"status": "inactive"},
			shouldPass: true,
		},
		{
			name:       "passes when no data map is provided",
			value:      "",
			data:       nil,
			shouldPass: true,
		},
		{
			name:       "passes when data map does not include condition field",
			value:      "",
			data:       map[string]any{},
			shouldPass: true,
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
				t.Errorf("expected failure but got success for value: %v", tt.value)
			}
		})
	}
}

func TestRequiredIfRule_InvalidParameters(t *testing.T) {
	tests := []struct {
		name   string
		params []string
	}{
		{"no parameters", []string{}},
		{"one parameter", []string{"field"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := conditional.NewRequiredIfRule(tt.params)
			if err == nil {
				t.Errorf("expected error for parameters: %v", tt.params)
			}
		})
	}
}
