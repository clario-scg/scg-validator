package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringRule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestAlphaNumRule(t *testing.T) {
	rule, err := stringRule.NewAlphaNumRule(nil)
	if err != nil {
		t.Fatalf("failed to create AlphaNumRule: %v", err)
	}

	tests := []struct {
		name  string
		input any
		valid bool
	}{
		{"letters and digits", "Hello123", true},
		{"letters only", "GoLang", true},
		{"digits only", "987654", true},
		{"unicode letters", "こんにちは123", true},
		{"contains space", "Hello 123", false},
		{"contains symbol", "Hello@World", false},
		{"contains dash", "Alpha-Num", false},
		{"contains underscore", "Alpha_Num", false},
		{"dot in input", "Alpha.Num", false},
		{"empty string", "", false},
		{"non-string (int)", 123, false},
		{"non-string (bool)", true, false},
		{"nil value", nil, false},
		{"slice input", []rune{'a', 'b'}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.input, nil, nil)
			err := rule.Validate(ctx)

			if tt.valid && err != nil {
				t.Errorf("[%s] expected valid, got error: %v", tt.name, err)
			}
			if !tt.valid && err == nil {
				t.Errorf("[%s] expected error, but got nil", tt.name)
			}
		})
	}
}
