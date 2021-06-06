package deepl

import (
	"context"
	"deepl-sdk-go/types"
)

// TranslateText calls the translate text API of the Deepl API.
func (c *Client) TranslateText(ctx context.Context, params *types.TranslateTextParams) (*types.TranslateTextResponse, *types.ErrorResponse, error) {
	var res types.TranslateTextResponse

	endpoint := c.EndpointBase + types.EndpointTranslateText
	params.SetAuthnKey(c.AuthenticationKey)
	requester := NewRequester(endpoint, params)

	errRes, err := requester.Exec(&res)
	if err != nil {
		return nil, nil, err
	}
	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
