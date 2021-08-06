package params_test

import (
	"io"
	"testing"

	"github.com/michimani/deepl-sdk-go/params"
	"github.com/michimani/deepl-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestLanguagesSetAuthnKey(t *testing.T) {
	k := "test-authn-key"
	p := params.LanguagesParams{}

	p.SetAuthnKey(k)

	assert.Equal(t, k, p.AuthKey)
}

func TestLanguagesBody(t *testing.T) {
	cases := []struct {
		name   string
		params params.LanguagesParams
		expect string
	}{
		{
			name:   "normal: langType=source",
			params: params.LanguagesParams{AuthKey: "test-authn-key", LangType: types.LangTypeSource},
			expect: "auth_key=test-authn-key&type=source",
		},
		{
			name:   "normal: langType=target",
			params: params.LanguagesParams{AuthKey: "test-authn-key", LangType: types.LangTypeTarget},
			expect: "auth_key=test-authn-key&type=target",
		},
		{
			name:   "normal: with white space",
			params: params.LanguagesParams{AuthKey: "test key", LangType: types.LangTypeSource},
			expect: "auth_key=test+key&type=source",
		},
		{
			name:   "normal: empty",
			params: params.LanguagesParams{},
			expect: "auth_key=&type=",
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
