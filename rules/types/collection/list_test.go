package collection_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/collection"
)

func TestListRule(t *testing.T) {
	t.Parallel()

	rule, err := collection.NewListRule(nil)
	if err != nil {
		t.Fatalf("failed to create ListRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// ✅ Valid cases
		{"valid - slice of strings", []string{"a", "b"}, true},
		{"valid - empty slice of ints", []int{}, true},
		{"valid - array of 2 ints", [2]int{1, 2}, true},

		// ❌ Invalid cases
		{"invalid - map", map[string]string{"key": "value"}, false},
		{"invalid - string", "not a list", false},
		{"invalid - nil", nil, false},
		{"invalid - int", 123, false},
		{"invalid - struct", struct{ Name string }{"Foo"}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := contract.NewValidationContext("list_field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass for value %#v, but got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure for value %#v, but got no error", tc.value)
			}
		})
	}
}
