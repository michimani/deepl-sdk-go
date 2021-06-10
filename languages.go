package deepl

import (
	"context"

	"github.com/michimani/deepl-sdk-go/types"
)

// Languages calls the languages API with type "target" of the Deepl API.
func (c *Client) TargetLanguages(ctx context.Context, params *types.LanguagesParams) (*types.TargetLanguagesResponse, *types.ErrorResponse, error) {
	res := types.TargetLanguagesResponse{}
	errRes, err := languages(ctx, c, params, &res)
	return &res, errRes, err
}

// Languages calls the languages API with type "source" of the Deepl API.
func (c *Client) SourceLanguages(ctx context.Context, params *types.LanguagesParams) (*types.SourceLanguagesResponse, *types.ErrorResponse, error) {
	res := types.SourceLanguagesResponse{}
	errRes, err := languages(ctx, c, params, &res)
	return &res, errRes, err
}

func languages(ctx context.Context, c *Client, params *types.LanguagesParams, res interface{}) (*types.ErrorResponse, error) {
	endpoint := c.EndpointBase + types.EndpointLanguages
	params.SetAuthnKey(c.AuthenticationKey)
	requester := NewRequester(endpoint, params)

	errRes, err := requester.Exec(res)
	if err != nil {
		return nil, err
	}
	if errRes != nil {
		return errRes, nil
	}

	return nil, nil
}
