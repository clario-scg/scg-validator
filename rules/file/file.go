// Package file provides validation rules for file inputs including MIME type checks.
package file

import (
	"errors"
	"mime/multipart"

	"github.com/next-trace/scg-validator/contract"
	"github.com/next-trace/scg-validator/rules/common"
)

const (
	fileRuleName              = "file"
	fileRuleDefaultMsg        = "the value must be a file"
	mimeTypeRuleName          = "mimetype"
	mimeTypeRuleDefaultMsg    = "the file must be of the specified MIME type"
	mimeTypeRuleMissingParams = "mimetype rule requires at least one parameter"
)

// Rule checks whether the input is a multipart file upload.
type Rule struct {
	common.BaseRule
}

// NewFileRule constructs a Rule instance.
func NewFileRule() (contract.Rule, error) {
	return &Rule{
		BaseRule: common.NewBaseRule(fileRuleName, fileRuleDefaultMsg, nil),
	}, nil
}

// Validate checks that the value is a *multipart.FileHeader.
func (r *Rule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	if _, ok := ctx.Value().(*multipart.FileHeader); !ok {
		return errors.New(fileRuleDefaultMsg)
	}
	return nil
}
func (r *Rule) Name() string {
	return fileRuleName
}
