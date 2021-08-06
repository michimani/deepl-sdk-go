package params

import (
	"net/url"
	"strings"
)

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
