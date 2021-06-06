package types

type TranslateTextResult struct {
	Text                   string   `json:"text"`
	DetectedSourceLanguage LangCode `json:"detected_source_language"`
}

type TranslateTextResponse struct {
	Translations []TranslateTextResult
}
