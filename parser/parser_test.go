package parser

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseParameters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty strings",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single parameter",
			input:    "5",
			expected: []string{"5"},
		},
		{
			name:     "multiple parameters",
			input:    "5,10",
			expected: []string{"5", "10"},
		},
		{
			name:     "parameters with spaces",
			input:    "5, 10",
			expected: []string{"5", "10"},
		},
		{
			name:     "quoted parameter",
			input:    `"value with, comma"`,
			expected: []string{"value with, comma"},
		},
		{
			name:     "mixed parameters",
			input:    `5,"value with, comma",10`,
			expected: []string{"5", "value with, comma", "10"},
		},
		{
			name:     "escaped comma",
			input:    `value1\,part2,value2`,
			expected: []string{"value1,part2", "value2"},
		},
		{
			name:     "escaped quote",
			input:    `value1,\"quoted\",value2`,
			expected: []string{"value1", "\"quoted\"", "value2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseParameters(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseParameters(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSplitRules(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty strings",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single rule",
			input:    "required",
			expected: []string{"required"},
		},
		{
			name:     "multiple rules",
			input:    "required|alpha",
			expected: []string{"required", "alpha"},
		},
		{
			name:     "rules with spaces",
			input:    "required | alpha",
			expected: []string{"required", "alpha"},
		},
		{
			name:     "escaped pipe",
			input:    "regex:\\|test|required",
			expected: []string{"regex:|test", "required"},
		},
		{
			name:     "multiple escaped pipes",
			input:    "regex:\\|\\|test|required",
			expected: []string{"regex:||test", "required"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitRules(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SplitRules(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseRules(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []ParsedRule
	}{
		{
			name:     "empty strings",
			input:    "",
			expected: nil,
		},
		{
			name:  "single rule",
			input: "required",
			expected: []ParsedRule{
				{Name: "required", Params: nil},
			},
		},
		{
			name:  "multiple rules",
			input: "required|alpha",
			expected: []ParsedRule{
				{Name: "required", Params: nil},
				{Name: "alpha", Params: nil},
			},
		},
		{
			name:  "rule with parameter",
			input: "min:5",
			expected: []ParsedRule{
				{Name: "min", Params: []string{"5"}},
			},
		},
		{
			name:  "multiple rules with parameters",
			input: "required|min:5|max:10",
			expected: []ParsedRule{
				{Name: "required", Params: nil},
				{Name: "min", Params: []string{"5"}},
				{Name: "max", Params: []string{"10"}},
			},
		},
		{
			name:  "rule with multiple parameters",
			input: "between:5,10",
			expected: []ParsedRule{
				{Name: "between", Params: []string{"5", "10"}},
			},
		},
		{
			name:  "complex rule strings",
			input: "required|numeric|gt:18",
			expected: []ParsedRule{
				{Name: "required", Params: nil},
				{Name: "numeric", Params: nil},
				{Name: "gt", Params: []string{"18"}},
			},
		},
		{
			name:  "rule with escaped pipe",
			input: "regex:\\|test|required",
			expected: []ParsedRule{
				{Name: "regex", Params: []string{"|test"}},
				{Name: "required", Params: nil},
			},
		},
		{
			name:  "rule with quoted parameters",
			input: `in:"value1,value2",value3`,
			expected: []ParsedRule{
				{Name: "in", Params: []string{"value1,value2", "value3"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseRules(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseRules(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDebugParseRules(t *testing.T) {
	inputs := []string{
		"required|alpha",
		"required|min:5|max:10",
		"required|numeric|gt:18",
	}

	for _, input := range inputs {
		t.Logf("\nInput: %q", input)

		// Print each character's ASCII value to check for hidden characters
		t.Logf("ASCII values: ")
		for i := 0; i < len(input); i++ {
			t.Logf("%d ", input[i])
		}

		// Manual split to check if the pipe character is recognized
		manualParts := strings.Split(input, "|")
		t.Logf("Manual split result: %#v", manualParts)
		t.Logf("Manual split number of parts: %d", len(manualParts))

		// Debug SplitRules function
		parts := SplitRules(input)
		t.Logf("SplitRules result: %#v", parts)
		t.Logf("Number of parts: %d", len(parts))

		// Debug ParseRules function
		result := ParseRules(input)
		t.Logf("ParseRules result: %+v", result)

		// Print each parsed rule
		for i, rule := range result {
			t.Logf("Rule %d: Name=%q, Params=%v", i, rule.Name, rule.Params)
		}
	}
}
