package contract

import "testing"

func TestValidationErrorsBasic(t *testing.T) {
	ve := NewValidationErrors()
	if !ve.IsValid() || ve.FirstError() != "" {
		t.Fatal("expected empty errors to be valid and no first error")
	}

	ve.AddError("name", "is required")
	ve.AddError("name", "too short")
	ve.AddError("age", "invalid")

	if ve.IsValid() {
		t.Fatal("expected not valid after adding errors")
	}
	if first := ve.FirstError(); first == "" {
		t.Fatal("expected some first error")
	}
	if !ve.HasFieldError("name") || ve.FieldError("name") != "is required" {
		t.Fatalf("unexpected field error: %q", ve.FieldError("name"))
	}
	if ve.HasFieldError("missing") || ve.FieldError("missing") != "" {
		t.Fatal("missing field should not have errors")
	}
	if _, ok := ve.Errors()["age"]; !ok {
		t.Fatal("expected age to be present in map")
	}
}
