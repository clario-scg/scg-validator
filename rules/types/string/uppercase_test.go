package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringrule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestUppercaseRule(t *testing.T) {
	rule, err := stringrule.NewUppercaseRule()
	if err != nil {
		t.Fatalf("failed to create UppercaseRule: %v", err)
	}

	testCases := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{name: "all uppercase", value: "UPPERCASE", shouldPass: true},
		{name: "capitalized", value: "Uppercase", shouldPass: false},
		{name: "uppercase with numbers", value: "UPPERCASE123", shouldPass: true},
		{name: "all lowercase", value: "uppercase", shouldPass: false},
		{name: "numeric input", value: 123, shouldPass: false},
		{name: "empty string", value: "", shouldPass: true},
		{name: "nil input", value: nil, shouldPass: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass but got error: %v", err)
			}

			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure but got success for value: %v", tc.value)
			}
		})
	}
}
