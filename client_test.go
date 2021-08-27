package deepl_test

import (
	"fmt"
	"testing"

	"github.com/michimani/deepl-sdk-go"
	"github.com/michimani/deepl-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	cases := []struct {
		name     string
		authnKey string
		planName string
		wantErr  bool
		expect   *deepl.Client
	}{
		{
			name:     "ok (free)",
			authnKey: "test-key",
			planName: "free",
			wantErr:  false,
			expect: &deepl.Client{
				AuthenticationKey: "test-key",
				APIPlanType:       types.APIPlanFree,
				EndpointBase:      fmt.Sprintf("https://%s/%s/", types.DomainFree, types.APIVersion),
			},
		},
		{
			name:     "ok (pro)",
			authnKey: "test-key",
			planName: "pro",
			wantErr:  false,
			expect: &deepl.Client{
				AuthenticationKey: "test-key",
				APIPlanType:       types.APIPlanPro,
				EndpointBase:      fmt.Sprintf("https://%s/%s/", types.DomainPro, types.APIVersion),
			},
		},
		{
			name:     "DEEPL_API_AUTHN_KEY not set",
			authnKey: "",
			planName: "free",
			wantErr:  true,
			expect:   nil,
		},
		{
			name:     "DEEPL_API_PLAN not set",
			authnKey: "test-key",
			planName: "",
			wantErr:  true,
			expect:   nil,
		},
		{
			name:     "DEEPL_API_PLAN is invalid",
			authnKey: "test-key",
			planName: "freeeeeeee",
			wantErr:  true,
			expect:   nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			if c.authnKey != "" {
				tt.Setenv("DEEPL_API_AUTHN_KEY", c.authnKey)
			}
			if c.planName != "" {
				tt.Setenv("DEEPL_API_PLAN", c.planName)
			}

			client, err := deepl.NewClient()

			if c.wantErr {
				assert.Error(tt, err)
				return
			}

			assert.Equal(tt, c.expect, client)
		})
	}
}
