package rules

import (
	"sort"
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

type dummyRule struct{}

func (d dummyRule) Name() string                          { return "dummy" }
func (d dummyRule) Validate(_ contract.RuleContext) error { return nil }

func TestRegistry_RegisterGetHasListClone(t *testing.T) {
	r := NewRegistry()
	if r.Has("x") {
		t.Fatal("expected empty registry")
	}

	creator := func(_ []string) (contract.Rule, error) { return dummyRule{}, nil }
	if err := r.Register("dummy", creator); err != nil {
		t.Fatalf("register error: %v", err)
	}

	if !r.Has("dummy") {
		t.Fatal("expected Has to be true")
	}

	got, ok := r.Get("dummy")
	if !ok || got == nil {
		t.Fatal("expected to get creator")
	}

	names := r.List()
	sort.Strings(names)
	if len(names) != 1 || names[0] != "dummy" {
		t.Fatalf("unexpected list: %#v", names)
	}

	if r.Count() != 1 {
		t.Fatalf("unexpected count: %d", r.Count())
	}

	clone := r.Clone().(*Registry)
	// mutate clone, ensure original unaffected
	_ = clone.Register("dummy2", creator)
	if r.Has("dummy2") {
		t.Fatal("original registry should not have new entry from clone")
	}
}
