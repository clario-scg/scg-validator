package file

import (
	"testing"

	"github.com/next-trace/scg-validator/utils"

	"github.com/next-trace/scg-validator/contract"
)

func TestFileRule(t *testing.T) {
	rule, err := NewFileRule()
	if err != nil {
		t.Fatalf("failed to create FileRule: %v", err)
	}

	tests := []struct {
		name    string
		value   any
		wantErr bool
	}{
		{"valid file input", utils.NewFileHeader("file.pdf"), false},
		{"invalid string input", "not a file", true},
		{"nil input", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := contract.NewValidationContext("file", tt.value, nil, nil)
			err := rule.Validate(ctx)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil for value: %v", tt.value)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}

func TestMimeTypeRule(t *testing.T) {
	t.Run("should fail with empty parameters", func(t *testing.T) {
		_, err := NewMimesRule([]string{})
		if err == nil {
			t.Error("expected error for missing parameters")
		}
	})

	tests := []struct {
		name    string
		value   any
		allowed []string
		wantErr bool
	}{
		{"valid extension", utils.NewFileHeaderWithMime("image.png", "image/png", 1024), []string{"png"}, false},
		{"invalid extension", utils.NewFileHeaderWithMime("doc.txt", "text/plain", 1024), []string{"png"}, true},
		{"non-file input", "not a file", []string{"png"}, true},
		{"nil input", nil, []string{"png"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule, err := NewMimesRule(tt.allowed)
			if err != nil {
				t.Fatalf("unexpected error constructing rule: %v", err)
			}

			ctx := contract.NewValidationContext("file", tt.value, nil, nil)
			err = rule.Validate(ctx)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil for value: %v", tt.value)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}
