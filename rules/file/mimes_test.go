package file_test

import (
	"testing"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/file"
	"github.com/next-trace/scg-validator/utils"
)

func TestMimesRule(t *testing.T) {
	rule, err := file.NewMimesRule([]string{"pdf", "docx"})
	if err != nil {
		t.Fatalf("failed to create rule: %v", err)
	}

	tests := []struct {
		name      string
		value     interface{}
		wantValid bool
	}{
		// Valid file extensions
		{"valid pdf file", utils.NewFileHeader("document.pdf"), true},
		{"valid docx file", utils.NewFileHeader("report.docx"), true},

		// Invalid file extensions
		{"invalid jpg file", utils.NewFileHeader("image.jpg"), false},
		{"invalid zip file", utils.NewFileHeader("archive.zip"), false},
		{"invalid mp3 file", utils.NewFileHeader("music.mp3"), false},

		// Non-file inputs
		{"non-file string input", "not-a-file", false},
		{"nil input", nil, false},
		{"integer input", 123, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use real ValidationContext instead of mock
			ctx := contract.NewValidationContext("file_field", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.wantValid && err != nil {
				t.Errorf("expected valid, got error: %v", err)
			}
			if !tt.wantValid && err == nil {
				t.Errorf("expected error, got none")
			}
		})
	}
}
