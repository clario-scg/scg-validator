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
	mimesRuleName       = "mimes"
	mimesRuleDefaultMsg = "the :attribute must be a file of type: :params"
	mimesRuleTypeError  = "the value must be a valid file"
)

type mimesRule struct {
	common.BaseRule
	allowedExts []string
}

func NewMimesRule(params []string) (contract.Rule, error) {
	if len(params) == 0 {
		return nil, errors.New("mimes rule requires at least one extension")
	}

	r := &mimesRule{
		allowedExts: params,
	}
	r.BaseRule = common.NewBaseRule(mimesRuleName, mimesRuleDefaultMsg, params)
	return r, nil
}

func (r *mimesRule) Validate(ctx contract.RuleContext) error {
	if r.ShouldSkipValidation(ctx.Value()) {
		return nil
	}

	file, ok := ctx.Value().(*multipart.FileHeader)
	if !ok {
		return errors.New(mimesRuleTypeError)
	}

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(file.Filename), "."))
	for _, allowed := range r.allowedExts {
		if ext == strings.ToLower(allowed) {
			return nil
		}
	}

	return errors.New(mimesRuleDefaultMsg)
}

func (r *mimesRule) Name() string {
	return mimesRuleName
}
