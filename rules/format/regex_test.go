package format_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/format"
)

func TestRegexRule_Validation(t *testing.T) {
	t.Parallel()

	rule, err := format.NewRegexRule([]string{"^hello$"})
	if err != nil {
		t.Fatalf("failed to create RegexRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"exact match", "hello", true},
		{"different word", "world", false},
		{"prefix only", "hell", false},
		{"suffix only", "hello world", false},
		{"different casing", "Hello", false},
		{"empty string", "", false},
		{"integer type", 123, false},
		{"nil input", nil, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := contract.NewValidationContext("greeting", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass for value %v, got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected fail for value %v, but passed", tc.value)
			}
		})
	}
}

func TestRegexRule_InvalidConstruction(t *testing.T) {
	t.Parallel()

	t.Run("missing parameters", func(t *testing.T) {
		t.Parallel()
		_, err := format.NewRegexRule([]string{})
		if err == nil {
			t.Error("expected error for missing regex parameter")
		}
	})

	t.Run("malformed regex pattern", func(t *testing.T) {
		t.Parallel()
		_, err := format.NewRegexRule([]string{"["}) // Invalid regex
		if err == nil {
			t.Error("expected error for invalid regex pattern")
		}
	})
}
