package format_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/format"
)

func TestJSONRule(t *testing.T) {
	rule, err := format.NewJSONRule([]string{})
	if err != nil {
		t.Fatalf("Failed to create JSONRule: %v", err)
	}

	tests := []struct {
		name       string
		value      interface{}
		shouldPass bool
	}{
		{"valid JSON object", `{"ipRuleName": "test"}`, true},
		{"valid JSON array", `[1, 2, 3]`, true},
		{"valid JSON string", `"hello"`, true},
		{"valid JSON number", `123`, true},
		{"valid JSON boolean", `true`, true},
		{"valid JSON null", `null`, true},
		{"invalid JSON malformed", `{"ipRuleName": test}`, false},
		{"invalid JSON incomplete", `{"ipRuleName":`, false},
		{"invalid non-JSON string", "not json", false},
		{"empty string", "", false},
		{"non-string type", 123, false},
		{"nil value", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.value, []string{}, map[string]any{})
			err := rule.Validate(ctx)

			if tt.shouldPass && err != nil {
				t.Errorf("Expected validation to pass for %v, but got error: %v", tt.value, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("Expected validation to fail for %v, but it passed", tt.value)
			}
		})
	}
}
