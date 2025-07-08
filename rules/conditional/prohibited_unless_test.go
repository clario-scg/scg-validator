package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/conditional"
)

func TestProhibitedUnlessRule(t *testing.T) {
	rule, err := conditional.NewProhibitedUnlessRule([]string{"other_field", "some_value"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name       string
		data       map[string]any
		field      string
		shouldPass bool
	}{
		{
			name:       "passes when field is present and other_field matches value",
			data:       map[string]any{"field": "value", "other_field": "some_value"},
			field:      "field",
			shouldPass: true,
		},
		{
			name:       "passes when field is absent and other_field matches value",
			data:       map[string]any{"other_field": "some_value"},
			field:      "field",
			shouldPass: true,
		},
		{
			name:       "fails when field is present and other_field has different value",
			data:       map[string]any{"field": "value", "other_field": "different_value"},
			field:      "field",
			shouldPass: false,
		},
		{
			name:       "fails when field is present and other_field is missing",
			data:       map[string]any{"field": "value"},
			field:      "field",
			shouldPass: false,
		},
		{
			name:       "passes when field is absent and other_field is missing",
			data:       map[string]any{},
			field:      "field",
			shouldPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := tt.data[tt.field]
			ctx := contract.NewValidationContext(tt.field, value, nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
