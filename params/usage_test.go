package params_test

import (
	"io"
	"testing"

	"github.com/michimani/deepl-sdk-go/params"
	"github.com/stretchr/testify/assert"
)

func TestUsageSetAuthnKey(t *testing.T) {
	k := "test-authn-key"
	p := params.UsageParams{}
	p.SetAuthnKey(k)

	assert.Equal(t, k, p.AuthKey)
}

func TestUsageBody(t *testing.T) {
	cases := []struct {
		name   string
		params params.UsageParams
		expect string
	}{
		{
			name:   "normal",
			params: params.UsageParams{AuthKey: "test-authn-key"},
			expect: "auth_key=test-authn-key",
		},
		{
			name:   "normal: with white space",
			params: params.UsageParams{AuthKey: "test key"},
			expect: "auth_key=test+key",
		},
		{
			name:   "normal: empty",
			params: params.UsageParams{},
			expect: "auth_key=",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()

			assert.NoError(tt, err)

			s := []byte{}
			for {
				b, err := r.ReadByte()
				if err == io.EOF {
					break
				}
				s = append(s, b)
			}

			assert.Equal(tt, c.expect, string(s))
		})
	}
}
