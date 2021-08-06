package params

import (
	"net/url"
	"strings"

	"github.com/michimani/deepl-sdk-go/types"
)

// LanguagesParams is parameters struct for Languages API.
type LanguagesParams struct {
	AuthKey  string
	LangType types.LangType
}

func (p *LanguagesParams) SetAuthnKey(k string) {
	p.AuthKey = k
}

func (p *LanguagesParams) Body() (*strings.Reader, error) {
	uv := url.Values{}

	uv.Add("auth_key", p.AuthKey)
	uv.Add("type", string(p.LangType))
	return strings.NewReader(uv.Encode()), nil
}
