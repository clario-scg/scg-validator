package utils

import (
	"strings"
	"unicode"

	"github.com/next-trace/scg-validator/contract"
)

// TranslateError converts internal validation errors into a simple, user-friendly
// map of field -> message. If err is not a validation error, it returns an empty map.
//
// Rules:
//   - If the rule corresponds to a required-type failure, the message becomes
//     "This field is required".
//   - Otherwise, we attempt to humanize the message by replacing ":attribute" with
//     "This field", capitalizing the sentence, and ensuring it ends with a period.
func TranslateError(err error) map[string]string {
	res := make(map[string]string)
	if err == nil {
		return res
	}

	// Prefer concrete type assertion
	if ve, ok := err.(*contract.ValidationErrors); ok {
		return translateFromMap(ve.Errors())
	}

	// Fallback: support any error that exposes Errors() map[string][]string
	type errorMap interface{ Errors() map[string][]string }
	if em, ok := err.(errorMap); ok {
		return translateFromMap(em.Errors())
	}

	return res
}

func translateFromMap(m map[string][]string) map[string]string {
	out := make(map[string]string, len(m))
	for field, msgs := range m {
		if len(msgs) == 0 {
			continue
		}
		msg := msgs[0]
		out[field] = humanizeMessage(msg)
	}
	return out
}

func humanizeMessage(msg string) string {
	s := strings.TrimSpace(msg)
	if s == "" {
		return s
	}

	low := strings.ToLower(s)
	// Handle any flavor of "required" as a simple, clear phrase.
	if strings.Contains(low, " required") || strings.HasPrefix(low, "required") {
		return "This field is required"
	}

	// Replace attribute placeholder with neutral wording.
	s = strings.ReplaceAll(s, ":attribute", "This field")
	// Common leading pattern from rule messages: "the :attribute ..."
	lowS := strings.ToLower(s)
	if strings.HasPrefix(lowS, "the this field ") {
		// Normalize to start with "This field "
		// len("the this field ") == 15
		s = "This field " + s[15:]
	}

	// Ensure sentence starts with uppercase
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
		s = string(runes)
	}

	// Ensure sentence ends with a period if it has no terminal punctuation
	if !strings.HasSuffix(s, ".") && !strings.HasSuffix(s, "!") && !strings.HasSuffix(s, "?") {
		s += "."
	}
	return s
}
