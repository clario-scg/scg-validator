package boolean_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/boolean"
)

func TestBooleanRule(t *testing.T) {
	t.Parallel()

	rule, err := boolean.NewBooleanRule()
	if err != nil {
		t.Fatalf("failed to create BooleanRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// ✅ Native bool types
		{"valid - bool true", true, true},
		{"valid - bool false", false, true},

		// ✅ Accepted string values
		{"valid - string 'true'", "true", true},
		{"valid - string 'false'", "false", true},
		{"valid - string '1'", "1", true},
		{"valid - string '0'", "0", true},
		{"valid - string 'yes'", "yes", true},
		{"valid - string 'no'", "no", true},
		{"valid - string 'on'", "on", true},
		{"valid - string 'off'", "off", true},
		{"valid - uppercase 'TRUE'", "TRUE", true},
		{"valid - uppercase 'FALSE'", "FALSE", true},

		// ✅ Integers
		{"valid - int 1", 1, true},
		{"valid - int 0", 0, true},
		{"invalid - int 2", 2, false},
		{"invalid - int -1", -1, false},

		// ✅ Unsigned integers
		{"valid - uint(1)", uint(1), true},
		{"valid - uint(0)", uint(0), true},
		{"invalid - uint(2)", uint(2), false},

		// ✅ Floats
		{"valid - float64(1.0)", 1.0, true},
		{"valid - float64(0.0)", 0.0, true},
		{"invalid - float64(1.5)", 1.5, false},

		// ❌ Invalid types
		{"invalid - string 'maybe'", "maybe", false},
		{"invalid - nil", nil, false},
		{"invalid - slice", []int{1, 0}, false},
		{"invalid - struct", struct{}{}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := contract.NewValidationContext("boolean_field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass for value %v, but got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure for value %v, but got none", tc.value)
			}
		})
	}
}
