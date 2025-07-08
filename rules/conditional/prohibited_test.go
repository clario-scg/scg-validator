package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/conditional"
)

func TestProhibitedRule(t *testing.T) {
	rule, err := conditional.NewProhibitedRule()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name       string
		field      string
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "fails when field is present",
			field:      "field",
			data:       map[string]any{"field": "value"},
			shouldPass: false,
		},
		{
			name:       "passes when field is not present",
			field:      "field",
			data:       map[string]any{"other_field": "value"},
			shouldPass: true,
		},
		{
			name:       "fails when field is present but nil",
			field:      "field",
			data:       map[string]any{"field": nil},
			shouldPass: false,
		},
		{
			name:       "fails when field is present but empty string",
			field:      "field",
			data:       map[string]any{"field": ""},
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext(tt.field, tt.data[tt.field], nil, tt.data)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, but got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Error("expected failure, but got no error")
			}
		})
	}
}
