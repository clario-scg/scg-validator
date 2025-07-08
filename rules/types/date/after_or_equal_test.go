package date_test

import (
	"testing"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/date"
)

func TestAfterOrEqualRule(t *testing.T) {
	t.Parallel()

	// Setup reference time
	now := time.Now().Truncate(time.Second) // truncate to avoid nanosecond mismatch
	nowStr := now.Format(time.RFC3339)
	yesterdayStr := now.AddDate(0, 0, -1).Format(time.RFC3339)
	tomorrowStr := now.AddDate(0, 0, 1).Format(time.RFC3339)

	// Create rule with 'now' as the reference point
	rule, err := date.NewAfterOrEqualRule([]string{nowStr})
	if err != nil {
		t.Fatalf("failed to create rule: %v", err)
	}

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"valid - after", tomorrowStr, true},
		{"valid - equal", nowStr, true},
		{"invalid - before", yesterdayStr, false},
		{"invalid - non-string type", 123, false},
		{"invalid - empty string", "", false},
		{"invalid - malformed string", "2023-13-99", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := contract.NewValidationContext("start_date", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.shouldPass && err != nil {
				t.Errorf("expected success for value %v, got error: %v", tc.value, err)
			}
			if !tc.shouldPass && err == nil {
				t.Errorf("expected failure for value %v, but got no error", tc.value)
			}
		})
	}
}
