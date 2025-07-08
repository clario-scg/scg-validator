package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

// GetAsFloat converts various types to a float64 for size comparison.
// For strings, it returns the rune count. For slices, arrays, and maps, it returns the length.
// For numeric types, it returns the float64 value.
func GetAsFloat(value interface{}) (float64, error) {
	if value == nil {
		return 0, nil
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		return float64(utf8.RuneCountInString(val.String())), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return float64(val.Len()), nil
	}

	return 0, fmt.Errorf("unsupported type for comparison: %T", value)
}

// GetAsNumeric converts various types to a float64 for numeric comparison.
// For strings, it attempts to parse them as numbers. For numeric types, it returns the float64 value.
// This is different from GetAsFloat which returns string length for strings.
func GetAsNumeric(value interface{}) (float64, error) {
	if value == nil {
		return 0, errors.New("cannot convert nil to numeric value")
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		str := val.String()
		if str == "" {
			return 0, errors.New("cannot convert empty string to numeric value")
		}
		parsed, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot convert string '%s' to numeric value", str)
		}
		return parsed, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	}

	return 0, fmt.Errorf("unsupported type for numeric comparison: %T", value)
}

// GetAsComparable converts various types to a float64 for comparison rules.
// It tries to parse strings as numbers first, but falls back to length if not numeric.
// For collections (slices, maps, arrays), it returns the length.
// For numeric types, it returns the numeric value.
func GetAsComparable(value interface{}) (float64, error) {
	if value == nil {
		return 0, errors.New("cannot convert nil to comparable value")
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		str := val.String()
		// Try to parse as number first
		if parsed, err := strconv.ParseFloat(str, 64); err == nil {
			return parsed, nil
		}
		// Fall back to string length (rune count)
		return float64(utf8.RuneCountInString(str)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return float64(val.Len()), nil
	}

	return 0, fmt.Errorf("unsupported type for comparison: %T", value)
}

func ReplacePlaceholder(msg string, i int, param string) string {
	placeholder := ":param" + strconv.Itoa(i)
	return strings.ReplaceAll(msg, placeholder, param)
}

func ExtractDomain(email string) string {
	at := strings.LastIndex(email, "@")
	if at == -1 || at+1 >= len(email) {
		return ""
	}
	return email[at+1:]
}

func IsValidDomain(domain string) bool {
	return domain != "" &&
		!strings.HasPrefix(domain, ".") &&
		!strings.HasSuffix(domain, ".") &&
		strings.Contains(domain, ".")
}

// ContainsDot checks whether the given domain part contains at least one dot,
// but not as the first or last character (e.g., ".example" and "example." are invalid).
func ContainsDot(domain string) bool {
	if len(domain) == 0 {
		return false
	}
	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return false
	}
	return strings.Contains(domain, ".")
}

func FloatToString(f float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.6f", f), "0"), ".")
}

func NewFileHeader(filename string) *multipart.FileHeader {
	return &multipart.FileHeader{
		Filename: filename,
		Header:   textproto.MIMEHeader{"Content-Type": []string{"application/octet-stream"}},
		Size:     1024,
	}
}

func NewFileHeaderWithMime(filename string, mimeType string, size int64) *multipart.FileHeader {
	return &multipart.FileHeader{
		Filename: filename,
		Size:     size,
		Header:   textproto.MIMEHeader{"Content-Type": []string{mimeType}},
	}
}
