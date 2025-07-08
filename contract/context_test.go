package contract

import "testing"

func TestValidationContext(t *testing.T) {
	ctx := NewValidationContext("email", "a@b.com", []string{"p1"}, map[string]any{"x": 1})
	if ctx.Field() != "email" || ctx.Value().(string) != "a@b.com" {
		t.Fatalf("unexpected ctx fields: %s %v", ctx.Field(), ctx.Value())
	}
	if len(ctx.Parameters()) != 1 || ctx.Data()["x"].(int) != 1 {
		t.Fatal("parameters or data not set")
	}
	if ctx.Attribute("email") != "email" {
		t.Fatal("default attribute should be field name")
	}
	ctx.SetAttribute("email", "Email Address")
	if ctx.Attribute("email") != "Email Address" {
		t.Fatal("custom attribute not applied")
	}
}
