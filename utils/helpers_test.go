package utils

import (
	"mime/multipart"
	"testing"
)

func TestGetAsFloat(t *testing.T) {
	cases := []struct {
		in      interface{}
		out     float64
		wantErr bool
	}{
		{"hello", 5, false},
		{[]int{1, 2, 3}, 3, false},
		{map[string]int{"a": 1, "b": 2}, 2, false},
		{int64(42), 42, false},
		{uint(7), 7, false},
		{3.14, 3.14, false},
		{nil, 0, false},
		{struct{}{}, 0, true},
	}
	for _, c := range cases {
		got, err := GetAsFloat(c.in)
		if c.wantErr {
			if err == nil {
				t.Fatalf("expected error for input %T, got nil", c.in)
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != c.out {
			t.Fatalf("GetAsFloat(%v) = %v, want %v", c.in, got, c.out)
		}
	}
}

func TestGetAsNumeric(t *testing.T) {
	cases := []struct {
		in      interface{}
		out     float64
		wantErr bool
	}{
		{"123.5", 123.5, false},
		{"", 0, true},
		{"abc", 0, true},
		{int32(10), 10, false},
		{uint16(9), 9, false},
		{float32(2.5), 2.5, false},
		{nil, 0, true},
		{[]int{1}, 0, true},
	}
	for _, c := range cases {
		got, err := GetAsNumeric(c.in)
		if c.wantErr {
			if err == nil {
				t.Fatalf("expected error for input %T, got nil", c.in)
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != c.out {
			t.Fatalf("GetAsNumeric(%v) = %v, want %v", c.in, got, c.out)
		}
	}
}

func TestGetAsComparable(t *testing.T) {
	cases := []struct {
		in      interface{}
		out     float64
		wantErr bool
	}{
		{"99.9", 99.9, false}, // numeric string should parse as number, not length
		{"go", 2, false},      // non-numeric string -> rune count
		{[]byte{1, 2, 3, 4}, 4, false},
		{map[string]int{"x": 1}, 1, false},
		{12, 12, false},
		{nil, 0, true},
		{struct{}{}, 0, true},
	}
	for _, c := range cases {
		got, err := GetAsComparable(c.in)
		if c.wantErr {
			if err == nil {
				t.Fatalf("expected error for input %T, got nil", c.in)
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != c.out {
			t.Fatalf("GetAsComparable(%v) = %v, want %v", c.in, got, c.out)
		}
	}
}

func TestReplacePlaceholder(t *testing.T) {
	msg := "The :attribute must be between :param0 and :param1"
	replaced := ReplacePlaceholder(msg, 0, "1")
	replaced = ReplacePlaceholder(replaced, 1, "5")
	if replaced != "The :attribute must be between 1 and 5" {
		t.Fatalf("unexpected replacement: %s", replaced)
	}
}

func TestExtractDomainAndValidation(t *testing.T) {
	if ExtractDomain("no-at-symbol") != "" {
		t.Fatal("expected empty domain when no @ present")
	}
	if ExtractDomain("a@b") != "b" {
		t.Fatal("expected domain 'b'")
	}
	if !IsValidDomain("example.com") {
		t.Fatal("expected example.com to be valid")
	}
	if IsValidDomain(".example") || IsValidDomain("example.") || IsValidDomain("") || IsValidDomain("nodot") {
		t.Fatal("invalid domains evaluated as valid")
	}
}

func TestContainsDot(t *testing.T) {
	if ContainsDot("") {
		t.Fatal("empty should be false")
	}
	if ContainsDot(".start") || ContainsDot("end.") {
		t.Fatal("leading or trailing dot should be false")
	}
	if !ContainsDot("a.b") {
		t.Fatal("a.b should contain dot")
	}
}

func TestFloatToString(t *testing.T) {
	cases := map[float64]string{
		1.23456789: "1.234568", // rounded to 6 then trimmed
		1.230000:   "1.23",
		2.000000:   "2",
	}
	for in, want := range cases {
		got := FloatToString(in)
		if got != want {
			t.Fatalf("FloatToString(%v) = %q, want %q", in, got, want)
		}
	}
}

func TestNewFileHeader(t *testing.T) {
	fh := NewFileHeader("file.bin")
	if fh == nil || fh.Filename != "file.bin" || fh.Size == 0 {
		t.Fatalf("unexpected file header: %+v", fh)
	}
	ct := fh.Header.Get("Content-Type")
	if ct == "" {
		t.Fatal("expected content-type to be set")
	}
}

func TestNewFileHeaderWithMime(t *testing.T) {
	var _ *multipart.FileHeader // silence import not used warning
	fh := NewFileHeaderWithMime("img.png", "image/png", 2048)
	if fh.Filename != "img.png" || fh.Size != 2048 {
		t.Fatalf("unexpected file header: %+v", fh)
	}
	if fh.Header.Get("Content-Type") != "image/png" {
		t.Fatalf("unexpected mime: %s", fh.Header.Get("Content-Type"))
	}
}
