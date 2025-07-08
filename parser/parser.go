package parser

import (
	"strings"
)

// ParsedRule represents a parsed validator rule with its name and parameters
type ParsedRule struct {
	Name   string   // Rule name (e.g., "required", "min", "between")
	Params []string // Rule parameters (e.g., ["5"] for "min:5")
}

// ConditionalRule represents a conditional validator rule
type ConditionalRule struct {
	Field    string       // Field to check
	Operator string       // Operator (=, !=, >, <, etc.)
	Value    string       // Value to compare against
	Rules    []ParsedRule // Rules to apply if condition is met
}

// ParseRules parses a rule strings into a slice of ParsedRule
// Supports Laravel-style rule syntax including:
// - Simple rules: "required|email"
// - Rules with parameters: "min:5|max:10"
// - Rules with multiple parameters: "between:5,10"
func ParseRules(ruleString string) []ParsedRule {
	if ruleString == "" {
		return nil
	}

	var parsedRules []ParsedRule

	// Split by pipe character, but handle escaped pipes
	ruleComponents := SplitRules(ruleString)

	for _, component := range ruleComponents {
		parsedRule := ParsedRule{}

		// Split rule name and parameters
		nameAndParams := strings.SplitN(component, ":", 2)
		parsedRule.Name = strings.TrimSpace(nameAndParams[0])

		if len(nameAndParams) == 2 {
			// Handle parameters with commas inside quotes
			parsedRule.Params = parseParameters(nameAndParams[1])
		}

		parsedRules = append(parsedRules, parsedRule)
	}

	return parsedRules
}

// SplitRules splits a rule strings by pipe character, respecting escaped pipes
func SplitRules(ruleString string) []string {
	// Return empty slice for empty strings
	if ruleString == "" {
		return []string{}
	}

	var parts []string
	var currentPart strings.Builder

	for i := 0; i < len(ruleString); i++ {
		char := ruleString[i]

		// Handle escape character
		if char == '\\' && i+1 < len(ruleString) && ruleString[i+1] == '|' {
			currentPart.WriteByte('|')
			i++ // Skip the next character (the pipe)
			continue
		}

		// If we hit a pipe, add part to list and reset
		if char == '|' {
			parts = append(parts, strings.TrimSpace(currentPart.String()))
			currentPart.Reset()
			continue
		}

		// Otherwise, add character to current part
		currentPart.WriteByte(char)
	}

	// Add the last part
	if currentPart.Len() > 0 {
		parts = append(parts, strings.TrimSpace(currentPart.String()))
	}

	return parts
}

// parseParameters parses rule parameters, respecting quoted values and escaped characters
func parseParameters(paramString string) []string {
	// Return empty slice for empty strings
	if paramString == "" {
		return []string{}
	}

	var params []string
	var currentParam strings.Builder
	inQuotes := false
	escaped := false

	for i := 0; i < len(paramString); i++ {
		char := paramString[i]

		// Handle escape character
		if char == '\\' && !escaped {
			escaped = true
			continue
		}

		// Handle quotes
		if char == '"' && !escaped {
			inQuotes = !inQuotes
			continue
		}

		// If we're in quotes, add character to current parameter
		if inQuotes {
			currentParam.WriteByte(char)
			escaped = false
			continue
		}

		// If we hit a comma and not escaped, add parameter to list and reset
		if char == ',' && !escaped {
			params = append(params, strings.TrimSpace(currentParam.String()))
			currentParam.Reset()
			escaped = false
			continue
		}

		// Otherwise, add character to current parameter
		currentParam.WriteByte(char)
		escaped = false
	}

	// Add the last parameter
	if currentParam.Len() > 0 {
		params = append(params, strings.TrimSpace(currentParam.String()))
	}

	return params
}
