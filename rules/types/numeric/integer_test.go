package numeric_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/numeric"
)

func TestIntegerRule(t *testing.T) {
	rule, err := numeric.NewIntegerRule()
	if err != nil {
		t.Fatalf("failed to create IntegerRule: %v", err)
	}

	cases := []struct {
		name  string
		input any
		valid bool
	}{
		// Signed integers
		{"int", int(123), true},
		{"int8", int8(123), true},
		{"int16", int16(123), true},
		{"int32", int32(123), true},
		{"int64", int64(123), true},

		// Unsigned integers
		{"uint", uint(123), true},
		{"uint8", uint8(123), true},
		{"uint16", uint16(123), true},
		{"uint32", uint32(123), true},
		{"uint64", uint64(123), true},

		// Valid strings
		{"string integer", "123", true},
		{"string negative", "-123", true},

		// Floats - valid
		{"float32 integer", float32(100.0), true},
		{"float64 integer", float64(100.0), true},

		// Floats - invalid
		{"float32 non-integer", float32(123.45), false},
		{"float64 non-integer", float64(123.45), false},

		// Invalid strings
		{"string float", "123.45", false},
		{"string non-numeric", "abc", false},

		// Other invalid types
		{"bool", true, false},
		{"nil", nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", c.input, nil, nil)
			err := rule.Validate(ctx)

			if c.valid && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !c.valid && err == nil {
				t.Errorf("expected fail, but passed: %v", c.input)
			}
		})
	}
}
