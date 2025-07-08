package contract

import (
	stdErrors "errors"
	"testing"

	validatorErrors "github.com/next-trace/scg-validator/errors"
)

func TestIsValidationFailedAndNewValidationError(t *testing.T) {
	if !IsValidationFailed(validatorErrors.ErrValidationFailed) {
		t.Fatal("expected IsValidationFailed to detect re-exported error")
	}
	if IsValidationFailed(stdErrors.New("other")) {
		t.Fatal("should not detect other errors as validation failed")
	}

	err := NewValidationError("field %s failed", "x")
	if err.Error() != "field x failed" {
		t.Fatalf("unexpected formatted error: %v", err)
	}
}
