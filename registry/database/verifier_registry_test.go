package database

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

type fakePresence struct{ exists, unique bool }

func (f fakePresence) Exists(_, _ string, _ any) (bool, error) { return f.exists, nil }
func (f fakePresence) Unique(_, _ string, _ any) (bool, error) { return f.unique, nil }

func TestPresenceVerifierRegistry(t *testing.T) {
	v := fakePresence{exists: true}
	RegisterPresenceVerifier("users", v)
	got, ok := FindPresenceVerifier("users")
	if !ok || got == nil {
		t.Fatal("expected registered verifier")
	}
	e, err := got.Exists("users", "email", "a@b.com")
	if err != nil || !e {
		t.Fatalf("unexpected exists: %v %v", e, err)
	}
	// missing table
	if _, ok := FindPresenceVerifier("missing"); ok {
		t.Fatal("unexpected ok for missing table")
	}
	// interface check
	var _ contract.PresenceVerifier = v
}
