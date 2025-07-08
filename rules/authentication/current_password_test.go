package authentication

import (
	"errors"
	"testing"

	"github.com/next-trace/scg-validator/registry/password"

	"github.com/next-trace/scg-validator/contract"
)

// Mock PasswordVerifier for testing purposes
type mockPasswordVerifier struct {
	correctPassword string
	shouldFail      bool
}

func (m *mockPasswordVerifier) Verify(password string) (bool, error) {
	if m.shouldFail {
		return false, errors.New("simulated verifier error")
	}
	return password == m.correctPassword, nil
}

// TestCurrentPasswordRule using a registry-based PasswordVerifier
func TestCurrentPasswordRule_RegistryBased(t *testing.T) {
	rule, _ := NewCurrentPasswordRule()

	// Register a mock verifier with a correct password
	verifier := &mockPasswordVerifier{correctPassword: "secret123"}
	password.RegisterPasswordVerifier("default", verifier)

	tests := []struct {
		name       string
		value      any
		shouldPass bool
	}{
		{"correct password", "secret123", true},
		{"wrong password", "wrong", false},
		{"non-string type", 12345, false},
		{"empty string", "", false},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use ValidationContext directly
			ctx := contract.NewValidationContext("field", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if (err == nil) != tt.shouldPass {
				t.Errorf("Validate(%v): got error = %v, want pass = %v", tt.value, err, tt.shouldPass)
			}
		})
	}
}

// TestCurrentPasswordRule when the verifier fails (e.g., DB error, etc.)
func TestCurrentPasswordRule_VerifierFails(t *testing.T) {
	rule, _ := NewCurrentPasswordRule()

	// Register a verifier that simulates a failure
	password.RegisterPasswordVerifier("default", &mockPasswordVerifier{shouldFail: true})

	// Use ValidationContext directly
	ctx := contract.NewValidationContext("field", "anything", nil, nil)
	err := rule.Validate(ctx)

	// Expect an error due to verifier failure
	if err == nil {
		t.Error("Expected verifier failure error, but got nil")
	}
}

// TestCurrentPasswordRule when no PasswordVerifier is registered
func TestCurrentPasswordRule_NoVerifierRegistered(t *testing.T) {
	// Clear registry by setting the PasswordVerifier to nil (simulate missing)
	password.RegisterPasswordVerifier("default", nil)

	rule, _ := NewCurrentPasswordRule()

	// Use ValidationContext directly
	ctx := contract.NewValidationContext("field", "something", nil, nil)
	err := rule.Validate(ctx)

	// Expect an error due to no verifier being available
	if err == nil {
		t.Error("Expected error due to missing verifier, got nil")
	}
}
