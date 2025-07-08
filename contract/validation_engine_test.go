package contract_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/engine"
)

// This test ensures the concrete engine satisfies the ValidationEngine interface
// used by the facade. It's a compile-time assertion guarded by a trivial test.
func TestEngineImplementsValidationEngine(t *testing.T) {
	var _ contract.ValidationEngine = (*engine.Engine)(nil)
}
