package types

type LangType string
type LangCode string
type TargetLangCode LangCode
type SourceLangCode LangCode

const (
	LangTypeTarget LangType = "target"
	LangTypeSource LangType = "source"
)

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

// Language codes that can be specified as source_lang
const (
	SourceLangBG SourceLangCode = SourceLangCode(BG)
	SourceLangCS SourceLangCode = SourceLangCode(CS)
	SourceLangDA SourceLangCode = SourceLangCode(DA)
	SourceLangDE SourceLangCode = SourceLangCode(DE)
	SourceLangEL SourceLangCode = SourceLangCode(EL)
	SourceLangEN SourceLangCode = SourceLangCode(EN)
	SourceLangES SourceLangCode = SourceLangCode(ES)
	SourceLangET SourceLangCode = SourceLangCode(ET)
	SourceLangFI SourceLangCode = SourceLangCode(FI)
	SourceLangFR SourceLangCode = SourceLangCode(FR)
	SourceLangHU SourceLangCode = SourceLangCode(HU)
	SourceLangIT SourceLangCode = SourceLangCode(IT)
	SourceLangJA SourceLangCode = SourceLangCode(JA)
	SourceLangLT SourceLangCode = SourceLangCode(LT)
	SourceLangLV SourceLangCode = SourceLangCode(LV)
	SourceLangNL SourceLangCode = SourceLangCode(NL)
	SourceLangPL SourceLangCode = SourceLangCode(PL)
	SourceLangPT SourceLangCode = SourceLangCode(PT)
	SourceLangRO SourceLangCode = SourceLangCode(RO)
	SourceLangRU SourceLangCode = SourceLangCode(RU)
	SourceLangSK SourceLangCode = SourceLangCode(SK)
	SourceLangSL SourceLangCode = SourceLangCode(SL)
	SourceLangSV SourceLangCode = SourceLangCode(SV)
	SourceLangZH SourceLangCode = SourceLangCode(ZH)
)

// Language codes that can be specified as target_lang
const (
	TargetLangBG   TargetLangCode = TargetLangCode(BG)
	TargetLangCS   TargetLangCode = TargetLangCode(CS)
	TargetLangDA   TargetLangCode = TargetLangCode(DA)
	TargetLangDE   TargetLangCode = TargetLangCode(DE)
	TargetLangEL   TargetLangCode = TargetLangCode(EL)
	TargetLangENGB TargetLangCode = TargetLangCode(ENGB)
	TargetLangENUS TargetLangCode = TargetLangCode(ENUS)
	TargetLangEN   TargetLangCode = TargetLangCode(EN)
	TargetLangES   TargetLangCode = TargetLangCode(ES)
	TargetLangET   TargetLangCode = TargetLangCode(ET)
	TargetLangFI   TargetLangCode = TargetLangCode(FI)
	TargetLangFR   TargetLangCode = TargetLangCode(FR)
	TargetLangHU   TargetLangCode = TargetLangCode(HU)
	TargetLangIT   TargetLangCode = TargetLangCode(IT)
	TargetLangJA   TargetLangCode = TargetLangCode(JA)
	TargetLangLT   TargetLangCode = TargetLangCode(LT)
	TargetLangLV   TargetLangCode = TargetLangCode(LV)
	TargetLangNL   TargetLangCode = TargetLangCode(NL)
	TargetLangPL   TargetLangCode = TargetLangCode(PL)
	TargetLangPTPT TargetLangCode = TargetLangCode(PTPT)
	TargetLangPTBR TargetLangCode = TargetLangCode(PTBR)
	TargetLangPT   TargetLangCode = TargetLangCode(PT)
	TargetLangRO   TargetLangCode = TargetLangCode(RO)
	TargetLangRU   TargetLangCode = TargetLangCode(RU)
	TargetLangSK   TargetLangCode = TargetLangCode(SK)
	TargetLangSL   TargetLangCode = TargetLangCode(SL)
	TargetLangSV   TargetLangCode = TargetLangCode(SV)
	TargetLangZH   TargetLangCode = TargetLangCode(ZH)
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
