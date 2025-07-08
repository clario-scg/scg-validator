package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/rules/conditional"

	"github.com/next-trace/scg-validator/contract"
)

func TestRequiredUnlessRule(t *testing.T) {
	rule, err := conditional.NewRequiredUnlessRule([]string{"status", "inactive"})
	if err != nil {
		t.Fatalf("Failed to create RequiredUnlessRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "passes when condition met (other field matches), value empty",
			value:      "",
			data:       map[string]any{"status": "inactive"},
			shouldPass: true,
		},
		{
			name:       "passes when condition met (other field matches), value nil",
			value:      nil,
			data:       map[string]any{"status": "inactive"},
			shouldPass: true,
		},
		{
			name:       "passes when condition not met, field provided",
			value:      "test",
			data:       map[string]any{"status": "active"},
			shouldPass: true,
		},
		{
			name:       "fails when condition not met, value empty",
			value:      "",
			data:       map[string]any{"status": "active"},
			shouldPass: false,
		},
		{
			name:       "fails when condition not met, value nil",
			value:      nil,
			data:       map[string]any{"status": "active"},
			shouldPass: false,
		},
		{
			name:       "fails when no data, value empty",
			value:      "",
			data:       nil,
			shouldPass: false,
		},
		{
			name:       "fails when field missing and value empty",
			value:      "",
			data:       map[string]any{},
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
				t.Errorf("expected fail but got pass for value: %v", tt.value)
			}
		})
	}
}
