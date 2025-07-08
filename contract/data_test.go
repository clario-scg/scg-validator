package contract

import "testing"

func TestSimpleDataProvider(t *testing.T) {
	data := map[string]any{"name": "Alice", "age": 30}
	dp := NewSimpleDataProvider(data)

	if !dp.Has("name") || dp.Has("missing") {
		t.Fatal("Has reported incorrectly")
	}
	if v, ok := dp.Get("age"); !ok || v.(int) != 30 {
		t.Fatalf("unexpected Get: %v %v", v, ok)
	}
	if all := dp.All(); len(all) != 2 || all["name"].(string) != "Alice" {
		t.Fatalf("unexpected All: %#v", all)
	}
}
