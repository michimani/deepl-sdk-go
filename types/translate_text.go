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
