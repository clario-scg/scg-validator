// Package engine contains the core validation execution engine.
//
// It parses rule expressions, evaluates them against the provided data, and
// produces a contract.Result. The engine is wired to registries, messages, and
// parsers and can be swapped through the contract.ValidationEngine interface.
package engine
