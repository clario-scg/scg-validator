package comparison_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/comparison"
)

func TestDifferentRule(t *testing.T) {
	tests := []struct {
		name       string
		value      any
		data       map[string]any
		shouldPass bool
	}{
		{
			name:       "values are different",
			value:      "foo",
			data:       map[string]any{"bar": "baz"},
			shouldPass: true,
		},
		{
			name:       "values are equal",
			value:      "foo",
			data:       map[string]any{"bar": "foo"},
			shouldPass: false,
		},
		{
			name:       "comparison field missing",
			value:      "foo",
			data:       map[string]any{},
			shouldPass: false,
		},
		{
			name:       "nil values (equal)",
			value:      nil,
			data:       map[string]any{"bar": nil},
			shouldPass: false,
		},
		{
			name:       "nil vs non-nil",
			value:      nil,
			data:       map[string]any{"bar": "something"},
			shouldPass: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rule, err := comparison.NewDifferentRule([]string{"bar"})
			if err != nil {
				t.Fatalf("Failed to create DifferentRule: %v", err)
			}

			ctx := contract.NewValidationContext("field", tc.value, []string{"bar"}, tc.data)
			err = rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("❌ %s: expected pass but got error: %v", tc.name, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("❌ %s: expected failure but passed", tc.name)
			}
		})
	}
}
