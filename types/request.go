package types

import (
	"errors"
	"net/url"
	"strings"
)

type RequestParams interface {
	SetAuthnKey(k string)
	Body() (*strings.Reader, error)
}

// UsageParams is parameters struct for Usage API.
type UsageParams struct {
	AuthKey string
}

func (p *UsageParams) SetAuthnKey(k string) {
	p.AuthKey = k
}

func (p *UsageParams) Body() (*strings.Reader, error) {
	uv := url.Values{}

	uv.Add("auth_key", p.AuthKey)
	return strings.NewReader(uv.Encode()), nil
}

// TranslateTextParams is parameters struct for TranslateText API
type TranslateTextParams struct {
	AuthKey            string
	Text               []string
	SourceLang         SourceLangCode
	TargetLang         TargetLangCode
	SplitSentences     interface{}
	PreserveFormatting interface{}
	Formality          string
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
		uv.Add("target_lang", string(p.SourceLang))
	}

	return strings.NewReader(uv.Encode()), nil
}

func (p *TranslateTextParams) IsValid() bool {
	return true
}
