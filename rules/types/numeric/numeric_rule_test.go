package numeric_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/numeric"
)

func TestNumericRule(t *testing.T) {
	rule, err := numeric.NewNumericRule()
	if err != nil {
		t.Fatalf("failed to create NumericRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// Integer values
		{"int", 123, true},
		{"int8", int8(-5), true},
		{"int16", int16(0), true},
		{"int32", int32(99999), true},
		{"int64", int64(-99999), true},
		{"uint", uint(123), true},
		{"uint8", uint8(255), true},
		{"uint16", uint16(65535), true},
		{"uint32", uint32(4294967295), true},
		{"uint64", uint64(18446744073709551615), true},

		// Float values
		{"float32", float32(123.45), true},
		{"float64", float64(-123.45), true},

		// Numeric strings
		{"string int", "123", true},
		{"string float", "123.456", true},
		{"string negative int", "-42", true},
		{"string negative float", "-3.14", true},
		{"string scientific", "1.23e3", true},

		// Invalid values
		{"string non-numeric", "abc", false},
		{"string mixed", "123abc", false},
		{"bool true", true, false},
		{"bool false", false, false},
		{"nil", nil, false},
		{"slice", []int{1, 2, 3}, false},
		{"map", map[string]int{"x": 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected to pass but failed: value=%v, error=%v", tt.value, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected to fail but passed: value=%v", tt.value)
			}
		})
	}
}
