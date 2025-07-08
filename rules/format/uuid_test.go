package format_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/format"
)

func TestUUIDRule(t *testing.T) {
	t.Parallel()

	rule, err := format.NewUUIDRule([]string{})
	if err != nil {
		t.Fatalf("failed to create UUIDRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		// ✅ Valid UUIDs
		{"valid - UUID v4", "550e8400-e29b-41d4-a716-446655440000", true},
		{"valid - UUID v1", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},

		// ❌ Invalid formats
		{"invalid - missing dashes", "550e8400e29b41d4a716446655440000", false},
		{"invalid - truncated", "550e8400-e29b-41d4-a716", false},
		{"invalid - invalid character", "550e8400-e29b-41d4-a716-44665544000g", false},
		{"invalid - random string", "not-a-uuid", false},

		// ❌ Type issues
		{"invalid - empty string", "", false},
		{"invalid - integer type", 123, false},
		{"invalid - nil", nil, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := contract.NewValidationContext("uuid_field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected pass for %v, but got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected fail for %v, but got no error", tc.value)
			}
		})
	}
}
