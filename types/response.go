package types

import (
	"encoding/json"
	"net/http"
)

type UsageResponse struct {
	CharacterCount int64 `json:"character_count"`
	CharacterLimit int64 `json:"character_limit"`
}

type TranslateTextResult struct {
	Text                   string   `json:"text"`
	DetectedSourceLanguage LangCode `json:"detected_source_language"`
}

type TranslateTextResponse struct {
	Translations []TranslateTextResult
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type RawResponse struct {
	StatusCode int
	Body       []byte
}

func (r *RawResponse) Unmarshal(i interface{}) (*ErrorResponse, error) {
	var err error
	if r.StatusCode != http.StatusOK {
		var errRes ErrorResponse
		err = json.Unmarshal(r.Body, &errRes)
		if err != nil {
			return nil, err
		}
		return &errRes, nil
	}

	err = json.Unmarshal(r.Body, i)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
