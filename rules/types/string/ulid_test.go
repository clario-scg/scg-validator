package string_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	stringrule "github.com/next-trace/scg-validator/rules/types/string"
)

func TestUlidRule(t *testing.T) {
	rule, err := stringrule.NewUlidRule()
	if err != nil {
		t.Fatalf("failed to create UlidRule: %v", err)
	}

	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{"valid - correct ULID", "01E4Z5Q6C8N9J1M2P3R4S5T6V7", true},
		{"invalid - lowercase ULID", "01e4z5q6c8n9j1m2p3r4s5t6v7", false},
		{"invalid - symbol included", "01E4Z5Q6C8N9J1M2P3R4S5T6V!", false},
		{"invalid - too short", "01E4Z5Q6C8N9J1M2P3R4S5T6V", false},
		{"invalid - too long", "01E4Z5Q6C8N9J1M2P3R4S5T6V7A", false},
		{"invalid - non-string", 123, false},
		{"invalid - empty string", "", false},
		{"invalid - nil input", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("ulid", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.want && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.want && err == nil {
				t.Errorf("expected fail, but passed: value=%v", tt.value)
			}
		})
	}
}
