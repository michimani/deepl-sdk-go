package deepl

import (
	"deepl-sdk-go/types"
	"fmt"
)

type Client struct {
	AuthenticationKey string
	IsPro             bool
	EndpointBase      string
}

// NewClient returns new DeepL API client with an authentication key.
func NewClient(authnKey string, isPro bool) *Client {
	var c Client
	c.AuthenticationKey = authnKey
	c.IsPro = isPro
	if isPro {
		c.EndpointBase = fmt.Sprintf("https://%s/%s/", types.DomainPro, types.APIVersion)
	} else {
		c.EndpointBase = fmt.Sprintf("https://%s/%s/", types.DomainFree, types.APIVersion)
	}

	return &c
}
