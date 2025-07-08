package control_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/control"
)

func TestPresentRule(t *testing.T) {
	rule, err := control.NewPresentRule()
	if err != nil {
		t.Fatalf("failed to create PresentRule: %v", err)
	}

	tests := []struct {
		name       string
		field      string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "field present with non-empty value",
			field:      "username",
			value:      "john",
			data:       map[string]any{"username": "john"},
			shouldPass: true,
		},
		{
			name:       "field present with empty string",
			field:      "username",
			value:      "",
			data:       map[string]any{"username": ""},
			shouldPass: true,
		},
		{
			name:       "field present with nil",
			field:      "username",
			value:      nil,
			data:       map[string]any{"username": nil},
			shouldPass: true,
		},
		{
			name:       "field missing from map",
			field:      "username",
			value:      "anything",
			data:       map[string]any{},
			shouldPass: false,
		},
		{
			name:       "nil data map",
			field:      "username",
			value:      "anything",
			data:       nil,
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext(tt.field, tt.value, nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected failure, but passed for value: %v", tt.value)
			}
		})
	}
}
