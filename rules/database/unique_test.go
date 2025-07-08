package database_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/registry/database"
	databaseRule "github.com/next-trace/scg-validator/rules/database"
)

type mockUniquePresenceVerifier struct {
	uniqueResult bool
	uniqueErr    error
}

func (m *mockUniquePresenceVerifier) Exists(_, _ string, _ any) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func (m *mockUniquePresenceVerifier) Unique(_, _ string, _ any) (bool, error) {
	return m.uniqueResult, m.uniqueErr
}

func TestUniqueRule(t *testing.T) {
	t.Run("missing parameters", func(t *testing.T) {
		rule, _ := databaseRule.NewUniqueRule()

		ctx := contract.NewValidationContext("email", "value", []string{"only_table"}, nil)

		if err := rule.Validate(ctx); err == nil {
			t.Error("expected error for missing field parameter")
		}
	})

	t.Run("verifier not registered", func(t *testing.T) {
		rule, _ := databaseRule.NewUniqueRule()

		ctx := contract.NewValidationContext("email", "value", []string{"unknown_table", "email"}, nil)

		if err := rule.Validate(ctx); err == nil {
			t.Error("expected error for missing presence verifier")
		}
	})

	t.Run("verifier returns error", func(t *testing.T) {
		rule, _ := databaseRule.NewUniqueRule()
		table := testTableName
		database.RegisterPresenceVerifier(table, &mockUniquePresenceVerifier{
			uniqueErr: errors.New("database failure"),
		})

		ctx := contract.NewValidationContext("email", "broken@domain.com", []string{table, "email"}, nil)

		err := rule.Validate(ctx)
		if err == nil || err.Error() != "database failure" {
			t.Errorf("expected 'database failure', got %v", err)
		}
	})

	t.Run("value is not unique", func(t *testing.T) {
		rule, _ := databaseRule.NewUniqueRule()
		table := testTableName
		database.RegisterPresenceVerifier(table, &mockUniquePresenceVerifier{
			uniqueResult: false,
		})

		ctx := contract.NewValidationContext("email", "duplicate@domain.com", []string{table, "email"}, nil)

		if err := rule.Validate(ctx); err == nil {
			t.Error("expected error for non-unique value")
		}
	})

	t.Run("value is unique", func(t *testing.T) {
		rule, _ := databaseRule.NewUniqueRule()
		table := testTableName
		database.RegisterPresenceVerifier(table, &mockUniquePresenceVerifier{
			uniqueResult: true,
		})

		ctx := contract.NewValidationContext("email", "unique@domain.com", []string{table, "email"}, nil)

		if err := rule.Validate(ctx); err != nil {
			t.Errorf("expected no error for unique value, got %v", err)
		}
	})
}
