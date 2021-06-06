package main

import (
	"context"
	"deepl-sdk-go"
	"deepl-sdk-go/types"
	"fmt"
	"os"
)

func main() {
	authnKey := os.Getenv("DEEPL_AUTHN_KEY")

	client := deepl.NewClient(authnKey)

	text := []string{
		"こんにちは",
		"これはサンプルテキストです。",
	}
	params := &types.TranslateTextParams{
		AuthKey:    client.AuthenticationKey,
		TargetLang: types.TargetLangEN,
		Text:       text,
	}

	res, err := client.TranslateText(context.TODO(), params)

	if err != nil {
		fmt.Println(err)
	}

	for i := range res.Translations {
		fmt.Printf("%s -> %s\n", text[i], res.Translations[i].Text)
	}
}
