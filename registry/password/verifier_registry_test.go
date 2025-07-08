package password

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

type fakePwdVerifier struct{ ok bool }

func (f fakePwdVerifier) Verify(_ string) (bool, error) { return f.ok, nil }

func TestPasswordVerifierRegistry(t *testing.T) {
	v := fakePwdVerifier{ok: true}
	RegisterPasswordVerifier("default", v)
	got, ok := FindPasswordVerifier("default")
	if !ok || got == nil {
		t.Fatal("expected registered verifier")
	}
	res, err := got.Verify("pass")
	if err != nil || !res {
		t.Fatalf("unexpected verify: %v %v", res, err)
	}
	// missing key
	if _, ok := FindPasswordVerifier("missing"); ok {
		t.Fatal("unexpected ok for missing key")
	}

	// interface compile-time check
	var _ contract.PasswordVerifier = v
}
