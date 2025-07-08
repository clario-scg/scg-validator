package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringrule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestDoesntEndWithRule(t *testing.T) {
	rule, err := stringrule.NewDoesntEndWithRule([]string{"foo", "bar"})
	if err != nil {
		t.Fatalf("failed to create DoesntEndWithRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"does not end with suffix", "hello", true},
		{"ends with foo", "myfoo", false},
		{"ends with bar", "testbar", false},
		{"suffix in middle", "foobar", false},
		{"empty string", "", true},
		{"only suffix", "foo", false},
		{"non-string type", 123, false},
		{"nil value", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass but got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected failure but passed: value=%v", tt.value)
			}
		})
	}
}
