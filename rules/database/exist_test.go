package database_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/registry/database"
	databaseRule "github.com/next-trace/scg-validator/rules/database"
)

const (
	testTableName = "users"
)

type mockExistPresenceVerifier struct {
	existsResult bool
	existsErr    error
}

func (m *mockExistPresenceVerifier) Exists(_, _ string, _ any) (bool, error) {
	return m.existsResult, m.existsErr
}

func (m *mockExistPresenceVerifier) Unique(_, _ string, _ any) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func TestExistRule(t *testing.T) {
	t.Run("should fail if no constructor parameters are passed", func(t *testing.T) {
		_, err := databaseRule.NewExistRule([]string{})
		if err == nil {
			t.Error("expected error when no parameters passed to NewExistRule")
		}
	})

	t.Run("should fail if context parameters are less than 2", func(t *testing.T) {
		rule, _ := databaseRule.NewExistRule([]string{testTableName})

		ctx := contract.NewValidationContext("email", "value", []string{"only_table"}, nil)

		if err := rule.Validate(ctx); err == nil {
			t.Error("expected error for missing field parameter")
		}
	})

	t.Run("should fail if verifier not registered", func(t *testing.T) {
		rule, _ := databaseRule.NewExistRule([]string{testTableName})

		ctx := contract.NewValidationContext("email", "test@example.com", []string{"unregistered_table", "email"}, nil)

		if err := rule.Validate(ctx); err == nil {
			t.Error("expected error for missing presence verifier")
		}
	})

	t.Run("should return error from verifier", func(t *testing.T) {
		rule, _ := databaseRule.NewExistRule([]string{testTableName})
		table := testTableName
		database.RegisterPresenceVerifier(table, &mockExistPresenceVerifier{
			existsErr: errors.New("db error"),
		})

		ctx := contract.NewValidationContext("email", "fail@example.com", []string{table, "email"}, nil)

		err := rule.Validate(ctx)
		if err == nil || err.Error() != "db error" {
			t.Errorf("expected db error, got: %v", err)
		}
	})

	t.Run("should fail if value does not exist", func(t *testing.T) {
		rule, _ := databaseRule.NewExistRule([]string{testTableName})
		table := testTableName
		database.RegisterPresenceVerifier(table, &mockExistPresenceVerifier{
			existsResult: false,
		})

		ctx := contract.NewValidationContext("email", "notfound@example.com", []string{table, "email"}, nil)

		if err := rule.Validate(ctx); err == nil {
			t.Error("expected error for non-existent value")
		}
	})

	t.Run("should pass if value exists", func(t *testing.T) {
		rule, _ := databaseRule.NewExistRule([]string{testTableName})
		table := testTableName
		database.RegisterPresenceVerifier(table, &mockExistPresenceVerifier{
			existsResult: true,
		})

		ctx := contract.NewValidationContext("email", "exists@example.com", []string{table, "email"}, nil)

		if err := rule.Validate(ctx); err != nil {
			t.Errorf("expected success, got error: %v", err)
		}
	})
}
