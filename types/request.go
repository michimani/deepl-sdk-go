package types

import (
	"errors"
	"net/url"
	"strings"
)

type TranslateTextParams struct {
	AuthKey            string
	Text               []string
	SourceLang         SourceLangCode
	TargetLang         TargetLangCode
	SplitSentences     interface{}
	PreserveFormatting interface{}
	Formality          string
}

func (r *TranslateTextParams) Body() (*strings.Reader, error) {
	if !r.IsValid() {
		return nil, errors.New("invalid params")
	}

	uv := url.Values{}

	uv.Add("auth_key", r.AuthKey)
	for _, t := range r.Text {
		uv.Add("text", t)
	}
	uv.Add("target_lang", string(r.TargetLang))
	if r.SourceLang != "" {
		uv.Add("target_lang", string(r.SourceLang))
	}

	return strings.NewReader(uv.Encode()), nil
}

func (r *TranslateTextParams) IsValid() bool {
	return true
}
