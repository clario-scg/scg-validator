package single

import (
	"errors"
	"testing"
)

func TestValidationError(t *testing.T) {
	err := NewValidationError("field %s is invalid", "name")
	if err == nil {
		t.Fatal("expected error")
	}
	if !IsValidationError(err) {
		t.Fatal("expected ValidationError type")
	}
	if err.Error() != "field name is invalid" {
		t.Fatalf("unexpected message: %s", err.Error())
	}
}

func TestIsValidationFailed(t *testing.T) {
	if !IsValidationFailed(ErrValidationFailed) {
		t.Fatal("expected ErrValidationFailed to be detected")
	}
	if IsValidationFailed(errors.New("other")) {
		t.Fatal("unexpected detection for other error")
	}
}
