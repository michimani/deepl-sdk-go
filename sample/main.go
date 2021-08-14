package main

import (
	"context"
	"fmt"

	"github.com/michimani/deepl-sdk-go"
	"github.com/michimani/deepl-sdk-go/params"
	"github.com/michimani/deepl-sdk-go/types"
)

func main() {
	client, err := deepl.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	usage(client)
	translateText(client)
	targetLanguages(client)
}

func usage(c *deepl.Client) {
	fmt.Println("UsageAPI sample")
	params := &params.UsageParams{}
	res, errRes, err := c.Usage(context.TODO(), params)

	if err != nil {
		fmt.Println(err)
	}

	if errRes != nil {
		fmt.Println("ErrorResponse", errRes.Message)
	}

	fmt.Printf("character_count: %d\n", res.CharacterCount)
	fmt.Printf("character_limit: %d\n", res.CharacterLimit)
}

func translateText(c *deepl.Client) {
	fmt.Println("TranslateTextAPI sample (JA to EN)")
	text := []string{
		"こんにちは",
		"これは\nサンプルテキストです。",
	}
	params := &params.TranslateTextParams{
		TargetLang:         types.TargetLangEN,
		Text:               text,
		SplitSentences:     types.SplitSentencesNoSplit,
		PreserveFormatting: types.PreserveFormattingDisabled,
	}

	res, errRes, err := c.TranslateText(context.TODO(), params)

	if err != nil {
		fmt.Println(err)
	}

	if errRes != nil {
		fmt.Println("ErrorResponse", errRes.Message)
	}

	for i := range res.Translations {
		fmt.Printf("%s -> %s\n", text[i], res.Translations[i].Text)
	}
}

func targetLanguages(c *deepl.Client) {
	fmt.Println("LanguagesAPI sample (get Target Languages)")
	res, errRes, err := c.TargetLanguages(context.TODO())

	if err != nil {
		fmt.Println(err)
	}

	if errRes != nil {
		fmt.Println("ErrorResponse", errRes.Message)
	}

	for _, tl := range *res {
		fmt.Printf("code:%s name:%s\n", tl.Language, tl.Name)
	}
}
