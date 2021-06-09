package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescription(t *testing.T) {
	cases := []struct {
		name      string
		errorCode DeeplAPIErrorCode
		expect    string
	}{
		{
			name:      "400",
			errorCode: DeeplAPIError400,
			expect:    "Bad request. Please check error message and your parameters.",
		},
		{
			name:      "403",
			errorCode: DeeplAPIError403,
			expect:    "Authorization failed. Please supply a valid auth_key parameter.",
		},
		{
			name:      "404",
			errorCode: DeeplAPIError404,
			expect:    "The requested resource could not be found.",
		},
		{
			name:      "413",
			errorCode: DeeplAPIError413,
			expect:    "The request size exceeds the limit.",
		},
		{
			name:      "414",
			errorCode: DeeplAPIError414,
			expect:    "The request URL is too long. You can avoid this error by using a POST request instead of a GET request.",
		},
		{
			name:      "429",
			errorCode: DeeplAPIError429,
			expect:    "Too many requests. Please wait and resend your request.",
		},
		{
			name:      "456",
			errorCode: DeeplAPIError456,
			expect:    "Quota exceeded. The character limit has been reached.",
		},
		{
			name:      "503",
			errorCode: DeeplAPIError503,
			expect:    "Resource currently unavailable. Try again later.",
		},
		{
			name:      "529",
			errorCode: DeeplAPIError529,
			expect:    "Too many requests. Please wait and resend your request.",
		},
		{
			name:      "500",
			errorCode: DeeplAPIError500,
			expect:    "Internal error",
		},
		{
			name:      "undefined code",
			errorCode: 0,
			expect:    "Internal error",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			desc := c.errorCode.Description()
			assert.Equal(t, c.expect, desc)
		})
	}
}
