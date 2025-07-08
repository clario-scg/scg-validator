package utils

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

func TestTranslateError_Nil(t *testing.T) {
	got := TranslateError(nil)
	if len(got) != 0 {
		t.Fatalf("expected empty map, got %v", got)
	}
}

func TestTranslateError_Required(t *testing.T) {
	ve := contract.NewValidationErrors()
	ve.AddError("name", "the :attribute field is required")

	got := TranslateError(ve)
	if got["name"] != "This field is required" {
		t.Fatalf("expected 'This field is required', got %q", got["name"])
	}
}

func TestTranslateError_GenericHumanize(t *testing.T) {
	ve := contract.NewValidationErrors()
	ve.AddError("accept", "the :attribute must be accepted when :other is :value")

	got := TranslateError(ve)
	expected := "This field must be accepted when :other is :value."
	if got["accept"] != expected {
		t.Fatalf("expected %q, got %q", expected, got["accept"])
	}
}

func TestTranslateError_TakesFirstErrorPerField(t *testing.T) {
	ve := contract.NewValidationErrors()
	ve.AddError("age", "the :attribute must be a number")
	ve.AddError("age", "the :attribute must be at least 18")

	got := TranslateError(ve)
	expected := "This field must be a number."
	if got["age"] != expected {
		t.Fatalf("expected %q, got %q", expected, got["age"])
	}
}
