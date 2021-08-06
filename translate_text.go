package deepl

import (
	"context"

	"github.com/michimani/deepl-sdk-go/params"
	"github.com/michimani/deepl-sdk-go/types"
)

// TranslateText calls the translate text API of the Deepl API.
func (c *Client) TranslateText(ctx context.Context, params *params.TranslateTextParams) (*types.TranslateTextResponse, *types.ErrorResponse, error) {
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
