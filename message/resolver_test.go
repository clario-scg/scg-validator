package message

import (
	"strings"
	"testing"
)

func TestResolver_DefaultAndCustomMessages(t *testing.T) {
	r := NewResolver()

	// default message with params
	msg := r.Resolve("between", "age", []string{"18", "65"})
	if !strings.Contains(msg, "age") || !strings.Contains(msg, "18") || !strings.Contains(msg, "65") {
		t.Fatalf("unexpected default message: %q", msg)
	}

	// custom message overrides
	r.SetCustomMessage("between", ":attribute must be in [:param0, :param1]")
	msg = r.Resolve("between", "age", []string{"18", "65"})
	if msg != "age must be in [18, 65]" {
		t.Fatalf("unexpected custom message: %q", msg)
	}

	// field-specific custom message
	r.SetCustomMessage("required.email", "Email is required")
	msg = r.Resolve("required", "email", nil)
	if msg != "Email is required" {
		t.Fatalf("unexpected field custom message: %q", msg)
	}
}

func TestResolver_CustomAttributeAndFallback(t *testing.T) {
	r := NewResolver()
	r.SetCustomAttribute("username", "User Name")

	msg := r.Resolve("required", "username", nil)
	if !strings.Contains(msg, "User Name") {
		t.Fatalf("custom attribute not applied: %q", msg)
	}

	// unknown rule -> ultimate fallback
	msg = r.Resolve("no_such_rule", "x", []string{"p"})
	if !strings.Contains(msg, "x") || !strings.Contains(msg, "invalid") {
		t.Fatalf("unexpected fallback message: %q", msg)
	}
}

func TestResolver_CloneIsolation(t *testing.T) {
	r := NewResolver()
	// custom message that uses :attribute so attribute substitution is testable
	r.SetCustomMessage("alpha", ":attribute alpha")
	r.SetCustomAttribute("field", "Field")

	clone := r.Clone().(*Resolver)
	// modify clone independently
	clone.SetCustomMessage("alpha", ":attribute clone alpha")
	clone.SetCustomAttribute("field", "Clone Field")

	// original should remain intact for formatting
	origMsg := r.Resolve("alpha", "field", nil)
	clMsg := clone.Resolve("alpha", "field", nil)
	if origMsg == clMsg {
		t.Fatalf("expected clone to be isolated from original")
	}
	if !strings.Contains(origMsg, "Field alpha") || !strings.Contains(clMsg, "Clone Field clone alpha") {
		t.Fatalf("attributes not isolated: orig=%q clone=%q", origMsg, clMsg)
	}
}
