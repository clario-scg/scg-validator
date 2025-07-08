package date_test

import (
	"testing"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/date"
)

func TestAfterRule(t *testing.T) {
	t.Parallel()

	now := time.Now().Truncate(time.Second) // ensure seconds match
	nowStr := now.Format(time.RFC3339)
	tomorrow := now.AddDate(0, 0, 1).Format(time.RFC3339)
	yesterday := now.AddDate(0, 0, -1).Format(time.RFC3339)

	rule, err := date.NewAfterRule([]string{nowStr})
	if err != nil {
		t.Fatalf("failed to create AfterRule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"valid - after", tomorrow, true},
		{"invalid - before", yesterday, false},
		{"invalid - equal", nowStr, false},
		{"invalid - int", 123, false},
		{"invalid - empty string", "", false},
		{"invalid - malformed date", "2023-13-99", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := contract.NewValidationContext("start_time", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected success for value %v, but got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure for value %v, but got no error", tc.value)
			}
		})
	}
}
