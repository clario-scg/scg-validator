package conditional_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/conditional"
)

func TestProhibitsRule(t *testing.T) {
	rule, err := conditional.NewProhibitsRule([]string{"other_field_1", "other_field_2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name string
		data map[string]any
		want bool // true means expected to pass (no error)
	}{
		{
			name: "✅ field is present, no prohibited fields → pass",
			data: map[string]any{"field": "value"},
			want: true,
		},
		{
			name: "❌ field is present, one prohibited field → fail",
			data: map[string]any{"field": "value", "other_field_1": "something"},
			want: false,
		},
		{
			name: "❌ field is present, all prohibited fields → fail",
			data: map[string]any{
				"field":         "value",
				"other_field_1": "x",
				"other_field_2": "y",
			},
			want: false,
		},
		{
			name: "✅ field is absent, prohibited fields present → pass",
			data: map[string]any{
				"other_field_1": "x",
				"other_field_2": "y",
			},
			want: true,
		},
		{
			name: "✅ neither field nor prohibited fields → pass",
			data: map[string]any{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use real ValidationContext instead of mock
			fieldValue, exists := tt.data["field"]
			if !exists {
				fieldValue = nil
			}

			ctx := contract.NewValidationContext("field", fieldValue, nil, tt.data)

			err := rule.Validate(ctx)
			passed := err == nil

			if passed != tt.want {
				t.Errorf("%s: expected pass=%v, got error=%v", tt.name, tt.want, err)
			}
		})
	}
}
