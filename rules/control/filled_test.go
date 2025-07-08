package control_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/control"
)

func TestFilledRule(t *testing.T) {
	rule, err := control.NewFilledRule()
	if err != nil {
		t.Fatalf("Failed to create FilledRule: %v", err)
	}

	tests := []struct {
		name       string
		field      string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "passes when field exists and is non-empty string",
			field:      "username",
			value:      "john",
			data:       map[string]any{"username": "john"},
			shouldPass: true,
		},
		{
			name:       "fails when field exists but value is empty string",
			field:      "username",
			value:      "",
			data:       map[string]any{"username": ""},
			shouldPass: false,
		},
		{
			name:       "fails when field exists but value is nil",
			field:      "username",
			value:      nil,
			data:       map[string]any{"username": nil},
			shouldPass: false,
		},
		{
			name:       "passes when field is missing from data",
			field:      "username",
			value:      "anything",
			data:       map[string]any{},
			shouldPass: true,
		},
		{
			name:       "passes when data map is nil",
			field:      "username",
			value:      "anything",
			data:       nil,
			shouldPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext(tt.field, tt.value, nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("Expected success, got error: %v", err)
			}

			if !tt.shouldPass && err == nil {
				t.Error("Expected failure, but validation passed")
			}
		})
	}
}
