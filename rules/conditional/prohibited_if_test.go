package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/conditional"
)

func TestProhibitedIfRule(t *testing.T) {
	rule, err := conditional.NewProhibitedIfRule([]string{"other_field", "some_value"})
	if err != nil {
		t.Fatalf("unexpected error creating rule: %v", err)
	}

	tests := []struct {
		name       string
		data       map[string]any
		field      string
		shouldPass bool
	}{
		{
			name:       "fails when field is present and other_field matches",
			field:      "field",
			data:       map[string]any{"field": "value", "other_field": "some_value"},
			shouldPass: false,
		},
		{
			name:       "passes when field is not present and other_field matches",
			field:      "field",
			data:       map[string]any{"other_field": "some_value"},
			shouldPass: true,
		},
		{
			name:       "passes when other_field has different value",
			field:      "field",
			data:       map[string]any{"field": "value", "other_field": "different_value"},
			shouldPass: true,
		},
		{
			name:       "passes when other_field is missing",
			field:      "field",
			data:       map[string]any{"field": "value"},
			shouldPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext(tt.field, tt.data[tt.field], nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass but got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected error but got none")
			}
		})
	}
}
