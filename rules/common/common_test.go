package common

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
)

type fakeCtx struct {
	field  string
	value  any
	params []string
	data   map[string]any
}

func (f fakeCtx) Field() string                 { return f.field }
func (f fakeCtx) Value() any                    { return f.value }
func (f fakeCtx) Parameters() []string          { return f.params }
func (f fakeCtx) Data() map[string]any          { return f.data }
func (f fakeCtx) Attribute(field string) string { return "attr:" + field }

func TestBaseRuleConfigAndSkip(t *testing.T) {
	r := NewBaseRule("required", "msg", []string{"p1", "p2"}, WithNullable(true), WithMessage("m"), WithStopOnFail(true))
	if r.Name() != "required" || r.GetMessage() != "m" {
		t.Fatalf("unexpected name/message: %s %s", r.Name(), r.GetMessage())
	}
	if r.IsNullable() != true {
		t.Fatal("expected nullable true")
	}
	if !r.ShouldSkipValidation(nil) {
		t.Fatal("expected skip when nil and nullable")
	}
	if r.ShouldSkipValidation(1) {
		t.Fatal("should not skip when value present")
	}
}

func TestSimpleRuleValidateAndMessage(t *testing.T) {
	base := NewBaseRule("min", "must be :param0", []string{"3"})
	called := 0
	r := &SimpleRule{BaseRule: base, validator: func(ctx contract.RuleContext) error {
		called++
		if ctx.Field() != "age" || ctx.Value().(int) != 2 || len(ctx.Parameters()) != 1 {
			t.Fatalf("unexpected ctx: %#v", ctx)
		}
		return nil
	}}

	msg := r.Message()
	if msg != "must be 3" {
		t.Fatalf("unexpected message: %q", msg)
	}

	_ = r.Validate(fakeCtx{field: "age", value: 2, params: []string{"3"}, data: map[string]any{"age": 2}})
	if called != 1 {
		t.Fatalf("validator not called, called=%d", called)
	}

	// nullable skip
	r.BaseRule = NewBaseRule("min", "", []string{"3"}, WithNullable(true))
	called = 0
	_ = r.Validate(fakeCtx{field: "age", value: nil})
	if called != 0 {
		t.Fatal("expected skip for nil value when nullable")
	}
}
