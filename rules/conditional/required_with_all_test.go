package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/rules/conditional"

	"github.com/next-trace/scg-validator/contract"
)

func TestRequiredWithAllRule(t *testing.T) {
	rule, err := conditional.NewRequiredWithAllRule([]string{"foo", "bar"})
	if err != nil {
		t.Fatalf("Failed to create RequiredWithAllRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "passes when not all fields are present (none present)",
			value:      "",
			data:       map[string]any{},
			shouldPass: true,
		},
		{
			name:       "passes when not all fields are present (one present)",
			value:      "",
			data:       map[string]any{"foo": "x"},
			shouldPass: true,
		},
		{
			name:       "fails when all fields present and value is empty string",
			value:      "",
			data:       map[string]any{"foo": "x", "bar": "y"},
			shouldPass: false,
		},
		{
			name:       "fails when all fields present and value is nil",
			value:      nil,
			data:       map[string]any{"foo": "x", "bar": "y"},
			shouldPass: false,
		},
		{
			name:       "passes when all fields present and value is not empty",
			value:      "abc",
			data:       map[string]any{"foo": "x", "bar": "y"},
			shouldPass: true,
		},
		{
			name:       "passes when data map is nil",
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
				t.Errorf("expected pass but got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected failure, but got pass for value: %#v", tt.value)
			}
		})
	}
}

func TestRequiredWithAllRule_InvalidParameters(t *testing.T) {
	_, err := conditional.NewRequiredWithAllRule([]string{})
	if err == nil {
		t.Error("expected error for no parameters, got nil")
	}
}
