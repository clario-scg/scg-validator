package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringRule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestPasswordRule_ValidationScenarios(t *testing.T) {
	t.Helper()

	tests := []struct {
		name       string
		params     []string
		value      any
		shouldPass bool
	}{
		// Min length
		{"fail - under min", []string{"min:10"}, "short", false},
		{"pass - exact min", []string{"min:10"}, "1234567890", true},

		// Symbols
		{"pass - has symbol", []string{"symbols"}, "pass@word", true},
		{"fail - no symbol", []string{"symbols"}, "password", false},

		// Numbers
		{"pass - has number", []string{"numbers"}, "password123", true},
		{"fail - no number", []string{"numbers"}, "password", false},

		// Letters
		{"pass - has letter", []string{"letters"}, "p@ssword123", true},
		{"fail - no letter", []string{"letters"}, "@123456789", false},

		// Mixedcase
		{"pass - mixed case", []string{"mixedcase"}, "Password", true},
		{"fail - all lowercase", []string{"mixedcase"}, "password", false},
		{"fail - all uppercase", []string{"mixedcase"}, "PASSWORD", false},

		// Uppercase
		{"pass - has uppercase", []string{"uppercase"}, "PASSword", true},
		{"fail - no uppercase", []string{"uppercase"}, "password", false},

		// Lowercase
		{"pass - has lowercase", []string{"lowercase"}, "password", true},
		{"fail - no lowercase", []string{"lowercase"}, "PASSWORD", false},

		// Complex combinations
		{"pass - all constraints", []string{"min:12", "mixedcase", "numbers", "symbols"}, "ComplexP@ss123", true},
		{"fail - missing symbol", []string{"min:12", "mixedcase", "numbers", "symbols"}, "ComplexPass123", false},
		{"fail - missing number", []string{"min:12", "mixedcase", "numbers", "symbols"}, "ComplexP@ssword", false},
		{"fail - missing mixed", []string{"min:12", "mixedcase", "numbers", "symbols"}, "complexp@ss123", false},
		{"fail - too short", []string{"min:12", "mixedcase", "numbers", "symbols"}, "CplxP@s1", false},

		// Invalid input
		{"fail - non-string input", []string{"min:8"}, 12345678, false},
		{"fail - nil input", []string{"min:8"}, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule, err := stringRule.NewPasswordRule(tt.params)
			if err != nil {
				t.Fatalf("failed to create password rule with %v: %v", tt.params, err)
			}

			ctx := contract.NewValidationContext("password", tt.value, nil, nil)
			err = rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("expected fail, but got success for value: %v", tt.value)
			}
		})
	}
}

func TestPasswordRule_InvalidMinParam(t *testing.T) {
	t.Helper()

	_, err := stringRule.NewPasswordRule([]string{"min:abc"})
	if err == nil {
		t.Error("expected error for invalid min parameter, got nil")
	}
}
