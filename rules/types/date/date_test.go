package date_test

import (
	"testing"
	"time"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/types/date"
)

func TestDateRule(t *testing.T) {
	rule, err := date.NewDateRule([]string{}) // defaults to RFC3339
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name    string
		value   any
		wantErr bool
	}{
		{
			name:    "valid RFC3339 date",
			value:   time.Now().Format(time.RFC3339),
			wantErr: false,
		},
		{
			name:    "invalid date format",
			value:   "invalid-date",
			wantErr: true,
		},
		{
			name:    "non-string type",
			value:   123,
			wantErr: true,
		},
		{
			name:    "empty string",
			value:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("birthdate", tt.value, nil, nil)
			err := rule.Validate(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}

	t.Run("valid custom format", func(t *testing.T) {
		customRule, err := date.NewDateRule([]string{"2006-01-02"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ctx := contract.NewValidationContext("release_date", "2023-12-25", nil, nil)
		if err := customRule.Validate(ctx); err != nil {
			t.Errorf("expected valid custom date format, got error: %v", err)
		}
	})

	t.Run("invalid custom format", func(t *testing.T) {
		customRule, err := date.NewDateRule([]string{"2006-01-02"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ctx := contract.NewValidationContext("release_date", "25/12/2023", nil, nil)
		if err := customRule.Validate(ctx); err == nil {
			t.Errorf("expected error for mismatched format, got nil")
		}
	})
}
