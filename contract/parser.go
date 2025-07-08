package contract

// This file is correctly named and domain-agnostic. No changes needed.

// Parser handles parsing of rule strings into structured rules
type Parser interface {
	// Parse converts a rule strings into a slice of ParsedRule structs
	Parse(ruleString string) []ParsedRule
}

// ParsedRule represents a parsed validator rule with its parameters
type ParsedRule struct {
	Name       string   `json:"name"`       // Rule name (e.g., "min", "required")
	Parameters []string `json:"parameters"` // Rule parameters (e.g., ["5"] for min:5)
}
