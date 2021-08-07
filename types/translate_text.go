package types

type SplitSentences string

const (
	// no splitting at all, whole input is treated as one sentence
	SplitSentencesNoSplit = "0"

	// splits on interpunction and on newlines (default)
	SplitSentencesSplit = "1"

	// splits on interpunction only, ignoring newlines
	SplitSentencesNoNewLines = "nonewlines"
)

func (s SplitSentences) Valid() bool {
	if s != SplitSentencesNoSplit && s != SplitSentencesSplit && s != SplitSentencesNoNewLines {
		return false
	}

	return true
}

type PreserveFormatting string

const (
	// default
	PreserveFormattingDisabled = "0"

	PreserveFormattingEnabled = "1"
)

func (p PreserveFormatting) Valid() bool {
	if p != PreserveFormattingDisabled && p != PreserveFormattingEnabled {
		return false
	}

	return true
}

type Formality string

const (
	FormalityDefault = "default"

	// for a more formal language
	FormalityMore = "more"

	// for a more informal language
	FormalityLess = "less"
)

var validTargetLanguages map[TargetLangCode]struct{} = map[TargetLangCode]struct{}{
	TargetLangDE:   {},
	TargetLangFR:   {},
	TargetLangIT:   {},
	TargetLangES:   {},
	TargetLangNL:   {},
	TargetLangPL:   {},
	TargetLangPTPT: {},
	TargetLangPTBR: {},
	TargetLangRU:   {},
}

func (f Formality) Valid(t TargetLangCode) bool {
	// This feature only works for certain target languages.
	if _, ok := validTargetLanguages[t]; !ok {
		return false
	}

	if f != FormalityDefault && f != FormalityMore && f != FormalityLess {
		return false
	}

	return true
}
