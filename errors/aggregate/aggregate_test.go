package aggregate

import (
	"strings"
	"testing"
)

func TestErrorsAddGetFirstHas(t *testing.T) {
	errs := Errors{}
	if errs.Has("name") {
		t.Fatal("expected empty map to not have key")
	}
	errs.Add("name", "is required")
	errs.Add("name", "too short")
	errs.Add("age", "invalid")

	if !errs.Has("name") || !errs.Has("age") {
		t.Fatal("expected keys to exist after Add")
	}

	got := errs.Get("name")
	if len(got) != 2 || got[0] != "is required" || got[1] != "too short" {
		t.Fatalf("unexpected get: %#v", got)
	}
	if first := errs.First("name"); first != "is required" {
		t.Fatalf("unexpected first: %q", first)
	}
	if first := errs.First("missing"); first != "" {
		t.Fatalf("expected empty for missing, got %q", first)
	}
}

func TestErrorsErrorString(t *testing.T) {
	errs := Errors{}
	errs.Add("email", "invalid format")
	errs.Add("email", "already taken")
	msg := errs.Error()
	if !strings.Contains(msg, "validator failed with the following errors:") ||
		!strings.Contains(msg, "field 'email'") || !strings.Contains(msg, "already taken") {
		t.Fatalf("unexpected error string: %q", msg)
	}
}
