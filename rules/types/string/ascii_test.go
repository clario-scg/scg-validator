package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringrule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestAsciiRule(t *testing.T) {
	rule, err := stringrule.NewASCIIRule()
	if err != nil {
		t.Fatalf("failed to create ASCIIRule: %v", err)
	}

	tests := []struct {
		name       string
		input      any
		shouldPass bool
	}{
		{"ASCII only lowercase", "ascii", true},
		{"ASCII with space and symbols", "ascii 123!@#", true},
		{"ASCII control chars", "\n\t\r", true},
		{"empty string", "", true},
		{"non-ASCII accented char", "café", false},
		{"emoji input", "hello 😊", false},
		{"mixed ASCII + non-ASCII", "testµ", false},
		{"non-string input int", 123, false},
		{"non-string input bool", true, false},
		{"nil input", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("ascii_field", tt.input, nil, nil)
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected fail, but passed: input=%v", tt.input)
			}
		})
	}
}
