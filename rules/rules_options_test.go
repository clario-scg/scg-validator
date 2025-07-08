package rules

import "testing"

func TestWithCustomMessage_Option(t *testing.T) {
	// Just ensure option function executes without panic to cover code path
	t.Helper()
	_ = NewRuleRegistry(WithCustomMessage("required", "custom"))
}
