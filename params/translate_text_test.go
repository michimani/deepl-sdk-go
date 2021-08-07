package params_test

import (
	"io"
	"testing"

	"github.com/michimani/deepl-sdk-go/params"
	"github.com/michimani/deepl-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestTranslateTextSetAuthnKey(t *testing.T) {
	k := "test-authn-key"
	p := params.TranslateTextParams{}

	p.SetAuthnKey(k)

	assert.Equal(t, k, p.AuthKey)
}

func TestTranslateTextBody(t *testing.T) {
	cases := []struct {
		name   string
		params params.TranslateTextParams
		expect string
	}{
		{
			name: "normal: all",
			params: params.TranslateTextParams{
				AuthKey:            "test-authn-key",
				Text:               []string{"text1", "text2"},
				SourceLang:         types.SourceLangEN,
				TargetLang:         types.TargetLangRU,
				SplitSentences:     types.SplitSentencesNoSplit,
				PreserveFormatting: types.PreserveFormattingDisabled,
				Formality:          types.FormalityMore,
			},
			expect: "auth_key=test-authn-key&formality=more&preserve_formatting=0&source_lang=EN&split_sentences=0&target_lang=RU&text=text1&text=text2",
		},
		{
			name: "normal: all with white space",
			params: params.TranslateTextParams{
				AuthKey:            "test key",
				Text:               []string{"text 1", "text 2"},
				SourceLang:         types.SourceLangEN,
				TargetLang:         types.TargetLangRU,
				SplitSentences:     types.SplitSentencesSplit,
				PreserveFormatting: types.PreserveFormattingEnabled,
				Formality:          types.FormalityMore,
			},
			expect: "auth_key=test+key&formality=more&preserve_formatting=1&source_lang=EN&split_sentences=1&target_lang=RU&text=text+1&text=text+2",
		},
		{
			name: "normal: all with invalid value",
			params: params.TranslateTextParams{
				AuthKey:            "test key",
				Text:               []string{"text 1", "text 2"},
				SourceLang:         types.SourceLangEN,
				TargetLang:         types.TargetLangRU,
				SplitSentences:     "invalid value",
				PreserveFormatting: "invalid value",
				Formality:          "invalid value",
			},
			expect: "auth_key=test+key&source_lang=EN&target_lang=RU&text=text+1&text=text+2",
		},
		{
			name: "normal: all ignore formality option",
			params: params.TranslateTextParams{
				AuthKey:            "test-authn-key",
				Text:               []string{"text1", "text2"},
				SourceLang:         types.SourceLangEN,
				TargetLang:         types.TargetLangJA,
				SplitSentences:     types.SplitSentencesNoSplit,
				PreserveFormatting: types.PreserveFormattingDisabled,
				Formality:          types.FormalityMore,
			},
			expect: "auth_key=test-authn-key&preserve_formatting=0&source_lang=EN&split_sentences=0&target_lang=JA&text=text1&text=text2",
		},
		{
			name:   "normal: empty",
			params: params.TranslateTextParams{},
			expect: "auth_key=&target_lang=",
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
