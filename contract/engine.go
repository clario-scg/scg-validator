package contract

import "context"

// This file is correctly named and domain-agnostic. No changes needed.

// Engine orchestrates the validator process
type Engine interface {
	// Execute performs validator using a data provider
	Execute(ctx context.Context, data DataProvider, rules map[string]string) Result
}
