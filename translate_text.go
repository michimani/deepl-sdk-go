package deepl

import (
	"context"
	"deepl-sdk-go/types"
	"encoding/json"
)

// TranslateText calls the translate text API of the Deepl API.
func (c *Client) TranslateText(ctx context.Context, params *types.TranslateTextParams) (*types.TranslateTextResponse, error) {
	var res types.TranslateTextResponse

	body, err := params.Body()
	if err != nil {
		return nil, err
	}

	endpoint := "https://api.deepl.com/v2/translate"
	bytes, err := POST(endpoint, body)
	if err != nil {
		return nil, err
	}

	jerr := json.Unmarshal(bytes, &res)
	if jerr != nil {
		return nil, jerr
	}

	return &res, nil
}
