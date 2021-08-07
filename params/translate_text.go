package params

import (
	"errors"
	"net/url"
	"strings"

	"github.com/michimani/deepl-sdk-go/types"
)

// TranslateTextParams is parameters struct for TranslateText API
type TranslateTextParams struct {
	AuthKey            string
	Text               []string
	SourceLang         types.SourceLangCode
	TargetLang         types.TargetLangCode
	SplitSentences     types.SplitSentences
	PreserveFormatting types.PreserveFormatting
	Formality          types.Formality
}

func (p *TranslateTextParams) SetAuthnKey(k string) {
	p.AuthKey = k
}

func (p *TranslateTextParams) Body() (*strings.Reader, error) {
	if !p.IsValid() {
		return nil, errors.New("invalid params")
	}

	uv := url.Values{}

	uv.Add("auth_key", p.AuthKey)
	for _, t := range p.Text {
		uv.Add("text", t)
	}
	uv.Add("target_lang", string(p.TargetLang))

	if p.SourceLang != "" {
		uv.Add("source_lang", string(p.SourceLang))
	}

	if p.SplitSentences.Valid() {
		uv.Add("split_sentences", string(p.SplitSentences))
	}

	if p.PreserveFormatting.Valid() {
		uv.Add("preserve_formatting", string(p.PreserveFormatting))
	}

	if p.Formality.Valid(p.TargetLang) {
		uv.Add("formality", string(p.Formality))
	}

	return strings.NewReader(uv.Encode()), nil
}

func (p *TranslateTextParams) IsValid() bool {
	return true
}
