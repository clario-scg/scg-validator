package date_test

import (
	"testing"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/date"
)

func TestBeforeRule(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	nowStr := now.Format(time.RFC3339)
	yesterdayStr := now.Add(-24 * time.Hour).Format(time.RFC3339)
	tomorrowStr := now.Add(24 * time.Hour).Format(time.RFC3339)

	rule, err := date.NewBeforeRule([]string{nowStr})
	if err != nil {
		t.Fatalf("failed to create BeforeRule: %v", err)
	}

	tests := []struct {
		name      string
		value     any
		wantValid bool
	}{
		{"valid: date before", yesterdayStr, true},
		{"invalid: date after", tomorrowStr, false},
		{"invalid: date equal", nowStr, false},
		{"invalid: non-string type", 42, false},
		{"invalid: empty string", "", false},
		{"invalid: malformed string", "not-a-date", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Use real ValidationContext instead of mock
			ctx := contract.NewValidationContext("date_field", tc.value, nil, nil)
			err := rule.Validate(ctx)

			if tc.wantValid && err != nil {
				t.Errorf("expected valid, got error: %v", err)
			}
			if !tc.wantValid && err == nil {
				t.Errorf("expected error, got none")
			}
		})
	}
}
