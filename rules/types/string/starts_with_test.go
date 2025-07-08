package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringRule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestStartsWithRule(t *testing.T) {
	tests := []struct {
		name       string
		params     []string
		value      any
		shouldPass bool
	}{
		// Valid cases
		{"pass - starts with foo", []string{"foo", "bar"}, "foobar", true},
		{"pass - starts with bar", []string{"foo", "bar"}, "barbaz", true},

		// Invalid cases
		{"fail - no matching prefix", []string{"foo", "bar"}, "bazqux", false},
		{"fail - empty string", []string{"foo", "bar"}, "", false},
		{"fail - non-string value", []string{"foo", "bar"}, 123, false},
		{"fail - nil input", []string{"foo", "bar"}, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule, err := stringRule.NewStartsWithRule(tt.params)
			if err != nil {
				t.Fatalf("failed to create rule with params %v: %v", tt.params, err)
			}

			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err = rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected fail, but passed: value=%v", tt.value)
			}
		})
	}
}
