package types_test

import (
	"testing"

	"github.com/michimani/deepl-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAPIPlanValid(t *testing.T) {
	cases := []struct {
		name   string
		p      types.APIPlan
		expect bool
	}{
		{
			name:   "API plan free",
			p:      types.APIPlanFree,
			expect: true,
		},
		{
			name:   "API plan pro",
			p:      types.APIPlanPro,
			expect: true,
		},
		{
			name:   "invalid API plan",
			p:      types.APIPlan(0),
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := c.p.Valid()
			assert.Equal(tt, c.expect, b)
		})
	}
}

func TestAPIPlanToString(t *testing.T) {
	cases := []struct {
		name   string
		p      types.APIPlan
		expect string
	}{
		{
			name:   "API plan free",
			p:      types.APIPlanFree,
			expect: "free",
		},
		{
			name:   "API plan pro",
			p:      types.APIPlanPro,
			expect: "pro",
		},
		{
			name:   "invalid API plan",
			p:      types.APIPlan(0),
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			s := c.p.ToString()
			assert.Equal(tt, c.expect, s)
		})
	}
}

func TestToAPIPlan(t *testing.T) {
	cases := []struct {
		name     string
		planName string
		wantErr  bool
		expect   types.APIPlan
	}{
		{
			name:     "plan name free",
			planName: "free",
			wantErr:  false,
			expect:   types.APIPlanFree,
		},
		{
			name:     "plan name free (include upper case)",
			planName: "FreE",
			wantErr:  false,
			expect:   types.APIPlanFree,
		},
		{
			name:     "plan name pro",
			planName: "pro",
			wantErr:  false,
			expect:   types.APIPlanPro,
		},
		{
			name:     "plan name pro (include upper case)",
			planName: "PrO",
			wantErr:  false,
			expect:   types.APIPlanPro,
		},
		{
			name:     "invalid plan name",
			planName: "invalid plan",
			wantErr:  true,
			expect:   types.APIPlanNone,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			a, err := types.ToAPIPlan(c.planName)
			if c.wantErr {
				assert.Error(tt, err)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, a)
		})
	}
}
