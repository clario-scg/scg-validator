package collection_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/collection"
)

func TestMapRule(t *testing.T) {
	t.Parallel()

	rule, err := collection.NewMapRule(nil)
	if err != nil {
		t.Fatalf("failed to create MapRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// ✅ Valid maps
		{"valid - string map", map[string]string{"key": "value"}, true},
		{"valid - empty map", map[string]int{}, true},
		{"valid - nested map", map[string]map[string]int{"outer": {"inner": 1}}, true},

		// ❌ Invalid values
		{"invalid - slice", []string{"a", "b"}, false},
		{"invalid - string", "not a map", false},
		{"invalid - nil", nil, false},
		{"invalid - int", 123, false},
		{"invalid - struct", struct{ A int }{1}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := contract.NewValidationContext("map_field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass for value %#v, got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure for value %#v, but got none", tc.value)
			}
		})
	}
}
