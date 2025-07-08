package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

func TestConfirmedRule(t *testing.T) {
	rule, err := comparison.NewConfirmedRule()
	if err != nil {
		t.Fatalf("Failed to create ConfirmedRule: %v", err)
	}

	tests := []struct {
		name       string
		fieldName  string
		value      interface{}
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "valid matching strings",
			fieldName:  "password",
			value:      "secret123",
			data:       map[string]any{"password": "secret123", "password_confirmation": "secret123"},
			shouldPass: true,
		},
		{
			name:       "invalid non-matching strings",
			fieldName:  "password",
			value:      "secret123",
			data:       map[string]any{"password": "secret123", "password_confirmation": "different"},
			shouldPass: false,
		},
		{
			name:       "valid matching numbers",
			fieldName:  "pin",
			value:      1234,
			data:       map[string]any{"pin": 1234, "pin_confirmation": 1234},
			shouldPass: true,
		},
		{
			name:       "invalid non-matching numbers",
			fieldName:  "pin",
			value:      1234,
			data:       map[string]any{"pin": 1234, "pin_confirmation": 5678},
			shouldPass: false,
		},
		{
			name:       "invalid missing confirmation field",
			fieldName:  "password",
			value:      "secret123",
			data:       map[string]any{"password": "secret123"},
			shouldPass: false,
		},
		{
			name:       "invalid with nil value",
			fieldName:  "password",
			value:      nil,
			data:       map[string]any{"password": nil, "password_confirmation": "secret123"},
			shouldPass: false,
		},
		{
			name:       "valid with nil values",
			fieldName:  "password",
			value:      nil,
			data:       map[string]any{"password": nil, "password_confirmation": nil},
			shouldPass: true,
		},
		{
			name:       "invalid with no data",
			fieldName:  "password",
			value:      "secret123",
			data:       nil,
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext(tt.fieldName, tt.value, nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("❌ %s: expected validation to pass, but got error: %v", tt.name, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("❌ %s: expected validation to fail, but passed", tt.name)
			}
		})
	}
}
