package date_test

import (
	"testing"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/date"
)

func TestDateEqualsRule(t *testing.T) {
	baseTime := time.Now().Truncate(time.Second)
	baseStr := baseTime.Format(time.RFC3339)
	nextDayStr := baseTime.AddDate(0, 0, 1).Format(time.RFC3339)

	rule, err := date.NewDateEqualsRule([]string{baseStr})
	if err != nil {
		t.Fatalf("failed to create rule: %v", err)
	}

	tests := []struct {
		name    string
		input   any
		wantErr bool
	}{
		{"exact match date", baseStr, false},
		{"different date", nextDayStr, true},
		{"non-string type", 123, true},
		{"empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("field", tt.input, nil, nil)
			err := rule.Validate(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
