package types

import (
	"strings"
)

type RequestParams interface {
	SetAuthnKey(k string)
	Body() (*strings.Reader, error)
}
