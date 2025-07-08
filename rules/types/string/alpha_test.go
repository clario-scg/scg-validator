package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringRule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestAlphaRule(t *testing.T) {
	rule, err := stringRule.NewAlphaRule(nil)
	if err != nil {
		t.Fatalf("failed to create AlphaRule: %v", err)
	}

	tests := []struct {
		name  string
		input any
		valid bool
	}{
		{"ASCII only", "Hello", true},
		{"Unicode letter", "Héllo", true},
		{"Unicode mark", "naïve", true},
		{"Korean characters", "안녕하세요", true},
		{"Arabic letters", "مرحبا", true},
		{"with emoji", "Hello😊", false},
		{"letters and digits", "abc123", false},
		{"with space", "Hello World", false},
		{"with punctuation", "Hi!", false},
		{"with symbols", "hello@", false},
		{"empty string", "", false},
		{"int value", 42, false},
		{"bool value", true, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.input, nil, nil)
			err := rule.Validate(ctx)

			if tt.valid && err != nil {
				t.Errorf("[%s] expected valid, got error: %v", tt.name, err)
			}
			if !tt.valid && err == nil {
				t.Errorf("[%s] expected error, got nil (input: %#v)", tt.name, tt.input)
			}
		})
	}
}
