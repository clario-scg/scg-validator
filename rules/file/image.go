package file

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	imageRuleName           = "image"
	imageRuleDefaultMsg     = "the file must be an image (jpg, jpeg, png, gif, bmp, webp)"
	imageRuleInvalidTypeMsg = "the value must be a file"
)

var (
	imageAllowedExtensions = []string{"jpg", "jpeg", "png", "gif", "bmp", "webp"}
)

// ImageRule checks if a file is an image by its extension.
type ImageRule struct {
	common.BaseRule
}

// NewImageRule creates a new instance of ImageRule.
func NewImageRule() (contract.Rule, error) {
	return &ImageRule{
		BaseRule: common.NewBaseRule(imageRuleName, imageRuleDefaultMsg, nil),
	}, nil
}

// Validate returns an error if the file does not have a valid image extension.
func (r *ImageRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	fh, ok := ctx.Value().(*multipart.FileHeader)
	if !ok {
		return errors.New(imageRuleInvalidTypeMsg)
	}

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(fh.Filename), "."))
	for _, allowed := range imageAllowedExtensions {
		if ext == allowed {
			return nil
		}
	}

	return errors.New(imageRuleDefaultMsg)
}
func (r *ImageRule) Name() string {
	return imageRuleName
}
