package types

type LangCode string

const (
	BG   LangCode = "BG"
	CS   LangCode = "CS"
	DA   LangCode = "DA"
	DE   LangCode = "DE"
	EL   LangCode = "EL"
	EN   LangCode = "EN"
	ENGB LangCode = "EN-GB"
	ENUS LangCode = "EN-US"
	ES   LangCode = "ES"
	ET   LangCode = "ET"
	FI   LangCode = "FI"
	FR   LangCode = "FR"
	HU   LangCode = "HU"
	IT   LangCode = "IT"
	JA   LangCode = "JA"
	LT   LangCode = "LT"
	LV   LangCode = "LV"
	NL   LangCode = "NL"
	PL   LangCode = "PL"
	PT   LangCode = "PT"
	PTBR LangCode = "PT-BR"
	PTPT LangCode = "PT-PT"
	RO   LangCode = "RO"
	RU   LangCode = "RU"
	SK   LangCode = "SK"
	SL   LangCode = "SL"
	SV   LangCode = "SV"
	ZH   LangCode = "ZH"
)

var langMap map[LangCode]string = map[LangCode]string{
	BG:   "Bulgarian",
	CS:   "Czech",
	DA:   "Danish",
	DE:   "German",
	EL:   "Greek",
	EN:   "English",
	ENGB: "English (British)",
	ENUS: "English (American)",
	ES:   "Spanish",
	ET:   "Estonian",
	FI:   "Finnish",
	FR:   "French",
	HU:   "Hungarian",
	IT:   "Italian",
	JA:   "Japanese",
	LT:   "Lithuanian",
	LV:   "Latvian",
	NL:   "Dutch",
	PL:   "Polish",
	PT:   "Portuguese",
	PTBR: "Portuguese (Brazilian)",
	PTPT: "Portuguese",
	RO:   "Romanian",
	RU:   "Russian",
	SK:   "Slovak",
	SL:   "Slovenian",
	SV:   "Swedish",
	ZH:   "Chinese",
}

func (l LangCode) Description() string {
	if desc, ok := langMap[l]; ok {
		return desc
	}
	return "undefined lang"
}
