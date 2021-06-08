package types

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal_Usage(t *testing.T) {
	cases := []struct {
		name       string
		statusCode int
		body       []byte
		wantError  bool
		expect     UsageResponse
	}{
		{
			name:       "usage response",
			statusCode: http.StatusOK,
			body:       []byte(`{"character_count":100,"character_limit":99999}`),
			wantError:  false,
			expect: UsageResponse{
				CharacterCount: 100,
				CharacterLimit: 99999,
			},
		},
		{
			name:       "error (invalid body)",
			statusCode: 0,
			body:       []byte(`not json string`),
			wantError:  true,
			expect:     UsageResponse{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			raw := RawResponse{
				StatusCode: c.statusCode,
				Body:       c.body,
			}

			var res UsageResponse
			errRes, err := raw.Unmarshal(&res)

			if c.wantError {
				assert.Error(t, err)
				return
			}

			if c.statusCode != http.StatusOK {
				assert.Equal(t, c.expect, errRes)
				return
			}

			assert.Equal(t, c.expect, res)
		})
	}
}

func TestUnmarshal_TranslateText(t *testing.T) {
	cases := []struct {
		name       string
		statusCode int
		body       []byte
		wantError  bool
		expect     TranslateTextResponse
	}{
		{
			name:       "translate_text response",
			statusCode: http.StatusOK,
			body:       []byte(`{"translations":[{"detected_source_language":"JA","text":"hello"}]}`),
			wantError:  false,
			expect: TranslateTextResponse{
				Translations: []TranslateTextResult{
					{
						Text:                   "hello",
						DetectedSourceLanguage: "JA",
					},
				},
			},
		},
		{
			name:       "error (invalid body)",
			statusCode: 0,
			body:       []byte(`not json string`),
			wantError:  true,
			expect:     TranslateTextResponse{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			raw := RawResponse{
				StatusCode: c.statusCode,
				Body:       c.body,
			}

			var res TranslateTextResponse
			errRes, err := raw.Unmarshal(&res)

			if c.wantError {
				assert.Error(t, err)
				return
			}

			if c.statusCode != http.StatusOK {
				assert.Equal(t, c.expect, errRes)
				return
			}

			assert.Equal(t, c.expect, res)
		})
	}
}

func TestUnmarshal_Error(t *testing.T) {
	cases := []struct {
		name       string
		statusCode int
		body       []byte
		wantError  bool
		expect     *ErrorResponse
	}{
		{
			name:       "error response",
			statusCode: http.StatusInternalServerError,
			body:       []byte(`{"message":"Error message."}`),
			wantError:  false,
			expect: &ErrorResponse{
				Message: "Error message.",
			},
		},
		{
			name:       "error (invalid body)",
			statusCode: 0,
			body:       []byte(`not json string`),
			wantError:  true,
			expect:     nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			raw := RawResponse{
				StatusCode: c.statusCode,
				Body:       c.body,
			}

			var res UsageResponse
			errRes, err := raw.Unmarshal(&res)

			if c.wantError {
				assert.Error(t, err)
				return
			}

			if c.statusCode != http.StatusOK {
				assert.Equal(t, c.expect, errRes)
				return
			}

			assert.Equal(t, c.expect, res)
		})
	}
}
