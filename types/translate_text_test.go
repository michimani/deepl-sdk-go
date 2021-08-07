package types_test

import (
	"testing"

	"github.com/michimani/deepl-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSplitSentencesValid(t *testing.T) {
	cases := []struct {
		name  string
		s     types.SplitSentences
		exect bool
	}{
		{
			name:  "SplitSentencesNoSplit",
			s:     types.SplitSentencesNoSplit,
			exect: true,
		},
		{
			name:  "SplitSentencesSplit",
			s:     types.SplitSentencesSplit,
			exect: true,
		},
		{
			name:  "SplitSentencesNoNewLines",
			s:     types.SplitSentencesNoNewLines,
			exect: true,
		},
		{
			name:  "invalid value",
			s:     types.SplitSentences("invalid value"),
			exect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := c.s.Valid()

			assert.Equal(tt, b, c.exect)
		})
	}
}

func TestPreserveFormattingValid(t *testing.T) {
	cases := []struct {
		name  string
		s     types.PreserveFormatting
		exect bool
	}{
		{
			name:  "PreserveFormattingDisabled",
			s:     types.PreserveFormattingDisabled,
			exect: true,
		},
		{
			name:  "PreserveFormattingEnabled",
			s:     types.PreserveFormattingEnabled,
			exect: true,
		},
		{
			name:  "invalid value",
			s:     types.PreserveFormatting("invalid value"),
			exect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := c.s.Valid()

			assert.Equal(tt, b, c.exect)
		})
	}
}

func TestFormalityValid(t *testing.T) {
	cases := []struct {
		name  string
		s     types.Formality
		l     types.TargetLangCode
		exect bool
	}{
		{
			name:  "FormalityDefault",
			s:     types.FormalityDefault,
			l:     types.TargetLangRU,
			exect: true,
		},
		{
			name:  "FormalityMore",
			s:     types.FormalityMore,
			l:     types.TargetLangRU,
			exect: true,
		},
		{
			name:  "FormalityLess",
			s:     types.FormalityLess,
			l:     types.TargetLangRU,
			exect: true,
		},
		{
			name:  "invalid value",
			s:     types.Formality("invalid value"),
			l:     types.TargetLangRU,
			exect: false,
		},
		{
			name:  "FormalityDefault with not supported target language",
			s:     types.FormalityDefault,
			l:     types.TargetLangJA,
			exect: false,
		},
		{
			name:  "FormalityMore with not supported target language",
			s:     types.FormalityMore,
			l:     types.TargetLangJA,
			exect: false,
		},
		{
			name:  "FormalityLess with not supported target language",
			s:     types.FormalityLess,
			l:     types.TargetLangJA,
			exect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := c.s.Valid(c.l)

			assert.Equal(tt, b, c.exect)
		})
	}
}
