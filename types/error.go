package types

import "net/http"

type DeeplAPIErrorCode int

const (
	DeeplAPIError400 DeeplAPIErrorCode = http.StatusBadRequest
	DeeplAPIError403 DeeplAPIErrorCode = http.StatusForbidden
	DeeplAPIError404 DeeplAPIErrorCode = http.StatusNotFound
	DeeplAPIError413 DeeplAPIErrorCode = http.StatusRequestEntityTooLarge
	DeeplAPIError414 DeeplAPIErrorCode = http.StatusRequestURITooLong
	DeeplAPIError429 DeeplAPIErrorCode = http.StatusTooManyRequests
	DeeplAPIError456 DeeplAPIErrorCode = 456
	DeeplAPIError503 DeeplAPIErrorCode = http.StatusServiceUnavailable
	DeeplAPIError529 DeeplAPIErrorCode = 529
	DeeplAPIError500 DeeplAPIErrorCode = http.StatusInternalServerError
)

var errorCodeMap map[DeeplAPIErrorCode]string = map[DeeplAPIErrorCode]string{
	DeeplAPIError400: "Bad request. Please check error message and your parameters.",
	DeeplAPIError403: "Authorization failed. Please supply a valid auth_key parameter.",
	DeeplAPIError404: "The requested resource could not be found.",
	DeeplAPIError413: "The request size exceeds the limit.",
	DeeplAPIError414: "The request URL is too long. You can avoid this error by using a POST request instead of a GET request.",
	DeeplAPIError429: "Too many requests. Please wait and resend your request.",
	DeeplAPIError456: "Quota exceeded. The character limit has been reached.",
	DeeplAPIError503: "Resource currently unavailable. Try again later.",
	DeeplAPIError529: "Too many requests. Please wait and resend your request.",
	DeeplAPIError500: "Internal error",
}

func (e DeeplAPIErrorCode) Description() string {
	if desc, ok := errorCodeMap[e]; ok {
		return desc
	}
	return errorCodeMap[DeeplAPIError500]
}
