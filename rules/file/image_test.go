package file

import (
	"testing"

	"github.com/next-trace/scg-validator/utils"

	"github.com/next-trace/scg-validator/contract"
)

func TestImageRule(t *testing.T) {
	rule, err := NewImageRule()
	if err != nil {
		t.Fatalf("failed to create ImageRule: %v", err)
	}

	t.Run("valid image extensions", func(t *testing.T) {
		validFilenames := []string{"image.jpg", "pic.jpeg", "logo.png", "banner.gif", "scan.bmp", "art.webp"}
		for _, name := range validFilenames {
			t.Run(name, func(t *testing.T) {
				ctx := contract.NewValidationContext("image", utils.NewFileHeader(name), nil, nil)
				if err := rule.Validate(ctx); err != nil {
					t.Errorf("expected valid image for %q, got error: %v", name, err)
				}
			})
		}
	})

	t.Run("invalid file extensions", func(t *testing.T) {
		invalidFilenames := []string{"file.txt", "data.pdf", "sound.mp3", "image.svg"}
		for _, name := range invalidFilenames {
			t.Run(name, func(t *testing.T) {
				ctx := contract.NewValidationContext("image", utils.NewFileHeader(name), nil, nil)
				if err := rule.Validate(ctx); err == nil {
					t.Errorf("expected validation to fail for %q, but passed", name)
				}
			})
		}
	})

	t.Run("non-file input", func(t *testing.T) {
		ctx := contract.NewValidationContext("image", "not-a-file", nil, nil)
		if err := rule.Validate(ctx); err == nil {
			t.Error("expected error for non-file input, but got nil")
		}
	})

	t.Run("nil input", func(t *testing.T) {
		ctx := contract.NewValidationContext("image", nil, nil, nil)
		if err := rule.Validate(ctx); err == nil {
			t.Error("expected error for nil input, but got nil")
		}
	})
}
